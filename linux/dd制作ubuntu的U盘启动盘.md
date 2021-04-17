## dd制作ubuntu的U盘启动盘:
    1. 下载ubuntu的ISO镜像到本地  https://ubuntu.com/download/desktop
        wget -c "http://mirrors.melbourne.co.uk/ubuntu-releases/18.04.4/ubuntu-18.04.4-desktop-amd64.iso"
        wget -c "http://releases.ubuntu.com/19.10/ubuntu-19.10-desktop-amd64.iso"
        wget -c 'https://mirrors.melbourne.co.uk/ubuntu-releases/20.04/ubuntu-20.04-desktop-amd64.iso'
    2. 如果U盘被自动挂载，请使用U盘设备号先umount: sudo umount /dev/sdb*
    3. 把u盘划分为一个分区  sudo fdisk /dev/sdb
    4. 格式化 
        sudo mkfs.vfat /dev/sdb -I
        sudo mkfs.ntfs /dev/sdb1
        sudo mkfs.ext4 /dev/sdb1
    5. 开始制作U盘启动盘: sudo dd if=/home/mtoou/下载/xubunut.iso of=/dev/sdb status=progress

## ffmpeg学习:
    ubuntu系统下载ffmpeg：sudo apt install ffmpeg ; ffmpeg -version

    推流命令：ffmpeg -re -i 本地视频.mp4 -vcodec copy -acodec copy -b:v 800k -b:a 32k -f flv "推流地址"

    拉流命令：ffplay "拉流地址"

    学习链接：http://linux.51yip.com/search/ffmpeg
http://mirror.ox.ac.uk/sites/releases.ubuntu.com/releases/19.10/ubuntu-19.10-desktop-amd64.iso
http://releases.ubuntu.com/19.10/ubuntu-19.10-desktop-amd64.iso?_ga=2.22020091.1998950585.1579500595-828916891.1579500595