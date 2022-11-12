# 七牛云——多人协作画板



***

#### 1. Quick Start

- [建议] Docker 部署:

  1. 需要安装docker
  2. 依次执行以下命令

  ```bash
  $ docker network create qiniu
  $ docker run --name mysql -p 3306:3306 --network=qiniu -v /mydata/conf/my.cnf:/etc/mysql/my.cnf -v /mydata/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=19990221 -d mysql
  $ docker exec -it mysql /bin/bash
  --- 进入mysql建表, sql文件在backend/model里 ---
  
  $ docker run -p 6379:6379 --name=redis -d  --network=qiniu redis:6.0.8
  
  $ docker run -d -p 9001:9000 --network=qiniu  suyame/qiniu-ws
  $ docker run -d -p 9002:9000 --network=qiniu  suyame/qiniu-ws
  $ docker run -d -p 9003:9000 --network=qiniu  suyame/qiniu-ws
  
  $ docker run -d -p 8080:8080 --network=qiniu suyame/qiniu-server
  
  ```

- 常规部署环境：

  1. golang 1.19.2
  2. mysql 8.0
  3. redis 6.0.8

- 运行步骤：

  1. 将项目克隆到本地：

     ```bash
     git clone https://github.com/YanShulinjj/qiniuProj.git
     ```

  2. 运行 ./back_end/model 里的两个sql文件

  3. 修改 ./back_end/etc/config.yml 文件中的相应字段

     ```yaml
     Name: qiniu-api
     Host: 0.0.0.0              # 主机IP
     Port: :8080                # 主机暴露端口
     Mode: dev              
     SVGPATH: ./data/svg
     Mysql:                     # mysql 连接地址
       DataSource: root:19990221@tcp(127.0.0.1:3306)/qiniu?  		charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai           
     
     CacheRedis:
       - Host: 127.0.0.1:6379   # redis 服务地址
         Type: node
         Pass:
     ```

  4. 编译+运行

     ```bash
     $ cd ws &&  go build start.go router.go
     $ start 127.0.0.1:9001
     $ start 127.0.0.1:9002
     $ start 127.0.0.1:9003
     $ cd ../back_end
     $ go build main.go router.go
     $ ./main
     ```

- 启动服务后, 可在浏览器访问: http:yourhost:8080/qiniu



#### 2.文件说明

```bash
│  Dockerfile									# HTTP服务Docker镜像文件
│  README.md
│  
├─back_end										# 后端文件夹
│  │  main.go									# 后端程序入口
│  │  router.go									# HTTP路由定义
│  │  
│  ├─config										# HTTP服务配置初始化
│  │      config.go
│  │      
│  ├─controller									# 控制层，主要定义路由函数
│  │  └─api
│  │      ├─response
│  │      │      response.go
│  │      │      
│  │      └─v1
│  │              common.go
│  │              page.go
│  │              user.go
│  │              
│  ├─dao										# dao层，封装对数据库基本数据操作
│  │      common.go
│  │      page.go
│  │      user.go
│  │      
│  ├─data										# 用于存储画板页面文件夹，数据的svgPath会链接到此
│  │  └─svg
│  │              
│  ├─etc										# 配置文件
│  │      config.yaml
│  │      
│  ├─model										# 对缓存和数据库基本操作
│  │      page.sql
│  │      pagemodel.go
│  │      pagemodel_gen.go
│  │      user.sql
│  │      usermodel.go
│  │      usermodel_gen.go
│  │      vars.go
│  │      
│  ├─pkg										# 外部依赖
│  │  ├─encryption							
│  │  │      md5.go								# 用于字符串加密
│  │  │      
│  │  ├─svg
│  │  │      svg.go								# 用于生成svgPath
│  │  │      
│  │  ├─verify
│  │  │      verify.go							# 验证特定字符串格式，例如IP格式
│  │  │      verify_test.go
│  │  │      
│  │  └─xerr									# 错误代码定义
│  │          errCode.go
│  │          errMsg.go
│  │          errors.go
│  │          
│  ├─service									# service层
│  │      page.go
│  │      user.go
│  │      
│  └─splitter									# 分流器，水平扩展，保证每个用户的所属页面都在一个ws服务器建立连接，但不保证不同用户的页面在同一个ws服务器建立连接
│      │  common.go
│      │  map.go
│      │  
│      ├─loadbalance							# 负载均衡，轮询方式
│      │      error.go
│      │      random.go
│      │      random_test.go
│      │      rr.go
│      │      rr_test.go
│      │      servers.go
│      │      test_example.go
│      │      
│      └─server			
│              server.go
│              
├─doc 
├─front_end										# 前端文件夹
│  │  index.html
│  │  
│  └─public
│      ├─css
│      │      iconfont.css
│      │      index.css
│      │      
│      ├─font
│      │      iconfont.json
│      │      iconfont.ttf
│      │      iconfont.woff
│      │      iconfont.woff2
│      │      
│      └─js
│              hidden.js
│              iconfont.js
│              index.js
│              server.js
│              text.js
│              ws.js
│              
└─ws											# WebSocket服务，用于实时同步画板
    │  Dockerfile								# WebSocket服务Docker部署
    │  router.go
    │  start.go
    │  
    ├─conf
    │      config.go
    │      
    ├─etc
    │      config.yaml
    │      
    ├─internal									# 建立连接逻辑，广播消息逻辑
    │      client.go	
    │      common.go
    │      handler.go
    │      main.go
    │      manager.go
    │      managergroup.go
    │      
    └─pkg
        └─xerr
                errCode.go
                errMsg.go
                errors.go
```



#### 3.Demo展示

*若有显示错误，[点我查看](https://410proj.oss-cn-chengdu.aliyuncs.com/qiniu/qiniu.mp4)

<video src="https://410proj.oss-cn-chengdu.aliyuncs.com/qiniu/qiniu.mp4" controls="controls" width="800" height="600">你的浏览器不支持播放该视频</video>

