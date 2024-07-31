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

func newAppAllow(db *gorm.DB, opts ...gen.DOOption) appAllow {
	_appAllow := appAllow{}

	_appAllow.appAllowDo.UseDB(db, opts...)
	_appAllow.appAllowDo.UseModel(&model.AppAllow{})

	tableName := _appAllow.appAllowDo.TableName()
	_appAllow.ALL = field.NewAsterisk(tableName)
	_appAllow.ID = field.NewInt64(tableName, "id")
	_appAllow.CreateAt = field.NewTime(tableName, "create_at")
	_appAllow.UpdateAt = field.NewTime(tableName, "update_at")
	_appAllow.DeleteAt = field.NewTime(tableName, "delete_at")
	_appAllow.AppID = field.NewInt64(tableName, "app_id")
	_appAllow.ServiceID = field.NewInt32(tableName, "service_id")
	_appAllow.ServiceAllow = field.NewTime(tableName, "service_allow")
	_appAllow.ServiceExp = field.NewTime(tableName, "service_exp")
	_appAllow.UseType = field.NewString(tableName, "use_type")
	_appAllow.Status = field.NewString(tableName, "status")

	_appAllow.fillFieldMap()

	return _appAllow
}

// appAllow 业务应用许可的服务项目和期限
type appAllow struct {
	appAllowDo appAllowDo

	ALL          field.Asterisk
	ID           field.Int64  // 数据id
	CreateAt     field.Time   // 创建时间
	UpdateAt     field.Time   // 更新时间
	DeleteAt     field.Time   // 删除时间
	AppID        field.Int64  // 应用id
	ServiceID    field.Int32  // 服务id
	ServiceAllow field.Time   // 服务许可时间
	ServiceExp   field.Time   // 服务过期时间
	UseType      field.String // 使用方式(1次数 2流量 3期限)
	Status       field.String // 1表示过期2禁用3正常4暂停

	fieldMap map[string]field.Expr
}

func (a appAllow) Table(newTableName string) *appAllow {
	a.appAllowDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appAllow) As(alias string) *appAllow {
	a.appAllowDo.DO = *(a.appAllowDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appAllow) updateTableName(table string) *appAllow {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreateAt = field.NewTime(table, "create_at")
	a.UpdateAt = field.NewTime(table, "update_at")
	a.DeleteAt = field.NewTime(table, "delete_at")
	a.AppID = field.NewInt64(table, "app_id")
	a.ServiceID = field.NewInt32(table, "service_id")
	a.ServiceAllow = field.NewTime(table, "service_allow")
	a.ServiceExp = field.NewTime(table, "service_exp")
	a.UseType = field.NewString(table, "use_type")
	a.Status = field.NewString(table, "status")

	a.fillFieldMap()

	return a
}

func (a *appAllow) WithContext(ctx context.Context) IAppAllowDo { return a.appAllowDo.WithContext(ctx) }

func (a appAllow) TableName() string { return a.appAllowDo.TableName() }

func (a appAllow) Alias() string { return a.appAllowDo.Alias() }

func (a appAllow) Columns(cols ...field.Expr) gen.Columns { return a.appAllowDo.Columns(cols...) }

func (a *appAllow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appAllow) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["delete_at"] = a.DeleteAt
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["service_id"] = a.ServiceID
	a.fieldMap["service_allow"] = a.ServiceAllow
	a.fieldMap["service_exp"] = a.ServiceExp
	a.fieldMap["use_type"] = a.UseType
	a.fieldMap["status"] = a.Status
}

func (a appAllow) clone(db *gorm.DB) appAllow {
	a.appAllowDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appAllow) replaceDB(db *gorm.DB) appAllow {
	a.appAllowDo.ReplaceDB(db)
	return a
}

type appAllowDo struct{ gen.DO }

type IAppAllowDo interface {
	gen.SubQuery
	Debug() IAppAllowDo
	WithContext(ctx context.Context) IAppAllowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppAllowDo
	WriteDB() IAppAllowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppAllowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppAllowDo
	Not(conds ...gen.Condition) IAppAllowDo
	Or(conds ...gen.Condition) IAppAllowDo
	Select(conds ...field.Expr) IAppAllowDo
	Where(conds ...gen.Condition) IAppAllowDo
	Order(conds ...field.Expr) IAppAllowDo
	Distinct(cols ...field.Expr) IAppAllowDo
	Omit(cols ...field.Expr) IAppAllowDo
	Join(table schema.Tabler, on ...field.Expr) IAppAllowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppAllowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppAllowDo
	Group(cols ...field.Expr) IAppAllowDo
	Having(conds ...gen.Condition) IAppAllowDo
	Limit(limit int) IAppAllowDo
	Offset(offset int) IAppAllowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppAllowDo
	Unscoped() IAppAllowDo
	Create(values ...*model.AppAllow) error
	CreateInBatches(values []*model.AppAllow, batchSize int) error
	Save(values ...*model.AppAllow) error
	First() (*model.AppAllow, error)
	Take() (*model.AppAllow, error)
	Last() (*model.AppAllow, error)
	Find() ([]*model.AppAllow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AppAllow, err error)
	FindInBatches(result *[]*model.AppAllow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AppAllow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppAllowDo
	Assign(attrs ...field.AssignExpr) IAppAllowDo
	Joins(fields ...field.RelationField) IAppAllowDo
	Preload(fields ...field.RelationField) IAppAllowDo
	FirstOrInit() (*model.AppAllow, error)
	FirstOrCreate() (*model.AppAllow, error)
	FindByPage(offset int, limit int) (result []*model.AppAllow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppAllowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appAllowDo) Debug() IAppAllowDo {
	return a.withDO(a.DO.Debug())
}

func (a appAllowDo) WithContext(ctx context.Context) IAppAllowDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appAllowDo) ReadDB() IAppAllowDo {
	return a.Clauses(dbresolver.Read)
}

func (a appAllowDo) WriteDB() IAppAllowDo {
	return a.Clauses(dbresolver.Write)
}

func (a appAllowDo) Session(config *gorm.Session) IAppAllowDo {
	return a.withDO(a.DO.Session(config))
}

func (a appAllowDo) Clauses(conds ...clause.Expression) IAppAllowDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appAllowDo) Returning(value interface{}, columns ...string) IAppAllowDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appAllowDo) Not(conds ...gen.Condition) IAppAllowDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appAllowDo) Or(conds ...gen.Condition) IAppAllowDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appAllowDo) Select(conds ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appAllowDo) Where(conds ...gen.Condition) IAppAllowDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appAllowDo) Order(conds ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appAllowDo) Distinct(cols ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appAllowDo) Omit(cols ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appAllowDo) Join(table schema.Tabler, on ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appAllowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appAllowDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appAllowDo) Group(cols ...field.Expr) IAppAllowDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appAllowDo) Having(conds ...gen.Condition) IAppAllowDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appAllowDo) Limit(limit int) IAppAllowDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appAllowDo) Offset(offset int) IAppAllowDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appAllowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppAllowDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appAllowDo) Unscoped() IAppAllowDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appAllowDo) Create(values ...*model.AppAllow) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appAllowDo) CreateInBatches(values []*model.AppAllow, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appAllowDo) Save(values ...*model.AppAllow) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appAllowDo) First() (*model.AppAllow, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppAllow), nil
	}
}

func (a appAllowDo) Take() (*model.AppAllow, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppAllow), nil
	}
}

func (a appAllowDo) Last() (*model.AppAllow, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppAllow), nil
	}
}

func (a appAllowDo) Find() ([]*model.AppAllow, error) {
	result, err := a.DO.Find()
	return result.([]*model.AppAllow), err
}

func (a appAllowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AppAllow, err error) {
	buf := make([]*model.AppAllow, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appAllowDo) FindInBatches(result *[]*model.AppAllow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appAllowDo) Attrs(attrs ...field.AssignExpr) IAppAllowDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appAllowDo) Assign(attrs ...field.AssignExpr) IAppAllowDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appAllowDo) Joins(fields ...field.RelationField) IAppAllowDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appAllowDo) Preload(fields ...field.RelationField) IAppAllowDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appAllowDo) FirstOrInit() (*model.AppAllow, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppAllow), nil
	}
}

func (a appAllowDo) FirstOrCreate() (*model.AppAllow, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppAllow), nil
	}
}

func (a appAllowDo) FindByPage(offset int, limit int) (result []*model.AppAllow, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appAllowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appAllowDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appAllowDo) Delete(models ...*model.AppAllow) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appAllowDo) withDO(do gen.Dao) *appAllowDo {
	a.DO = *do.(*gen.DO)
	return a
}