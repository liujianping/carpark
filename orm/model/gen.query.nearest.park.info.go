package model

import (
	"database/sql"
	"fmt"
	"github.com/auto-program/db-orm/orm"
	"strings"
	"time"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type NearestParkInfo struct {
	Address       string  `db:"address" json:"address"`
	Latitude      float64 `db:"latitude" json:"latitude"`
	Longitude     float64 `db:"longitude" json:"longitude"`
	TotalLots     int32   `db:"total_lots" json:"total_lots"`
	AvailableLots int32   `db:"available_lots" json:"available_lots"`
}

type _NearestParkInfoMgr struct {
}

var NearestParkInfoMgr *_NearestParkInfoMgr

func (m *_NearestParkInfoMgr) NewNearestParkInfo() *NearestParkInfo {
	return &NearestParkInfo{}
}

type _NearestParkInfoDBMgr struct {
	db orm.DB
}

func (m *_NearestParkInfoMgr) DB(db orm.DB) *_NearestParkInfoDBMgr {
	return NearestParkInfoDBMgr(db)
}

func NearestParkInfoDBMgr(db orm.DB) *_NearestParkInfoDBMgr {
	if db == nil {
		panic(fmt.Errorf("NearestParkInfoDBMgr init need db"))
	}
	return &_NearestParkInfoDBMgr{db: db}
}

func (m *_NearestParkInfoDBMgr) QueryBySQL(q string, args ...interface{}) (results []*NearestParkInfo, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("NearestParkInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result NearestParkInfo
		err = rows.Scan(&(result.Address), &(result.Latitude), &(result.Longitude), &(result.TotalLots), &(result.AvailableLots))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("NearestParkInfo fetch result error: %v", err)
	}
	return
}
