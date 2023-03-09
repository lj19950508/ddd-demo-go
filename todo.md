1.evenbus的接口以及  实现
2.提出httphandler
7.自定义validator 翻译器 
8.使用httpmux 替代 httphandler 实现多路复用


//1.优化grpc代码 使其独立8081
//2.优化客户端grpc代码 使其默认连接
//3.grpc转为eventbus
//4.新增异步事件中心 （发送日志）

asnycEventBus
eventbus

cd $WORK
go work init
go work use ./tools/ ./tools/gopls/