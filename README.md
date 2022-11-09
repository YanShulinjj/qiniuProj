#### 七牛云——多人协作画板

***

- 运行环境：

  1. golang 1.19.2

  2. mysql 8.0

     ```bash
     mysql -uroot -p
     ...
     <mysql> source ./back_end/model/user.sql
     <mysql> source /back_end/model/page.sql
     ```

  3. redis 

  4. edge | chrome | ..

- 运行步骤：

  1. 需要修改 ./back_end/etc/config.yml
  
  2. ```bash
     $ cd ..
   go run main.go
     ```
  
  3. 浏览器输入地址 127.0.0.1:8080/qiniu

