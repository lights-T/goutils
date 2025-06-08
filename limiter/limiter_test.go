package limiter

import (
	"testing"
)

func Test_limiter(t *testing.T) {
	//func Route(r *gin.Engine, frontend embed.FS) error {
	//	r.Use(CorsMiddleware())
	//
	//	// 创建一个限流器：每秒允许 1 次请求
	//	limitedHandler := RateLimitMiddleware(time.Second, 1)
	//
	//	ddd := r.Group("/api")
	//{
	//	//job管理
	//	ddd.GET("/job/refreshSecJob", limitedHandler, secJob.RefreshSecJob)
	//	ddd.GET("/job/updateMouseMoveEvent", secJob.UpdateMouseMoveEvent)
	//}
	//	return nil
	//}
	t.Run("", func(t *testing.T) {
		//在router中：如对接口/job/refreshSecJob进行限制
	})
}
