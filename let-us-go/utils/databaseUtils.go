package utils

import (
	"fmt"
	// "context"

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

// func (db *ContextDB) NewSession(ctx context.Context) *xorm.Session {
// 	session := db.Engine.NewSession()

// 	func(session interface{}, ctx context.Context) {
// 		if tmpSession, ok := session.(interface{ SetContext(context.Context) }); ok {
// 			tmpSession.SetContext(ctx)
// 		}
// 	}(session, ctx)

// 	return session
// }

/*
	The middleware functions
*/

func ContextDB(xormEngine *xorm.Engine, redisClient *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(serverContext echo.Context) error {
			serverRequest := serverContext.Request()
			contextPtr := serverContext.Context()

			serverRequest.SetRequest(serverRequest.WithContext(context.WithValue(serverContext, ContextDBName)))

			switch req.Method {
			case "POST", "PUT", "DELETE", "PATCH":
				if err := session.Begin(); err != nil {
					log.Println(err)
				}
				if err := next(c); err != nil {
					session.Rollback()
					return err
				}
				if c.Response().Status >= 500 {
					session.Rollback()
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			default:
				return next(c)
			}

			return nil
		}
	}
}