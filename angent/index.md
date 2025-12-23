port: 7890
socks-port: 7891
allow-lan: false
mode: rule
log-level: info
geodata-mode: true
geo-auto-update: true
geodata-loader: standard
geo-update-interval: 24
geox-url:
  geoip: https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.dat
  geosite: >-
    https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geosite.dat
  mmdb: >-
    https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/country.mmdb
  asn: >-
    https://github.com/xishang0128/geoip/releases/download/latest/GeoLite2-ASN.mmdb
rule-providers:
  category-ai-!cn:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/category-ai-!cn.mrs
    path: ./ruleset/category-ai-!cn.mrs
    interval: 86400
  youtube:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/youtube.mrs
    path: ./ruleset/youtube.mrs
    interval: 86400
  google:
    type: http
    format: mrs
    behavior: ipcidr
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geoip/google.mrs
    path: ./ruleset/google.mrs
    interval: 86400
  geolocation-cn:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/geolocation-cn.mrs
    path: ./ruleset/geolocation-cn.mrs
    interval: 86400
  cn:
    type: http
    format: mrs
    behavior: ipcidr
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geoip/cn.mrs
    path: ./ruleset/cn.mrs
    interval: 86400
  github:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/github.mrs
    path: ./ruleset/github.mrs
    interval: 86400
  gitlab:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/gitlab.mrs
    path: ./ruleset/gitlab.mrs
    interval: 86400
  geolocation-!cn:
    type: http
    format: mrs
    behavior: domain
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geosite/geolocation-!cn.mrs
    path: ./ruleset/geolocation-!cn.mrs
    interval: 86400
  private:
    type: http
    format: mrs
    behavior: ipcidr
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geoip/private.mrs
    path: ./ruleset/private.mrs
    interval: 86400
  telegram:
    type: http
    format: mrs
    behavior: ipcidr
    url: >-
      https://gh-proxy.com/https://github.com/MetaCubeX/meta-rules-dat/raw/refs/heads/meta/geo/geoip/telegram.mrs
    path: ./ruleset/telegram.mrs
    interval: 86400
dns:
  enable: true
  ipv6: true
  respect-rules: true
  enhanced-mode: fake-ip
  nameserver:
    - https://120.53.53.53/dns-query
    - https://223.5.5.5/dns-query
  proxy-server-nameserver:
    - https://120.53.53.53/dns-query
    - https://223.5.5.5/dns-query
  nameserver-policy:
    geosite:cn,private:
      - https://120.53.53.53/dns-query
      - https://223.5.5.5/dns-query
    geosite:geolocation-!cn:
      - https://dns.cloudflare.com/dns-query
      - https://dns.google/dns-query
proxies:
  - name: rural
    type: hysteria2
    server: 198.23.236.149
    port: 8811
    password: xy.78904453
    sni: bing.com
    skip-cert-verify: true
proxy-groups:
  - type: select
    name: 🚀 节点选择
    proxies:
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - name: ⚡ 自动选择
    type: url-test
    proxies:
      - rural
    url: https://www.gstatic.com/generate_204
    interval: 300
    lazy: false
  - type: select
    name: 💬 AI 服务
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 📹 油管视频
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🔍 谷歌服务
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🏠 私有网络
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🔒 国内服务
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 📲 电报消息
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🐱 Github
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🌐 非中国
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
  - type: select
    name: 🐟 漏网之鱼
    proxies:
      - 🚀 节点选择
      - DIRECT
      - REJECT
      - ⚡ 自动选择
      - rural
rules:
  - RULE-SET,category-ai-!cn,rural
  - RULE-SET,youtube,rural
  - RULE-SET,google,rural
  - RULE-SET,geolocation-cn,DIRECT
  - RULE-SET,cn,DIRECT
  - RULE-SET,github,rural
  - RULE-SET,gitlab,rural
  - RULE-SET,geolocation-!cn,rural
  - RULE-SET,google,🔍 谷歌服务,no-resolve
  - RULE-SET,private,🏠 私有网络,no-resolve
  - RULE-SET,cn,DIRECT,no-resolve
  - RULE-SET,telegram,📲 电报消息,no-resolve
  - MATCH,🐟 漏网之鱼