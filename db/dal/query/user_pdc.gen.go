// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"coredx/db/dal/model"
)

func newUserPdc(db *gorm.DB, opts ...gen.DOOption) userPdc {
	_userPdc := userPdc{}

	_userPdc.userPdcDo.UseDB(db, opts...)
	_userPdc.userPdcDo.UseModel(&model.UserPdc{})

	tableName := _userPdc.userPdcDo.TableName()
	_userPdc.ALL = field.NewAsterisk(tableName)
	_userPdc.ID = field.NewInt64(tableName, "id")
	_userPdc.CreateAt = field.NewTime(tableName, "create_at")
	_userPdc.UpdateAt = field.NewTime(tableName, "update_at")
	_userPdc.DeleteAt = field.NewTime(tableName, "delete_at")
	_userPdc.UserID = field.NewInt64(tableName, "user_id")
	_userPdc.SpaceTotal = field.NewInt32(tableName, "space_total")
	_userPdc.SpaceUse = field.NewFloat64(tableName, "space_use")
	_userPdc.SpaceAvailable = field.NewFloat64(tableName, "space_available")
	_userPdc.Status = field.NewString(tableName, "status")

	_userPdc.fillFieldMap()

	return _userPdc
}

// userPdc 用户数据空间变化记录表
type userPdc struct {
	userPdcDo userPdcDo

	ALL            field.Asterisk
	ID             field.Int64   // 数据id
	CreateAt       field.Time    // 创建时间
	UpdateAt       field.Time    // 更新时间
	DeleteAt       field.Time    // 删除时间
	UserID         field.Int64   // 用户id
	SpaceTotal     field.Int32   // 默认单位MB
	SpaceUse       field.Float64 // 使用空间单位字节
	SpaceAvailable field.Float64 // 剩余空间单位字节
	Status         field.String  // 1正常 2禁用

	fieldMap map[string]field.Expr
}

func (u userPdc) Table(newTableName string) *userPdc {
	u.userPdcDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userPdc) As(alias string) *userPdc {
	u.userPdcDo.DO = *(u.userPdcDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userPdc) updateTableName(table string) *userPdc {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreateAt = field.NewTime(table, "create_at")
	u.UpdateAt = field.NewTime(table, "update_at")
	u.DeleteAt = field.NewTime(table, "delete_at")
	u.UserID = field.NewInt64(table, "user_id")
	u.SpaceTotal = field.NewInt32(table, "space_total")
	u.SpaceUse = field.NewFloat64(table, "space_use")
	u.SpaceAvailable = field.NewFloat64(table, "space_available")
	u.Status = field.NewString(table, "status")

	u.fillFieldMap()

	return u
}

func (u *userPdc) WithContext(ctx context.Context) IUserPdcDo { return u.userPdcDo.WithContext(ctx) }

func (u userPdc) TableName() string { return u.userPdcDo.TableName() }

func (u userPdc) Alias() string { return u.userPdcDo.Alias() }

func (u userPdc) Columns(cols ...field.Expr) gen.Columns { return u.userPdcDo.Columns(cols...) }

func (u *userPdc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userPdc) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["create_at"] = u.CreateAt
	u.fieldMap["update_at"] = u.UpdateAt
	u.fieldMap["delete_at"] = u.DeleteAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["space_total"] = u.SpaceTotal
	u.fieldMap["space_use"] = u.SpaceUse
	u.fieldMap["space_available"] = u.SpaceAvailable
	u.fieldMap["status"] = u.Status
}

func (u userPdc) clone(db *gorm.DB) userPdc {
	u.userPdcDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userPdc) replaceDB(db *gorm.DB) userPdc {
	u.userPdcDo.ReplaceDB(db)
	return u
}

type userPdcDo struct{ gen.DO }

type IUserPdcDo interface {
	gen.SubQuery
	Debug() IUserPdcDo
	WithContext(ctx context.Context) IUserPdcDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserPdcDo
	WriteDB() IUserPdcDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserPdcDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserPdcDo
	Not(conds ...gen.Condition) IUserPdcDo
	Or(conds ...gen.Condition) IUserPdcDo
	Select(conds ...field.Expr) IUserPdcDo
	Where(conds ...gen.Condition) IUserPdcDo
	Order(conds ...field.Expr) IUserPdcDo
	Distinct(cols ...field.Expr) IUserPdcDo
	Omit(cols ...field.Expr) IUserPdcDo
	Join(table schema.Tabler, on ...field.Expr) IUserPdcDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserPdcDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserPdcDo
	Group(cols ...field.Expr) IUserPdcDo
	Having(conds ...gen.Condition) IUserPdcDo
	Limit(limit int) IUserPdcDo
	Offset(offset int) IUserPdcDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserPdcDo
	Unscoped() IUserPdcDo
	Create(values ...*model.UserPdc) error
	CreateInBatches(values []*model.UserPdc, batchSize int) error
	Save(values ...*model.UserPdc) error
	First() (*model.UserPdc, error)
	Take() (*model.UserPdc, error)
	Last() (*model.UserPdc, error)
	Find() ([]*model.UserPdc, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserPdc, err error)
	FindInBatches(result *[]*model.UserPdc, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserPdc) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserPdcDo
	Assign(attrs ...field.AssignExpr) IUserPdcDo
	Joins(fields ...field.RelationField) IUserPdcDo
	Preload(fields ...field.RelationField) IUserPdcDo
	FirstOrInit() (*model.UserPdc, error)
	FirstOrCreate() (*model.UserPdc, error)
	FindByPage(offset int, limit int) (result []*model.UserPdc, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserPdcDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userPdcDo) Debug() IUserPdcDo {
	return u.withDO(u.DO.Debug())
}

func (u userPdcDo) WithContext(ctx context.Context) IUserPdcDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userPdcDo) ReadDB() IUserPdcDo {
	return u.Clauses(dbresolver.Read)
}

func (u userPdcDo) WriteDB() IUserPdcDo {
	return u.Clauses(dbresolver.Write)
}

func (u userPdcDo) Session(config *gorm.Session) IUserPdcDo {
	return u.withDO(u.DO.Session(config))
}

func (u userPdcDo) Clauses(conds ...clause.Expression) IUserPdcDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userPdcDo) Returning(value interface{}, columns ...string) IUserPdcDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userPdcDo) Not(conds ...gen.Condition) IUserPdcDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userPdcDo) Or(conds ...gen.Condition) IUserPdcDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userPdcDo) Select(conds ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userPdcDo) Where(conds ...gen.Condition) IUserPdcDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userPdcDo) Order(conds ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userPdcDo) Distinct(cols ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userPdcDo) Omit(cols ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userPdcDo) Join(table schema.Tabler, on ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userPdcDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userPdcDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userPdcDo) Group(cols ...field.Expr) IUserPdcDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userPdcDo) Having(conds ...gen.Condition) IUserPdcDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userPdcDo) Limit(limit int) IUserPdcDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userPdcDo) Offset(offset int) IUserPdcDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userPdcDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserPdcDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userPdcDo) Unscoped() IUserPdcDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userPdcDo) Create(values ...*model.UserPdc) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userPdcDo) CreateInBatches(values []*model.UserPdc, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userPdcDo) Save(values ...*model.UserPdc) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userPdcDo) First() (*model.UserPdc, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPdc), nil
	}
}

func (u userPdcDo) Take() (*model.UserPdc, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPdc), nil
	}
}

func (u userPdcDo) Last() (*model.UserPdc, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPdc), nil
	}
}

func (u userPdcDo) Find() ([]*model.UserPdc, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserPdc), err
}

func (u userPdcDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserPdc, err error) {
	buf := make([]*model.UserPdc, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userPdcDo) FindInBatches(result *[]*model.UserPdc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userPdcDo) Attrs(attrs ...field.AssignExpr) IUserPdcDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userPdcDo) Assign(attrs ...field.AssignExpr) IUserPdcDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userPdcDo) Joins(fields ...field.RelationField) IUserPdcDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userPdcDo) Preload(fields ...field.RelationField) IUserPdcDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userPdcDo) FirstOrInit() (*model.UserPdc, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPdc), nil
	}
}

func (u userPdcDo) FirstOrCreate() (*model.UserPdc, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPdc), nil
	}
}

func (u userPdcDo) FindByPage(offset int, limit int) (result []*model.UserPdc, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userPdcDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userPdcDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userPdcDo) Delete(models ...*model.UserPdc) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userPdcDo) withDO(do gen.Dao) *userPdcDo {
	u.DO = *do.(*gen.DO)
	return u
}