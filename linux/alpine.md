alpine 提供了非常好用的apk软件包管理工具，通过apk –help命令查看完整的包管理命令。

## apk --help  apk查询命令
    apk-tools 2.10.4, compiled for x86_64.

    Installing and removing packages:
      add       Add PACKAGEs to 'world' and install (or upgrade) them, while ensuring that all dependencies are met
      del       Remove PACKAGEs from 'world' and uninstall them

    System maintenance:
      fix       Repair package or upgrade it without modifying main dependencies
      update    Update repository indexes from all remote repositories
      upgrade   Upgrade currently installed packages to match repositories
      cache     Download missing PACKAGEs to cache and/or delete unneeded files from cache

    Querying information about packages:
      info      Give detailed information about PACKAGEs or repositories
      list      List packages by PATTERN and other criteria
      dot       Generate graphviz graphs
      policy    Show repository policy for packages

    Repository maintenance:
      index     Create repository index file from FILEs
      fetch     Download PACKAGEs from global repositories to a local directory
      verify    Verify package integrity and signature
      manifest  Show checksums of package contents

    Use apk <command> --help for command-specific help.
    Use apk --help --verbose for a full command listing.

    This apk has coffee making abilities.


## 更新索引
    update：从远程镜像源中更新本地镜像源索引，update命令会从各个镜像源列表下载APKINDEX.tar.gz并存储到本地缓存，一般在/var/cache/apk/(Alpine在该目录下)、 /var/lib/apk/ 、/etc/apk/cache/下
    apk update
    
## 安装包
    add：命令从仓库中安装最新软件包，并自动安装必须的依赖包,也可以从第三方仓库添加软件包。add:安装PACKAGES并自动解决依赖关系。
    
    apk add openssh openntp vim busybox-extras curl 
    apk add --no-cache mysql-client
    apk add docker --update-cache --repository http://mirrors.ustc.edu.cn/alpine/v3.4/main/ --allow-untrusted
    
    安装指定版本软件包
    apk add asterisk=1.6.0.21-r0
    apk add 'asterisk<1.6.1'
    apk add 'asterisk>1.6.1
    
## 卸载
    del：卸载并删除PACKAGES
    apk del openssh openntp vim
    
## 升级
    upgrade命令升级系统已安装的所以软件包（一般包括内核），当然也可指定仅升级部分软件包（通过-u或–upgrade选择指定）。

    apk update #更新最新本地镜像源

    apk upgrade #升级软件

    apk add --upgrade busybox #指定升级部分软件包
    
## 搜索
    search命令搜索可用软件包，-v参数输出描述内容，支出通配符，-d或–description参数指定通过软件包描述查询。
    apk search #查找所以可用软件包
    apk search -v #查找所以可用软件包及其描述内容
    apk search -v 'acf*' #通过软件包名称查找软件包
    apk search -v -d 'docker' #通过描述文件查找特定的软件包
    
## 查看包信息
    info命令用于显示软件包的信息。
    apk info #列出所有已安装的软件包
    apk info -a zlib #显示完整的软件包信息
    apk info --who-owns /sbin/lbu #显示指定文件属于的包
