- [CHAINS KEY](#org2e3cab0)
  - [══════════════════════════════════════════](#org2f3e08d)
- [TIER 1 — WEEKEND PROJECTS (1–2 days each)](#org21f4f02)
  - [══════════════════════════════════════════](#org80cfe0b)
  - [[NET-W1] Set Up Kali Linux VM](#orga6ee71a)
  - [[NET-W2] Home Network Recon with Nmap](#org981cbeb)
  - [[NET-W3] Wireshark Traffic Analysis](#org4e1a385)
  - [[NET-W4] Crack a WPA2 Handshake (Own Router)](#org49b4eef)
  - [[NET-W5] Netcat Fundamentals — Shells & Transfers](#org191cbd8)
  - [[WEB-W1] SQLi & XSS on DVWA](#org8bd2d6c)
  - [[WEB-W2] Burp Suite Interception Basics](#orgb8531a1)
  - [[WEB-W3] Command Injection & File Inclusion on DVWA](#org373f764)
  - [[WEB-W4] Subdomain Enumeration & Recon](#org849609b)
  - [[OSINT-W1] Google Dorking & Shodan Recon](#org32ccdcf)
  - [[OSINT-W2] Email & Username OSINT](#org6e63353)
  - [[OSINT-W3] DNS Enumeration](#orge1528d1)
  - [[SCRPT-W1] Python or Go Port Scanner](#org65be947)
  - [[SCRPT-W2] Caesar Cipher → Basic Crypto in Python/go](#org4d007a1)
  - [[CRYPT-W1] Hash Cracking with Hashcat](#org4e13731)
  - [[CRYPT-W2] Steganography — Hide & Find](#org685dc19)
  - [[DEF-W1] SSH Hardening + Key Auth](#orgcc42e78)
  - [[DEF-W2] Firewall Rules with UFW/iptables](#orga7cdedf)
  - [[MAL-W1] Static Malware Analysis](#orgf13104b)
  - [[MAL-W2] Analyze a Malicious PCAP](#org0c194b3)
  - [[RE-W1] Linux Privilege Escalation Basics](#org71b9286)
  - [[DEF-W3] GPG Encryption — Files & Email](#org294986f)
  - [[NET-W6] Set Up WireGuard VPN](#orgdaf6082)
  - [[OSINT-W4] Digital Forensics — File Recovery](#org237e59a)
  - [[WEB-W5] JWT Attack Lab](#orgb76f27c)
  - [[CTF-W1] Complete 5 picoCTF Beginner Challenges](#org837b8c3)
  - [[SCRPT-W3] Log Parser in Python](#org5a142f3)
  - [[NET-W7] Proxy Chains + Tor Setup](#org1c7b29a)
  - [[WEB-W6] HTTP Security Headers Audit Tool](#orgee9b839)
  - [[DEF-W4] Set Up Basic Honeypot (Cowrie)](#org89d7a3a)
  - [[RE-W2] Reverse a Simple Crackme Binary](#org03a0949)
  - [[AD-W1] Active Directory Concepts + Lab Setup](#orga4fdf46)
  - [[CLOUD-W1] AWS Free Tier — IAM Misconfig Hunt](#orga14244e)
  - [[SCRPT-W4] Build HTTP Header Fuzzer](#orge35c309)
  - [[NET-W8] TryHackMe — Complete 2 Beginner Rooms](#org3e22c14)
  - [══════════════════════════════════════════](#org9c3093e)
- [TIER 2 — WEEK PROJECTS (3–7 days each)](#orga100c6e)
  - [══════════════════════════════════════════](#orgd7ddb70)
    - [[SCRPT-WK1] Full Go or Python Recon Framework](#org73eb690)
    - [[NET-WK1] Full Pentest: VulnHub Beginner Machine](#org6153845)
    - [[WEB-WK1] Complete OWASP Top 10 on WebGoat](#org0ee2f8e)
    - [[NET-WK2] Man-in-the-Middle Attack Lab](#orgc7f1b0b)
    - [[DEF-WK1] Set Up ELK Stack SIEM](#org61ffe04)
    - [[DEF-WK2] Honeypot + Log Pipeline](#org6c8e5c0)
    - [[MAL-WK1] Dynamic Malware Analysis in Sandbox](#orgf24fced)
    - [[MAL-WK2] Reverse Engineering with Ghidra](#org469d990)
    - [[NET-WK3] Network Pivoting Lab](#orgdd86167)
    - [[CRYPT-WK1] Implement Crypto Attacks in Python](#orgd4795c6)
    - [[WEB-WK2] SQLi Scanner + SSRF + XXE Lab](#org8c74f76)
    - [[AD-WK1] Active Directory Attack Lab](#orgb433c49)
    - [[SCRPT-WK2] Build a C2 Beaconing Script (Lab Only)](#orga1b2f86)
    - [[DEF-WK3] Set Up Suricata + Zeek IDS](#orgc644dd4)
    - [[WEB-WK3] Full Subdomain + Dir Recon Automation](#org50236b4)
    - [[CTF-WK1] Complete HackTheBox Starting Point (3 Machines)](#org839f371)
    - [[OSINT-WK1] OSINT Framework in Python](#org726ec56)
    - [[NET-WK4] Metasploit Deep Dive](#org7375da2)
    - [[CLOUD-WK1] AWS Misconfig & Container Security Lab](#org371718d)
    - [[CRYPT-WK2] Build a PKI from Scratch](#orgcb0b788)
    - [[MAL-WK3] YARA Rules — Write & Test](#org5e669e5)
    - [[WEB-WK4] Android App Security Testing](#org445d13f)
    - [[DEF-WK4] Incident Response Lab](#orgc0f3c32)
    - [[NET-WK5] WPA2 PMKID Attack + Evil Twin AP](#org16aea04)
    - [[SCRPT-WK3] Vulnerability Scanner in Python](#orgcca7f54)
    - [[OSINT-WK2] Threat Intel Aggregator](#orgd2fa4b8)
    - [[RE-WK1] Buffer Overflow 101](#org04e23c4)
    - [[WEB-WK5] Build a Vulnerable Web App (for CTF)](#org6566c69)
    - [[DEF-WK5] Zero Trust Network Lab](#orga31003b)
    - [[NET-WK6] Analyze a Real-World CVE + Write PoC](#orgc5bd272)
    - [[CLOUD-WK2] Serverless + API Security Testing](#org32bbf4f)
  - [══════════════════════════════════════════](#org3b4ccbb)
- [TIER 3 — MONTH PROJECTS (3–4 weeks each)](#org923e413)
  - [══════════════════════════════════════════](#orgfa0bb31)
    - [[DEF-MO1] Build a Full Home SOC](#org24efb69)
    - [[NET-MO1] OSCP-Style Multi-Machine Lab + Report](#org2b8348b)
    - [[WEB-MO1] Full Web App Pentest Automation Suite](#orgab96889)
    - [[MAL-MO1] Full Malware Analysis Report on Real Sample](#org9751e45)
    - [[AD-MO1] Active Directory Full Attack + Defense Lab](#orgdfdca13)
    - [[OSINT-MO1] Automated OSINT Platform](#org34f8015)
    - [[CRYPT-MO1] Cryptography Attack Suite + PKI System](#org5e22e86)
    - [[CLOUD-MO1] Cloud Security Audit + Hardening](#org6f956be)
    - [[SCRPT-MO1] Full Pentest Automation Suite (CLI Tool)](#org3b26d9a)
    - [[CTF-MO1] Host a Public CTF Competition](#org899b684)
  - [══════════════════════════════════════════](#org8be5a88)
- [CHAIN COMBO MAPS — SUGGESTED PATHS](#org2be59f2)
  - [══════════════════════════════════════════](#orgfe04580)
    - [PATH A: Network Pentester (Offensive)](#orge2ca6fd)
    - [PATH B: Web/Bug Bounty Hunter](#org0f68b3a)
    - [PATH C: Blue Team / SOC Analyst](#org5ffdcfd)
    - [PATH D: Malware Analyst / Threat Intel](#orgacc0a8f)
    - [PATH E: Security Engineer / Tool Builder](#org1931468)
  - [══════════════════════════════](#org36f3caf)
- [RESOURCES](#orgc95ab10)
  - [══════════════════════════════](#org8c3f7e1)
    - [Platforms And Websites](#org0c6995d)
    - [Essential Tools](#orgf3d9be7)
    - [Free Certifications to Stack](#org9359209)
    - [Time Budget Suggestion](#org13d584a)



<a id="org2e3cab0"></a>

# CHAINS KEY

Projects grouped in 8 chains. Chain projects stack — weekend → week → month. Finish weekend before week. Finish week before month. Cross-chain combos noted.

| Chain | Theme                  | Color Tag   |
|----- |---------------------- |----------- |
| NET   | Network Recon & Attack | :network:   |
| WEB   | Web App Security       | :web:       |
| MAL   | Malware & RE           | :malware:   |
| CRYPT | Cryptography           | :crypto:    |
| DEF   | Defense & Blue Team    | :defense:   |
| SCRPT | Scripting & Tooling    | :scripting: |
| OSINT | OSINT & Recon          | :osint:     |
| AD    | Active Directory       | :ad:        |


<a id="org2f3e08d"></a>

## ══════════════════════════════════════════


<a id="org21f4f02"></a>

# TIER 1 — WEEKEND PROJECTS (1–2 days each)


<a id="org80cfe0b"></a>

## ══════════════════════════════════════════


<a id="orga6ee71a"></a>

## TODO [NET-W1] Set Up Kali Linux VM     :network:

-   Install VirtualBox or VMware
-   Download Kali ISO, install, snapshot clean state
-   Learn basic terminal nav, update system
-   Install guest additions

**\*** Combo: Base for every other project. Do first.


<a id="org981cbeb"></a>

## TODO [NET-W2] Home Network Recon with Nmap     :network:

-   Scan local subnet: `nmap -sn 192.168.x.0/24`
-   Service scan: `nmap -sV -sC -A <target>`
-   Output to XML, read it
-   Understand open ports on your own devices

**\*** Combo: Feeds into Python scanner (SCRPT-W1) and SIEM setup (DEF-WK1)


<a id="org4e1a385"></a>

## TODO [NET-W3] Wireshark Traffic Analysis     :network:

-   Capture live traffic on home net
-   Filter HTTP, DNS, ARP
-   Find plaintext credentials in pcap (test pcap from online)
-   Export objects from HTTP stream

**\*** Combo: Pairs with MitM week project (NET-WK4)


<a id="org49b4eef"></a>

## TODO [NET-W4] Crack a WPA2 Handshake (Own Router)     :network:

-   Use monitor mode + airodump-ng to capture 4-way handshake
-   Deauth a client to force reconnect
-   Crack with hashcat + rockyou.txt
-   Change your WiFi password after. Learn why WPA3 matters.

**\*** NOTE: Own network only. Legal line clear.


<a id="org191cbd8"></a>

## TODO [NET-W5] Netcat Fundamentals — Shells & Transfers     :network:

-   Open listeners, connect clients
-   Send files with nc
-   Reverse shell: `nc -e /bin/bash`
-   Bind shell vs reverse shell — understand difference

**\*** Combo: Foundation for all post-exploitation work


<a id="org8bd2d6c"></a>

## TODO [WEB-W1] SQLi & XSS on DVWA     :web:

-   Install DVWA (Docker or XAMPP)
-   Complete SQLi: manual + sqlmap
-   Complete XSS: reflected, stored, DOM
-   Toggle security levels low→medium→high

**\*** Combo: Directly builds into full OWASP week (WEB-WK1)


<a id="orgb8531a1"></a>

## TODO [WEB-W2] Burp Suite Interception Basics     :web:

-   Set up browser proxy through Burp
-   Intercept, modify, replay requests
-   Use Repeater on DVWA login
-   Use Intruder for basic brute force

**\*** Combo: Essential tool for all web projects


<a id="org373f764"></a>

## TODO [WEB-W3] Command Injection & File Inclusion on DVWA     :web:

-   Command injection: OS commands through web input
-   LFI: read `/etc/passwd` via vulnerable param
-   RFI: include remote malicious file
-   CSRF: forge requests, steal sessions


<a id="org849609b"></a>

## TODO [WEB-W4] Subdomain Enumeration & Recon     :web:osint:

-   Use subfinder, amass on a target (HackerOne public programs)
-   Certificate transparency: crt.sh
-   Directory bruteforce with ffuf: `ffuf -w wordlist -u https://target/FUZZ`
-   Document findings in structured notes


<a id="org32ccdcf"></a>

## TODO [OSINT-W1] Google Dorking & Shodan Recon     :osint:

-   Learn 10 key Google dork operators
-   Find exposed login panels, open dirs, config files
-   Shodan: search for services by banner, CVE
-   Build a personal dork cheatsheet


<a id="org6e63353"></a>

## TODO [OSINT-W2] Email & Username OSINT     :osint:

-   theHarvester: email harvest from domain
-   holehe or Sherlock: username across platforms
-   Have I Been Pwned API lookup
-   Build a target profile (use yourself as test subject)


<a id="orge1528d1"></a>

## TODO [OSINT-W3] DNS Enumeration     :osint:network:

-   dnsrecon, dnsenum on practice domains
-   Zone transfer attempt
-   MX, TXT, NS record analysis
-   Reverse DNS lookup sweep


<a id="org65be947"></a>

## TODO [SCRPT-W1] Python or Go Port Scanner     :scripting:

-   Socket-based TCP port scanner
-   Add threading for speed
-   Service banner grabbing
-   Output to JSON/CSV

**\*** Combo: Base for full recon tool (SCRPT-WK1)


<a id="org4d007a1"></a>

## TODO [SCRPT-W2] Caesar Cipher → Basic Crypto in Python/go     :scripting:crypto:

-   Implement Caesar, Vigenere, XOR cipher
-   Brute-force Caesar without key
-   Frequency analysis for Vigenere
-   Understand why these fail


<a id="org4e13731"></a>

## TODO [CRYPT-W1] Hash Cracking with Hashcat     :crypto:

-   Identify hash types (hash-identifier, hashid)
-   Crack MD5, SHA1, bcrypt with rockyou.txt
-   Rules-based attack with hashcat rules
-   Dictionary vs brute vs combo attack modes


<a id="org685dc19"></a>

## TODO [CRYPT-W2] Steganography — Hide & Find     :crypto:

-   Hide text in image: steghide, LSB
-   Extract: steghide extract, stegsolve
-   Audio steganography: MP3Stego
-   Solve 3 stego CTF challenges


<a id="orgcc42e78"></a>

## TODO [DEF-W1] SSH Hardening + Key Auth     :defense:

-   Disable password auth, enable key-only
-   Change default port, restrict users
-   Set up fail2ban for SSH brute protection
-   Test hardening with nmap from Kali


<a id="orga7cdedf"></a>

## TODO [DEF-W2] Firewall Rules with UFW/iptables     :defense:

-   Default deny inbound policy
-   Allow only necessary ports
-   Log dropped packets
-   Test rules from external VM


<a id="orgf13104b"></a>

## TODO [MAL-W1] Static Malware Analysis     :malware:

-   strings, file, xxd on a safe malware sample (MalwareBazaar)
-   Extract IPs/domains/registry keys from strings
-   PE header analysis with PEview or pestudio
-   Identify packing/obfuscation signs


<a id="org0c194b3"></a>

## TODO [MAL-W2] Analyze a Malicious PCAP     :malware:network:

-   Download malware traffic pcap (malware-traffic-analysis.net)
-   Identify C2 beaconing patterns
-   Extract indicators of compromise (IOCs)
-   Write a short analysis report


<a id="org71b9286"></a>

## TODO [RE-W1] Linux Privilege Escalation Basics     :re:

-   GTFOBins: SUID binary abuse
-   Writable /etc/passwd, cron abuse
-   sudo -l misconfigs
-   linPEAS on a VulnHub machine


<a id="org294986f"></a>

## TODO [DEF-W3] GPG Encryption — Files & Email     :defense:crypto:

-   Generate GPG keypair
-   Encrypt/decrypt files
-   Sign and verify
-   Export/import public keys


<a id="orgdaf6082"></a>

## TODO [NET-W6] Set Up WireGuard VPN     :network:defense:

-   Install WireGuard on a VPS or local VM
-   Generate peer keys, configure tunnels
-   Route traffic through tunnel
-   Verify with Wireshark — confirm encryption


<a id="org237e59a"></a>

## TODO [OSINT-W4] Digital Forensics — File Recovery     :osint:

-   Create a disk image with dd
-   Recover deleted files with autopsy + foremost
-   Analyze file metadata (exiftool)
-   Build a basic forensics checklist


<a id="orgb76f27c"></a>

## TODO [WEB-W5] JWT Attack Lab     :web:

-   Decode JWT (jwt.io)
-   none algorithm attack
-   Brute force weak HS256 secret (hashcat)
-   Key confusion attack (RS256→HS256)


<a id="org837b8c3"></a>

## TODO [CTF-W1] Complete 5 picoCTF Beginner Challenges     :web:crypto:re:

-   Pick challenges across: crypto, forensics, web, general skills
-   Document solve methodology for each
-   Learn to use CyberChef
-   Join a CTF Discord for hints


<a id="org5a142f3"></a>

## TODO [SCRPT-W3] Log Parser in Python     :scripting:defense:

-   Parse /var/log/auth.log for failed logins
-   Count IPs, flag threshold breaches
-   Output alert summary
-   Extend to syslog, apache access logs


<a id="org1c7b29a"></a>

## TODO [NET-W7] Proxy Chains + Tor Setup     :network:

-   Install tor + proxychains
-   Route nmap through proxychains
-   Understand Tor limitations for pentesting
-   Test anonymity with whatismyip


<a id="orgee9b839"></a>

## TODO [WEB-W6] HTTP Security Headers Audit Tool     :web:scripting:

-   Python script: fetch headers from any URL
-   Check: CSP, HSTS, X-Frame-Options, CORS
-   Score and report missing headers
-   Run against 10 real sites (ethically)


<a id="org89d7a3a"></a>

## TODO [DEF-W4] Set Up Basic Honeypot (Cowrie)     :defense:

-   Install Cowrie SSH honeypot
-   Expose on a VPS or local VM
-   Watch logs for hit attempts
-   Extract attacker IPs and commands


<a id="org03a0949"></a>

## TODO [RE-W2] Reverse a Simple Crackme Binary     :re:

-   Download crackme from crackmes.one (easy level)
-   Use ltrace/strace first
-   Open in Ghidra — find password check logic
-   Patch binary to bypass check


<a id="orga4fdf46"></a>

## TODO [AD-W1] Active Directory Concepts + Lab Setup     :ad:

-   Install Windows Server eval VM
-   Promote to domain controller
-   Create OUs, users, groups
-   Join a Windows 10 VM to the domain


<a id="orga14244e"></a>

## TODO [CLOUD-W1] AWS Free Tier — IAM Misconfig Hunt     :cloud:

-   Create AWS free tier account
-   Create intentionally misconfigured IAM (for lab)
-   Use ScoutSuite or Prowler to audit
-   Enumerate with AWS CLI using overprivileged user


<a id="orge35c309"></a>

## TODO [SCRPT-W4] Build HTTP Header Fuzzer     :scripting:web:

-   Python requests — iterate custom headers
-   Fuzz Host, X-Forwarded-For, Content-Type
-   Look for 500 errors or behavioral changes
-   Test on DVWA or local lab app


<a id="org3e22c14"></a>

## TODO [NET-W8] TryHackMe — Complete 2 Beginner Rooms     :network:

-   Recommended: &ldquo;Basic Pentesting&rdquo;, &ldquo;Startup&rdquo;
-   Document methodology: recon → exploit → flags
-   Note tools used and commands
-   Subscribe to free tier


<a id="org9c3093e"></a>

## ══════════════════════════════════════════


<a id="orga100c6e"></a>

# TIER 2 — WEEK PROJECTS (3–7 days each)


<a id="orgd7ddb70"></a>

## ══════════════════════════════════════════


<a id="org73eb690"></a>

### TODO [SCRPT-WK1] Full Go or Python Recon Framework     :scripting:network:osint:

-   Combine: port scanner + subdomain enum + DNS recon + header check
-   Single CLI tool with argparse
-   Output to JSON report + markdown summary
-   Add screenshot capability (selenium headless)

**\*** Combo: Ports directly into full pentest suite (SCRPT-MO1)


<a id="org6153845"></a>

### TODO [NET-WK1] Full Pentest: VulnHub Beginner Machine     :network:web:re:

-   Download: Mr-Robot, Kioptrix, or Basic Pentesting 1
-   Recon → foothold → privesc → root
-   Document every step in markdown
-   Write a mini pentest report

**\*** Combo: Chain 3+ machines → OSCP prep (NET-MO1)


<a id="org0ee2f8e"></a>

### TODO [WEB-WK1] Complete OWASP Top 10 on WebGoat     :web:

-   Install WebGoat (Java or Docker)
-   Complete all OWASP Top 10 lessons
-   A01 Broken Access Control through A10 SSRF
-   Write one-pager summary per vuln

**\*** Combo: Unlocks web pentest automation month project


<a id="orgc7f1b0b"></a>

### TODO [NET-WK2] Man-in-the-Middle Attack Lab     :network:

-   ARP spoofing: arpspoof + Wireshark in isolated VM lab
-   SSL stripping with bettercap
-   Capture credentials from HTTP traffic
-   Defend: static ARP + HTTPS enforcement


<a id="org61ffe04"></a>

### TODO [DEF-WK1] Set Up ELK Stack SIEM     :defense:

-   Install Elasticsearch + Logstash + Kibana (Docker)
-   Ship syslog, auth.log, firewall logs via Filebeat
-   Build 3 dashboards: failed logins, port scans, outbound traffic
-   Write 2 detection rules

**\*** Combo: Core of home SOC (DEF-MO1)


<a id="org6c8e5c0"></a>

### TODO [DEF-WK2] Honeypot + Log Pipeline     :defense:

-   Ship Cowrie logs into ELK
-   Dashboard: attacker IPs, commands run, passwords tried
-   Cross-reference IPs with threat intel feeds (AbuseIPDB API)
-   Alert on new attacker commands


<a id="orgf24fced"></a>

### TODO [MAL-WK1] Dynamic Malware Analysis in Sandbox     :malware:

-   Set up FlareVM or REMnux
-   Run safe malware sample in isolated VM
-   Monitor: procmon, Wireshark, regshot
-   Document: file drops, registry changes, network IOCs


<a id="org469d990"></a>

### TODO [MAL-WK2] Reverse Engineering with Ghidra     :malware:re:

-   Install Ghidra
-   Decompile a simple CTF binary — find hardcoded key
-   Decompile crackme — patch jump condition
-   Analyze a real open-source malware (TinyShell)
-   Annotate functions in Ghidra


<a id="orgdd86167"></a>

### TODO [NET-WK3] Network Pivoting Lab     :network:

-   3-VM lab: attacker | pivot | inner target
-   Compromise pivot, use it to reach inner
-   SSH tunneling: local/remote/dynamic port forward
-   Metasploit route + socks proxy

**\*** Combo: Essential for AD month project


<a id="orgd4795c6"></a>

### TODO [CRYPT-WK1] Implement Crypto Attacks in Python     :crypto:scripting:

-   Padding oracle attack (against vulnerable Flask app you write)
-   Length extension attack on SHA1
-   ECB mode block detection (CBC vs ECB oracle)
-   RSA small e attack (cube root)


<a id="org8c74f76"></a>

### TODO [WEB-WK2] SQLi Scanner + SSRF + XXE Lab     :web:scripting:

-   Write Python SQLi error-based scanner
-   SSRF: reach internal metadata endpoint (cloud lab)
-   XXE: read /etc/passwd via XML input
-   Test all three on deliberately vulnerable apps


<a id="orgb433c49"></a>

### TODO [AD-WK1] Active Directory Attack Lab     :ad:

-   AS-REP Roasting (GetNPUsers.py)
-   Kerberoasting (GetUserSPNs.py)
-   Pass-the-Hash with Mimikatz (isolated lab)
-   BloodHound: visualize attack paths

**\*** Combo: Full AD pentest chains into month project


<a id="orga1b2f86"></a>

### TODO [SCRPT-WK2] Build a C2 Beaconing Script (Lab Only)     :scripting:malware:

-   Python agent: beacon home every N seconds
-   Server: receive beacon, send back commands
-   Encode traffic in base64
-   Add jitter to beaconing interval

**\*** NOTE: Lab/VM only. Learn detection via DEF-WK1.


<a id="orgc644dd4"></a>

### TODO [DEF-WK3] Set Up Suricata + Zeek IDS     :defense:network:

-   Install Suricata, load ET Open rules
-   Generate test alerts (nmap scan, exploit traffic)
-   Install Zeek, read conn.log and dns.log
-   Feed both into ELK (from DEF-WK1)


<a id="org50236b4"></a>

### TODO [WEB-WK3] Full Subdomain + Dir Recon Automation     :web:osint:scripting:

-   Chain: subfinder → httpx → ffuf → nuclei
-   Bash/Python pipeline: one command does all
-   Output: live subdomains, interesting endpoints, known CVE hits
-   Run against HackerOne bug bounty target


<a id="org839f371"></a>

### TODO [CTF-WK1] Complete HackTheBox Starting Point (3 Machines)     :network:web:

-   Tier 0–1 Starting Point machines
-   No walkthroughs until truly stuck (30 min rule)
-   Write report-style writeup for each
-   Focus: methodology, not just flags


<a id="org726ec56"></a>

### TODO [OSINT-WK1] OSINT Framework in Python     :osint:scripting:

-   Inputs: email, username, domain, IP
-   Lookups: WHOIS, DNS, breach check, social, Shodan
-   Output: markdown profile report
-   Add screenshot of profiles (selenium)


<a id="org7375da2"></a>

### TODO [NET-WK4] Metasploit Deep Dive     :network:

-   Exploit VulnHub machine fully through Metasploit
-   Post-exploitation: hashdump, meterpreter, persistence
-   Pivoting with Metasploit route
-   Write custom resource script to automate


<a id="org371718d"></a>

### TODO [CLOUD-WK1] AWS Misconfig & Container Security Lab     :cloud:

-   Deploy intentionally vulnerable app (Damn Vulnerable Cloud App)
-   Find: public S3, overprivileged IAM, exposed metadata
-   Docker escape: privileged container lab
-   Kubernetes: exposed dashboard, RBAC bypass


<a id="orgcb0b788"></a>

### TODO [CRYPT-WK2] Build a PKI from Scratch     :crypto:defense:

-   Create root CA with openssl
-   Issue intermediate CA, end-entity certs
-   Configure Apache/Nginx with custom cert
-   Implement CRL (certificate revocation list)


<a id="org5e669e5"></a>

### TODO [MAL-WK3] YARA Rules — Write & Test     :malware:defense:

-   Learn YARA syntax
-   Write rules for 5 malware families from IOCs
-   Test against malware samples (MalwareBazaar)
-   Integrate YARA scan into Python script


<a id="org445d13f"></a>

### TODO [WEB-WK4] Android App Security Testing     :web:re:

-   Decompile APK: jadx, apktool
-   Static: hardcoded keys, exported activities
-   Dynamic: MobSF, Frida hook
-   Intercept traffic with Burp on Android emulator


<a id="orgc0f3c32"></a>

### TODO [DEF-WK4] Incident Response Lab     :defense:

-   Simulate: attacker compromises web server VM
-   IR process: detection → containment → eradication
-   Collect artifacts: memory dump (volatility), disk image
-   Write incident report


<a id="org16aea04"></a>

### TODO [NET-WK5] WPA2 PMKID Attack + Evil Twin AP     :network:

-   PMKID attack with hcxdumptool (no client needed)
-   Set up evil twin with hostapd-wpe
-   Capture MSCHAPv2 credentials
-   Crack with hashcat mode 5500

**\*** NOTE: Own network lab only.


<a id="orgcca7f54"></a>

### TODO [SCRPT-WK3] Vulnerability Scanner in Python     :scripting:network:web:

-   Port scan → service detect → CVE lookup (NVD API)
-   Web: check SQLi, XSS, open redirect, headers
-   Output: severity-ranked HTML report
-   Diff reports: detect new vulns between scans


<a id="orgd2fa4b8"></a>

### TODO [OSINT-WK2] Threat Intel Aggregator     :osint:defense:scripting:

-   Pull from: AlienVault OTX, AbuseIPDB, VirusTotal API
-   IOC lookup: IP, domain, hash
-   Feed matches into ELK SIEM alerts
-   Daily digest email report (smtplib)


<a id="org04e23c4"></a>

### TODO [RE-WK1] Buffer Overflow 101     :re:

-   Compile vulnerable C program (strcpy, no canary)
-   Find offset with pattern<sub>create</sub> (Metasploit)
-   Control EIP, redirect to shellcode
-   Bypass NX with ret2libc


<a id="org6566c69"></a>

### TODO [WEB-WK5] Build a Vulnerable Web App (for CTF)     :web:scripting:

-   Flask app with intentional vulns: SQLi, XSS, IDOR, path traversal
-   Write challenge descriptions + flags
-   Host for friends or local CTF

**\*** Combo: CTF hosting = teaches both attack & defense mindset


<a id="orga31003b"></a>

### TODO [DEF-WK5] Zero Trust Network Lab     :defense:network:

-   Segment home lab into trust zones
-   WireGuard + firewall rules enforce zone boundaries
-   Service identity via mTLS (mutual TLS)
-   Verify: no lateral movement possible between zones


<a id="orgc5bd272"></a>

### TODO [NET-WK6] Analyze a Real-World CVE + Write PoC     :network:scripting:

-   Pick recent CVE (Log4Shell, ProxyLogon class)
-   Read: NVD, GitHub advisory, patch diff
-   Set up vulnerable version in Docker
-   Write Python PoC or adapt existing one
-   Document: vuln class, impact, patch


<a id="org32bbf4f"></a>

### TODO [CLOUD-WK2] Serverless + API Security Testing     :cloud:web:

-   Deploy Lambda function with IDOR vuln
-   Test: broken auth, over-privileged role, unvalidated input
-   API Gateway: enumerate endpoints, find undocumented
-   Use Postman + manual testing


<a id="org3b4ccbb"></a>

## ══════════════════════════════════════════


<a id="org923e413"></a>

# TIER 3 — MONTH PROJECTS (3–4 weeks each)


<a id="orgfa0bb31"></a>

## ══════════════════════════════════════════


<a id="org24efb69"></a>

### TODO [DEF-MO1] Build a Full Home SOC     :defense:network:

-   ELK Stack SIEM with real dashboards
-   Suricata + Zeek feeding into ELK
-   Cowrie honeypot logging live attacks
-   Wazuh or OSSEC host-based IDS on all VMs
-   PagerDuty/email alerts on critical events
-   Weekly threat digest auto-report

**\*** Showcase: This alone is a real portfolio piece


<a id="org2b8348b"></a>

### TODO [NET-MO1] OSCP-Style Multi-Machine Lab + Report     :network:web:re:

-   Set up 5+ VulnHub/HackTheBox machines in lab
-   Full pentest each: recon → exploit → privesc → persist
-   Write a professional pentest report (executive summary + technical)
-   Include: scope, findings, risk ratings, remediation
-   Simulate: time-boxed (72h per machine)

**\*** Showcase: Submit to eJPT or use as OSCP prep


<a id="orgab96889"></a>

### TODO [WEB-MO1] Full Web App Pentest Automation Suite     :web:scripting:

-   Chain: subfinder → httpx → nuclei → custom SQLi/XSS scanner
-   Auto-screenshot interesting pages
-   Deduplicate + triage findings by severity
-   HTML report with evidence screenshots
-   Submit findings to HackerOne bug bounty

**\*** Showcase: Use on real bug bounty targets (HackerOne/Bugcrowd)


<a id="org9751e45"></a>

### TODO [MAL-MO1] Full Malware Analysis Report on Real Sample     :malware:re:

-   Pick a notable open malware sample (emotet, njRAT)
-   Full static: PE analysis, string extraction, Ghidra decompile
-   Full dynamic: FlareVM, behavioral analysis
-   Extract all IOCs: IPs, domains, hashes, registry keys, mutexes
-   Write professional malware analysis report (15+ pages)
-   Publish on GitHub + LinkedIn

**\*** Showcase: Top 1% of junior candidates have this


<a id="orgdfdca13"></a>

### TODO [AD-MO1] Active Directory Full Attack + Defense Lab     :ad:network:defense:

-   Red: full AD attack chain (recon → foothold → lateral → DA)
-   AS-REP, Kerberoasting, DCSync, Golden Ticket
-   Blue: deploy Microsoft Defender for Identity, Sentinel
-   Detection rules for each attack technique (SIEM alerts)
-   Harden: tiering model, LAPS, privileged access workstations

**\*** Showcase: Directly maps to enterprise pentest + SOC roles


<a id="org34f8015"></a>

### TODO [OSINT-MO1] Automated OSINT Platform     :osint:scripting:

-   Web UI (Flask/FastAPI + React) for OSINT investigations
-   Modules: person, domain, IP, company
-   Data sources: Shodan, HaveIBeenPwned, WHOIS, crt.sh, LinkedIn
-   Store results in SQLite, export PDF reports
-   Graph visualization of relationships (networkx + d3.js)

**\*** Showcase: Open-source on GitHub — recruiter magnet


<a id="org5e22e86"></a>

### TODO [CRYPT-MO1] Cryptography Attack Suite + PKI System     :crypto:scripting:

-   Full PKI: root CA → intermediate → end-entity certs
-   Attack demonstrations: padding oracle, length extension, timing attack
-   Implement Diffie-Hellman, RSA, ECDSA from scratch in Python
-   Blog post explaining each attack with diagrams

**\*** Showcase: Shows you understand crypto deeply, not just tools


<a id="org6f956be"></a>

### TODO [CLOUD-MO1] Cloud Security Audit + Hardening     :cloud:defense:

-   Full AWS audit: IAM, S3, EC2, Lambda, RDS
-   Find and document all misconfigurations (ScoutSuite report)
-   Remediate each finding + document steps
-   Implement: CloudTrail, GuardDuty, Config Rules, SCPs
-   Terraform IaC for hardened baseline deployment

**\*** Showcase: AWS/GCP security skills are very hireable


<a id="org3b26d9a"></a>

### TODO [SCRPT-MO1] Full Pentest Automation Suite (CLI Tool)     :scripting:network:web:

-   Modules: recon, web scan, vuln check, exploit assist, report gen
-   Plugin architecture — easy to extend
-   Config file support, rate limiting, scope enforcement
-   Full documentation, README, example output
-   Publish on GitHub, write a blog/Medium post

**\*** Showcase: If this gets GitHub stars, it opens doors


<a id="org899b684"></a>

### TODO [CTF-MO1] Host a Public CTF Competition     :web:network:crypto:re:

-   Design 15–20 challenges across categories
-   Categories: web, crypto, forensics, RE, pwn, OSINT
-   Deploy CTFd platform (free)
-   Announce on Reddit/Discord, run for 48h
-   Write post-mortems + solution writeups after

**\*** Showcase: Organizing = leadership. Recruiting loves this.


<a id="org8be5a88"></a>

## ══════════════════════════════════════════


<a id="org2be59f2"></a>

# CHAIN COMBO MAPS — SUGGESTED PATHS


<a id="orgfe04580"></a>

## ══════════════════════════════════════════


<a id="orge2ca6fd"></a>

### PATH A: Network Pentester (Offensive)

Weekend → Week → Month NET-W1 → NET-W2 → NET-W5 → RE-W1 → NET-WK1 → NET-WK3 → NET-WK4 → AD-WK1 → NET-MO1 (OSCP-style lab)


<a id="org0f68b3a"></a>

### PATH B: Web/Bug Bounty Hunter

Weekend → Week → Month WEB-W1 → WEB-W2 → WEB-W3 → WEB-W4 → WEB-W5 → WEB-WK1 → WEB-WK2 → WEB-WK3 → WEB-MO1 (pentest suite + bounty submission)


<a id="org5ffdcfd"></a>

### PATH C: Blue Team / SOC Analyst

Weekend → Week → Month DEF-W1 → DEF-W2 → DEF-W4 → MAL-W1 → DEF-WK1 → DEF-WK2 → DEF-WK3 → DEF-WK4 → DEF-MO1 (home SOC) → AD-MO1 (AD defense)


<a id="orgacc0a8f"></a>

### PATH D: Malware Analyst / Threat Intel

Weekend → Week → Month MAL-W1 → MAL-W2 → RE-W2 → MAL-WK1 → MAL-WK2 → MAL-WK3 → RE-WK1 → MAL-MO1 (full malware analysis report)


<a id="org1931468"></a>

### PATH E: Security Engineer / Tool Builder

Weekend → Week → Month SCRPT-W1 → SCRPT-W2 → SCRPT-W3 → SCRPT-WK1 → SCRPT-WK2 → SCRPT-WK3 → SCRPT-MO1 (pentest suite)


<a id="org36f3caf"></a>

## ══════════════════════════════


<a id="orgc95ab10"></a>

# RESOURCES


<a id="org8c3f7e1"></a>

## ══════════════════════════════


<a id="org0c6995d"></a>

### Platforms And Websites

-   TryHackMe: <https://tryhackme.com> (beginner friendly)
-   HackTheBox: <https://hackthebox.com> (intermediate)
-   VulnHub: <https://vulnhub.com> (free VMs)
-   picoCTF: <https://picoctf.org> (CTF practice)
-   HackerOne: <https://hackerone.com> (bug bounty)
-   MalwareBazaar: <https://bazaar.abuse.ch>
-   crackmes.one: <https://crackmes.one>


<a id="orgf3d9be7"></a>

### Essential Tools

-   Kali Linux / ParrotOS
-   Burp Suite Community
-   Wireshark + tshark
-   Metasploit Framework
-   Ghidra (NSA, free)
-   BloodHound + SharpHound
-   Nuclei + ProjectDiscovery tools
-   Python 3 or Go (your best weapons)


<a id="org9359209"></a>

### Free Certifications to Stack

-   Google Cybersecurity Certificate (Coursera)
-   TryHackMe Jr Penetration Tester path
-   eJPT (INE, $200 — worth it after month projects)
-   CompTIA Security+ (entry-level industry standard)


<a id="org13d584a"></a>

### Time Budget Suggestion

```
   Month 1-2:  All Weekend projects (pick 2/weekend)
   Month 3-4:  Week projects (1 per week)
   Month 5-8:  Month projects (1 per month)
   Month 9+:   Bug bounty, certs, job apps
```
