
## spectrum-blockchain
基于区块链的频谱资源交易平台，频谱资源管理中心将频谱资源授权给相关机构，机构可将频谱资源出售或出租给其他机构。

#### 所需环境
1. 系统： Linux（如Ubuntu16以上版本），必须连网
2. Docker和Docker Compose

#### 安装步骤
1. 环境初始化：首先，清理之前部署spectrum-blockchain时建立的镜像，运行/deploy中的./clear_images.sh；其次，
   清理Ubuntu中所有可能残留的文件，点击Ubuntu界面左下角图标，选择Disk Usage Analyzer,进入ubuntu,
   清理之前搭建spectrum-blockchain时，残留的文件；
   
2. 将项目拷贝到Linux相关目录下，并给予项目权限，执行 sudo chmod -R +x ./spectrum-blockchain/，此外非root用户需要加入docker组，
   sudo usermod -aG docker username；
3. 安装区块链底层环境：打开/deploy，运行./start.sh；
4. 安装区块链浏览器：打开/deploy/explorer，运行./start.sh；
5. 配置UI界面环境：打开/vue，运行./build.sh；
6. 建立应用环境：打开/application，运行./build.sh；
7. 启动应用：打开/application，运行./start.sh。

#### 操作步骤
1. 在浏览器中输入http://localhost:8000/web/ ，可打开系统界面；
2. 点击右上角的小人图标，选择区块链浏览器，可查看区块信息，区块链浏览器账户为admin，密码为123456。
