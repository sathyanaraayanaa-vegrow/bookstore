// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"BookStore/model"
)

func newAuthor(db *gorm.DB, opts ...gen.DOOption) author {
	_author := author{}

	_author.authorDo.UseDB(db, opts...)
	_author.authorDo.UseModel(&model.Author{})

	tableName := _author.authorDo.TableName()
	_author.ALL = field.NewAsterisk(tableName)
	_author.ID = field.NewInt64(tableName, "id")
	_author.Name = field.NewString(tableName, "name")
	_author.CreatedAt = field.NewTime(tableName, "created_at")
	_author.UpdatedAt = field.NewTime(tableName, "updated_at")

	_author.fillFieldMap()

	return _author
}

type author struct {
	authorDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a author) Table(newTableName string) *author {
	a.authorDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a author) As(alias string) *author {
	a.authorDo.DO = *(a.authorDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *author) updateTableName(table string) *author {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Name = field.NewString(table, "name")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *author) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *author) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a author) clone(db *gorm.DB) author {
	a.authorDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a author) replaceDB(db *gorm.DB) author {
	a.authorDo.ReplaceDB(db)
	return a
}

type authorDo struct{ gen.DO }

type IAuthorDo interface {
	gen.SubQuery
	Debug() IAuthorDo
	WithContext(ctx context.Context) IAuthorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthorDo
	WriteDB() IAuthorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthorDo
	Not(conds ...gen.Condition) IAuthorDo
	Or(conds ...gen.Condition) IAuthorDo
	Select(conds ...field.Expr) IAuthorDo
	Where(conds ...gen.Condition) IAuthorDo
	Order(conds ...field.Expr) IAuthorDo
	Distinct(cols ...field.Expr) IAuthorDo
	Omit(cols ...field.Expr) IAuthorDo
	Join(table schema.Tabler, on ...field.Expr) IAuthorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthorDo
	Group(cols ...field.Expr) IAuthorDo
	Having(conds ...gen.Condition) IAuthorDo
	Limit(limit int) IAuthorDo
	Offset(offset int) IAuthorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthorDo
	Unscoped() IAuthorDo
	Create(values ...*model.Author) error
	CreateInBatches(values []*model.Author, batchSize int) error
	Save(values ...*model.Author) error
	First() (*model.Author, error)
	Take() (*model.Author, error)
	Last() (*model.Author, error)
	Find() ([]*model.Author, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Author, err error)
	FindInBatches(result *[]*model.Author, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Author) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthorDo
	Assign(attrs ...field.AssignExpr) IAuthorDo
	Joins(fields ...field.RelationField) IAuthorDo
	Preload(fields ...field.RelationField) IAuthorDo
	FirstOrInit() (*model.Author, error)
	FirstOrCreate() (*model.Author, error)
	FindByPage(offset int, limit int) (result []*model.Author, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authorDo) Debug() IAuthorDo {
	return a.withDO(a.DO.Debug())
}

func (a authorDo) WithContext(ctx context.Context) IAuthorDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authorDo) ReadDB() IAuthorDo {
	return a.Clauses(dbresolver.Read)
}

func (a authorDo) WriteDB() IAuthorDo {
	return a.Clauses(dbresolver.Write)
}

func (a authorDo) Session(config *gorm.Session) IAuthorDo {
	return a.withDO(a.DO.Session(config))
}

func (a authorDo) Clauses(conds ...clause.Expression) IAuthorDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authorDo) Returning(value interface{}, columns ...string) IAuthorDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authorDo) Not(conds ...gen.Condition) IAuthorDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authorDo) Or(conds ...gen.Condition) IAuthorDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authorDo) Select(conds ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authorDo) Where(conds ...gen.Condition) IAuthorDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authorDo) Order(conds ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authorDo) Distinct(cols ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authorDo) Omit(cols ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authorDo) Join(table schema.Tabler, on ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authorDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authorDo) Group(cols ...field.Expr) IAuthorDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authorDo) Having(conds ...gen.Condition) IAuthorDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authorDo) Limit(limit int) IAuthorDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authorDo) Offset(offset int) IAuthorDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthorDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authorDo) Unscoped() IAuthorDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authorDo) Create(values ...*model.Author) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authorDo) CreateInBatches(values []*model.Author, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authorDo) Save(values ...*model.Author) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authorDo) First() (*model.Author, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Take() (*model.Author, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Last() (*model.Author, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Find() ([]*model.Author, error) {
	result, err := a.DO.Find()
	return result.([]*model.Author), err
}

func (a authorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Author, err error) {
	buf := make([]*model.Author, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authorDo) FindInBatches(result *[]*model.Author, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authorDo) Attrs(attrs ...field.AssignExpr) IAuthorDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authorDo) Assign(attrs ...field.AssignExpr) IAuthorDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authorDo) Joins(fields ...field.RelationField) IAuthorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authorDo) Preload(fields ...field.RelationField) IAuthorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authorDo) FirstOrInit() (*model.Author, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) FirstOrCreate() (*model.Author, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) FindByPage(offset int, limit int) (result []*model.Author, count int64, err error) {
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

func (a authorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authorDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authorDo) Delete(models ...*model.Author) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authorDo) withDO(do gen.Dao) *authorDo {
	a.DO = *do.(*gen.DO)
	return a
}
