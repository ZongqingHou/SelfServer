package utils

import (
	"fmt"
	"log"
	"context"
	"net/http"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql" // mysql
	// "github.com/Shopify/sarama"		   // kafka
	"github.com/go-redis/redis"		   // redis
	"github.com/go-xorm/xorm"		   // engine

	"../modules"

)

func NewSession(xormEngine *xorm.Engine, ctx context.Context) *xorm.Session {
	session := xormEngine.NewSession()

	func(session interface{}, ctx context.Context) {
		if tmpSession, ok := session.(interface{ SetContext(context.Context) }); ok {
			tmpSession.SetContext(ctx)
		}
	}(session, ctx)

	return session
}