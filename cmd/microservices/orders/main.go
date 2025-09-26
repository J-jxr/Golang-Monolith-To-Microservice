package main

import (
	"log"
	"os"
	"fmt"
	"net/http"

)

func main() {
	log.Println("Starting the orders microservice...")

	// 创建上下文对象
	// 推测cmd.Context()是一个自定义函数，返回一个可被取消的context
	// 作用：监听系统终止信号（如Ctrl+C、kill命令），当收到信号时触发"取消"
	ctx := cmd.Context()

	// 返回一个路由器和一个关闭函数
	// r：HTTP请求路由器，用于注册接口路由和处理函数
	r, closeFu := createOrdersMicroservice()

	defer closeFu()

	// 配置HTTP服务器
	server := &http.Server{
		// 监听地址从环境变量中获取
		Addr:    os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"),
		// 路由器作为处理器
		Handler: r,
	}

	// 在单独的goroutine中启动HTTP服务器（避免阻塞主线程）
	go func() {
		// 启动服务器，开始监听并处理HTTP请求
		// ListenAndServe()是阻塞调用，直到服务器关闭才会返回
		if err := server.ListenAndServe(); err != http.ErrServerClosed  {
			// 如果返回的错误不是"服务器已关闭"（正常关闭），则触发panic
			panic(err)
		}

	}()

	// 【逻辑问题】由于上面的ListenAndServe()是阻塞的，这行代码永远不会执行
	// 意图：等待上下文被取消（即收到终止信号）
	<-ctx.Done()

	log.Println("closing the orders service ...")

	if err := server.Close(); err != nil {
		panic(err)
	}

}

func createOrdersMicroservice()(router *chi.Mux, closeFunc func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.CreateRouter()

	orders_public_http.AddRouter(r, ordersService, ordersRepo)
	orders_private_http.AddRouter(r, ordersService, ordersRepo)

	return r, func(){
		
	}
}
