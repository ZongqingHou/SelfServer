package controllers

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"../utils"

	"github.com/labstack/echo"
)

const WECHAT_TOKEN = "hdDA1lab"
const APPID = "wx64792e91e7a194e2"
const APPSECRET = "3c394068fafc3d78e3d654a5301a75ae"

type WechatController struct {
}

type Image_Message struct {
	MediaId string `xml: "MediaId"`
}

type xml struct {
	ToUserName   string         `xml: "ToUserName"`
	FromUserName string         `xml: "FromUserName"`
	CreateTime   int            `xml: "CreateTime"`
	MsgType      string         `xml: "MsgType"`
	Content      string         `xml: "Content"`
	Image        *Image_Message `xml: "Image"`
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
			message_recieve := new(xml)

			if err := context.Bind(message_recieve); err != nil {
				return err
			}

			if message_recieve.MsgType == "text" {
				image_field := &Image_Message{
					MediaId: "JvzoATjTVQn0kj58AwVwXqTY4Izh0Ejmhe1d2-mP3IITZh6FBtKNZEshxJ3klvWE",
				}

				message_response := &xml{
					ToUserName:   message_recieve.FromUserName,
					FromUserName: message_recieve.ToUserName,
					CreateTime:   int(time.Now().Unix()),
					MsgType:      message_recieve.MsgType,
					Image:        image_field,
				}

				return context.XML(http.StatusCreated, message_response)
			}
		}
	}

	return context.String(http.StatusForbidden, "Invalid Source")
}
