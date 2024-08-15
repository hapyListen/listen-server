# listen-server
Music Center Server.
## 环境搭建

### Golang
- 版本： 1.20.12

### Mysql
- 版本： 5.7.44
- 配置： 
    ```shell
    docker pull mysql:5.7.44
    docker run -itd --name mysql-01 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.44
    ```

### Redis
- 版本：7.0.12
- 配置：
    ```shell
    docker pull redis:7.0.12
    docker run -p 6379:6379 --name redis-01 -d redis:7.0.12  --requirepass "123456"

    ```