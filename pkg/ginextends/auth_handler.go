package ginextends

import (
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

var TokenWithPermissionHandler = func(ctx *gin.Context) {
	//ctx.GetPermission()
	// if( admin && ctx.GetPermission 包含 当前资源则可以访问)
	//admin PERMISSION has->ctx.url
	//userapi notpermission   =》ctx.GET(userid)
	//user 
	// tokenString := ctx.GetHeader("Authorization")
	// // Bearer
	// token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

	// 	//需要用私钥解密
	// 	// since we only use the one private key to sign the tokens,
	// 	// we also only use its public counter part to verify
	// 	return "", nil
	// })
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, nil)
	// }
	// claim, ok := token.Claims.(jwt.MapClaims)
	// // a:=claim["userId"]

	// //adminId本质是一种数据权限的行为
	// adminId := claim["adminId"]
	// if adminId == nil {
	// 	ctx.JSON(http.StatusUnauthorized, nil)
	// }
	// ctx.Set("currentAdminId", adminId)
	// ctx.GetInt()

	//TODO 权限问题
	//1. admin > user >noauth   接口是都可以通用 管理员可以访问到 与用户无强关联的接口 （获取用户信息） ， （不能获取当前用户信息）
	// 1.管理员调用 获取当前用户信息（可以报错） userId=nil error
	// 1.管理员调用 用户信息（没问题）
	// 1.用户调用获取全部用户 (按需)
	// 1.管理员调用获取订单详情BYid，没问题可以直接访问，  用户调用获取订单详情BYID，不行得验证是自身订单
	// admin -> api /admin -> applicationSevice  根据ID获取订单  不用权限
	// user  -> api /v1  -> applicationService    根据ID获取订单  要判断是当前用户的订单 在applicationService判断 所以有时候api不能复用 service也不能复用
	// 但是根据ID获取文章这种  admin,user都可以（）
	// 所以 admin > login > noauth
	// admin|user > login >noauth   admin专属接口 ， user专属接口 > 登录通用接口, 未登录接口
	//CRUD 都是admin的特权 可能会不会从路由上分也不错？ /admin/v1/users/1
	//admin 如果想调用通用接口  aritcle 上面的通用就不得劲了  根据requirelevel吧
	// 2> 1 

	//数据权限呢
	//Find的时候传入列条件过滤即可

	//2. 注入ctx让用户自己业务处理。

	//PERMISSION  GET /usr/set/sdjfkl
	// b:=claim["permission"]
	// if !ok {
	// 	ctx.JSON(http.StatusUnauthorized, nil)
	// }
	// // claim.GetSubject()
	// claim.GetIssuer()
	// claim.GetAudience() 受众
	// c
	// claims := token.Claims.(*CustomClaimsExample)
	// fmt.Println(claims.CustomerInfo.Name)

}

var TokenHandler = func(ctx *gin.Context) {
}
