# Golang 六边形架构模板（基于Gin）

## 测试

```
golangci-lint run ./...
go generate ./...
go test -cover ./...
```

## 依赖

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/gin-gonic/gin
go get -u github.com/go-playground/validator/v10
go get -u github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@v1.6.0
```

## 六边形架构

所有数据处理全部由repository适配成实体，逻辑都是领域服务、聚合、实体的行为。

分多少层和平层、跨层调用本身也不存在

```
adapter 端口适配器
    - driver 输入适配器
        - RESTHandler
        - MQHandler
    - driven 输出适配器
domain 领域层 (定义应用程序的域和业务逻辑的地方)
    - aggregate 聚合
    - entity 领域实体（接口，并实现）
    - enums 枚举字段
    - repository 存储库 (类似三层架构的DAO接口，不包括实现；只管CRUD，不管业务逻辑)   
    - vo 视图对象
infra 基础设施层 (为其他层提供通用的技术能力：业务平台，编程框架，持久化机制，消息机制，第三方库的封装，通用算法，等等)
    - config 配置
    - db 数据库
    - errors 错误
    - log 日志
    - middleware 中间件
    - model 模型
    - persistence 持久化 (对domain层的repository接口的实现)
    - requests
    - responses
    - utils
server
    - config
        - config.yaml
    - main.go
```

概念：

- VO（View Object）：视图对象，用于展示层，它的作用是把某个指定页面（或组件）的所有数据封装起来。
- DTO（Data Transfer
  Object）：数据传输对象，这个概念来源于J2EE的设计模式，原来的目的是为了EJB的分布式应用提供粗粒度的数据实体，以减少分布式调用的次数，从而提高分布式调用的性能和降低网络负载，但在这里，我泛指用于展示层与服务层之间的数据传输对象。
- DO（Domain Object）：领域对象，就是从现实世界中抽象出来的有形或无形的业务实体。
- PO（Persistent Object）：持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么，数据表中的每个字段（或若干个）就对应PO的一个（或若干个）属性。

```
用户发出请求（可能是填写表单），表单的数据在展示层被匹配为VO。
展示层把VO转换为服务层对应方法所要求的DTO，传送给服务层。
服务层首先根据DTO的数据构造（或重建）一个DO，调用DO的业务方法完成具体业务。
服务层把DO转换为持久层对应的PO（可以使用ORM工具，也可以不用），调用持久层的持久化方法，把PO传递给它，完成持久化操作。
对于一个逆向操作，如读取数据，也是用类似的方式转换和传递，略。
```

## 数据库

```
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `go_project` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci */;
```
