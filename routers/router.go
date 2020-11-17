// @author  dreamlu
package routers

import (
	"dnovel/controllers"
	"dnovel/services"
	str2 "dnovel/util/cons"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/util/cons"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

const prefix = ""

var Router = SetRouter()

var V = Router.Group(prefix)

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//router := gin.Default()
	router := gin.New()
	// router.MaxMultipartMemory = // 默认32M
	//router.Use(CorsMiddleware())

	// 过滤器
	router.Use(Filter())
	router.Use(Recover)
	//权限中间件
	// load the casbin model and policy from files, database is also supported.
	//e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	//router.Use(authz.NewAuthorizer(e))

	//cookie session
	//store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	//redis session
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//组的路由,version
	v1 := router.Group(prefix)
	{
		v := v1

		// 静态目录
		// relativePath:请求路径
		// root:静态文件所在目录
		v.Static("static", "static")
		// v.GET("/statics/file", file.StaticFile)
		inc := controllers.IndexController{Service: services.NewBookService()}
		v.GET("/search", inc.GetSearch)
		v.GET("/info", inc.GetInfo)
		v.GET("/chapters", inc.GetChapters)
		v.GET("/read", inc.GetRead)
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":    "接口不存在->('.')/请求方法不存在",
		})
	})
	return router
}

// 登录失效验证
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 127请求且本地开发且为dev时无需验证,方便自己测试
		if strings.Contains(c.Request.RemoteAddr, "127.0.0.1") &&
			str2.DevMode == cons.Dev {
			c.Next()
			return
		}

		if c.Request.Method == "GET" {
			c.Next()
			return
		}
		path := c.Request.URL.String()

		if !strings.Contains(path, "login") && !strings.Contains(path, "/static/file") {
			r := c.Request
			token := r.Header.Get("token")
			if token == "" {
				c.Abort()
				c.JSON(http.StatusOK, result.GetError("缺少token"))
				return
			}
			ca := cache.NewCache()
			log.Println("[token]:", token)
			cam, err := ca.Get(token)
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, result.MapNoAuth)
				return
			}
			// 延长token对应时间
			_ = ca.Set(token, cam)
		}
	}
}

// 异常捕获
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			ss := strings.Split(string(debug.Stack()), "\n\t")
			res := make(map[string]string)
			for _, v := range ss {
				ks := strings.Split(v, "\n")
				res[ks[0]] = ks[1]
			}
			c.JSON(http.StatusOK, result.GetError(res))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// 处理跨域请求,支持options访问
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		//fmt.Println(method)
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}
