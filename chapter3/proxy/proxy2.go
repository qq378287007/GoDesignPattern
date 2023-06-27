package main

import "fmt"

//定义主体服务器接口
type Server interface {
	HandleRequest(string, string) (int, string)
}

//定义真实主体类
type Application struct {
}

//处理请求
func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/user/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/user/login" && method == "POST" {
		return 201, "User Login"
	}
	return 404, "Not Ok"
}

//Apache类
type Apache struct {
	Application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

//创建Apache服务器
func NewApacheServer() *Apache {
	return &Apache{
		Application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

//检查频率限制
func (n *Apache) CheckRateLimiting(url string) bool {
	n.rateLimiter[url] += 1
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	return true
}

//处理请求
func (n *Apache) HandleRequest(url, method string) (int, string) {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.Application.HandleRequest(url, method)
}

func main() {
	//初始化Apache服务器
	ApacheServer := NewApacheServer()

	userStatusURL := "/user/status"
	//发送一个GET请求
	httpCode, body := ApacheServer.HandleRequest(userStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//发送一个POST请求
	httpCode, body = ApacheServer.HandleRequest(userStatusURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	userLoginURL := "/user/login"
	//发送一个GET请求
	httpCode, body = ApacheServer.HandleRequest(userLoginURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//发送一个POST请求
	httpCode, body = ApacheServer.HandleRequest(userLoginURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//多次请求
	httpCode, body = ApacheServer.HandleRequest(userLoginURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)
}
