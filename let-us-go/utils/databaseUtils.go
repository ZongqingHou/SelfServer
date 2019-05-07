package utils

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type ContextDB struct {
	*xorm.Engine
}

func New(db *xorm.Engine, service string, config kafka.Config) *ContextDB {
	db.ShowExecTime()
	if len(config.Brokers) != 0 {
		if producer, err := kafka.NewProducer(config.Brokers, config.Topic,
			kafka.WithDefault(),
			kafka.WithTLS(config.SSL)); err == nil {
			db.SetLogger(&dbLogger{serviceName: service, Producer: producer})
			db.ShowSQL()
		}
	}

	return &ContextDB{Engine: db}
}

func (db *ContextDB) NewSession(ctx context.Context) *xorm.Session {
	session := db.Engine.NewSession()

	func(session interface{}, ctx context.Context) {
		if s, ok := session.(interface{ SetContext(context.Context) }); ok {
			s.SetContext(ctx)
		}
	}(session, ctx)

	return session
}

func initDB(driver, conntection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, conntection)
	if err != nil {
		return nil, err
	}

	if err = db.Sync(new(modules.Test)); err != nil {
		panic(err)
	}
	return db, nil
}

func ContextDB(service string, xormEngine *xorm.Engine, kafkaConfig kafka.Config) echo.MiddlewareFunc {
	db := ctxdb.New(xormEngine, service)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			session := db.NewSession(ctx)
			defer session.Close()

			c.SetRequest(req.WithContext(context.WithValue(ctx, ContextDBName, session)))

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