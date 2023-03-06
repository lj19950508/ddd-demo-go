基础框架(ioc/基础事务/)
### 项目结构
aggregation
  - adater
    - in
        - httpapi  
        - websocket
        - mqconsumer
    - out 
        - mqproducer
        - mysql
        - repositoryImpl
        - redis
        - queryService
        - queryServiceImpl
        - config
        - eventbusimpl
  - application
    - command
      - cmdservice
      - dto
    - query
      - querysvc
      - ->querydto/ resultdto->
    - factory
  - domain
    - root/entity/valueobject
    - repository
    - domainService
    - event （ID，type,status?,time，data）
    - eventbus ? 这个interface放哪一层呢
  - pkg  （可独立出来的包）
    - eventbus 抽象eventbus
  
  mq调用handler 同步 读库的操作还没想好怎么实现 先不管
  可靠消息最终一致性 依赖于事务消息   首先消息的存取都是事务的  回滚则要补偿（反操作）

### DDD项目规范 
api
 在api层本身定义自身路由，然后httphandler构造的时候会去register路由
out
  repoimpl 实现的是领域层定义的接口，然后再注入db做具体操作
  queryimpl 实现的是application.query层的接口， 可以注入多种数据源， redis es db 等做具体查询操作
  eventHandler 监听commandservice发出的领域事件处理   通过 eventbus.dispture(eventhandler) 实现 ，  其中eventhandler会绑定一个事件 event 这样就实现 eventhandler处理event
application  这里将service拆成 query和command
  command  
     实现cqrs command无返回值(错误不算返回值), 
     这里可以注入writedb.tx事务操作 从Load到Save 做事务操作 也可以引入eventbus.send(event) 这里主要是做调度功能各个领域功能。不做具体业务操作逻辑  event（这里应该也可以引入领域事件）
     如果这里有通过调用其他服务的场景，如果是通过 http/grpc框架 能够返回一个 bizerror，可以通过这个bizerror 判断是否回滚
     如果是通过消息队列发送的话，如果有事务一执行的需求，可以使用补偿操作。 这里就不打算引入分布式事务框架了
  query 只定义了 qeuryservice的接口和 query_dto,query result, 由 queryimpl实现对应的操作逻辑
config 理论上可以移动到out/config
domain
  entity/聚合根  不要成为一个贫血模型哦 把业务操作写在这 让后通用command调度
  repo 仓储定义
  error 错误定义
  domainservice 领域服务 需要多个当前领域对象用的 无法用domain单独描述的才写到这， 这个类极少用到
  event定义 定义了领域事件

--------------------

领域 子域 限界上下文 战略设计

实体 值对象 聚合 工厂 repo 领域服务 领域事件 战术设计

上下文映射图 => 服务（上下文）依赖方向  

一个上下文一个子域 是一个微服务



领域 子域 上下文
实现边界上下文 = 一个子域 每个子域 等于一个服务
区分聚合根  功能独立  订单 订单子项（）
值对象理解为（属性）

调用授权和鉴权的功能交给网关
授权=>统一登陆
鉴权教给拦截器




### 总结一下
多用户类型用前缀分组如
/api（客户，游客）  解析token（在本体,不存在query远程调用的情况）
/merchant (商户)   解析token  在本体
/admin (管理员)    解析token  
/物流配送 (如果配送物流是自己管理的话)
之所以这么粉因为不同类型的用户 调用的接口基本上都是不一样的 从query上来说

//登录command想返回 token need
//TODO存不存在跨服务调用需要返回值的情况
//command 生成token到数据库 
//然后根据 id查询token

用户登录的场景中
用户根据 登录凭证 生成jwt
command可以返回值 
用户 ->提交登录凭证 ——> query
command 附带excuteResult or return
eventbus.handler->c
repo 变成 add和save方法