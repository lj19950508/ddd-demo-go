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


通常，我们希望将子域一对一地对应到限界上下文。 
上下文映射图 => 服务（上下文）依赖方向  

一个限界上下文就是一个项目 一个jar包？  希望能对应到子域？  一一对应  
upms 一个身份和访问上下文系统   而不是用户系统 权限系统 这根据业务来的。




在微服务中落地DDD

一个上下文一个子域 是一个微服务
微服务下有多个聚合
比如CMS服务下
有upms权限系统中有  user role permission 等  这种情况就没聚合了。。
如订单系统中   有订单项，订单子项  都有id , 这种情况据昂 order会聚合orderitem order就是 

其实冥冥之中已经实现了DDD了。。 就是引入一个 上下文 最好是通过引入包 或者 http调用的方式实现 。
比如我现在遇到的权限问题，我把handler写在每一个引入的上下文里 这样是不对的，得想想怎么解决


用户身份访问系统
认证用户身份（通过用户密码）颁发登录凭证
凭借这个登录凭借访问其他上下文又不经过

就得把接口抽象成resource 这个会损耗性能很多

1.鉴权网关 所有接口都走这个
 解析token获取用户信息
 用户 查看用户角色 资源
 系统中的资源都注册到 数据库里
 如果是http访问资源 则从系统中访问这个资源信息， 判断这个资源是否需要鉴权 如果不需要直接放过  如果需要 对比用户权限和资源权限是否符合 
 1.网关变成了流量焦点
 2.每次查询网关中负责解析token并，去查询数据信息（或redis） 这个都可以优化从redis中取 否则才找d
 3.资源节点都得依次添加到数据库
 4.依赖网关才能正常运行

优点
1.这个系统可以独立售卖了 符合DDD
token生命周期明确不会在 user1服务走完  服务2就过期
//依赖于身份认证系统 所有系统都用同一套用户体系
//id role dept 这些通用的
admin ->user     这些业务相关则自己写
custoemr ->user
merchant ->user
骑手 -> user      

 
如果按我这种做法
1.授权中心负责颁发用户凭证
2.每个服务都是独立鉴权并获取用户信息
3.服务间调用都要多一层token
4.不用把资源节点都注册到服务器 服务内拦截器判断即可
-------------------
好处
不用添加资源节点
多种用户类型比较方便
不需要网关即可启动 自身可行

坏处
每个服务都耦合的鉴权上下文，服务间调用都要传递token 并解析用户信息
不用走api网关 对网关压力小
不符合DDD 除非 
不能随便更改
要传递token信息

领域 子域 上下文
实现边界上下文 = 一个子域 每个子域 等于一个服务
区分聚合根  功能独立  订单 订单子项（）
值对象理解为（属性）

调用授权和鉴权的功能交给网关
授权=>统一登陆
鉴权教给拦截器


//1.是否交给api网关调用鉴权
//2.统一用户中心怎么做

认证 授权 鉴权 权限管理
认证 (登录凭证)   
account/pwd  
mobile/code 
第三方认证
Oauth2单点登录
二维码登录。。

授权 赋予执行者权限 （token）
鉴权 网关鉴权。过滤器鉴权
权限管理 UPMS
securiy 系统
security securityid groupid roleid
关联表 securityid - userid moduleid

group -moduleid
role -moduleid 
resource 
module merchant admin customer 配置

login_cert securityid  account pwd,mobile.cert
认证 根据login_cret匹配  
  | Login(accountpwd)
授权 生成（带资源集）token给user
  | return token
鉴权 根据token 判断访问权限  通常可以在网关拦截器调用
  |

网关调用 认证地址...
         鉴权地址...

admin->admin -> auth  <-token
顾客同一登录 商家统一登录 admin统一登录 有必要吗。
customer nickname 
merchant nickname
admin nickname

换种想法

认证服务生成 token 返回给用户
用户拿token 访问网关， 网关只判断 token是否合法 并解析放在 头部中，
合法以后则  应用服务根据传入的 userresource 来操作，
如果 应用服务内某个资源 需要权限，角色的判断 ，在拦截器中判断用户是否拥有这个权限

gateway->auth ->token
token->gateway-> 解析 -> appservice ->handler（拦截判断） 这样避免要有一个接口资源的定义


页面/菜单权限 有没有这个菜单
操作权限  能否执行这个接口操作 - 资源就得注册到上面 （如何注册（手动添加,自动注册））
数据权限 （通过context注入）
 - 行权限  (租户id)
 - 列权限  (需要配置)行下的列 不配置就是全部 ，

GATEWAY

资源中心
接口  GET /api 
菜单 menu 

服务间交互用eventbus
先不做权限了 ，
token解析好继续ing 
直接访问都要鉴权 CMD

得有一个统一用户中心 才好被  通用模块使用

cas统一认证中心 就是让 不同系统的认证集中在一起 不同域名子系统[实现单点登录]  防止这里登录一下 那里登录一下 。 目前没必要
oatuh 开放给第三方客户端 接入的能力 返回给第三方客户端登录token
认证中心
统一用户中心
upms
商家有独立的登录接口，登录凭证
买家也有独立的登录接口 和登录凭证
也不是统一用户了
/merchant/login
/user/login
//登录凭证表
认证
直接在用户服务里处理登录操作即可 ,也可提出一个auth概念，但是这个auth与user是紧密相关的 不与其他类型的user相关 
比如在 user 购买外卖这个上下文中   auth只处理user
在  商家上传商品这个上下文中       auth只处理商家
在 骑手注册这个上下文总            auth只处理骑手
如果引入了 oauth
appid 用户端 appsecret 
appid 商家端 
oatuh一般用在 用户端的auth  开放用户信息给其他使用 
比如 饿了么登录， 不会是骑手信息登录 也不会是商家把 除非商家是买家的一种属性。 这种情况auth就是一个了