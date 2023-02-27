package app

//区分什么是不运行就关掉的   redis不必须， mysql 必须  ,mq（可选）必须。。等等  必须的 要监听错误的话得关闭当前整个服务以便于被观测到， 或者得发送通知，
//只会在两个地方register  wire和main
//注册的东西有service, repository , mysql
//在api wire service,或者 也是在adapter wire别的东西，

func wire() {

}
