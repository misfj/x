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

func newSysAuthorityMenu(db *gorm.DB, opts ...gen.DOOption) sysAuthorityMenu {
	_sysAuthorityMenu := sysAuthorityMenu{}

	_sysAuthorityMenu.sysAuthorityMenuDo.UseDB(db, opts...)
	_sysAuthorityMenu.sysAuthorityMenuDo.UseModel(&model.SysAuthorityMenu{})

	tableName := _sysAuthorityMenu.sysAuthorityMenuDo.TableName()
	_sysAuthorityMenu.ALL = field.NewAsterisk(tableName)
	_sysAuthorityMenu.ID = field.NewInt64(tableName, "id")
	_sysAuthorityMenu.CreateAt = field.NewTime(tableName, "create_at")
	_sysAuthorityMenu.UpdateAt = field.NewTime(tableName, "update_at")
	_sysAuthorityMenu.DeleteAt = field.NewString(tableName, "delete_at")
	_sysAuthorityMenu.AuthorityID = field.NewInt32(tableName, "authority_id")
	_sysAuthorityMenu.MenuID = field.NewInt32(tableName, "menu_id")

	_sysAuthorityMenu.fillFieldMap()

	return _sysAuthorityMenu
}

// sysAuthorityMenu 角色菜单绑定表
type sysAuthorityMenu struct {
	sysAuthorityMenuDo sysAuthorityMenuDo

	ALL         field.Asterisk
	ID          field.Int64  // 数据id
	CreateAt    field.Time   // 创建时间
	UpdateAt    field.Time   // 更新时间
	DeleteAt    field.String // 删除时间
	AuthorityID field.Int32  // 角色ID
	MenuID      field.Int32  // 菜单ID

	fieldMap map[string]field.Expr
}

func (s sysAuthorityMenu) Table(newTableName string) *sysAuthorityMenu {
	s.sysAuthorityMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysAuthorityMenu) As(alias string) *sysAuthorityMenu {
	s.sysAuthorityMenuDo.DO = *(s.sysAuthorityMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysAuthorityMenu) updateTableName(table string) *sysAuthorityMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CreateAt = field.NewTime(table, "create_at")
	s.UpdateAt = field.NewTime(table, "update_at")
	s.DeleteAt = field.NewString(table, "delete_at")
	s.AuthorityID = field.NewInt32(table, "authority_id")
	s.MenuID = field.NewInt32(table, "menu_id")

	s.fillFieldMap()

	return s
}

func (s *sysAuthorityMenu) WithContext(ctx context.Context) ISysAuthorityMenuDo {
	return s.sysAuthorityMenuDo.WithContext(ctx)
}

func (s sysAuthorityMenu) TableName() string { return s.sysAuthorityMenuDo.TableName() }

func (s sysAuthorityMenu) Alias() string { return s.sysAuthorityMenuDo.Alias() }

func (s sysAuthorityMenu) Columns(cols ...field.Expr) gen.Columns {
	return s.sysAuthorityMenuDo.Columns(cols...)
}

func (s *sysAuthorityMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysAuthorityMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["create_at"] = s.CreateAt
	s.fieldMap["update_at"] = s.UpdateAt
	s.fieldMap["delete_at"] = s.DeleteAt
	s.fieldMap["authority_id"] = s.AuthorityID
	s.fieldMap["menu_id"] = s.MenuID
}

func (s sysAuthorityMenu) clone(db *gorm.DB) sysAuthorityMenu {
	s.sysAuthorityMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysAuthorityMenu) replaceDB(db *gorm.DB) sysAuthorityMenu {
	s.sysAuthorityMenuDo.ReplaceDB(db)
	return s
}

type sysAuthorityMenuDo struct{ gen.DO }

type ISysAuthorityMenuDo interface {
	gen.SubQuery
	Debug() ISysAuthorityMenuDo
	WithContext(ctx context.Context) ISysAuthorityMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysAuthorityMenuDo
	WriteDB() ISysAuthorityMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysAuthorityMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysAuthorityMenuDo
	Not(conds ...gen.Condition) ISysAuthorityMenuDo
	Or(conds ...gen.Condition) ISysAuthorityMenuDo
	Select(conds ...field.Expr) ISysAuthorityMenuDo
	Where(conds ...gen.Condition) ISysAuthorityMenuDo
	Order(conds ...field.Expr) ISysAuthorityMenuDo
	Distinct(cols ...field.Expr) ISysAuthorityMenuDo
	Omit(cols ...field.Expr) ISysAuthorityMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo
	Group(cols ...field.Expr) ISysAuthorityMenuDo
	Having(conds ...gen.Condition) ISysAuthorityMenuDo
	Limit(limit int) ISysAuthorityMenuDo
	Offset(offset int) ISysAuthorityMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysAuthorityMenuDo
	Unscoped() ISysAuthorityMenuDo
	Create(values ...*model.SysAuthorityMenu) error
	CreateInBatches(values []*model.SysAuthorityMenu, batchSize int) error
	Save(values ...*model.SysAuthorityMenu) error
	First() (*model.SysAuthorityMenu, error)
	Take() (*model.SysAuthorityMenu, error)
	Last() (*model.SysAuthorityMenu, error)
	Find() ([]*model.SysAuthorityMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysAuthorityMenu, err error)
	FindInBatches(result *[]*model.SysAuthorityMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysAuthorityMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysAuthorityMenuDo
	Assign(attrs ...field.AssignExpr) ISysAuthorityMenuDo
	Joins(fields ...field.RelationField) ISysAuthorityMenuDo
	Preload(fields ...field.RelationField) ISysAuthorityMenuDo
	FirstOrInit() (*model.SysAuthorityMenu, error)
	FirstOrCreate() (*model.SysAuthorityMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysAuthorityMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysAuthorityMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysAuthorityMenuDo) Debug() ISysAuthorityMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysAuthorityMenuDo) WithContext(ctx context.Context) ISysAuthorityMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysAuthorityMenuDo) ReadDB() ISysAuthorityMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysAuthorityMenuDo) WriteDB() ISysAuthorityMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysAuthorityMenuDo) Session(config *gorm.Session) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysAuthorityMenuDo) Clauses(conds ...clause.Expression) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysAuthorityMenuDo) Returning(value interface{}, columns ...string) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysAuthorityMenuDo) Not(conds ...gen.Condition) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysAuthorityMenuDo) Or(conds ...gen.Condition) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysAuthorityMenuDo) Select(conds ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysAuthorityMenuDo) Where(conds ...gen.Condition) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysAuthorityMenuDo) Order(conds ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysAuthorityMenuDo) Distinct(cols ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysAuthorityMenuDo) Omit(cols ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysAuthorityMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysAuthorityMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysAuthorityMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysAuthorityMenuDo) Group(cols ...field.Expr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysAuthorityMenuDo) Having(conds ...gen.Condition) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysAuthorityMenuDo) Limit(limit int) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysAuthorityMenuDo) Offset(offset int) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysAuthorityMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysAuthorityMenuDo) Unscoped() ISysAuthorityMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysAuthorityMenuDo) Create(values ...*model.SysAuthorityMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysAuthorityMenuDo) CreateInBatches(values []*model.SysAuthorityMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysAuthorityMenuDo) Save(values ...*model.SysAuthorityMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysAuthorityMenuDo) First() (*model.SysAuthorityMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysAuthorityMenu), nil
	}
}

func (s sysAuthorityMenuDo) Take() (*model.SysAuthorityMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysAuthorityMenu), nil
	}
}

func (s sysAuthorityMenuDo) Last() (*model.SysAuthorityMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysAuthorityMenu), nil
	}
}

func (s sysAuthorityMenuDo) Find() ([]*model.SysAuthorityMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysAuthorityMenu), err
}

func (s sysAuthorityMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysAuthorityMenu, err error) {
	buf := make([]*model.SysAuthorityMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysAuthorityMenuDo) FindInBatches(result *[]*model.SysAuthorityMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysAuthorityMenuDo) Attrs(attrs ...field.AssignExpr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysAuthorityMenuDo) Assign(attrs ...field.AssignExpr) ISysAuthorityMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysAuthorityMenuDo) Joins(fields ...field.RelationField) ISysAuthorityMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysAuthorityMenuDo) Preload(fields ...field.RelationField) ISysAuthorityMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysAuthorityMenuDo) FirstOrInit() (*model.SysAuthorityMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysAuthorityMenu), nil
	}
}

func (s sysAuthorityMenuDo) FirstOrCreate() (*model.SysAuthorityMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysAuthorityMenu), nil
	}
}

func (s sysAuthorityMenuDo) FindByPage(offset int, limit int) (result []*model.SysAuthorityMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysAuthorityMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysAuthorityMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysAuthorityMenuDo) Delete(models ...*model.SysAuthorityMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysAuthorityMenuDo) withDO(do gen.Dao) *sysAuthorityMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}