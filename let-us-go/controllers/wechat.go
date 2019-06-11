package controllers

import (
	"net/http"
	"sort"
	"strings"

	"../utils"

	"github.com/labstack/echo"
)

const WECHAT_TOKEN = "hdDA1lab"

type WechatController struct {
}

func (controller WechatController) Init(group *echo.Group) {
	group.GET("", controller.Index)
}

func (WechatController) Index(context echo.Context) error {
	signature := context.QueryParam("signature")
	timestamp := context.QueryParam("timestamp")
	nonce := context.QueryParam("nonce")
	echostr := context.QueryParam("echostr")

	// println(context.QueryParams())
	// if context.QueryParams() == nil {
	// 	return nil
	// }
	string_collection := []string{WECHAT_TOKEN, timestamp, nonce}
	sort.Strings(string_collection)

	cmp_string := strings.Join(string_collection, "")
	cmp_string = utils.Sha1(cmp_string)

	if cmp_string == signature {
		return context.String(http.StatusOK, echostr)
	}
	println(cmp_string)
	println(WECHAT_TOKEN)
	println(signature)
	println(timestamp)
	println(nonce)
	println(echostr)

	return context.String(http.StatusOK, "Hello, World!")
}
