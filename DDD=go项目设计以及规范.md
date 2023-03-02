基础框架(ioc/基础事务/)
domain层 不依赖于外部框架，如 gorm/等 ， 不与表对等
service层 不依赖外部框架 负责编排领域，控制事务（事务可能要采取原声事物interface）
adapter/api层 负责 接受参数，解析参数，返回参数 与 耦合  与第三方耦合
     
adapter/persistent/gorm 负责具体的查询

in为用户输入和返回给用户  如http socket
out为输出到第三方主键和返回给当前程序  如mq  


====-=-=-=
1/解决ioc的问题  uber/fx 不错不错
cqrs要不要实现呢？