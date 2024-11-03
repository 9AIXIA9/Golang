package api

import "life/internal/controllers"

// 未定义页面路由
func setUpNoRoute() {
	r.NoRoute(controllers.NoRouteHandler)
}
