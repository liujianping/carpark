package model

//! conf.mysql
import (
	"sync"
	"time"

	"github.com/auto-program/db-orm/orm"
)

var (
	_mysql_store *orm.DBStore
	_mysql_cfg   MySQLConfig
	_mysql_once  sync.Once
)

type MySQLConfig struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	Database        string
	PoolSize        int
	ConnMaxLifeTime time.Duration
}

func MySQLSetup(cf *MySQLConfig) {
	_mysql_cfg = *cf
}

func MySQL() orm.DB {
	var err error
	_mysql_once.Do(func() {
		_mysql_store, err = orm.NewDBStore("mysql",
			_mysql_cfg.Host,
			_mysql_cfg.Port,
			_mysql_cfg.Database,
			_mysql_cfg.UserName,
			_mysql_cfg.Password)
		if err != nil {
			panic(err)
		}
		_mysql_store.SetConnMaxLifetime(time.Hour)
		if _mysql_cfg.ConnMaxLifeTime > 0 {
			_mysql_store.SetConnMaxLifetime(_mysql_cfg.ConnMaxLifeTime)
		}
		_mysql_store.SetMaxIdleConns(_mysql_cfg.PoolSize)
		_mysql_store.SetMaxOpenConns(_mysql_cfg.PoolSize)
	})
	return _mysql_store
}
