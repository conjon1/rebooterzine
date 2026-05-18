#!/usr/bin/env sh

# ─── Phone Tunnel Config ────────────────────────────────────────────────────
GNET_USER="u0_a242"
GNET_PORT="8022"
GNET_PROXY="1080"
GNET_REDSOCKS="12345"

_gnet_wan_iface() { ip route show default | awk 'NR==1 {print $5}'; }
_gnet_lan_iface() {
    if [[ -n "$GNET_LAN_IFACE" ]]; then
        echo "$GNET_LAN_IFACE"
        return
    fi
    local wan
    wan=$(_gnet_wan_iface)
    ip -o link show | awk -F': ' '{print $2}' \
        | grep -E '^(wlan|wlp)' \
        | grep -v "^${wan}$" \
        | head -1
}

_tunnel_active() { ss -tuln | grep -q ":${GNET_PROXY}"; }

_redsocks_config() {
    cat <<EOF > /tmp/redsocks.conf
base {
    log_debug = off;
    log_info  = off;
    log       = "syslog:daemon";
    daemon    = on;
    redirector = iptables;
}
redsocks {
    local_ip   = 0.0.0.0;
    local_port = ${GNET_REDSOCKS};
    ip         = 127.0.0.1;
    port       = ${GNET_PROXY};
    type       = socks5;
}
EOF
}

_dnscrypt_config() {
    local DNSSEC="${1:-false}"
    cat <<EOF > /tmp/dnscrypt-proxy.toml
listen_addresses  = ['127.0.0.1:53']
server_names      = ['mullvad-doh']
max_clients       = 250
ipv4_servers      = true
ipv6_servers      = false
dnscrypt_servers  = false
doh_servers       = true
require_dnssec    = ${DNSSEC}
require_nolog     = true
require_nofilter  = true
log_level         = 0

[sources.public-resolvers]
  urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/public-resolvers.md']
  cache_file = '/tmp/dnscrypt-resolvers.md'
  minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
  refresh_delay = 72
EOF
}

_iptables_up() {
    local LAN_IFACE WAN_IFACE
    LAN_IFACE=$(_gnet_lan_iface)
    WAN_IFACE=$(_gnet_wan_iface)

    sudo iptables -t nat -N REDSOCKS 2>/dev/null

    for net in 0.0.0.0/8 10.0.0.0/8 127.0.0.0/8 169.254.0.0/16 \
               172.16.0.0/12 192.168.0.0/16 224.0.0.0/4 240.0.0.0/4; do
        sudo iptables -t nat -A REDSOCKS -d "$net" -j RETURN
    done
    sudo iptables -t nat -A REDSOCKS -p tcp -j REDIRECT --to-ports "${GNET_REDSOCKS}"

    sudo iptables -t nat -A OUTPUT -p tcp -j REDSOCKS

    if [[ -n "$LAN_IFACE" ]]; then
        echo "lan      ${LAN_IFACE} → redsocks"

        sudo iptables -t nat -A PREROUTING -i "${LAN_IFACE}" -p tcp -j REDSOCKS

        sudo iptables -t nat -A POSTROUTING -o "${WAN_IFACE}" -j MASQUERADE

        sudo iptables -A FORWARD -i "${LAN_IFACE}" -j ACCEPT
        sudo iptables -A FORWARD -o "${LAN_IFACE}" -m state --state RELATED,ESTABLISHED -j ACCEPT

        sudo sysctl -w net.ipv4.ip_forward=1 > /dev/null
        echo "ip_forward enabled"
    else
        echo "warning: LAN iface not found -- LAN sharing disabled"
        echo "         set GNET_LAN_IFACE=wlan0 (or your iface) to override"
    fi
}

_iptables_down() {
    local LAN_IFACE WAN_IFACE
    LAN_IFACE=$(_gnet_lan_iface)
    WAN_IFACE=$(_gnet_wan_iface)

    sudo iptables -t nat -D OUTPUT -p tcp -j REDSOCKS 2>/dev/null

    if [[ -n "$LAN_IFACE" ]]; then
        sudo iptables -t nat -D PREROUTING -i "${LAN_IFACE}" -p tcp -j REDSOCKS 2>/dev/null
        sudo iptables -t nat -D POSTROUTING -o "${WAN_IFACE}" -j MASQUERADE 2>/dev/null
        sudo iptables -D FORWARD -i "${LAN_IFACE}" -j ACCEPT 2>/dev/null
        sudo iptables -D FORWARD -o "${LAN_IFACE}" -m state --state RELATED,ESTABLISHED -j ACCEPT 2>/dev/null
    fi

    sudo iptables -t nat -F REDSOCKS 2>/dev/null
    sudo iptables -t nat -X REDSOCKS 2>/dev/null

    sudo sysctl -w net.ipv4.ip_forward=0 > /dev/null
}

function gnet() {
    # ── flag parsing ─────────────────────────────────────────────────────────
    local MODE="default"
    for arg in "$@"; do
        case "$arg" in
            -fast) MODE="fast" ;;
            -safe) MODE="safe" ;;
            *) echo "usage: gnet [-fast|-safe]"; return 1 ;;
        esac
    done

    case "$MODE" in
        fast) echo "mode     fast (lower latency, DNS unencrypted)" ;;
        safe) echo "mode     safe (UDP blocked, DNSSEC enforced)" ;;
        *)    echo "mode     default" ;;
    esac

    # ── iface detection (awk, not fragile read) ───────────────────────────────
    local WAN_IFACE PHONE_IP
    WAN_IFACE=$(_gnet_wan_iface)
    PHONE_IP=$(ip route show default | awk 'NR==1 {print $3}')

    if [[ -z "$WAN_IFACE" ]]; then
        echo "error: no internet connection found -- is USB tethering on?"
        return 1
    fi
    if [[ -z "$PHONE_IP" ]]; then
        echo "error: could not determine phone gateway IP"
        return 1
    fi

    echo "wan      ${WAN_IFACE}"
    echo "lan      $(_gnet_lan_iface || echo 'not detected')"
    echo "phone    ${PHONE_IP}"

    # ── ssh tunnel ───────────────────────────────────────────────────────────
    if _tunnel_active; then
        echo "tunnel already active on port ${GNET_PROXY}"
    else
        echo "checking port ${GNET_PORT}..."
        if ! timeout 3 bash -c "echo >/dev/tcp/${PHONE_IP}/${GNET_PORT}" 2>/dev/null; then
            echo "error: port ${GNET_PORT} unreachable on ${PHONE_IP}"
            echo "       is sshd running in Termux? is a VPN blocking it?"
            return 1
        fi

        echo "opening socks5 tunnel on port ${GNET_PROXY}..."

        # -fast: chacha20 — better on ARM phones without AES-NI
        # default/safe: aes128-gcm — faster on x86 with AES-NI
        local CIPHER="aes128-gcm@openssh.com"
        [[ "$MODE" == "fast" ]] && CIPHER="chacha20-poly1305@openssh.com"

        ssh -f -N -D "${GNET_PROXY}" -p "${GNET_PORT}" \
            -c "${CIPHER}" \
            -o Compression=no \
            -o IPQoS=throughput \
            -o ConnectTimeout=5 \
            "${GNET_USER}@${PHONE_IP}" || {
            echo "error: ssh failed to start tunnel"
            return 1
        }

        sleep 1
        if ! _tunnel_active; then
            echo "error: tunnel did not open on port ${GNET_PROXY}"
            return 1
        fi
    fi

    echo "starting redsocks..."
    _redsocks_config
    sudo redsocks -c /tmp/redsocks.conf || {
        echo "error: redsocks failed to start"
        return 1
    }

    echo "applying iptables rules..."
    _iptables_up

    if [[ "$MODE" == "safe" ]]; then
        echo "blocking udp (safe mode)..."
        sudo iptables -A OUTPUT -p udp ! -d 127.0.0.0/8 -j REJECT
        local LAN_IFACE
        LAN_IFACE=$(_gnet_lan_iface)
        [[ -n "$LAN_IFACE" ]] && \
            sudo iptables -A FORWARD -i "${LAN_IFACE}" -p udp -j REJECT
    fi

    if [[ "$MODE" == "fast" ]]; then
        echo "dns      phone gateway (fast mode -- DoH skipped)"
        sudo cp /etc/resolv.conf /tmp/resolv.conf.bak
    else
        local DNSSEC="false"
        [[ "$MODE" == "safe" ]] && DNSSEC="true"

        echo "starting dnscrypt-proxy..."
        _dnscrypt_config "$DNSSEC"
        sudo dnscrypt-proxy -config /tmp/dnscrypt-proxy.toml \
            -pidfile /tmp/dnscrypt-proxy.pid \
            > /tmp/dnscrypt-proxy.log 2>&1 &
        sleep 2
        if ! pgrep -x dnscrypt-proxy > /dev/null; then
            echo "error: dnscrypt-proxy failed to start (check /tmp/dnscrypt-proxy.log)"
            return 1
        fi

        sudo cp /etc/resolv.conf /tmp/resolv.conf.bak
        echo "nameserver 127.0.0.1" | sudo tee /etc/resolv.conf > /dev/null
    fi

    echo "disabling ipv6..."
    sudo sysctl -w net.ipv6.conf.all.disable_ipv6=1     > /dev/null
    sudo sysctl -w net.ipv6.conf.default.disable_ipv6=1 > /dev/null

    echo "tunnel active -- mode: ${MODE}"
}

function gdown() {
    echo "tearing down..."

    _iptables_down
    echo "iptables rules removed"

    if pgrep -x redsocks > /dev/null; then
        sudo pkill redsocks && echo "redsocks stopped"
    fi

    local TUNNEL_PID
    TUNNEL_PID=$(ss -lptn "sport = :${GNET_PROXY}" \
        | awk -F'pid=' 'NR>1 { split($2, a, ","); print a[1]; exit }')

    if [[ -n "$TUNNEL_PID" ]]; then
        kill "$TUNNEL_PID" \
            && echo "tunnel stopped (pid $TUNNEL_PID)" \
            || echo "error: could not kill tunnel pid $TUNNEL_PID"
    else
        echo "no active tunnel on port ${GNET_PROXY}"
    fi

    if pgrep -x dnscrypt-proxy > /dev/null; then
        sudo pkill dnscrypt-proxy && echo "dnscrypt-proxy stopped"
    fi

    if [[ -f /tmp/resolv.conf.bak ]]; then
        sudo cp /tmp/resolv.conf.bak /etc/resolv.conf
        sudo rm /tmp/resolv.conf.bak
        echo "dns restored"
    fi

    sudo iptables -D OUTPUT -p udp ! -d 127.0.0.0/8 -j REJECT 2>/dev/null
    local LAN_IFACE
    LAN_IFACE=$(_gnet_lan_iface)
    [[ -n "$LAN_IFACE" ]] && \
        sudo iptables -D FORWARD -i "${LAN_IFACE}" -p udp -j REJECT 2>/dev/null

    sudo sysctl -w net.ipv6.conf.all.disable_ipv6=0     > /dev/null
    sudo sysctl -w net.ipv6.conf.default.disable_ipv6=0 > /dev/null
    echo "ipv6 restored"
}

function gstatus() {
    local WAN_IFACE PHONE_IP LAN_IFACE TUNNEL_PID

    WAN_IFACE=$(_gnet_wan_iface)
    PHONE_IP=$(ip route show default | awk 'NR==1 {print $3}')
    LAN_IFACE=$(_gnet_lan_iface)

    echo "[ ifaces ]"
    echo "  wan      ${WAN_IFACE:-unknown}  (phone USB tether)"
    echo "  lan      ${LAN_IFACE:-not found}  (Pi AP / homelab)"
    echo "  phone    ${PHONE_IP:-unknown}"
    echo "  forward  $(sysctl -n net.ipv4.ip_forward)  (1=LAN sharing active)"

    echo "[ ssh tunnel ]"
    if _tunnel_active; then
        TUNNEL_PID=$(ss -lptn "sport = :${GNET_PROXY}" \
            | awk -F'pid=' 'NR>1 { split($2, a, ","); print a[1]; exit }')
        echo "  status   active"
        echo "  port     ${GNET_PROXY}"
        echo "  pid      ${TUNNEL_PID:-unknown}"
        echo "  via      ${PHONE_IP:-unknown}  (${WAN_IFACE:-unknown})"
    else
        echo "  status   inactive"
    fi

    echo "[ redsocks ]"
    if pgrep -x redsocks > /dev/null; then
        echo "  status   active"
        echo "  port     ${GNET_REDSOCKS}"
        if sudo iptables -t nat -L OUTPUT -n | grep -q REDSOCKS; then
            echo "  OUTPUT   -> active  (Arch local traffic)"
        else
            echo "  OUTPUT   -> missing"
        fi
        if [[ -n "$LAN_IFACE" ]] && sudo iptables -t nat -L PREROUTING -n | grep -q REDSOCKS; then
            echo "  PREROUTING -> active  (LAN sharing on: Pi/Talos get internet)"
        else
            echo "  PREROUTING -> missing (LAN sharing off)"
        fi
    else
        echo "  status   inactive"
    fi

    echo "[ ip ]"
    local PROXY_IP REAL_IP
    PROXY_IP=$(curl -s --max-time 5 --socks5-hostname "127.0.0.1:${GNET_PROXY}" ifconfig.me 2>/dev/null)
    REAL_IP=$(curl -s --max-time 5 --interface "${WAN_IFACE}" ifconfig.me 2>/dev/null)
    echo "  tunnel   ${PROXY_IP:-unavailable}"
    echo "  real     ${REAL_IP:-unavailable}"
    if [[ -n "$PROXY_IP" && -n "$REAL_IP" && "$PROXY_IP" != "$REAL_IP" ]]; then
        echo "  masked   yes"
    elif [[ -z "$PROXY_IP" ]]; then
        echo "  masked   unknown (tunnel unreachable)"
    else
        echo "  masked   no -- traffic may not route through tunnel"
    fi

    echo "[ dns ]"
    if pgrep -x dnscrypt-proxy > /dev/null; then
        echo "  dnscrypt active  (127.0.0.1 → Mullvad DoH → tunnel)"
        # FIX: test resolver works, not compare to exit IP (those are different things)
        local DNS_TEST
        DNS_TEST=$(dig +short +time=3 cloudflare.com @127.0.0.1 2>/dev/null | grep -E '^[0-9]+\.' | head -1)
        if [[ -n "$DNS_TEST" ]]; then
            echo "  resolving    ok  (${DNS_TEST})"
        else
            echo "  resolving    FAIL -- dnscrypt may be broken (check /tmp/dnscrypt-proxy.log)"
        fi
    else
        echo "  dnscrypt inactive -- DNS may be leaking"
        local DNS_SERVER
        DNS_SERVER=$(awk '/^nameserver/{print $2; exit}' /etc/resolv.conf)
        echo "  server   ${DNS_SERVER:-unknown}"
    fi

    echo "[ system ]"
    echo "  iface    ${WAN_IFACE:-unknown}"
    echo "  ttl      $(sysctl -n net.ipv4.ip_default_ttl)"
}

. "$HOME/.local/bin/env"



