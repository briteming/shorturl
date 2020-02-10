/**
 * @auther:  zu1k
 * @date:    2020/2/10
 */
package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	domain = "zll.us/"
	length = 4
)

func ginShort(c *gin.Context) {
	originUrl := c.Query("url")
	if originUrl == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "param url not found",
		})
	} else if !strings.HasPrefix(originUrl, "http") {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "url should start with http",
			"url":  originUrl,
		})
	} else {
		shorted := short(originUrl)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"url":  domain + shorted,
		})
	}
}

func short(urlOrigin string) (shorted string) {
	shorted, err := get(urlOrigin)
	if err == nil { //之前设置过
		return
	}
	shorted = RandString(length)
	_, err = get(shorted)
	for err == nil { //已经使用了
		shorted = RandString(length)
		_, err = get(shorted)
	}
	set(shorted, urlOrigin)
	set(urlOrigin, shorted)
	return
}
