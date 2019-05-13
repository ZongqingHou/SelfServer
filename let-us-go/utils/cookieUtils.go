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

func 