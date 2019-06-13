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
	group.Any("", controller.Index)
}

func (WechatController) Index(context echo.Context) error {
	signature := context.QueryParam("signature")
	timestamp := context.QueryParam("timestamp")
	nonce := context.QueryParam("nonce")

	string_collection := []string{WECHAT_TOKEN, timestamp, nonce}
	sort.Strings(string_collection)

	cmp_string := strings.Join(string_collection, "")
	cmp_string = utils.Sha1(cmp_string)

	if cmp_string == signature {
		if context.Request().Method == "GET" {
			echostr := context.QueryParam("echostr")
			return context.String(http.StatusOK, echostr)
		} else if context.Request().Method == "POST" {
			println(context)
		}
	}

	return context.String(http.StatusForbidden, "Invalid Source")
}
