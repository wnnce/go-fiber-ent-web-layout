**程序内部目录**

- [common](common)：全局函数、工具函数或者一些通用服务的包
- [conf](conf)：定义配置、解析配置
- [data](data)：处理持久层，包括创建数据库连接客户端，业务的持久层操作等，需要实现[usercase](usercase)中对应的持久层接口
- [factory](factory)：工厂包，工厂模式使用
- [middlewares](middlewares)：存放需要使用的中间件
- [service](service)：业务层，供`API`层调用，需要实现[usercase](usercase)中对应的业务层接口
- [usercase](usercase)：定义实体类（使用`ent`框架无需定义），持久层和业务层接口

除了工厂包以外，其余如果需要创建单例的结构体或接口，都需要通过向对应包中的`InjectSet`添加结构体构造函数的方式来进行依赖注入