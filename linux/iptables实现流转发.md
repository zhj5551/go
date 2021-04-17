# iptables实现流转发
```

iptables -t nat -A PREROUTING -p tcp -m tcp --dport 1935 -j DNAT --to-destination 119.3.55.161
iptables -t nat -A POSTROUTING -p tcp -m tcp --dport 1935 -j SNAT --to-source 192.168.1.82
iptables -t filter -A INPUT -p tcp -m state --state NEW -m tcp --dport 1935 -j ACCEPT

119.3.55.161：直播推流地址

192.168.1.82：本机内网IP
```

