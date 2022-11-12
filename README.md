# 七牛云——多人协作画板



***

- *Docker 部署:

  1. 需要安装docker

  ```bash
  $ docker network create qiniu
  $ docker run --name mysql -p 3306:3306 --network=qiniu -v /mydata/conf/my.cnf:/etc/mysql/my.cnf -v /mydata/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=19990221 -d mysql
  $ docker exec -it mysql /bin/bash
  --- 进入mysql建表 ---
  
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

  1. 运行 ./back_end/model 里的两个sql文件

  2. 修改 ./back_end/etc/config.yml 文件中的相应字段

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

  3. 编译+运行

     ```bash
     $ cd ws &&  go build start.go router.go
     $ start 127.0.0.1:9001
     $ start 127.0.0.1:9002
$ start 127.0.0.1:9003
     $ cd ../back_end
     $ go build main.go router.go
     $ ./main
     ```
     
  
- 更多请查看docs

