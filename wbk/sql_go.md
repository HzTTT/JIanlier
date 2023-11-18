# 1.设计数据库

## 1.数据库设计网址

https://dbdiagram.io/home

他可以用来架构数据库，并且能把架构另存为PDF或者PNG

也可以生成sql代码，可以是PostgreSQL、Mysql、SQL sever

### 1.自增类型 bigserial

### 2.主键 pk

示例代码：

```sql
 id bigserial [pk]
```

# 2.Docker+Postgres+TablePlus

## 1.docker安装(linux)

```linux
#打开终端
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg

sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o/etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update

# List the available versions:
apt-cache madison docker-ce | awk '{ print $3 }'

5:24.0.0-1~ubuntu.22.04~jammy
5:23.0.6-1~ubuntu.22.04~jammy
<...>

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

```linux
 #测试一下装成功没
 sudo docker run hello-world
```

然后就emmm，就莫名装好了

## 2.拉取 Postgres 镜像

```linux
#打开终端

sudo docker pull postgres:12-alpine
sudo docker run postgres12

#Docker 端口映射
sudo docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

#访问 Postgres 控制台
sudo docker exec -it postgres12 psql -U root
sudo docker exec -it postgres12 /bin/sh
#控制台可做的操作
select now();  //查看当前时间
/q     //退出控制台

#查看容器日志
sudo docker logs postgres12

#启动容器
sudo docker start (容器name)

#创建容器中创数据库
sudo createdb --username=root --owner=root simple_bank
#删除
sudo dropdb simple_bank
```

> 映像（image）与容器（container）

容器是映像中包含的应用程序的 1 个实例，由 `docker run` 命令启动。

1 个映像可以启动多个容器。

```
#查看容器
sudo docker ps
```

# 3.数据库迁移

安装

```
sudo curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
sudo echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
sudo apt-get update
sudo apt-get install -y migrate
```



不知道迁移是为了做什么

```sql
#向上迁移
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
#向下迁移
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
```

# 4.SQL 生成 CRUD Golang 代码

安装sqlc

```
sudo snap install sqlc
```

装成功了,执行一下下面的命令，可以看说明书

```
sqlc help
```

（哎，不会用makefile，就直接执行原始命令好了）

![image-20231117124820980](sql_go.assets/image-20231117124820980-17002006860782.png)

OK，这里报错error的没啥事，解释一下先

## 1.sqlc init

初始化一个 sqlc.yaml 文件在数据库的文件夹里（在数据库的文件夹的目录下执行的命令）

然后往里面写

```
version: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: false
    emit_interface: false
    emit_exact_table_names: false
```

这些代码可以在文档里找到

[Getting started with MySQL — sqlc 1.23.0 documentation --- MySQL 入门 — sqlc 1.23.0 文档](https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html)

（我复制一下原话翻译）

这里的 `name` 选项是告诉 sqlc 将要生成的 Go 包的名称是什么。我认为这 `db` 是一个很好的包名。

接下来，我们必须指定 `path` to 文件夹来存储生成的 golang 代码文件。我将在该文件夹中创建一个新文件夹 `sqlc` ，并将此 `path` 字符串更改为 `./db/sqlc` . `db`

然后， `queries` 我们可以选择告诉 sqlc 在哪里查找 SQL 查询文件。让我们在文件夹内创建一个新 `db` 文件夹 `query` 。然后将此值更改为 `./db/query`

同样，此架构选项应指向包含数据库架构或迁移文件的文件夹。在我们的例子中，它是 `./db/migration`

下一个选项是 `engine` 告诉 sqlc 我们要使用什么数据库引擎。我们正在用于 `Postgresql` 我们的简单银行项目。如果要试验 MySQL，可以将此值更改为 `mysql`

在这里， `emit_json_tags` 我们设置 to `true` ，因为我们希望 sqlc 将 JSON 标记添加到生成的结构中

告诉 `emit_prepared_queries` sqlc 生成使用预处理语句的代码。目前，我们还不需要优化性能，因此让我们将其 `false` 设置为简单

然后， `emit_interface` 选择告诉 sqlc 为生成的包生成 `Querier` 接口。如果我们想模拟 db 来测试更高级别的函数，它可能会很有用。现在，让我们将其设置为 `false`

最后一个选项是 `emit_exact_table_names` 。默认情况下，此值为 `false` 。Sqlc 将尝试将表名单数化以用作模型结构名称。例如 `accounts` ，table 将变为 `Account` struct。如果将此选项设置为 true，则结构名称将改为 `Accounts` 。我认为单数名称更好，因为复数形式的 1 个类型的 `Accounts` 对象可能会混淆为多个对象

<hr>

## 2.sqlc generate

执行这代码前，先去这个文件夹目录下生成一个query文件夹

![image-20231117125709585](sql_go.assets/image-20231117125709585.png)

然后再在query文件夹中写这些东东（其实就是一个基础的插入查询语句）

值得注意的是

`RETURNING *` 子句用于告诉 Postgres 在将记录插入账户表后返回所有列的值（包括 `id` 和 `created_at` ）。

![image-20231117125906974](sql_go.assets/image-20231117125906974.png)

OK，搞定完这些就可以执行命令了

执行完会出现3个go文件

## 3. go mod init github.com/techschool/simplebank

## 4.go mod tidy

安装依赖

## 5.小小的尾巴

如果要改代码，不要对那3个生成的go文件进行修改，改的应该是那个sql文件，然后再次生成文件（sqlc generate）

# 5.测试

在 Golang 中，我们有一个约定，将测试文件与代码放在同一个文件夹中，测试文件的名称应该以后 `_test` 缀结尾。

Go 中的每个单元测试函数都必须以前缀（大写字母 T）开头

## 1.代码覆盖率怎么弄出来？

![image-20231117134331314](sql_go.assets/image-20231117134331314.png)
