package utils

import (
	"context"

	"github.com/labstack/echo"
)

const SESSION = "session"
const MESSAGPROCESSER = "messageprocesser"

// type Session struct {
// 	Driver *redis.Client
// 	Name   string
// 	TTL    int64
// }

// type SessionInformation struct {
// 	SessionID  string    `json:"sessionId" bson:"sessionId"`
// 	CreateTime time.Time `json:"-" bson:"createTime"`
// 	UpdateTime time.Time `json:"-" bson:"updateTime"`
// 	Expires    time.Time `json:"-" bson:"expires"`
// 	Locale     string    `json:"-" bson:"locale"`
// }

// func (this *Session) NewSession(ctx context.Context) *SessionInformation {
// 	// dirver := this.Driver
// 	// session := redisClient.

// 	// func(session interface{}, ctx context.Context) {
// 	// 	if tmpSession, ok := session.(interface{ SetContext(context.Context) }); ok {
// 	// 		tmpSession.SetContext(ctx)
// 	// 	}
// 	// }(session, ctx)

// 	return &SessionInformation{}
// }

/*
	The middleware functions
*/
func ContextSession(message *MessageProcess) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(serverContext echo.Context) error {
			serverRequest := serverContext.Request()
			contextPtr := serverRequest.Context()

			serverContext.SetRequest(serverRequest.WithContext(context.WithValue(contextPtr, MESSAGPROCESSER, message)))

			sessionid, err := readCookie(serverContext)
			if err != nil {
				return next(serverContext)
			}

			if message.redisClient.Exists(sessionid).Val() == 1 {
				return next(serverContext)
			}

			serverContext.SetRequest(serverRequest.WithContext(context.WithValue(contextPtr, SESSION, sessionid)))

			switch serverRequest.Method {
			case "POST", "PUT", "DELETE", "PATCH":
				if err := next(serverContext); err != nil {
					return err
				}
				if serverContext.Response().Status >= 500 {
					return nil
				}
			default:
				return next(serverContext)
			}

			return nil
		}
	}
}
