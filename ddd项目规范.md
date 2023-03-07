### DDD项目规范 
不同子域调用要通过eventbus,这里由于单体架构比较方便，先把子域区分提到文件夹内，如果是作为微服务则是按把子域提到项目名  如adapter/in/api/user ->  user/adapter/in/api
在application和domain层不会出现实现操作，都只是体现流程和算法，具体操作交给out/in  
Command可以有结果值但是要写在CommandDto 里面 一个ExcuteResult