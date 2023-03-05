基础框架(ioc/基础事务/)

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

---------------------





