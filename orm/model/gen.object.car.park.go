package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/auto-program/db-orm/orm"
	"gopkg.in/go-playground/validator.v9"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
	_ validator.Validate
)

type CarPark struct {
	CarParkNo            string  `db:"car_park_no"`
	Address              string  `db:"address"`
	Latitude             float64 `db:"latitude"`
	Longitude            float64 `db:"longitude"`
	CarParkType          string  `db:"car_park_type"`
	TypeOfParkingSystem  string  `db:"type_of_parking_system"`
	ShortTermParking     string  `db:"short_term_parking"`
	ShortTermParkingFrom int64   `db:"short_term_parking_from"`
	ShortTermParkingTo   int64   `db:"short_term_parking_to"`
	FreeParking          string  `db:"free_parking"`
	NightParking         bool    `db:"night_parking"`
	CarParkDecks         int32   `db:"car_park_decks"`
	GantryHeight         float64 `db:"gantry_height"`
	CarParkBasement      bool    `db:"car_park_basement"`
}

var CarParkColumns = struct {
	CarParkNo            string
	Address              string
	Latitude             string
	Longitude            string
	CarParkType          string
	TypeOfParkingSystem  string
	ShortTermParking     string
	ShortTermParkingFrom string
	ShortTermParkingTo   string
	FreeParking          string
	NightParking         string
	CarParkDecks         string
	GantryHeight         string
	CarParkBasement      string
}{
	"car_park_no",
	"address",
	"latitude",
	"longitude",
	"car_park_type",
	"type_of_parking_system",
	"short_term_parking",
	"short_term_parking_from",
	"short_term_parking_to",
	"free_parking",
	"night_parking",
	"car_park_decks",
	"gantry_height",
	"car_park_basement",
}

type _CarParkMgr struct {
}

var CarParkMgr *_CarParkMgr

func (m *_CarParkMgr) NewCarPark() *CarPark {
	return &CarPark{}
}

//! object function

func (obj *CarPark) GetNameSpace() string {
	return "model"
}

func (obj *CarPark) GetClassName() string {
	return "CarPark"
}

func (obj *CarPark) GetTableName() string {
	return "carparks"
}

func (obj *CarPark) GetColumns() []string {
	columns := []string{
		"`car_park_no`",
		"`address`",
		"`latitude`",
		"`longitude`",
		"`car_park_type`",
		"`type_of_parking_system`",
		"`short_term_parking`",
		"`short_term_parking_from`",
		"`short_term_parking_to`",
		"`free_parking`",
		"`night_parking`",
		"`car_park_decks`",
		"`gantry_height`",
		"`car_park_basement`",
	}
	return columns
}

func (obj *CarPark) GetNoneIncrementColumns() []string {
	columns := []string{
		"`car_park_no`",
		"`address`",
		"`latitude`",
		"`longitude`",
		"`car_park_type`",
		"`type_of_parking_system`",
		"`short_term_parking`",
		"`short_term_parking_from`",
		"`short_term_parking_to`",
		"`free_parking`",
		"`night_parking`",
		"`car_park_decks`",
		"`gantry_height`",
		"`car_park_basement`",
	}
	return columns
}

func (obj *CarPark) GetPrimaryKey() PrimaryKey {
	pk := CarParkMgr.NewPrimaryKey()
	pk.CarParkNo = obj.CarParkNo
	return pk
}

func (obj *CarPark) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type CarParkNoOfCarParkPK struct {
	CarParkNo string
}

func (m *_CarParkMgr) NewPrimaryKey() *CarParkNoOfCarParkPK {
	return &CarParkNoOfCarParkPK{}
}

func (u *CarParkNoOfCarParkPK) Key() string {
	strs := []string{
		"CarParkNo",
		fmt.Sprint(u.CarParkNo),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *CarParkNoOfCarParkPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vCarParkNo, ok := kv["CarParkNo"]
	if !ok {
		return fmt.Errorf("key (%s) without (CarParkNo) field", key)
	}
	if err := orm.StringScan(vCarParkNo, &(u.CarParkNo)); err != nil {
		return err
	}
	return nil
}

func (u *CarParkNoOfCarParkPK) SQLFormat() string {
	conditions := []string{
		"`car_park_no` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *CarParkNoOfCarParkPK) SQLParams() []interface{} {
	return []interface{}{
		u.CarParkNo,
	}
}

func (u *CarParkNoOfCarParkPK) Columns() []string {
	return []string{
		"`car_park_no`",
	}
}

//! uniques

//! indexes

type ShortTermParkingFromShortTermParkingToOfCarParkIDX struct {
	ShortTermParkingFrom int64
	ShortTermParkingTo   int64
	offset               int
	limit                int
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) Key() string {
	strs := []string{
		"ShortTermParkingFrom",
		fmt.Sprint(u.ShortTermParkingFrom),
		"ShortTermParkingTo",
		fmt.Sprint(u.ShortTermParkingTo),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`short_term_parking_from` = ?",
		"`short_term_parking_to` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) SQLParams() []interface{} {
	return []interface{}{
		u.ShortTermParkingFrom,
		u.ShortTermParkingTo,
	}
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) Limit(n int) {
	u.limit = n
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) Offset(n int) {
	u.offset = n
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *ShortTermParkingFromShortTermParkingToOfCarParkIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type _CarParkDBMgr struct {
	db orm.DB
}

func (m *_CarParkMgr) DB(db orm.DB) *_CarParkDBMgr {
	return CarParkDBMgr(db)
}

func CarParkDBMgr(db orm.DB) *_CarParkDBMgr {
	if db == nil {
		panic(fmt.Errorf("CarParkDBMgr init need db"))
	}
	return &_CarParkDBMgr{db: db}
}

func (m *_CarParkDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*CarPark, error) {
	obj := CarParkMgr.NewCarPark()
	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*CarPark, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarPark))
	}
	return results, nil
}

func (m *_CarParkDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*CarPark, error) {
	obj := CarParkMgr.NewCarPark()
	q := fmt.Sprintf("SELECT %s FROM carparks %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	objs, err := m.FetchBySQL(q, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*CarPark, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarPark))
	}
	return results, nil
}

func (m *_CarParkDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_CarParkDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_CarParkDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("CarPark fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result CarPark
		err = rows.Scan(&(result.CarParkNo), &(result.Address), &(result.Latitude), &(result.Longitude), &(result.CarParkType), &(result.TypeOfParkingSystem), &(result.ShortTermParking), &(result.ShortTermParkingFrom), &(result.ShortTermParkingTo), &(result.FreeParking), &(result.NightParking), &(result.CarParkDecks), &(result.GantryHeight), &(result.CarParkBasement))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("CarPark fetch result error: %v", err)
	}
	return
}
func (m *_CarParkDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

func (m *_CarParkDBMgr) Fetch(pk PrimaryKey) (*CarPark, error) {
	obj := CarParkMgr.NewCarPark()
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*CarPark), nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkDBMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*CarPark, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*CarParkNoOfCarParkPK).CarParkNo))
	}
	obj := CarParkMgr.NewCarPark()
	query := fmt.Sprintf("SELECT %s FROM carparks WHERE `car_park_no` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*CarPark, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarPark))
	}
	return results, nil
}

func (m *_CarParkDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkDBMgr) FindOneFetch(unique Unique) (*CarPark, error) {
	obj := CarParkMgr.NewCarPark()
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*CarPark), nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_CarParkDBMgr) FindFetch(index Index) (int64, []*CarPark, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := CarParkMgr.NewCarPark()
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	results := make([]*CarPark, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarPark))
	}
	return total, results, nil
}

func (m *_CarParkDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_CarParkDBMgr) RangeFetch(scope Range) (int64, []*CarPark, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := CarParkMgr.NewCarPark()
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	results := make([]*CarPark, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarPark))
	}
	return total, results, nil
}

func (m *_CarParkDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_CarParkDBMgr) RangeRevertFetch(scope Range) (int64, []*CarPark, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_CarParkDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := CarParkMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM carparks %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("CarPark query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := CarParkMgr.NewPrimaryKey()
		err = rows.Scan(&(result.CarParkNo))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("CarPark query limit result error: %v", err)
	}
	return
}

func (m *_CarParkDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`car_park_no`) FROM carparks %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("CarPark query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			m.db.SetError(err)
			return 0, err
		}
		break
	}
	return count, nil
}

func (m *_CarParkDBMgr) BatchCreate(objs []*CarPark) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*14)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(14, "?"), ",")))
		values = append(values, obj.CarParkNo)
		values = append(values, obj.Address)
		values = append(values, obj.Latitude)
		values = append(values, obj.Longitude)
		values = append(values, obj.CarParkType)
		values = append(values, obj.TypeOfParkingSystem)
		values = append(values, obj.ShortTermParking)
		values = append(values, obj.ShortTermParkingFrom)
		values = append(values, obj.ShortTermParkingTo)
		values = append(values, obj.FreeParking)
		values = append(values, obj.NightParking)
		values = append(values, obj.CarParkDecks)
		values = append(values, obj.GantryHeight)
		values = append(values, obj.CarParkBasement)
	}
	query := fmt.Sprintf("INSERT INTO carparks(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (m *_CarParkDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE carparks SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE carparks SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkDBMgr) Create(obj *CarPark) (int64, error) {
	params := orm.NewStringSlice(14, "?")
	q := fmt.Sprintf("INSERT INTO carparks(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 14)
	values = append(values, obj.CarParkNo)
	values = append(values, obj.Address)
	values = append(values, obj.Latitude)
	values = append(values, obj.Longitude)
	values = append(values, obj.CarParkType)
	values = append(values, obj.TypeOfParkingSystem)
	values = append(values, obj.ShortTermParking)
	values = append(values, obj.ShortTermParkingFrom)
	values = append(values, obj.ShortTermParkingTo)
	values = append(values, obj.FreeParking)
	values = append(values, obj.NightParking)
	values = append(values, obj.CarParkDecks)
	values = append(values, obj.GantryHeight)
	values = append(values, obj.CarParkBasement)
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkDBMgr) Update(obj *CarPark) (int64, error) {
	columns := []string{
		"`address` = ?",
		"`latitude` = ?",
		"`longitude` = ?",
		"`car_park_type` = ?",
		"`type_of_parking_system` = ?",
		"`short_term_parking` = ?",
		"`short_term_parking_from` = ?",
		"`short_term_parking_to` = ?",
		"`free_parking` = ?",
		"`night_parking` = ?",
		"`car_park_decks` = ?",
		"`gantry_height` = ?",
		"`car_park_basement` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE carparks SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 14-1)
	values = append(values, obj.Address)
	values = append(values, obj.Latitude)
	values = append(values, obj.Longitude)
	values = append(values, obj.CarParkType)
	values = append(values, obj.TypeOfParkingSystem)
	values = append(values, obj.ShortTermParking)
	values = append(values, obj.ShortTermParkingFrom)
	values = append(values, obj.ShortTermParkingTo)
	values = append(values, obj.FreeParking)
	values = append(values, obj.NightParking)
	values = append(values, obj.CarParkDecks)
	values = append(values, obj.GantryHeight)
	values = append(values, obj.CarParkBasement)
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkDBMgr) Save(obj *CarPark) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_CarParkDBMgr) Delete(obj *CarPark) (int64, error) {
	pk := obj.GetPrimaryKey()
	return m.DeleteByPrimaryKey(pk)
}

func (m *_CarParkDBMgr) DeleteByPrimaryKey(pk PrimaryKey) (int64, error) {
	q := fmt.Sprintf("DELETE FROM carparks %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM carparks")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM carparks WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
