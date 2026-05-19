- [CHAINS KEY](#org217c9b8)
  - [══════════════════════════════════════════](#org9c96600)
- [TIER 1 — WEEKEND PROJECTS (1–2 days each)](#orgaa48ca8)
  - [══════════════════════════════════════════](#org4ad0167)
  - [[NET-W1] Set Up Kali Linux VM](#org34c528b)
  - [[NET-W2] Home Network Recon with Nmap](#orge7fe1c7)
  - [[NET-W3] Wireshark Traffic Analysis](#orgda1b2cc)
  - [[NET-W4] Crack a WPA2 Handshake (Own Router)](#org5f49a65)
  - [[NET-W5] Netcat Fundamentals — Shells & Transfers](#org3a8dec9)
  - [[WEB-W1] SQLi & XSS on DVWA](#org4124123)
  - [[WEB-W2] Burp Suite Interception Basics](#orgfda88f4)
  - [[WEB-W3] Command Injection & File Inclusion on DVWA](#orgc7bcf6f)
  - [[WEB-W4] Subdomain Enumeration & Recon](#org721de5e)
  - [[OSINT-W1] Google Dorking & Shodan Recon](#org39d5eb9)
  - [[OSINT-W2] Email & Username OSINT](#org4acaa8e)
  - [[OSINT-W3] DNS Enumeration](#orga1b4704)
  - [[SCRPT-W1] Python or Go Port Scanner](#org2d815a4)
  - [[SCRPT-W2] Caesar Cipher → Basic Crypto in Python/go](#orgb662738)
  - [[CRYPT-W1] Hash Cracking with Hashcat](#org96c9e02)
  - [[CRYPT-W2] Steganography — Hide & Find](#orgb6a14d4)
  - [[DEF-W1] SSH Hardening + Key Auth](#orge92331f)
  - [[DEF-W2] Firewall Rules with UFW/iptables](#org83196c5)
  - [[MAL-W1] Static Malware Analysis](#org65d9c7b)
  - [[MAL-W2] Analyze a Malicious PCAP](#org2901e9c)
  - [[RE-W1] Linux Privilege Escalation Basics](#org1bf2b1d)
  - [[DEF-W3] GPG Encryption — Files & Email](#org7bb4658)
  - [[NET-W6] Set Up WireGuard VPN](#org410028e)
  - [[OSINT-W4] Digital Forensics — File Recovery](#orge27d84a)
  - [[WEB-W5] JWT Attack Lab](#org69c9316)
  - [[CTF-W1] Complete 5 picoCTF Beginner Challenges](#org6ddfc4f)
  - [[SCRPT-W3] Log Parser in Python](#orge082d60)
  - [[NET-W7] Proxy Chains + Tor Setup](#org5bb5936)
  - [[WEB-W6] HTTP Security Headers Audit Tool](#orgf5a445e)
  - [[DEF-W4] Set Up Basic Honeypot (Cowrie)](#org3a517c0)
  - [[RE-W2] Reverse a Simple Crackme Binary](#org418a00d)
  - [[AD-W1] Active Directory Concepts + Lab Setup](#org831a748)
  - [[CLOUD-W1] AWS Free Tier — IAM Misconfig Hunt](#orgb5c33ae)
  - [[SCRPT-W4] Build HTTP Header Fuzzer](#org3bed862)
  - [[NET-W8] TryHackMe — Complete 2 Beginner Rooms](#orged9caa6)
  - [══════════════════════════════════════════](#org783104b)
- [TIER 2 — WEEK PROJECTS (3–7 days each)](#orgc42c3db)
  - [══════════════════════════════════════════](#orgf12346d)
    - [[SCRPT-WK1] Full Go or Python Recon Framework](#org084acf1)
    - [[NET-WK1] Full Pentest: VulnHub Beginner Machine](#org37236f2)
    - [[WEB-WK1] Complete OWASP Top 10 on WebGoat](#org1ed5a6a)
    - [[NET-WK2] Man-in-the-Middle Attack Lab](#org2d49f34)
    - [[DEF-WK1] Set Up ELK Stack SIEM](#org4d8da00)
    - [[DEF-WK2] Honeypot + Log Pipeline](#org12b426d)
    - [[MAL-WK1] Dynamic Malware Analysis in Sandbox](#org5edd089)
    - [[MAL-WK2] Reverse Engineering with Ghidra](#org194cff4)
    - [[NET-WK3] Network Pivoting Lab](#orgcf351d3)
    - [[CRYPT-WK1] Implement Crypto Attacks in Python](#org210bcc3)
    - [[WEB-WK2] SQLi Scanner + SSRF + XXE Lab](#org332ba26)
    - [[AD-WK1] Active Directory Attack Lab](#org5eba004)
    - [[SCRPT-WK2] Build a C2 Beaconing Script (Lab Only)](#orgea48ce5)
    - [[DEF-WK3] Set Up Suricata + Zeek IDS](#org4690052)
    - [[WEB-WK3] Full Subdomain + Dir Recon Automation](#org9e81ccf)
    - [[CTF-WK1] Complete HackTheBox Starting Point (3 Machines)](#org5fc6ae6)
    - [[OSINT-WK1] OSINT Framework in Python](#orgd001b83)
    - [[NET-WK4] Metasploit Deep Dive](#org95e646d)
    - [[CLOUD-WK1] AWS Misconfig & Container Security Lab](#org5d43644)
    - [[CRYPT-WK2] Build a PKI from Scratch](#org13b2a8d)
    - [[MAL-WK3] YARA Rules — Write & Test](#orge956d74)
    - [[WEB-WK4] Android App Security Testing](#org4142b5a)
    - [[DEF-WK4] Incident Response Lab](#org4bca2a7)
    - [[NET-WK5] WPA2 PMKID Attack + Evil Twin AP](#org6e3037e)
    - [[SCRPT-WK3] Vulnerability Scanner in Python](#org95c55ff)
    - [[OSINT-WK2] Threat Intel Aggregator](#org56ba6c7)
    - [[RE-WK1] Buffer Overflow 101](#orgf4d9654)
    - [[WEB-WK5] Build a Vulnerable Web App (for CTF)](#orgc763128)
    - [[DEF-WK5] Zero Trust Network Lab](#org64a3ba0)
    - [[NET-WK6] Analyze a Real-World CVE + Write PoC](#orgb129e4e)
    - [[CLOUD-WK2] Serverless + API Security Testing](#orgba565df)
  - [══════════════════════════════════════════](#org3cd2da4)
- [TIER 3 — MONTH PROJECTS (3–4 weeks each)](#org03d4a97)
  - [══════════════════════════════════════════](#org93bd9d5)
    - [[DEF-MO1] Build a Full Home SOC](#org000df4b)
    - [[NET-MO1] OSCP-Style Multi-Machine Lab + Report](#org81d9d1c)
    - [[WEB-MO1] Full Web App Pentest Automation Suite](#org71e5893)
    - [[MAL-MO1] Full Malware Analysis Report on Real Sample](#org7bbd52d)
    - [[AD-MO1] Active Directory Full Attack + Defense Lab](#org0a2e65c)
    - [[OSINT-MO1] Automated OSINT Platform](#orgc1e4e8f)
    - [[CRYPT-MO1] Cryptography Attack Suite + PKI System](#orgbb96d8d)
    - [[CLOUD-MO1] Cloud Security Audit + Hardening](#org7e42fcc)
    - [[SCRPT-MO1] Full Pentest Automation Suite (CLI Tool)](#org30bd316)
    - [[CTF-MO1] Host a Public CTF Competition](#org15ccd19)
  - [══════════════════════════════════════════](#orgd2e3408)
- [CHAIN COMBO MAPS — SUGGESTED PATHS](#org8760742)
  - [══════════════════════════════════════════](#orgf1fb5b1)
    - [PATH A: Network Pentester (Offensive)](#org5708a97)
    - [PATH B: Web/Bug Bounty Hunter](#org8049081)
    - [PATH C: Blue Team / SOC Analyst](#org10b8603)
    - [PATH D: Malware Analyst / Threat Intel](#orgd8878cb)
    - [PATH E: Security Engineer / Tool Builder](#org69aae1e)
  - [══════════════════════════════](#org0c97b83)
- [RESOURCES](#org361750f)
  - [══════════════════════════════](#orgd6f6a44)
    - [Platforms And Websites](#org5b76cfa)
    - [Essential Tools](#orgd9d2072)
    - [Free Certifications to Stack](#org36f2d88)
    - [Time Budget Suggestion](#org6ffbb79)



<a id="org217c9b8"></a>

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


<a id="org9c96600"></a>

## ══════════════════════════════════════════


<a id="orgaa48ca8"></a>

# TIER 1 — WEEKEND PROJECTS (1–2 days each)


<a id="org4ad0167"></a>

## ══════════════════════════════════════════


<a id="org34c528b"></a>

## TODO [NET-W1] Set Up Kali Linux VM     :network:

-   Install VirtualBox or VMware
-   Download Kali ISO, install, snapshot clean state
-   Learn basic terminal nav, update system
-   Install guest additions

**\*** Combo: Base for every other project. Do first.


<a id="orge7fe1c7"></a>

## TODO [NET-W2] Home Network Recon with Nmap     :network:

-   Scan local subnet: `nmap -sn 192.168.x.0/24`
-   Service scan: `nmap -sV -sC -A <target>`
-   Output to XML, read it
-   Understand open ports on your own devices

**\*** Combo: Feeds into Python scanner (SCRPT-W1) and SIEM setup (DEF-WK1)


<a id="orgda1b2cc"></a>

## TODO [NET-W3] Wireshark Traffic Analysis     :network:

-   Capture live traffic on home net
-   Filter HTTP, DNS, ARP
-   Find plaintext credentials in pcap (test pcap from online)
-   Export objects from HTTP stream

**\*** Combo: Pairs with MitM week project (NET-WK4)


<a id="org5f49a65"></a>

## TODO [NET-W4] Crack a WPA2 Handshake (Own Router)     :network:

-   Use monitor mode + airodump-ng to capture 4-way handshake
-   Deauth a client to force reconnect
-   Crack with hashcat + rockyou.txt
-   Change your WiFi password after. Learn why WPA3 matters.

**\*** NOTE: Own network only. Legal line clear.


<a id="org3a8dec9"></a>

## TODO [NET-W5] Netcat Fundamentals — Shells & Transfers     :network:

-   Open listeners, connect clients
-   Send files with nc
-   Reverse shell: `nc -e /bin/bash`
-   Bind shell vs reverse shell — understand difference

**\*** Combo: Foundation for all post-exploitation work


<a id="org4124123"></a>

## TODO [WEB-W1] SQLi & XSS on DVWA     :web:

-   Install DVWA (Docker or XAMPP)
-   Complete SQLi: manual + sqlmap
-   Complete XSS: reflected, stored, DOM
-   Toggle security levels low→medium→high

**\*** Combo: Directly builds into full OWASP week (WEB-WK1)


<a id="orgfda88f4"></a>

## TODO [WEB-W2] Burp Suite Interception Basics     :web:

-   Set up browser proxy through Burp
-   Intercept, modify, replay requests
-   Use Repeater on DVWA login
-   Use Intruder for basic brute force

**\*** Combo: Essential tool for all web projects


<a id="orgc7bcf6f"></a>

## TODO [WEB-W3] Command Injection & File Inclusion on DVWA     :web:

-   Command injection: OS commands through web input
-   LFI: read `/etc/passwd` via vulnerable param
-   RFI: include remote malicious file
-   CSRF: forge requests, steal sessions


<a id="org721de5e"></a>

## TODO [WEB-W4] Subdomain Enumeration & Recon     :web:osint:

-   Use subfinder, amass on a target (HackerOne public programs)
-   Certificate transparency: crt.sh
-   Directory bruteforce with ffuf: `ffuf -w wordlist -u https://target/FUZZ`
-   Document findings in structured notes


<a id="org39d5eb9"></a>

## TODO [OSINT-W1] Google Dorking & Shodan Recon     :osint:

-   Learn 10 key Google dork operators
-   Find exposed login panels, open dirs, config files
-   Shodan: search for services by banner, CVE
-   Build a personal dork cheatsheet


<a id="org4acaa8e"></a>

## TODO [OSINT-W2] Email & Username OSINT     :osint:

-   theHarvester: email harvest from domain
-   holehe or Sherlock: username across platforms
-   Have I Been Pwned API lookup
-   Build a target profile (use yourself as test subject)


<a id="orga1b4704"></a>

## TODO [OSINT-W3] DNS Enumeration     :osint:network:

-   dnsrecon, dnsenum on practice domains
-   Zone transfer attempt
-   MX, TXT, NS record analysis
-   Reverse DNS lookup sweep


<a id="org2d815a4"></a>

## TODO [SCRPT-W1] Python or Go Port Scanner     :scripting:

-   Socket-based TCP port scanner
-   Add threading for speed
-   Service banner grabbing
-   Output to JSON/CSV

**\*** Combo: Base for full recon tool (SCRPT-WK1)


<a id="orgb662738"></a>

## TODO [SCRPT-W2] Caesar Cipher → Basic Crypto in Python/go     :scripting:crypto:

-   Implement Caesar, Vigenere, XOR cipher
-   Brute-force Caesar without key
-   Frequency analysis for Vigenere
-   Understand why these fail


<a id="org96c9e02"></a>

## TODO [CRYPT-W1] Hash Cracking with Hashcat     :crypto:

-   Identify hash types (hash-identifier, hashid)
-   Crack MD5, SHA1, bcrypt with rockyou.txt
-   Rules-based attack with hashcat rules
-   Dictionary vs brute vs combo attack modes


<a id="orgb6a14d4"></a>

## TODO [CRYPT-W2] Steganography — Hide & Find     :crypto:

-   Hide text in image: steghide, LSB
-   Extract: steghide extract, stegsolve
-   Audio steganography: MP3Stego
-   Solve 3 stego CTF challenges


<a id="orge92331f"></a>

## TODO [DEF-W1] SSH Hardening + Key Auth     :defense:

-   Disable password auth, enable key-only
-   Change default port, restrict users
-   Set up fail2ban for SSH brute protection
-   Test hardening with nmap from Kali


<a id="org83196c5"></a>

## TODO [DEF-W2] Firewall Rules with UFW/iptables     :defense:

-   Default deny inbound policy
-   Allow only necessary ports
-   Log dropped packets
-   Test rules from external VM


<a id="org65d9c7b"></a>

## TODO [MAL-W1] Static Malware Analysis     :malware:

-   strings, file, xxd on a safe malware sample (MalwareBazaar)
-   Extract IPs/domains/registry keys from strings
-   PE header analysis with PEview or pestudio
-   Identify packing/obfuscation signs


<a id="org2901e9c"></a>

## TODO [MAL-W2] Analyze a Malicious PCAP     :malware:network:

-   Download malware traffic pcap (malware-traffic-analysis.net)
-   Identify C2 beaconing patterns
-   Extract indicators of compromise (IOCs)
-   Write a short analysis report


<a id="org1bf2b1d"></a>

## TODO [RE-W1] Linux Privilege Escalation Basics     :re:

-   GTFOBins: SUID binary abuse
-   Writable /etc/passwd, cron abuse
-   sudo -l misconfigs
-   linPEAS on a VulnHub machine


<a id="org7bb4658"></a>

## TODO [DEF-W3] GPG Encryption — Files & Email     :defense:crypto:

-   Generate GPG keypair
-   Encrypt/decrypt files
-   Sign and verify
-   Export/import public keys


<a id="org410028e"></a>

## TODO [NET-W6] Set Up WireGuard VPN     :network:defense:

-   Install WireGuard on a VPS or local VM
-   Generate peer keys, configure tunnels
-   Route traffic through tunnel
-   Verify with Wireshark — confirm encryption


<a id="orge27d84a"></a>

## TODO [OSINT-W4] Digital Forensics — File Recovery     :osint:

-   Create a disk image with dd
-   Recover deleted files with autopsy + foremost
-   Analyze file metadata (exiftool)
-   Build a basic forensics checklist


<a id="org69c9316"></a>

## TODO [WEB-W5] JWT Attack Lab     :web:

-   Decode JWT (jwt.io)
-   none algorithm attack
-   Brute force weak HS256 secret (hashcat)
-   Key confusion attack (RS256→HS256)


<a id="org6ddfc4f"></a>

## TODO [CTF-W1] Complete 5 picoCTF Beginner Challenges     :web:crypto:re:

-   Pick challenges across: crypto, forensics, web, general skills
-   Document solve methodology for each
-   Learn to use CyberChef
-   Join a CTF Discord for hints


<a id="orge082d60"></a>

## TODO [SCRPT-W3] Log Parser in Python     :scripting:defense:

-   Parse /var/log/auth.log for failed logins
-   Count IPs, flag threshold breaches
-   Output alert summary
-   Extend to syslog, apache access logs


<a id="org5bb5936"></a>

## TODO [NET-W7] Proxy Chains + Tor Setup     :network:

-   Install tor + proxychains
-   Route nmap through proxychains
-   Understand Tor limitations for pentesting
-   Test anonymity with whatismyip


<a id="orgf5a445e"></a>

## TODO [WEB-W6] HTTP Security Headers Audit Tool     :web:scripting:

-   Python script: fetch headers from any URL
-   Check: CSP, HSTS, X-Frame-Options, CORS
-   Score and report missing headers
-   Run against 10 real sites (ethically)


<a id="org3a517c0"></a>

## TODO [DEF-W4] Set Up Basic Honeypot (Cowrie)     :defense:

-   Install Cowrie SSH honeypot
-   Expose on a VPS or local VM
-   Watch logs for hit attempts
-   Extract attacker IPs and commands


<a id="org418a00d"></a>

## TODO [RE-W2] Reverse a Simple Crackme Binary     :re:

-   Download crackme from crackmes.one (easy level)
-   Use ltrace/strace first
-   Open in Ghidra — find password check logic
-   Patch binary to bypass check


<a id="org831a748"></a>

## TODO [AD-W1] Active Directory Concepts + Lab Setup     :ad:

-   Install Windows Server eval VM
-   Promote to domain controller
-   Create OUs, users, groups
-   Join a Windows 10 VM to the domain


<a id="orgb5c33ae"></a>

## TODO [CLOUD-W1] AWS Free Tier — IAM Misconfig Hunt     :cloud:

-   Create AWS free tier account
-   Create intentionally misconfigured IAM (for lab)
-   Use ScoutSuite or Prowler to audit
-   Enumerate with AWS CLI using overprivileged user


<a id="org3bed862"></a>

## TODO [SCRPT-W4] Build HTTP Header Fuzzer     :scripting:web:

-   Python requests — iterate custom headers
-   Fuzz Host, X-Forwarded-For, Content-Type
-   Look for 500 errors or behavioral changes
-   Test on DVWA or local lab app


<a id="orged9caa6"></a>

## TODO [NET-W8] TryHackMe — Complete 2 Beginner Rooms     :network:

-   Recommended: &ldquo;Basic Pentesting&rdquo;, &ldquo;Startup&rdquo;
-   Document methodology: recon → exploit → flags
-   Note tools used and commands
-   Subscribe to free tier


<a id="org783104b"></a>

## ══════════════════════════════════════════


<a id="orgc42c3db"></a>

# TIER 2 — WEEK PROJECTS (3–7 days each)


<a id="orgf12346d"></a>

## ══════════════════════════════════════════


<a id="org084acf1"></a>

### TODO [SCRPT-WK1] Full Go or Python Recon Framework     :scripting:network:osint:

-   Combine: port scanner + subdomain enum + DNS recon + header check
-   Single CLI tool with argparse
-   Output to JSON report + markdown summary
-   Add screenshot capability (selenium headless)

**\*** Combo: Ports directly into full pentest suite (SCRPT-MO1)


<a id="org37236f2"></a>

### TODO [NET-WK1] Full Pentest: VulnHub Beginner Machine     :network:web:re:

-   Download: Mr-Robot, Kioptrix, or Basic Pentesting 1
-   Recon → foothold → privesc → root
-   Document every step in markdown
-   Write a mini pentest report

**\*** Combo: Chain 3+ machines → OSCP prep (NET-MO1)


<a id="org1ed5a6a"></a>

### TODO [WEB-WK1] Complete OWASP Top 10 on WebGoat     :web:

-   Install WebGoat (Java or Docker)
-   Complete all OWASP Top 10 lessons
-   A01 Broken Access Control through A10 SSRF
-   Write one-pager summary per vuln

**\*** Combo: Unlocks web pentest automation month project


<a id="org2d49f34"></a>

### TODO [NET-WK2] Man-in-the-Middle Attack Lab     :network:

-   ARP spoofing: arpspoof + Wireshark in isolated VM lab
-   SSL stripping with bettercap
-   Capture credentials from HTTP traffic
-   Defend: static ARP + HTTPS enforcement


<a id="org4d8da00"></a>

### TODO [DEF-WK1] Set Up ELK Stack SIEM     :defense:

-   Install Elasticsearch + Logstash + Kibana (Docker)
-   Ship syslog, auth.log, firewall logs via Filebeat
-   Build 3 dashboards: failed logins, port scans, outbound traffic
-   Write 2 detection rules

**\*** Combo: Core of home SOC (DEF-MO1)


<a id="org12b426d"></a>

### TODO [DEF-WK2] Honeypot + Log Pipeline     :defense:

-   Ship Cowrie logs into ELK
-   Dashboard: attacker IPs, commands run, passwords tried
-   Cross-reference IPs with threat intel feeds (AbuseIPDB API)
-   Alert on new attacker commands


<a id="org5edd089"></a>

### TODO [MAL-WK1] Dynamic Malware Analysis in Sandbox     :malware:

-   Set up FlareVM or REMnux
-   Run safe malware sample in isolated VM
-   Monitor: procmon, Wireshark, regshot
-   Document: file drops, registry changes, network IOCs


<a id="org194cff4"></a>

### TODO [MAL-WK2] Reverse Engineering with Ghidra     :malware:re:

-   Install Ghidra
-   Decompile a simple CTF binary — find hardcoded key
-   Decompile crackme — patch jump condition
-   Analyze a real open-source malware (TinyShell)
-   Annotate functions in Ghidra


<a id="orgcf351d3"></a>

### TODO [NET-WK3] Network Pivoting Lab     :network:

-   3-VM lab: attacker | pivot | inner target
-   Compromise pivot, use it to reach inner
-   SSH tunneling: local/remote/dynamic port forward
-   Metasploit route + socks proxy

**\*** Combo: Essential for AD month project


<a id="org210bcc3"></a>

### TODO [CRYPT-WK1] Implement Crypto Attacks in Python     :crypto:scripting:

-   Padding oracle attack (against vulnerable Flask app you write)
-   Length extension attack on SHA1
-   ECB mode block detection (CBC vs ECB oracle)
-   RSA small e attack (cube root)


<a id="org332ba26"></a>

### TODO [WEB-WK2] SQLi Scanner + SSRF + XXE Lab     :web:scripting:

-   Write Python SQLi error-based scanner
-   SSRF: reach internal metadata endpoint (cloud lab)
-   XXE: read /etc/passwd via XML input
-   Test all three on deliberately vulnerable apps


<a id="org5eba004"></a>

### TODO [AD-WK1] Active Directory Attack Lab     :ad:

-   AS-REP Roasting (GetNPUsers.py)
-   Kerberoasting (GetUserSPNs.py)
-   Pass-the-Hash with Mimikatz (isolated lab)
-   BloodHound: visualize attack paths

**\*** Combo: Full AD pentest chains into month project


<a id="orgea48ce5"></a>

### TODO [SCRPT-WK2] Build a C2 Beaconing Script (Lab Only)     :scripting:malware:

-   Python agent: beacon home every N seconds
-   Server: receive beacon, send back commands
-   Encode traffic in base64
-   Add jitter to beaconing interval

**\*** NOTE: Lab/VM only. Learn detection via DEF-WK1.


<a id="org4690052"></a>

### TODO [DEF-WK3] Set Up Suricata + Zeek IDS     :defense:network:

-   Install Suricata, load ET Open rules
-   Generate test alerts (nmap scan, exploit traffic)
-   Install Zeek, read conn.log and dns.log
-   Feed both into ELK (from DEF-WK1)


<a id="org9e81ccf"></a>

### TODO [WEB-WK3] Full Subdomain + Dir Recon Automation     :web:osint:scripting:

-   Chain: subfinder → httpx → ffuf → nuclei
-   Bash/Python pipeline: one command does all
-   Output: live subdomains, interesting endpoints, known CVE hits
-   Run against HackerOne bug bounty target


<a id="org5fc6ae6"></a>

### TODO [CTF-WK1] Complete HackTheBox Starting Point (3 Machines)     :network:web:

-   Tier 0–1 Starting Point machines
-   No walkthroughs until truly stuck (30 min rule)
-   Write report-style writeup for each
-   Focus: methodology, not just flags


<a id="orgd001b83"></a>

### TODO [OSINT-WK1] OSINT Framework in Python     :osint:scripting:

-   Inputs: email, username, domain, IP
-   Lookups: WHOIS, DNS, breach check, social, Shodan
-   Output: markdown profile report
-   Add screenshot of profiles (selenium)


<a id="org95e646d"></a>

### TODO [NET-WK4] Metasploit Deep Dive     :network:

-   Exploit VulnHub machine fully through Metasploit
-   Post-exploitation: hashdump, meterpreter, persistence
-   Pivoting with Metasploit route
-   Write custom resource script to automate


<a id="org5d43644"></a>

### TODO [CLOUD-WK1] AWS Misconfig & Container Security Lab     :cloud:

-   Deploy intentionally vulnerable app (Damn Vulnerable Cloud App)
-   Find: public S3, overprivileged IAM, exposed metadata
-   Docker escape: privileged container lab
-   Kubernetes: exposed dashboard, RBAC bypass


<a id="org13b2a8d"></a>

### TODO [CRYPT-WK2] Build a PKI from Scratch     :crypto:defense:

-   Create root CA with openssl
-   Issue intermediate CA, end-entity certs
-   Configure Apache/Nginx with custom cert
-   Implement CRL (certificate revocation list)


<a id="orge956d74"></a>

### TODO [MAL-WK3] YARA Rules — Write & Test     :malware:defense:

-   Learn YARA syntax
-   Write rules for 5 malware families from IOCs
-   Test against malware samples (MalwareBazaar)
-   Integrate YARA scan into Python script


<a id="org4142b5a"></a>

### TODO [WEB-WK4] Android App Security Testing     :web:re:

-   Decompile APK: jadx, apktool
-   Static: hardcoded keys, exported activities
-   Dynamic: MobSF, Frida hook
-   Intercept traffic with Burp on Android emulator


<a id="org4bca2a7"></a>

### TODO [DEF-WK4] Incident Response Lab     :defense:

-   Simulate: attacker compromises web server VM
-   IR process: detection → containment → eradication
-   Collect artifacts: memory dump (volatility), disk image
-   Write incident report


<a id="org6e3037e"></a>

### TODO [NET-WK5] WPA2 PMKID Attack + Evil Twin AP     :network:

-   PMKID attack with hcxdumptool (no client needed)
-   Set up evil twin with hostapd-wpe
-   Capture MSCHAPv2 credentials
-   Crack with hashcat mode 5500

**\*** NOTE: Own network lab only.


<a id="org95c55ff"></a>

### TODO [SCRPT-WK3] Vulnerability Scanner in Python     :scripting:network:web:

-   Port scan → service detect → CVE lookup (NVD API)
-   Web: check SQLi, XSS, open redirect, headers
-   Output: severity-ranked HTML report
-   Diff reports: detect new vulns between scans


<a id="org56ba6c7"></a>

### TODO [OSINT-WK2] Threat Intel Aggregator     :osint:defense:scripting:

-   Pull from: AlienVault OTX, AbuseIPDB, VirusTotal API
-   IOC lookup: IP, domain, hash
-   Feed matches into ELK SIEM alerts
-   Daily digest email report (smtplib)


<a id="orgf4d9654"></a>

### TODO [RE-WK1] Buffer Overflow 101     :re:

-   Compile vulnerable C program (strcpy, no canary)
-   Find offset with pattern<sub>create</sub> (Metasploit)
-   Control EIP, redirect to shellcode
-   Bypass NX with ret2libc


<a id="orgc763128"></a>

### TODO [WEB-WK5] Build a Vulnerable Web App (for CTF)     :web:scripting:

-   Flask app with intentional vulns: SQLi, XSS, IDOR, path traversal
-   Write challenge descriptions + flags
-   Host for friends or local CTF

**\*** Combo: CTF hosting = teaches both attack & defense mindset


<a id="org64a3ba0"></a>

### TODO [DEF-WK5] Zero Trust Network Lab     :defense:network:

-   Segment home lab into trust zones
-   WireGuard + firewall rules enforce zone boundaries
-   Service identity via mTLS (mutual TLS)
-   Verify: no lateral movement possible between zones


<a id="orgb129e4e"></a>

### TODO [NET-WK6] Analyze a Real-World CVE + Write PoC     :network:scripting:

-   Pick recent CVE (Log4Shell, ProxyLogon class)
-   Read: NVD, GitHub advisory, patch diff
-   Set up vulnerable version in Docker
-   Write Python PoC or adapt existing one
-   Document: vuln class, impact, patch


<a id="orgba565df"></a>

### TODO [CLOUD-WK2] Serverless + API Security Testing     :cloud:web:

-   Deploy Lambda function with IDOR vuln
-   Test: broken auth, over-privileged role, unvalidated input
-   API Gateway: enumerate endpoints, find undocumented
-   Use Postman + manual testing


<a id="org3cd2da4"></a>

## ══════════════════════════════════════════


<a id="org03d4a97"></a>

# TIER 3 — MONTH PROJECTS (3–4 weeks each)


<a id="org93bd9d5"></a>

## ══════════════════════════════════════════


<a id="org000df4b"></a>

### TODO [DEF-MO1] Build a Full Home SOC     :defense:network:

-   ELK Stack SIEM with real dashboards
-   Suricata + Zeek feeding into ELK
-   Cowrie honeypot logging live attacks
-   Wazuh or OSSEC host-based IDS on all VMs
-   PagerDuty/email alerts on critical events
-   Weekly threat digest auto-report

**\*** Showcase: This alone is a real portfolio piece


<a id="org81d9d1c"></a>

### TODO [NET-MO1] OSCP-Style Multi-Machine Lab + Report     :network:web:re:

-   Set up 5+ VulnHub/HackTheBox machines in lab
-   Full pentest each: recon → exploit → privesc → persist
-   Write a professional pentest report (executive summary + technical)
-   Include: scope, findings, risk ratings, remediation
-   Simulate: time-boxed (72h per machine)

**\*** Showcase: Submit to eJPT or use as OSCP prep


<a id="org71e5893"></a>

### TODO [WEB-MO1] Full Web App Pentest Automation Suite     :web:scripting:

-   Chain: subfinder → httpx → nuclei → custom SQLi/XSS scanner
-   Auto-screenshot interesting pages
-   Deduplicate + triage findings by severity
-   HTML report with evidence screenshots
-   Submit findings to HackerOne bug bounty

**\*** Showcase: Use on real bug bounty targets (HackerOne/Bugcrowd)


<a id="org7bbd52d"></a>

### TODO [MAL-MO1] Full Malware Analysis Report on Real Sample     :malware:re:

-   Pick a notable open malware sample (emotet, njRAT)
-   Full static: PE analysis, string extraction, Ghidra decompile
-   Full dynamic: FlareVM, behavioral analysis
-   Extract all IOCs: IPs, domains, hashes, registry keys, mutexes
-   Write professional malware analysis report (15+ pages)
-   Publish on GitHub + LinkedIn

**\*** Showcase: Top 1% of junior candidates have this


<a id="org0a2e65c"></a>

### TODO [AD-MO1] Active Directory Full Attack + Defense Lab     :ad:network:defense:

-   Red: full AD attack chain (recon → foothold → lateral → DA)
-   AS-REP, Kerberoasting, DCSync, Golden Ticket
-   Blue: deploy Microsoft Defender for Identity, Sentinel
-   Detection rules for each attack technique (SIEM alerts)
-   Harden: tiering model, LAPS, privileged access workstations

**\*** Showcase: Directly maps to enterprise pentest + SOC roles


<a id="orgc1e4e8f"></a>

### TODO [OSINT-MO1] Automated OSINT Platform     :osint:scripting:

-   Web UI (Flask/FastAPI + React) for OSINT investigations
-   Modules: person, domain, IP, company
-   Data sources: Shodan, HaveIBeenPwned, WHOIS, crt.sh, LinkedIn
-   Store results in SQLite, export PDF reports
-   Graph visualization of relationships (networkx + d3.js)

**\*** Showcase: Open-source on GitHub — recruiter magnet


<a id="orgbb96d8d"></a>

### TODO [CRYPT-MO1] Cryptography Attack Suite + PKI System     :crypto:scripting:

-   Full PKI: root CA → intermediate → end-entity certs
-   Attack demonstrations: padding oracle, length extension, timing attack
-   Implement Diffie-Hellman, RSA, ECDSA from scratch in Python
-   Blog post explaining each attack with diagrams

**\*** Showcase: Shows you understand crypto deeply, not just tools


<a id="org7e42fcc"></a>

### TODO [CLOUD-MO1] Cloud Security Audit + Hardening     :cloud:defense:

-   Full AWS audit: IAM, S3, EC2, Lambda, RDS
-   Find and document all misconfigurations (ScoutSuite report)
-   Remediate each finding + document steps
-   Implement: CloudTrail, GuardDuty, Config Rules, SCPs
-   Terraform IaC for hardened baseline deployment

**\*** Showcase: AWS/GCP security skills are very hireable


<a id="org30bd316"></a>

### TODO [SCRPT-MO1] Full Pentest Automation Suite (CLI Tool)     :scripting:network:web:

-   Modules: recon, web scan, vuln check, exploit assist, report gen
-   Plugin architecture — easy to extend
-   Config file support, rate limiting, scope enforcement
-   Full documentation, README, example output
-   Publish on GitHub, write a blog/Medium post

**\*** Showcase: If this gets GitHub stars, it opens doors


<a id="org15ccd19"></a>

### TODO [CTF-MO1] Host a Public CTF Competition     :web:network:crypto:re:

-   Design 15–20 challenges across categories
-   Categories: web, crypto, forensics, RE, pwn, OSINT
-   Deploy CTFd platform (free)
-   Announce on Reddit/Discord, run for 48h
-   Write post-mortems + solution writeups after

**\*** Showcase: Organizing = leadership. Recruiting loves this.


<a id="orgd2e3408"></a>

## ══════════════════════════════════════════


<a id="org8760742"></a>

# CHAIN COMBO MAPS — SUGGESTED PATHS


<a id="orgf1fb5b1"></a>

## ══════════════════════════════════════════


<a id="org5708a97"></a>

### PATH A: Network Pentester (Offensive)

Weekend → Week → Month NET-W1 → NET-W2 → NET-W5 → RE-W1 → NET-WK1 → NET-WK3 → NET-WK4 → AD-WK1 → NET-MO1 (OSCP-style lab)


<a id="org8049081"></a>

### PATH B: Web/Bug Bounty Hunter

Weekend → Week → Month WEB-W1 → WEB-W2 → WEB-W3 → WEB-W4 → WEB-W5 → WEB-WK1 → WEB-WK2 → WEB-WK3 → WEB-MO1 (pentest suite + bounty submission)


<a id="org10b8603"></a>

### PATH C: Blue Team / SOC Analyst

Weekend → Week → Month DEF-W1 → DEF-W2 → DEF-W4 → MAL-W1 → DEF-WK1 → DEF-WK2 → DEF-WK3 → DEF-WK4 → DEF-MO1 (home SOC) → AD-MO1 (AD defense)


<a id="orgd8878cb"></a>

### PATH D: Malware Analyst / Threat Intel

Weekend → Week → Month MAL-W1 → MAL-W2 → RE-W2 → MAL-WK1 → MAL-WK2 → MAL-WK3 → RE-WK1 → MAL-MO1 (full malware analysis report)


<a id="org69aae1e"></a>

### PATH E: Security Engineer / Tool Builder

Weekend → Week → Month SCRPT-W1 → SCRPT-W2 → SCRPT-W3 → SCRPT-WK1 → SCRPT-WK2 → SCRPT-WK3 → SCRPT-MO1 (pentest suite)


<a id="org0c97b83"></a>

## ══════════════════════════════


<a id="org361750f"></a>

# RESOURCES


<a id="orgd6f6a44"></a>

## ══════════════════════════════


<a id="org5b76cfa"></a>

### Platforms And Websites

-   TryHackMe: <https://tryhackme.com> (beginner friendly)
-   HackTheBox: <https://hackthebox.com> (intermediate)
-   VulnHub: <https://vulnhub.com> (free VMs)
-   picoCTF: <https://picoctf.org> (CTF practice)
-   HackerOne: <https://hackerone.com> (bug bounty)
-   MalwareBazaar: <https://bazaar.abuse.ch>
-   crackmes.one: <https://crackmes.one>


<a id="orgd9d2072"></a>

### Essential Tools

-   Kali Linux / ParrotOS
-   Burp Suite Community
-   Wireshark + tshark
-   Metasploit Framework
-   Ghidra (NSA, free)
-   BloodHound + SharpHound
-   Nuclei + ProjectDiscovery tools
-   Python 3 or Go (your best weapons)


<a id="org36f2d88"></a>

### Free Certifications to Stack

-   Google Cybersecurity Certificate (Coursera)
-   TryHackMe Jr Penetration Tester path
-   eJPT (INE, $200 — worth it after month projects)
-   CompTIA Security+ (entry-level industry standard)


<a id="org6ffbb79"></a>

### Time Budget Suggestion

```
   Month 1-2:  All Weekend projects (pick 2/weekend)
   Month 3-4:  Week projects (1 per week)
   Month 5-8:  Month projects (1 per month)
   Month 9+:   Bug bounty, certs, job apps
```
