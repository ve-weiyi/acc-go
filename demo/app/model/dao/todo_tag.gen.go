// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"acc/app/model/entity"
)

func newTodoTag(db *gorm.DB, opts ...gen.DOOption) todoTag {
	_todoTag := todoTag{}

	_todoTag.todoTagDo.UseDB(db, opts...)
	_todoTag.todoTagDo.UseModel(&entity.TodoTag{})

	tableName := _todoTag.todoTagDo.TableName()
	_todoTag.ALL = field.NewAsterisk(tableName)
	_todoTag.ID = field.NewInt(tableName, "id")
	_todoTag.TodoID = field.NewInt(tableName, "todo_id")
	_todoTag.TagID = field.NewInt(tableName, "tag_id")

	_todoTag.fillFieldMap()

	return _todoTag
}

type todoTag struct {
	todoTagDo

	ALL    field.Asterisk
	ID     field.Int
	TodoID field.Int
	TagID  field.Int

	fieldMap map[string]field.Expr
}

func (t todoTag) Table(newTableName string) *todoTag {
	t.todoTagDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t todoTag) As(alias string) *todoTag {
	t.todoTagDo.DO = *(t.todoTagDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *todoTag) updateTableName(table string) *todoTag {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt(table, "id")
	t.TodoID = field.NewInt(table, "todo_id")
	t.TagID = field.NewInt(table, "tag_id")

	t.fillFieldMap()

	return t
}

func (t *todoTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *todoTag) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["todo_id"] = t.TodoID
	t.fieldMap["tag_id"] = t.TagID
}

func (t todoTag) clone(db *gorm.DB) todoTag {
	t.todoTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t todoTag) replaceDB(db *gorm.DB) todoTag {
	t.todoTagDo.ReplaceDB(db)
	return t
}

type todoTagDo struct{ gen.DO }

func (t todoTagDo) Debug() *todoTagDo {
	return t.withDO(t.DO.Debug())
}

func (t todoTagDo) WithContext(ctx context.Context) *todoTagDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t todoTagDo) ReadDB() *todoTagDo {
	return t.Clauses(dbresolver.Read)
}

func (t todoTagDo) WriteDB() *todoTagDo {
	return t.Clauses(dbresolver.Write)
}

func (t todoTagDo) Session(config *gorm.Session) *todoTagDo {
	return t.withDO(t.DO.Session(config))
}

func (t todoTagDo) Clauses(conds ...clause.Expression) *todoTagDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t todoTagDo) Returning(value interface{}, columns ...string) *todoTagDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t todoTagDo) Not(conds ...gen.Condition) *todoTagDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t todoTagDo) Or(conds ...gen.Condition) *todoTagDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t todoTagDo) Select(conds ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t todoTagDo) Where(conds ...gen.Condition) *todoTagDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t todoTagDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *todoTagDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t todoTagDo) Order(conds ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t todoTagDo) Distinct(cols ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t todoTagDo) Omit(cols ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t todoTagDo) Join(table schema.Tabler, on ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t todoTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t todoTagDo) RightJoin(table schema.Tabler, on ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t todoTagDo) Group(cols ...field.Expr) *todoTagDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t todoTagDo) Having(conds ...gen.Condition) *todoTagDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t todoTagDo) Limit(limit int) *todoTagDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t todoTagDo) Offset(offset int) *todoTagDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t todoTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *todoTagDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t todoTagDo) Unscoped() *todoTagDo {
	return t.withDO(t.DO.Unscoped())
}

func (t todoTagDo) Create(values ...*entity.TodoTag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t todoTagDo) CreateInBatches(values []*entity.TodoTag, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t todoTagDo) Save(values ...*entity.TodoTag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t todoTagDo) First() (*entity.TodoTag, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TodoTag), nil
	}
}

func (t todoTagDo) Take() (*entity.TodoTag, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TodoTag), nil
	}
}

func (t todoTagDo) Last() (*entity.TodoTag, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TodoTag), nil
	}
}

func (t todoTagDo) Find() ([]*entity.TodoTag, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TodoTag), err
}

func (t todoTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TodoTag, err error) {
	buf := make([]*entity.TodoTag, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t todoTagDo) FindInBatches(result *[]*entity.TodoTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t todoTagDo) Attrs(attrs ...field.AssignExpr) *todoTagDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t todoTagDo) Assign(attrs ...field.AssignExpr) *todoTagDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t todoTagDo) Joins(fields ...field.RelationField) *todoTagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t todoTagDo) Preload(fields ...field.RelationField) *todoTagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t todoTagDo) FirstOrInit() (*entity.TodoTag, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TodoTag), nil
	}
}

func (t todoTagDo) FirstOrCreate() (*entity.TodoTag, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TodoTag), nil
	}
}

func (t todoTagDo) FindByPage(offset int, limit int) (result []*entity.TodoTag, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t todoTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t todoTagDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t todoTagDo) Delete(models ...*entity.TodoTag) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *todoTagDo) withDO(do gen.Dao) *todoTagDo {
	t.DO = *do.(*gen.DO)
	return t
}
