package repository

import (
	"demoapi/internal/conf"
	"demoapi/internal/core/mysqlx"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ActivityDB struct {
	ActivityR *gorm.DB
	ActivityW *gorm.DB
}

func NewActivityDB(c *conf.Data, logger log.Logger) (*ActivityDB, func(), error) {

	_log := log.NewHelper(log.With(logger, "x_module", "data/NewActivityDB"))

	dbR_ := mysqlx.NewMysqlDB(c.ActivityRead, logger)
	dbW_ := mysqlx.NewMysqlDB(c.ActivityWrite, logger)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")

		if dbR_ != nil {
			db, _ := dbR_.DB()
			_ = db.Close()
		}
		if dbW_ != nil {
			db, _ := dbW_.DB()
			_ = db.Close()
		}

		_log.Info("closing the data resources")
	}

	return &ActivityDB{
		ActivityR: dbR_,
		ActivityW: dbW_,
	}, cleanup, nil

}
