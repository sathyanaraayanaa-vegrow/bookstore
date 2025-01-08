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

func newBorrowingRecord(db *gorm.DB, opts ...gen.DOOption) borrowingRecord {
	_borrowingRecord := borrowingRecord{}

	_borrowingRecord.borrowingRecordDo.UseDB(db, opts...)
	_borrowingRecord.borrowingRecordDo.UseModel(&model.BorrowingRecord{})

	tableName := _borrowingRecord.borrowingRecordDo.TableName()
	_borrowingRecord.ALL = field.NewAsterisk(tableName)
	_borrowingRecord.ID = field.NewInt64(tableName, "id")
	_borrowingRecord.BooksID = field.NewInt64(tableName, "books_id")
	_borrowingRecord.UsersID = field.NewInt64(tableName, "users_id")
	_borrowingRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_borrowingRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_borrowingRecord.ReturnedAt = field.NewString(tableName, "returned_at")

	_borrowingRecord.fillFieldMap()

	return _borrowingRecord
}

type borrowingRecord struct {
	borrowingRecordDo

	ALL        field.Asterisk
	ID         field.Int64
	BooksID    field.Int64
	UsersID    field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	ReturnedAt field.String

	fieldMap map[string]field.Expr
}

func (b borrowingRecord) Table(newTableName string) *borrowingRecord {
	b.borrowingRecordDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b borrowingRecord) As(alias string) *borrowingRecord {
	b.borrowingRecordDo.DO = *(b.borrowingRecordDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *borrowingRecord) updateTableName(table string) *borrowingRecord {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.BooksID = field.NewInt64(table, "books_id")
	b.UsersID = field.NewInt64(table, "users_id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.ReturnedAt = field.NewString(table, "returned_at")

	b.fillFieldMap()

	return b
}

func (b *borrowingRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *borrowingRecord) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 6)
	b.fieldMap["id"] = b.ID
	b.fieldMap["books_id"] = b.BooksID
	b.fieldMap["users_id"] = b.UsersID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["returned_at"] = b.ReturnedAt
}

func (b borrowingRecord) clone(db *gorm.DB) borrowingRecord {
	b.borrowingRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b borrowingRecord) replaceDB(db *gorm.DB) borrowingRecord {
	b.borrowingRecordDo.ReplaceDB(db)
	return b
}

type borrowingRecordDo struct{ gen.DO }

type IBorrowingRecordDo interface {
	gen.SubQuery
	Debug() IBorrowingRecordDo
	WithContext(ctx context.Context) IBorrowingRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBorrowingRecordDo
	WriteDB() IBorrowingRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBorrowingRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBorrowingRecordDo
	Not(conds ...gen.Condition) IBorrowingRecordDo
	Or(conds ...gen.Condition) IBorrowingRecordDo
	Select(conds ...field.Expr) IBorrowingRecordDo
	Where(conds ...gen.Condition) IBorrowingRecordDo
	Order(conds ...field.Expr) IBorrowingRecordDo
	Distinct(cols ...field.Expr) IBorrowingRecordDo
	Omit(cols ...field.Expr) IBorrowingRecordDo
	Join(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo
	Group(cols ...field.Expr) IBorrowingRecordDo
	Having(conds ...gen.Condition) IBorrowingRecordDo
	Limit(limit int) IBorrowingRecordDo
	Offset(offset int) IBorrowingRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBorrowingRecordDo
	Unscoped() IBorrowingRecordDo
	Create(values ...*model.BorrowingRecord) error
	CreateInBatches(values []*model.BorrowingRecord, batchSize int) error
	Save(values ...*model.BorrowingRecord) error
	First() (*model.BorrowingRecord, error)
	Take() (*model.BorrowingRecord, error)
	Last() (*model.BorrowingRecord, error)
	Find() ([]*model.BorrowingRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BorrowingRecord, err error)
	FindInBatches(result *[]*model.BorrowingRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BorrowingRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBorrowingRecordDo
	Assign(attrs ...field.AssignExpr) IBorrowingRecordDo
	Joins(fields ...field.RelationField) IBorrowingRecordDo
	Preload(fields ...field.RelationField) IBorrowingRecordDo
	FirstOrInit() (*model.BorrowingRecord, error)
	FirstOrCreate() (*model.BorrowingRecord, error)
	FindByPage(offset int, limit int) (result []*model.BorrowingRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBorrowingRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b borrowingRecordDo) Debug() IBorrowingRecordDo {
	return b.withDO(b.DO.Debug())
}

func (b borrowingRecordDo) WithContext(ctx context.Context) IBorrowingRecordDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b borrowingRecordDo) ReadDB() IBorrowingRecordDo {
	return b.Clauses(dbresolver.Read)
}

func (b borrowingRecordDo) WriteDB() IBorrowingRecordDo {
	return b.Clauses(dbresolver.Write)
}

func (b borrowingRecordDo) Session(config *gorm.Session) IBorrowingRecordDo {
	return b.withDO(b.DO.Session(config))
}

func (b borrowingRecordDo) Clauses(conds ...clause.Expression) IBorrowingRecordDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b borrowingRecordDo) Returning(value interface{}, columns ...string) IBorrowingRecordDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b borrowingRecordDo) Not(conds ...gen.Condition) IBorrowingRecordDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b borrowingRecordDo) Or(conds ...gen.Condition) IBorrowingRecordDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b borrowingRecordDo) Select(conds ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b borrowingRecordDo) Where(conds ...gen.Condition) IBorrowingRecordDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b borrowingRecordDo) Order(conds ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b borrowingRecordDo) Distinct(cols ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b borrowingRecordDo) Omit(cols ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b borrowingRecordDo) Join(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b borrowingRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b borrowingRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b borrowingRecordDo) Group(cols ...field.Expr) IBorrowingRecordDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b borrowingRecordDo) Having(conds ...gen.Condition) IBorrowingRecordDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b borrowingRecordDo) Limit(limit int) IBorrowingRecordDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b borrowingRecordDo) Offset(offset int) IBorrowingRecordDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b borrowingRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBorrowingRecordDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b borrowingRecordDo) Unscoped() IBorrowingRecordDo {
	return b.withDO(b.DO.Unscoped())
}

func (b borrowingRecordDo) Create(values ...*model.BorrowingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b borrowingRecordDo) CreateInBatches(values []*model.BorrowingRecord, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b borrowingRecordDo) Save(values ...*model.BorrowingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b borrowingRecordDo) First() (*model.BorrowingRecord, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BorrowingRecord), nil
	}
}

func (b borrowingRecordDo) Take() (*model.BorrowingRecord, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BorrowingRecord), nil
	}
}

func (b borrowingRecordDo) Last() (*model.BorrowingRecord, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BorrowingRecord), nil
	}
}

func (b borrowingRecordDo) Find() ([]*model.BorrowingRecord, error) {
	result, err := b.DO.Find()
	return result.([]*model.BorrowingRecord), err
}

func (b borrowingRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BorrowingRecord, err error) {
	buf := make([]*model.BorrowingRecord, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b borrowingRecordDo) FindInBatches(result *[]*model.BorrowingRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b borrowingRecordDo) Attrs(attrs ...field.AssignExpr) IBorrowingRecordDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b borrowingRecordDo) Assign(attrs ...field.AssignExpr) IBorrowingRecordDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b borrowingRecordDo) Joins(fields ...field.RelationField) IBorrowingRecordDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b borrowingRecordDo) Preload(fields ...field.RelationField) IBorrowingRecordDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b borrowingRecordDo) FirstOrInit() (*model.BorrowingRecord, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BorrowingRecord), nil
	}
}

func (b borrowingRecordDo) FirstOrCreate() (*model.BorrowingRecord, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BorrowingRecord), nil
	}
}

func (b borrowingRecordDo) FindByPage(offset int, limit int) (result []*model.BorrowingRecord, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b borrowingRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b borrowingRecordDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b borrowingRecordDo) Delete(models ...*model.BorrowingRecord) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *borrowingRecordDo) withDO(do gen.Dao) *borrowingRecordDo {
	b.DO = *do.(*gen.DO)
	return b
}
