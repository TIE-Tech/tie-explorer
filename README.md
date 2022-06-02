##TieChain 浏览器

`免责声明: 仅供技术交流学习,不适用于商业用途`

本代码分前，后台. 所具的功能是针对 Polygon-edge 区块链的,
与 Ethereum 链的功能不完全一致.

后台代码 使用 `Beego` 框架开发. 
前台代码 使用 `Vue` 开发.

数据缓存采用 `Postgresql` 数据库.

浏览器后台功能:

- 解析区块信息
- 解析交易信息 (解析交易Logs,细化交易详情)
- 统计用户信息
- 统计每日交易数据
- 统计链基本信息

浏览器页面功能:

- 首页
- 区块列表 & 区块详情
- 交易列表 & 交易详情
- 账户地址信息

###运行

#### 启动数据采集功能

```shell
go run main.go
或
go build -o indexer main.go
./indexer
```

#### 启动浏览器页面
```shell
cd web
npm install
npm run serve
```


### Docker 运行

#### 手动安装方式
第1步: docker 安装 并运行 Postgres 数据库
```shell
docker pull postgres:13.6

docker run --name postgres -e POSTGRES_PASSWORD='' -e POSTGRES_USER='postgres' -e POSTGRES_HOST_AUTH_METHOD='trust' -p 5432:5432 -d postgres:13.6
```

第2步: 创建数据库 `(可跳过)`
```shell
docker exec -it postgres /bin/bash
psql -U postgres
CREATE DATABASE explorer;
```

第3步: docker 安装 golang
```shell
#修改 conf/app.conf
#将 Postgresql 数据库连接信息 调整为对应信息

docker build . -f docker/server/Dockerfile -t server

docker run --name server -p 8081:8081 -d server
```

第4步: docker 安装 Nginx
```shell
docker build . -f docker/web/Dockerfile -t web

docker run --name web -p 8080:8080 -d
```


#### Docker-compose 方式
```shell
docker-compose up
```
