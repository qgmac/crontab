package master

import (
	"net"
	"net/http"
	"time"
)

var(
	//单例对象
	G_apiServer *ApiServer
)

//任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

//保存任务接口
func handleJobSave(response http.ResponseWriter,r *http.Request){
	//任务保存到ETCD中
}

//初始化服务
func InitApiServer()(err error	) {
	var(
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
	)
	//配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/savd",handleJobSave)

	//启动tcp监听
	if listener ,err = net.Listen("tcp",":"+string(G_config.ApiPort));err!=nil {
		return
	}

	//创建一个HTTP服务

	httpServer = &http.Server{
		ReadTimeout:time.Duration(G_config.ApiReadTiemout)*time.Millisecond,
		WriteTimeout:time.Duration(G_config.ApiWriteTimeout)*time.Millisecond,
		Handler:mux,
	}

	G_apiServer = &ApiServer{
		httpServer:httpServer,
	}

	//启动服务端
	go httpServer.Serve(listener)

	return
}
