# Micro相关笔记
## 技术栈
demo项目主要为了体验微服务框架中相关组件，中间件的使用方式和体验
涉及到的组件/中间件如下：  
---
* 注册中心：[consul官网](https://www.consul.io/docs/install) | [go-micro/plugins](https://github.com/asim/go-micro/tree/master/plugins/registry/consul)
* 链路追踪：[Jaeger官网](https://www.jaegertracing.io/)
* 熔断：[hystrix-go github](https://github.com/afex/hystrix-go)
* 负载均衡：[roundrobin](https://github.com/asim/go-micro/tree/master/plugins/wrapper/select/roundrobin)
* 限流：[ratelimiter/uber](https://github.com/asim/go-micro/tree/master/plugins/wrapper/ratelimiter/uber)
* 系统监控：[promethues官网](https://prometheus.io/)
* 日志系统：ELK（Elasticsearch , Logstash, Kibana , Beats）
* 支付模块对接的是 [paypal](https://www.paypal.com/) | [沙盒环境](https://developer.paypal.com/) | [使用虚拟账号查看流水](https://www.sandbox.paypal.com/ )
* 路由组件：[gin](https://github.com/gin-gonic/gin)
---

## 项目
payment模块使用的组件最全，基本上都涵盖了。其中payment是服务端，paymentapi是客户端提供对外的http接口
### 启动方式
* docker-compose up -d 启动 consul，jaeger，grafana，prometheus，hystrix-dashboard
* 进入consul（默认）控制台配置mysql
```
// 样例，key:micro/config/mysql
{"host":"127.0.0.1","port":3306,"user":"","password":"","database":""}
```
* 启动日志系统 : ```docker-compose -f ./docker-elk/docker-stack.yml up```
* 进入到payment模块 运行go run main.go
* 进入到paymentapi模块 运行go run main.go
* 通过查看consul确定是否启动成功
* 访问demo接口
---
## 资源
* [micro github](https://github.com/micro/micro)
* [go-micro github](https://github.com/asim/go-micro)
* [docker 命令文档](https://skyao.io/learning-docker/docs/introduction.html)