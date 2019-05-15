package utils

import (
	"fmt"

	// kafka
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"        // redis
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/go-xorm/xorm"          // engine

	"../modules"
)

/*
	The definations
*/
const ContextMySQLName = "mysql"
const ContextRedisName = "redis"

var Address = []string{"0.0.0.0:9092"}

type MessageProcess struct {
	mysqlEngine    *xorm.Engine
	redisClient    *redis.Client
	kafkaAProducer *sarama.AsyncProducer
	kafkaSProducer *sarama.SyncProducer
}

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

func MessageClose(message *MessageProcess) {
	if message.mysqlEngine != nil {
		defer message.mysqlEngine.Close()
	}

	if message.redisClient != nil {
		defer message.redisClient.Close()
	}

	if message.kafkaAProducer != nil {
		defer (*message.kafkaAProducer).Close()
	}

	if message.kafkaSProducer != nil {
		defer (*message.kafkaSProducer).Close()
	}
}

func InitMessageProcesser() *MessageProcess {
	db := InitMySQL("mysql", "root:123456@tcp(127.0.0.1:3306)/hdd?charset=utf8")
	rd := InitRedis("localhost:6379", "", 0)
	producer := InitKafkaAsyncProducer(Address)

	return &MessageProcess{
		mysqlEngine:    db,
		redisClient:    rd,
		kafkaAProducer: producer,
		kafkaSProducer: nil,
	}
}
