# 命名规范
 - 包名用小学字母无下划线
 - 文件名用 小写+下划线
 - 结构体用驼峰 大小写取决于权限
 - 参数名驼峰 首字母小写
 - 返回值驼峰 首字母小写
 - 单函数接口命名一般以 er结尾  Reader Writer   
 - 常量命名使用驼峰
 - 接受者使用单一字母如(t) 
  

## 指针使用规范 *  &
 - 使用原则 避免重复的值拷贝
 - 只读文件则使用结构体
 - 正常的struct 构造函数都使用指针
 - 在需要使用原对象引用的方法 声明才使用指针，不需要则用结构体，同理 传入判断是否是指针 以被传入的方法为准
 - compoent这种单例对象的构造使用指针存入ioc
## 变量类型规范
 - 使用int 而不是 int64 int32 ?
 - 不使用枚举类
## 错误使用规范以及返回规范
 - 不在服务中直接处理异常，返回到controller中处理
 - 只有在一些强依赖的服务错误时才panic
 - 在web容器api的错误都不panic，并且外层recover
 - errors.wrapper(在应用层的最后一层（产生的地方）使用包装，不在基础层使用，也不处处使用)
## 单元测试规范
 - 单元测试用 _test.go ， 函数名以Test开头
 - 单元测试写在相同的包
 - 保证覆盖率
## 引入标准
 - 【规则1.3.2】禁止使用相对路径导入（./subpackage），所有导入路径必须符合 go get 标准。
 - 【建议1.3.3】建议使用goimports工具或者IDE工具来管理多行import
 - module 以全仓库命名 module github.com/evrone/go-clean-template


更多规范 //https://www.jianshu.com/p/20861de6332c
错误处理 https://zhuanlan.zhihu.com/p/328591249


————————————————
https://blog.csdn.net/u010524722/article/details/124512369