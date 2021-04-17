由于yum源中没有我们想要的nginx，那么我们就需要创建一个“/etc/yum.repos.d/nginx.repo”的文件，其实就是新增一个yum源

## 添加nginx.repo 文件：

    [root@localhost ~]# vim /etc/yum.repos.d/nginx.repo

    [nginx]
    name=nginx repo
    baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
    gpgcheck=0
    enabled=1

## 查看yum列表是否成功
    yum list |grep nginx

## 安装nginx
    yum -y install nginx

## 检查是否安装成功
    rpm -q nginx
