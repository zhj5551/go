# 微服务日志规范
1. 日志文件管理
2. 日志收集管理
3. 报警
4. 日志查询

## 日志文件管理
虽然很多日志库支持日志轮询，但如果直接使用会造成日志管理与特定的项目耦合。 [Logrus](https://github.com/sirupsen/logrus)日志库的文档中有指出，建议使用[logrotate](https://linux.die.net/man/8/logrotate)进行日志轮询。

### Logrotate
Logrotate是个十分有用的工具，它可以自动对日志进行截断（或轮循）、压缩以及删除旧的日志文件。例如，你可以设置logrotate，让/var/log/foo日志文件每天轮循，并删除超过7天的日志。配置完后，logrotate的运作完全自动化，不必进行任何进一步的人为干预。另外，旧日志也可以通过电子邮件发送。


### 安装
主流Linux发行版上都默认安装有logrotate包，如果出于某种原因，logrotate没有出现在里头，你可以使用apt-get/yum/apk命令来安装。
```
debian:  yum install logrotate crontabs 

alpine:  apk add logrotate
```
logrotate的cron任务也会在安装时自动创建：
```
debian: cat /etc/cron.daily/logrotate 

alpine: cat /etc/periodic/daily/logrotate
```
### 配置
logrotate的配置文件是/etc/logrotate.conf，通常不需要对它进行修改。日志文件的轮循设置在独立的配置文件中，它（们）放在/etc/logrotate.d/目录下。

如需要对/data/main.log和/data/access.log进行日志轮询，只要在/etc/logrotate.d目录下创建svc_notify配置文件:
```
/data/myun/svc_notify/logs/main.log {
    missingok
    notifempty
    daily
    copytruncate
    dateext
    create 0777 root root
    rotate 7
}

/data/myun/svc_notify/logs/access.log {
    missingok
    notifempty
    daily
    copytruncate
    dateext
    create 0777 root root
    rotate 7
}

这里：

missingok: 在日志轮循期间，任何错误将被忽略，例如“文件无法找到”之类的错误。
notifempty: 如果日志文件为空，轮循不会进行。
daily: 日志文件将按天轮循。其它可用值为‘weekly’，‘monthly’或者‘yearly’。
copytruncate: 用于还在打开中的日志文件，把当前日志备份并截断；是先拷贝再清空的方式，拷贝和清空之间有一个时间差，可能会丢失部分日志数据。因为目睹云微服务只打开一次日志文件，如果不使用该选项，新日志还会写到备份的日志文件中。
dateext: 让旧日志文件以创建日期命名。
rotate 7: 一次将存储7个归档日志。对于第8个归档，时间最久的归档将被删除。
create 0777 root root: 以指定的权限创建全新的日志文件，同时logrotate也会重命名原始日志文件。

此外还有：

size: 文件大小限制。
compress: 在轮循任务完成后，已轮循的归档将使用gzip进行压缩。
delaycompress: 总是与compress选项一起用，delaycompress选项指示logrotate不要将最近的归档压缩，压缩将在下一次轮循周期进行。这在你或任何软件仍然需要读取最新归档时很有用。
postrotate/endscript: 在所有其它指令完成后，postrotate和endscript里面指定的命令将被执行。

每个微服务的日志空间默认只有10G,请根据空间大小自行设置日志切割规则和日志保留时间
```


### 调试
立刻运行/etc/lograte.d/目录下的所有日志配置：

logrotate /etc/logrotate.conf 
立刻运行特定的配置：

logrotate /etc/logrotate.d/svc_notify
调试过程中的最佳选择是使用‘-d’选项以预演方式运行logrotate。要进行验证，不用实际轮循任何日志文件，可以模拟演练日志轮循并显示其输出。

即使轮循条件没有满足，我们也可以通过使用‘-f’选项来强制logrotate轮循日志文件，‘-v’参数提供了详细的输出。

Dockerfile Logrotate 配置
1. 在项目中创建conf文件,并且移动到docker容器的/etc/logrotate.d/目录,重命名为main

2. 把容器中的/etc/logrotate.d/main权限修改为0644

3. 配置定时任务,按天执行的放在/etc/periodic/daily/logrotate,按小时执行的放在/etc/periodic/hourly/

4. 示例:
```
COPY logrotate/conf /etc/logrotate.d/main
RUN chmod +x /bin/start.sh \ 
    && chmod 0644 /etc/logrotate.d/main \
    && mv /etc/periodic/daily/logrotate /etc/periodic/hourly/
```

## 日志收集
### 日志服务部署文档
我们的微服务一般都是放在alpine中，下面介绍alpine中安装fluentd，并且收集日志的办法。

### 目的
掌握配置dockerfile 和 日志收集文件的方法，以完成相关日志的收集
日志收集的前提：需要把日志打印到日志文件中，并且打印到日志文件中的日志格式必须是json
第一步：编写fluent.conf文件
现在假设我们要收集vp中/data/main.log文件中的日志

在项目的根目录建立fluentd文件夹
在fluentd文件夹下创建fluentd.conf中
编辑fluentd.conf文件内容:
```
<source>
    @type tail
    format json
    time_key time
    time_format %Y-%m-%dT%H:%M:%S.%N%z
    keep_time_key true
    path /data/main.log
    tag svc_notify_main.log
    pos_file /fluentd/main.log.pos
</source>
<match svc_notify_main.*>
    @type elasticsearch 
    host "#{ENV['FLUENTD_HOST']}"
    port 9200
    include_tag_key true
    tag_key @log_name
    logstash_format true
    logstash_prefix myun.svc-notify.main
    flush_interval 2s
    reconnect_on_error true
    reload_on_failure true
    reload_connections false
    buffer_chunk_limit 100m
    buffer_queue_limit 1024
</match>

<source>
    @type tail
    format json
    time_key time
    time_format %Y-%m-%dT%H:%M:%S.%N%z
    keep_time_key true
    path /data/access.log
    tag svc_notify_access.log
    pos_file /fluentd/access.log.pos
</source>
<match svc_notify_access.*>
    @type elasticsearch 
    host "#{ENV['FLUENTD_HOST']}"
    port 9200
    include_tag_key true
    tag_key @log_name
    logstash_format true
    logstash_prefix myun.svc-notify.access
    flush_interval 2s
    reconnect_on_error true
    reload_on_failure true
    reload_connections false
    buffer_chunk_limit 100m
    buffer_queue_limit 1024
</match>
```



### 线上报警
微服务 ---打印日志----> 本地文件 ----收集---> fluentd ----发送---> elk

当前的微服务都要求使用logrotate来切分日志, 所有微服务打印的日志都会写到本地的文件中,一般放在/data/下, 同时和微服务一起部署的fluentd会使用tail的方式收集日志,并且发送到elk, 日志报警程序会定期根据日志等级查询elk中的日志,查询到日志后会通过钉钉报警机器人发送到钉钉



### 日志查询




