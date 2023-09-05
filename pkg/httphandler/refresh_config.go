package httphandler

import (
	"net/http"
)

// TODO(cat): 修改成启动一个协程定时刷新配置, 不要通过 Http 实现了, 而且基于 atomic.Value 实现从而不需要加锁
func refreshConfig(w http.ResponseWriter, r *http.Request) {
	// refreshResult := map[string]interface{}{
	// 	"is_success": true,
	// 	"error_code": 0,
	// 	"error_msg":  "",
	// }
}

func muxRefreshConfig(m *http.ServeMux) {
	m.HandleFunc("/refresh-config", refreshConfig)
}
