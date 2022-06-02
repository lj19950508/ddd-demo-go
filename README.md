# 命名规范
包名用小学字母无下划线
单元测试用 _test.go ， 函数名以Test开头
文件名用 小写+下划线
结构体用驼峰 大小写取决于权限
单函数接口命名一般以 er结尾  Reader Writer   
常量命名则是 驼峰
枚举 推荐不用 使用驼峰
参数名驼峰 首字母小写 
返回值驼峰 首字母小写 
接受者不实用 me,this, self,1-2 字母，如果方法不饮用则 省略掉 _
错误处理层层返回到controler 层处理 ， recover层只是为了防止意外。一般是
if(err!=nil){
  return err
}

....


【规则1.3.2】禁止使用相对路径导入（./subpackage），所有导入路径必须符合 go get 标准。
【建议1.3.3】建议使用goimports工具或者IDE工具来管理多行import

更多规范 //https://www.jianshu.com/p/20861de6332c
