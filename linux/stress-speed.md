## vegeta压测
- 安装： 
```
wget -c https://github.com/tsenart/vegeta/releases/download/v12.8.3/vegeta-12.8.3-linux-amd64.tar.gz
tar -xzf vegeta-12.8.3-linux-amd64.tar.gz -C /usr/local/bin/
vegeta --version

```
- [vegeta的github教程](https://github.com/tsenart/vegeta)
- http压测命令
```
echo "GET http://192.168.179.34:5000"| vegeta attack -rate=500000 -duration=10s | tee results.bin | vegeta report
```
### 常见压测返回的报错
- Get "http://192.168.179.12/index.html": dial tcp 0.0.0.0:0->192.168.179.12:80: bind: address already in use
- Get "http://192.168.179.12/index.html": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
## ab压测
- 安装：
```
yum install -y httpd-tools
```
- 压测命令
```
ab -n 100000 -c 5000 http://aaaaa.testqibo.mudu.tv/watch/4m9v367v
-n 请求数
-c 多少个客户端同时发起请求
```
## 查看外网ip命令
- curl cip.cc
- curl ifconfig.me

## speedtest测网速工具

## iPerf查看带宽
- [使用iPerf进行网络吞吐量测试](https://www.jianshu.com/p/15f888309c72)

## iftop网速实时监控
### 安装
- apt install iftop -y
- [网络实时流量监测工具iftop](https://blog.csdn.net/kairui123/article/details/76014660)