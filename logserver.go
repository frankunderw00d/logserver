package main

import (
	"io"
	"jarvis/base/log"
	"jarvis/base/network"
	logModule "logserver/module/log"
	"os"
	"os/signal"
	"syscall"
)

const (
	// Socket 监听地址
	SocketListenAddress = ":10000"
	// WebSocket 监听地址
	WebSocketListenAddress = ":10001"
	// gRPC 监听地址
	GRPCListenAddress = ":10002"
)

var (
	service network.Service
	fh      io.WriteCloser
)

func init() {
	// 设置文件输出
	log.SetFlag(log.FlagNone)
	nfh := log.NewDefaultFileHook()
	fh = nfh
	log.SetHook(fh)

	// 初始化服务
	service = network.NewService(
		network.DefaultMaxConnection,
		network.DefaultIntoStreamSize,
	)
}

func main() {
	// 1.注册模块
	err := service.RegisterModule(logModule.NewModule())
	if err != nil {
		log.ErrorF("Register module error : %s", err.Error())
		return
	}

	// 2.启动
	err = service.Run(
		network.NewSocketGate(SocketListenAddress),
		network.NewWebSocketGate(WebSocketListenAddress),
		network.NewGRPCGate(GRPCListenAddress),
	)
	if err != nil {
		log.ErrorF("Run error : %s", err.Error())
		return
	}

	// 3.监听退出信号
	monitorSystemSignal()
}

// 监听系统信号
// kill -SIGQUIT [进程号] : 杀死当前进程
func monitorSystemSignal() {
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGQUIT)
	select {
	case <-sc:
		log.InfoF("Done")
		_ = fh.Close()
	}
}
