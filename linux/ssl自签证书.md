## Let's Encrypt 泛域名证书申请及配置
- [Let's Encrypt 泛域名证书申请及配置](https://segmentfault.com/a/1190000015354547)

- 制作阿里的域名证书: 
```
.acme.sh/acme.sh --issue --dns dns_ali -d godruoyi.com -d *.godruoyi.com
```