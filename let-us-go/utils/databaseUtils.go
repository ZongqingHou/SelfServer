package utils

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	// "github.com/Shopify/sarama"		   // kafka
	"github.com/go-redis/redis" // redis
	"github.com/go-xorm/xorm"   // engine

	"../modules"
)

/*
	The definations
*/
const ContextMySQLName = "mysql"
const ContextRedisName = "redis"

/*
	The initilation functions
*/

func InitMySQL(driver, conntection string) *xorm.Engine {
	db, err := xorm.NewEngine(driver, conntection)
	if err != nil {
		panic(err)
	}

	if err = db.Sync(new(modules.Test)); err != nil {
		panic(err)
	}

	return db
}

func InitRedis(addr string, pswd string, device int) *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pswd,
		DB:       0,
	})

	ping, err := rd.Ping().Result()
	fmt.Printf("ping result: %s\n", ping)
	if err != nil {
		panic(err)
	}

	return rd
}
