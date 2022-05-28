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

	"github.com/shiningrush/fastflow/store/gorm/model"
)

func newTaskInstance(db *gorm.DB) taskInstance {
	_taskInstance := taskInstance{}

	_taskInstance.taskInstanceDo.UseDB(db)
	_taskInstance.taskInstanceDo.UseModel(&model.TaskInstance{})

	tableName := _taskInstance.taskInstanceDo.TableName()
	_taskInstance.ALL = field.NewField(tableName, "*")
	_taskInstance.ID = field.NewInt32(tableName, "id")
	_taskInstance.UID = field.NewString(tableName, "uid")
	_taskInstance.TaskUID = field.NewString(tableName, "task_uid")
	_taskInstance.DagInstanceUID = field.NewString(tableName, "dag_instance_uid")
	_taskInstance.Name = field.NewString(tableName, "name")
	_taskInstance.ActionName = field.NewString(tableName, "action_name")
	_taskInstance.TimeoutSecs = field.NewInt32(tableName, "timeout_secs")
	_taskInstance.Params = field.NewString(tableName, "params")
	_taskInstance.Status = field.NewString(tableName, "status")
	_taskInstance.Reason = field.NewString(tableName, "reason")
	_taskInstance.Precheck = field.NewString(tableName, "precheck")

	_taskInstance.fillFieldMap()

	return _taskInstance
}

type taskInstance struct {
	taskInstanceDo taskInstanceDo

	ALL            field.Field
	ID             field.Int32
	UID            field.String
	TaskUID        field.String
	DagInstanceUID field.String
	Name           field.String
	ActionName     field.String
	TimeoutSecs    field.Int32
	Params         field.String
	Status         field.String
	Reason         field.String
	Precheck       field.String

	fieldMap map[string]field.Expr
}

func (t taskInstance) Table(newTableName string) *taskInstance {
	t.taskInstanceDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskInstance) As(alias string) *taskInstance {
	t.taskInstanceDo.DO = *(t.taskInstanceDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskInstance) updateTableName(table string) *taskInstance {
	t.ALL = field.NewField(table, "*")
	t.ID = field.NewInt32(table, "id")
	t.UID = field.NewString(table, "uid")
	t.TaskUID = field.NewString(table, "task_uid")
	t.DagInstanceUID = field.NewString(table, "dag_instance_uid")
	t.Name = field.NewString(table, "name")
	t.ActionName = field.NewString(table, "action_name")
	t.TimeoutSecs = field.NewInt32(table, "timeout_secs")
	t.Params = field.NewString(table, "params")
	t.Status = field.NewString(table, "status")
	t.Reason = field.NewString(table, "reason")
	t.Precheck = field.NewString(table, "precheck")

	t.fillFieldMap()

	return t
}

func (t *taskInstance) WithContext(ctx context.Context) *taskInstanceDo {
	return t.taskInstanceDo.WithContext(ctx)
}

func (t taskInstance) TableName() string { return t.taskInstanceDo.TableName() }

func (t taskInstance) Alias() string { return t.taskInstanceDo.Alias() }

func (t *taskInstance) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskInstance) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["uid"] = t.UID
	t.fieldMap["task_uid"] = t.TaskUID
	t.fieldMap["dag_instance_uid"] = t.DagInstanceUID
	t.fieldMap["name"] = t.Name
	t.fieldMap["action_name"] = t.ActionName
	t.fieldMap["timeout_secs"] = t.TimeoutSecs
	t.fieldMap["params"] = t.Params
	t.fieldMap["status"] = t.Status
	t.fieldMap["reason"] = t.Reason
	t.fieldMap["precheck"] = t.Precheck
}

func (t taskInstance) clone(db *gorm.DB) taskInstance {
	t.taskInstanceDo.ReplaceDB(db)
	return t
}

type taskInstanceDo struct{ gen.DO }

func (t taskInstanceDo) Debug() *taskInstanceDo {
	return t.withDO(t.DO.Debug())
}

func (t taskInstanceDo) WithContext(ctx context.Context) *taskInstanceDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskInstanceDo) Clauses(conds ...clause.Expression) *taskInstanceDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskInstanceDo) Returning(value interface{}, columns ...string) *taskInstanceDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskInstanceDo) Not(conds ...gen.Condition) *taskInstanceDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskInstanceDo) Or(conds ...gen.Condition) *taskInstanceDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskInstanceDo) Select(conds ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskInstanceDo) Where(conds ...gen.Condition) *taskInstanceDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskInstanceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *taskInstanceDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t taskInstanceDo) Order(conds ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskInstanceDo) Distinct(cols ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskInstanceDo) Omit(cols ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskInstanceDo) Join(table schema.Tabler, on ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskInstanceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskInstanceDo) RightJoin(table schema.Tabler, on ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskInstanceDo) Group(cols ...field.Expr) *taskInstanceDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskInstanceDo) Having(conds ...gen.Condition) *taskInstanceDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskInstanceDo) Limit(limit int) *taskInstanceDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskInstanceDo) Offset(offset int) *taskInstanceDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskInstanceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *taskInstanceDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskInstanceDo) Unscoped() *taskInstanceDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskInstanceDo) Create(values ...*model.TaskInstance) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskInstanceDo) CreateInBatches(values []*model.TaskInstance, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskInstanceDo) Save(values ...*model.TaskInstance) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskInstanceDo) First() (*model.TaskInstance, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskInstance), nil
	}
}

func (t taskInstanceDo) Take() (*model.TaskInstance, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskInstance), nil
	}
}

func (t taskInstanceDo) Last() (*model.TaskInstance, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskInstance), nil
	}
}

func (t taskInstanceDo) Find() ([]*model.TaskInstance, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskInstance), err
}

func (t taskInstanceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskInstance, err error) {
	buf := make([]*model.TaskInstance, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskInstanceDo) FindInBatches(result *[]*model.TaskInstance, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskInstanceDo) Attrs(attrs ...field.AssignExpr) *taskInstanceDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskInstanceDo) Assign(attrs ...field.AssignExpr) *taskInstanceDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskInstanceDo) Joins(fields ...field.RelationField) *taskInstanceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskInstanceDo) Preload(fields ...field.RelationField) *taskInstanceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskInstanceDo) FirstOrInit() (*model.TaskInstance, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskInstance), nil
	}
}

func (t taskInstanceDo) FirstOrCreate() (*model.TaskInstance, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskInstance), nil
	}
}

func (t taskInstanceDo) FindByPage(offset int, limit int) (result []*model.TaskInstance, count int64, err error) {
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

func (t taskInstanceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *taskInstanceDo) withDO(do gen.Dao) *taskInstanceDo {
	t.DO = *do.(*gen.DO)
	return t
}
