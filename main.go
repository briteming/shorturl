/**
 * @auther:  zu1k
 * @date:    2020/2/10
 */
package main

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port string

func main() {
	go ginStart()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ExitFunc()
		}
	}
}

func ginStart() {
	gin.SetMode("release")
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/:short", jump)
	_ = r.Run(":" + port)
}

func jump(c *gin.Context) {
	short := c.Param("short")
	if strings.HasPrefix(short, "short") {
		ginShort(c)
		return
	}
	originUrl, err := get(short)
	if err != nil {
		c.String(http.StatusNotFound, "短连接不存在:"+err.Error())
	} else {
		c.Redirect(http.StatusTemporaryRedirect, originUrl)
	}
}

func ExitFunc() {
	db.Close()
	os.Exit(0)
}
