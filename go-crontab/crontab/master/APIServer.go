package master

import (
	"net"
	"net/http"
	"time"
)

// APIServer 任务的HTTP接口
type APIServer struct {
	httpServer *http.Server
}

var (
	// GAPIServer 单例对象
	GAPIServer *APIServer
)

// 保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

// InitAPIServer 初始化服务
func InitAPIServer() (err error) {
	// 配置路由
	mux := http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动TCP监听
	listener, err := net.Listen("tcp", "8070")
	if err != nil {
		return err
	}

	// 创建一个HTTP服务
	httpServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}

	GAPIServer = &APIServer{
		httpServer: httpServer,
	}

	go httpServer.Serve(listener)

	return
}
