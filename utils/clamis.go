package utils

//// 从Gin的Context中获取从jwt解析出来的用户信息
//func GetUser(c *gin.Context) *systemReq.UserCache {
//	if claims, exists := c.Get("claims"); !exists {
//		logger.Error("从Gin的Context中获取从jwt解析出来的用户失败, 请检查路由是否使用jwt中间件!")
//		return nil
//	} else {
//		waitUse := claims.(*systemReq.UserCache)
//		return waitUse
//	}
//}
