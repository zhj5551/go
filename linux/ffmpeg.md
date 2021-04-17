    
## 安装ffmpeg
- sudo apt install ffmpeg ; ffmpeg -version
- 推流地址： ffmpeg -re -i 测试.mp4 -vcodec copy -acodec copy -b:v 800k -b:a 32k -f flv ""
- 拉流地址： ffplay "流地址"

### 在线中文手册
- [ffmpeg linux 命令 在线中文手册](http://linux.51yip.com/search/ffmpeg)