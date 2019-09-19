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

type CarParkStatus struct {
	CarParkNo     string `db:"car_park_no"`
	TotalLots     int32  `db:"total_lots"`
	AvailableLots int32  `db:"available_lots"`
	ReportAt      int64  `db:"report_at"`
	CreatedAt     int64  `db:"created_at"`
}

var CarParkStatusColumns = struct {
	CarParkNo     string
	TotalLots     string
	AvailableLots string
	ReportAt      string
	CreatedAt     string
}{
	"car_park_no",
	"total_lots",
	"available_lots",
	"report_at",
	"created_at",
}

type _CarParkStatusMgr struct {
}

var CarParkStatusMgr *_CarParkStatusMgr

func (m *_CarParkStatusMgr) NewCarParkStatus() *CarParkStatus {
	return &CarParkStatus{}
}

//! object function

func (obj *CarParkStatus) GetNameSpace() string {
	return "model"
}

func (obj *CarParkStatus) GetClassName() string {
	return "CarParkStatus"
}

func (obj *CarParkStatus) GetTableName() string {
	return "carpark_status"
}

func (obj *CarParkStatus) GetColumns() []string {
	columns := []string{
		"`car_park_no`",
		"`total_lots`",
		"`available_lots`",
		"`report_at`",
		"`created_at`",
	}
	return columns
}

func (obj *CarParkStatus) GetNoneIncrementColumns() []string {
	columns := []string{
		"`car_park_no`",
		"`total_lots`",
		"`available_lots`",
		"`report_at`",
		"`created_at`",
	}
	return columns
}

func (obj *CarParkStatus) GetPrimaryKey() PrimaryKey {
	pk := CarParkStatusMgr.NewPrimaryKey()
	pk.CarParkNo = obj.CarParkNo
	return pk
}

func (obj *CarParkStatus) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type CarParkNoOfCarParkStatusPK struct {
	CarParkNo string
}

func (m *_CarParkStatusMgr) NewPrimaryKey() *CarParkNoOfCarParkStatusPK {
	return &CarParkNoOfCarParkStatusPK{}
}

func (u *CarParkNoOfCarParkStatusPK) Key() string {
	strs := []string{
		"CarParkNo",
		fmt.Sprint(u.CarParkNo),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *CarParkNoOfCarParkStatusPK) Parse(key string) error {
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

func (u *CarParkNoOfCarParkStatusPK) SQLFormat() string {
	conditions := []string{
		"`car_park_no` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *CarParkNoOfCarParkStatusPK) SQLParams() []interface{} {
	return []interface{}{
		u.CarParkNo,
	}
}

func (u *CarParkNoOfCarParkStatusPK) Columns() []string {
	return []string{
		"`car_park_no`",
	}
}

//! uniques

//! indexes

type AvailableLotsOfCarParkStatusIDX struct {
	AvailableLots int32
	offset        int
	limit         int
}

func (u *AvailableLotsOfCarParkStatusIDX) Key() string {
	strs := []string{
		"AvailableLots",
		fmt.Sprint(u.AvailableLots),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *AvailableLotsOfCarParkStatusIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`available_lots` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *AvailableLotsOfCarParkStatusIDX) SQLParams() []interface{} {
	return []interface{}{
		u.AvailableLots,
	}
}

func (u *AvailableLotsOfCarParkStatusIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *AvailableLotsOfCarParkStatusIDX) Limit(n int) {
	u.limit = n
}

func (u *AvailableLotsOfCarParkStatusIDX) Offset(n int) {
	u.offset = n
}

func (u *AvailableLotsOfCarParkStatusIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *AvailableLotsOfCarParkStatusIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type _CarParkStatusDBMgr struct {
	db orm.DB
}

func (m *_CarParkStatusMgr) DB(db orm.DB) *_CarParkStatusDBMgr {
	return CarParkStatusDBMgr(db)
}

func CarParkStatusDBMgr(db orm.DB) *_CarParkStatusDBMgr {
	if db == nil {
		panic(fmt.Errorf("CarParkStatusDBMgr init need db"))
	}
	return &_CarParkStatusDBMgr{db: db}
}

func (m *_CarParkStatusDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*CarParkStatus, error) {
	obj := CarParkStatusMgr.NewCarParkStatus()
	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*CarParkStatus, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarParkStatus))
	}
	return results, nil
}

func (m *_CarParkStatusDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*CarParkStatus, error) {
	obj := CarParkStatusMgr.NewCarParkStatus()
	q := fmt.Sprintf("SELECT %s FROM carpark_status %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	objs, err := m.FetchBySQL(q, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*CarParkStatus, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarParkStatus))
	}
	return results, nil
}

func (m *_CarParkStatusDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_CarParkStatusDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_CarParkStatusDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("CarParkStatus fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result CarParkStatus
		err = rows.Scan(&(result.CarParkNo), &(result.TotalLots), &(result.AvailableLots), &(result.ReportAt), &(result.CreatedAt))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("CarParkStatus fetch result error: %v", err)
	}
	return
}
func (m *_CarParkStatusDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

func (m *_CarParkStatusDBMgr) Fetch(pk PrimaryKey) (*CarParkStatus, error) {
	obj := CarParkStatusMgr.NewCarParkStatus()
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*CarParkStatus), nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkStatusDBMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*CarParkStatus, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*CarParkNoOfCarParkStatusPK).CarParkNo))
	}
	obj := CarParkStatusMgr.NewCarParkStatus()
	query := fmt.Sprintf("SELECT %s FROM carpark_status WHERE `car_park_no` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*CarParkStatus, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarParkStatus))
	}
	return results, nil
}

func (m *_CarParkStatusDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkStatusDBMgr) FindOneFetch(unique Unique) (*CarParkStatus, error) {
	obj := CarParkStatusMgr.NewCarParkStatus()
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*CarParkStatus), nil
	}
	return nil, orm.NoRecord
}

func (m *_CarParkStatusDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_CarParkStatusDBMgr) FindFetch(index Index) (int64, []*CarParkStatus, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := CarParkStatusMgr.NewCarParkStatus()
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	results := make([]*CarParkStatus, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarParkStatus))
	}
	return total, results, nil
}

func (m *_CarParkStatusDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_CarParkStatusDBMgr) RangeFetch(scope Range) (int64, []*CarParkStatus, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := CarParkStatusMgr.NewCarParkStatus()
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	results := make([]*CarParkStatus, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*CarParkStatus))
	}
	return total, results, nil
}

func (m *_CarParkStatusDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_CarParkStatusDBMgr) RangeRevertFetch(scope Range) (int64, []*CarParkStatus, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_CarParkStatusDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := CarParkStatusMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM carpark_status %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("CarParkStatus query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := CarParkStatusMgr.NewPrimaryKey()
		err = rows.Scan(&(result.CarParkNo))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("CarParkStatus query limit result error: %v", err)
	}
	return
}

func (m *_CarParkStatusDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`car_park_no`) FROM carpark_status %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("CarParkStatus query count error: %v", err)
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

func (m *_CarParkStatusDBMgr) BatchCreate(objs []*CarParkStatus) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*5)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(5, "?"), ",")))
		values = append(values, obj.CarParkNo)
		values = append(values, obj.TotalLots)
		values = append(values, obj.AvailableLots)
		values = append(values, obj.ReportAt)
		values = append(values, obj.CreatedAt)
	}
	query := fmt.Sprintf("INSERT INTO carpark_status(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_CarParkStatusDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE carpark_status SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE carpark_status SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkStatusDBMgr) Create(obj *CarParkStatus) (int64, error) {
	params := orm.NewStringSlice(5, "?")
	q := fmt.Sprintf("INSERT INTO carpark_status(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 5)
	values = append(values, obj.CarParkNo)
	values = append(values, obj.TotalLots)
	values = append(values, obj.AvailableLots)
	values = append(values, obj.ReportAt)
	values = append(values, obj.CreatedAt)
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkStatusDBMgr) Update(obj *CarParkStatus) (int64, error) {
	columns := []string{
		"`total_lots` = ?",
		"`available_lots` = ?",
		"`report_at` = ?",
		"`created_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE carpark_status SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 5-1)
	values = append(values, obj.TotalLots)
	values = append(values, obj.AvailableLots)
	values = append(values, obj.ReportAt)
	values = append(values, obj.CreatedAt)
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkStatusDBMgr) Save(obj *CarParkStatus) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_CarParkStatusDBMgr) Delete(obj *CarParkStatus) (int64, error) {
	pk := obj.GetPrimaryKey()
	return m.DeleteByPrimaryKey(pk)
}

func (m *_CarParkStatusDBMgr) DeleteByPrimaryKey(pk PrimaryKey) (int64, error) {
	q := fmt.Sprintf("DELETE FROM carpark_status %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_CarParkStatusDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM carpark_status")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM carpark_status WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
