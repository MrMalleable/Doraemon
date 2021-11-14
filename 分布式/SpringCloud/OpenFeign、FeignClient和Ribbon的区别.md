## Ribbon
Ribbon 是 Netflix开源的基于HTTP和TCP等协议负载均衡组件

Ribbon 可以用来做客户端负载均衡，调用注册中心的服务

Ribbon的使用需要代码里手动调用目标服务，请参考官方示例：https://github.com/Netflix/ribbon

项目中多见RestTemplate和Ribbon相结合,如以下代码：
```java
@Configuration
public class SysConfiguration {

    @Bean
    @LoadBalanced
    public RestTemplate getRestTemplate(){
        return  new RestTemplate();
    }
}


@RestController
public class ConsumerController {

    @Resource
    private RestTemplate restTemplate;

    /**
     * 服务提供的地址：即服务提供方的spring.application.name 转大写
     * 也可以理解为在Eureka服务注册中心，注册的application值。
     */
    final String urlPath = "http://ORDER-PROVIDER";

    @GetMapping("/consumer/make/order")
    public String makeOrder(@RequestParam("name") String name){
        String template = restTemplate.getForObject(urlPath + "/order-server/create/order", String.class);
        return name + template;
    }
}
```

Ribbon实现负载均衡（**进程式负载均衡**）是在客户端实现的，而常见的以nginx作负载均衡（**集中式负载均衡**）时，是请求先到nginx，由nginx的负载均衡策略来决定去请求后台部署的哪台服务器。

## Feign
Feign是Spring Cloud组件中的一个轻量级RESTful的HTTP服务客户端

**Feign内置了Ribbon**，用来做客户端负载均衡，去调用服务注册中心的服务。

Feign的使用方式是：使用Feign的注解定义接口，调用这个接口，就可以调用服务注册中心的服务

Feign支持的注解和用法请参考官方文档：https://github.com/OpenFeign/feign

Feign本身不支持Spring MVC的注解，它有一套自己的注解

## OpenFeign
OpenFeign是Spring Cloud 在Feign的基础上支持了Spring MVC的注解，如@RequesMapping等等。

OpenFeign的@FeignClient可以解析SpringMVC的@RequestMapping注解下的接口，

并通过动态代理的方式产生实现类，实现类中做负载均衡并调用其他服务。

**需要注意，@RequesMapping不能在类名上与@FeignClient同时使用**

---
参考：
[Ribbon、Feign和OpenFeign的区别](https://blog.csdn.net/zimou5581/article/details/89949852)