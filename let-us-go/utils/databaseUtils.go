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

/*
	The definations
*/

const ContextMySQLName = "mysql"
const ContextRedisName = "redis"

type ContextDB struct {
	*xorm.Engine
}

/*
	The initilation functions
*/

func InitMySQL(driver, conntection string) (*xorm.Engine) {
	db, err := xorm.NewEngine(driver, conntection)
	if err != nil {
		panic(err)
	}

	if err = db.Sync(new(modules.Test)); err != nil {
		panic(err)
	}

	return db
}

func InitRedis(addr string, pswd string, device int) (*redis.Client){
	rd := redis.NewClient(&redis.Options{
			Addr: 		addr,
			Password: 	pswd,
			DB: 		0,
		})

	ping, err := rd.Ping().Result()
	fmt.Printf("ping result: %s\n", ping)
	if err != nil{
		panic(err)
	}

	return rd
}

/*
	The Session related functions
*/



// func NewRedisSession(redisClient *redis.Client, ctx context.Context) *redis.Session {
// 	session := redisClient.NewSession()


// }

/*
	The middleware functions
*/
func SaveSessionID(redisClient *redis.Client) echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(serverContext echo.Context) error {
			return nil
		}
	}
}


func ContextMySQL(xormEngine *xorm.Engine, redisClient *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(serverContext echo.Context) error {
			serverRequest := serverContext.Request()
			contextPtr := serverRequest.Context()

			session := NewSession(xormEngine, contextPtr)
			defer session.Close()

			serverContext.SetRequest(serverRequest.WithContext(context.WithValue(contextPtr, ContextMySQLName, session)))

			switch serverRequest.Method{
				case "POST", "PUT", "DELETE", "PATCH":
					if err := session.Begin; err != nil {
						log.Println(err)
					}
					if err := next(serverContext); err != nil {
						session.Rollback()
						return err
					}
					if serverContext.Response().Status >= 500 {
						session.Rollback()
						return nil
					}
					if err := session.Commit(); err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
					}
				default:
					return next(serverContext)
			}

			return nil
		}
	}
}