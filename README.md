# Go-fiber-ent-web-layout  GoWeb应用快速开发模板项目

使用`go-fiber`Web框架和`ent`ORM框架，集成了`wire`依赖注入和`yaml`配置文件读取，文件目录参照bilibili开源的`go-kratos`微服务开发框架

## 项目简介

此模板项目主要面向`Golang`单体项目后端开发，数据交互均使用`json`格式，集成了错误处理，`panic`捕获、通用返回消息、权限验证、依赖注入、配置文件读取等大部分创建项目时需要重复使用的代码。

### 使用到的工具

- go-fiber：Go最快的`web`开发框架
- ent：`facebook`开源的`ORM`框架
- golang-jwt：`JWT`的`Go`实现
- sonic：字节跳动开源的序列化工具
- yaml：读取`.yaml`配置文件
- validator：结构体参数验证

### 通用返回消息

```go
// /cmd/internal/tools/res.go

type Result struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
```

### 错误处理

```go
// /cmd/internal/tools/tools.go

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code, message := http.StatusInternalServerError, "server error"
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	result := Fail(code, message)
	return ctx.Status(code).JSON(result)
}
```

### 依赖注入

```go
// /cmd/go-fiber-ent-web-layout

func wireApp(*conf.Data, *conf.Jwt) (*fiber.App, func(), error) {
	panic(wire.Build(api.InjectSet, data.InjectSet, service.InjectSet, common.InjectSet, middlewares.NewAuthMiddleware, newApp))
}
```

## 快速开始

> 在使用此模板之前，请确保`golang`版本大于等于`1.21`并且已经启用了`go mod`管理依赖

`Clone`此项目

```bash
git clone https://github.com/wnnce/go-fiber-ent-web-layout.git
```

获取依赖

```bash
cd go-fiber-ent-web-layout

go mod tidy
```

安装`ent`和`wire`工具

```bash
go install entgo.io/ent/cmd/ent@latest

go install github.com/google/wire/cmd/wire@latest
```

生成依赖注入代码

```bash
go generate ./cmd/go-fiber-ent-web-layou/
```

编译项目，**配置文件不会被打包，运行二进制文件时可以使用 -conf 参数指定配置文件路径**

```bash
go build .\cmd\go-fiber-ent-web-layout\ 

# 打包完成后运行 需要指定配置文件路径
./go-fiber-ent-web-layout.exe -conf .\config.yaml
```

> 如果需要添加实体类或数据表可以参考`ent`框架的文档