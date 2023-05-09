# Siri IPMI WEB

## Install

```bash
git clone https://github.com/XRSec/Siri_IPMI_WEB.git
cd Siri_IPMI_WEB
```

```bash
apt-get install ipmitool wakeonlan

make USER="admin" PASSWORD="passwd" IPADDRESS="192.168.0.1" TOKEN="hhTp5eUSsc7iS5" MAC='CF:A8:59:57:67:A5'
```
## USE

快捷指令 ShortCuts [Apple iCloud Siri IPMI WEB](https://www.icloud.com/shortcuts/8eb138cf2a68451d982ba2c089b5e0fa)

 URL: `http://[IPaddress]/power?type=status&token=hhTp5eUSsc7iS5`

## LIST

- chassis power Commands: status, on, off, cycle, reset, diag, soft
