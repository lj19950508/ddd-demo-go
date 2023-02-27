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
 - 避免重复的值拷贝，所以结构体尽量使用指针类型
 - 只读文件则使用结构体
## 变量类型规范
 - 使用int 而不是 int64 int32 
 - 使用基础类型Or nulltype
 - 不使用枚举类
## 错误使用规范以及返回规范
 - 不在服务中直接panic异常，返回到controller中处理
 - errors.wrapper 只在产生的地方使用
## 单元测试规范
 - 单元测试用 _test.go ， 函数名以Test开头
 - 单元测试写在相同的包
 - 保证覆盖率
## 引入标准
 - 【规则1.3.2】禁止使用相对路径导入（./subpackage），所有导入路径必须符合 go get 标准。
 - 【建议1.3.3】建议使用goimports工具或者IDE工具来管理多行import


//泛型
https://taoshu.in/go/generics/design.html
更多规范 //https://www.jianshu.com/p/20861de6332c
错误处理 https://zhuanlan.zhihu.com/p/328591249

