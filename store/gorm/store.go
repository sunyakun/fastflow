package gorm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gen/field"
	"gorm.io/gorm"

	"github.com/shiningrush/fastflow/pkg/entity"
	"github.com/shiningrush/fastflow/pkg/mod"
	"github.com/shiningrush/fastflow/pkg/utils"
	"github.com/shiningrush/fastflow/store/gorm/model"
	"github.com/shiningrush/fastflow/store/gorm/query"
)

type Store struct {
	db      *gorm.DB
	query   *query.Query
	context context.Context
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db:      db,
		query:   query.Use(db),
		context: context.Background(),
	}
}

func (s *Store) marshalAndSet(value interface{}, field **string) error {
	v, err := s.Marshal(value)
	if err != nil {
		return err
	}

	if len(v) == 0 {
		*field = nil
	} else {
		*field = utils.BytesToStringPtr(v)
	}

	return nil
}

func (s *Store) toDagModel(dag *entity.Dag) (*model.Dag, error) {
	m := &model.Dag{
		UID:       dag.ID,
		Name:      dag.Name,
		Desc:      &dag.Desc,
		Cron:      &dag.Cron,
		Status:    string(dag.Status),
		CreatedAt: time.Unix(dag.CreatedAt, 0),
		UpdatedAt: time.Unix(dag.UpdatedAt, 0),
	}
	if err := s.marshalAndSet(dag.Vars, &m.Vars); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Store) fromDagModel(dagModel *model.Dag) (*entity.Dag, error) {
	dag := &entity.Dag{
		BaseInfo: entity.BaseInfo{
			ID:        dagModel.UID,
			CreatedAt: dagModel.CreatedAt.Unix(),
			UpdatedAt: dagModel.UpdatedAt.Unix(),
		},
		Name:   dagModel.Name,
		Desc:   utils.StringPtrToVal(dagModel.Desc, ""),
		Cron:   utils.StringPtrToVal(dagModel.Cron, ""),
		Status: entity.DagStatus(dagModel.Status),
	}
	if err := s.Unmarshal([]byte(utils.StringPtrToVal(dagModel.Vars, "")), &dag.Vars); err != nil {
		return nil, err
	}
	return dag, nil
}

func (s *Store) toTaskModel(task *entity.Task, dagUID string) (*model.Task, error) {
	taskModel := model.Task{
		UID:         task.ID,
		DagUID:      dagUID,
		Name:        task.Name,
		ActionName:  task.ActionName,
		TimeoutSecs: utils.IntToInt32Ptr(task.TimeoutSecs),
	}
	if err := s.marshalAndSet(task.Params, &taskModel.Params); err != nil {
		return nil, err
	}
	if err := s.marshalAndSet(task.PreChecks, &taskModel.Prechecks); err != nil {
		return nil, err
	}
	if err := s.marshalAndSet(task.DependOn, &taskModel.DependOn); err != nil {
		return nil, err
	}
	return &taskModel, nil
}

func (s *Store) fromTaskModel(taskModel *model.Task) (*entity.Task, error) {
	task := &entity.Task{
		ID:          taskModel.UID,
		Name:        taskModel.Name,
		DependOn:    nil, // filled to later
		ActionName:  taskModel.ActionName,
		TimeoutSecs: int(*taskModel.TimeoutSecs),
		Params:      nil, // filled to later
		PreChecks:   nil, // filled to later
	}
	if err := s.Unmarshal([]byte(utils.StringPtrToVal(taskModel.Params, "")), &task.Params); err != nil {
		return nil, err
	}
	if err := s.Unmarshal([]byte(utils.StringPtrToVal(taskModel.Prechecks, "")), &task.PreChecks); err != nil {
		return nil, err
	}
	if err := s.Unmarshal([]byte(utils.StringPtrToVal(taskModel.DependOn, "")), &task.DependOn); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Store) toTaskInstanceModel(taskIns *entity.TaskInstance) (*model.TaskInstance, error) {
	taskInsModel := model.TaskInstance{
		UID:            taskIns.ID,
		TaskUID:        taskIns.TaskID,
		DagInstanceUID: taskIns.DagInsID,
		Name:           taskIns.Name,
		ActionName:     taskIns.ActionName,
		TimeoutSecs:    utils.IntToInt32Ptr(taskIns.TimeoutSecs),
		Status:         string(taskIns.Status),
		Reason:         utils.StringPtr(taskIns.Reason),
		CreatedAt:      time.Unix(taskIns.CreatedAt, 0),
		UpdatedAt:      time.Unix(taskIns.UpdatedAt, 0),
	}

	if err := s.marshalAndSet(taskIns.Params, &taskInsModel.Params); err != nil {
		return nil, err
	}
	if err := s.marshalAndSet(taskIns.PreChecks, &taskInsModel.Precheck); err != nil {
		return nil, err
	}
	if err := s.marshalAndSet(taskIns.DependOn, &taskInsModel.DependOn); err != nil {
		return nil, err
	}
	return &taskInsModel, nil
}

func (s *Store) fromTaskInstanceModel(taskInsModel *model.TaskInstance) (*entity.TaskInstance, error) {
	taskIns := &entity.TaskInstance{
		BaseInfo: entity.BaseInfo{
			ID:        taskInsModel.UID,
			CreatedAt: taskInsModel.CreatedAt.Unix(),
			UpdatedAt: taskInsModel.UpdatedAt.Unix(),
		},
		TaskID:      taskInsModel.TaskUID,
		DagInsID:    taskInsModel.DagInstanceUID,
		Name:        taskInsModel.Name,
		DependOn:    nil, // filled to later
		ActionName:  taskInsModel.ActionName,
		TimeoutSecs: int(*taskInsModel.TimeoutSecs),
		Params:      nil, // filled in later
		Traces:      nil,
		Status:      entity.TaskInstanceStatus(taskInsModel.Status),
		Reason:      utils.StringPtrToVal(taskInsModel.Reason, ""),
		PreChecks:   nil, // filled in later
		Patch:       nil, // TODO: implement
		Context:     nil, // TODO: implement
	}
	err := s.Unmarshal([]byte(utils.StringPtrToVal(taskInsModel.Params, "")), &taskIns.Params)
	if err != nil {
		return nil, err
	}
	err = s.Unmarshal([]byte(utils.StringPtrToVal(taskInsModel.Precheck, "")), &taskIns.PreChecks)
	if err != nil {
		return nil, err
	}
	return taskIns, nil
}

func (s *Store) toDagInstanceModel(dagIns *entity.DagInstance) (*model.DagInstance, error) {
	dagInsModel := model.DagInstance{
		UID:       dagIns.ID,
		DagUID:    dagIns.DagID,
		Trigger:   string(dagIns.Trigger),
		Worker:    dagIns.Worker,
		Status:    string(dagIns.Status),
		Reason:    utils.StringPtr(dagIns.Reason),
		CreatedAt: time.Unix(dagIns.CreatedAt, 0),
		UpdatedAt: time.Unix(dagIns.UpdatedAt, 0),
	}

	if err := s.marshalAndSet(dagIns.Cmd, &dagInsModel.Cmd); err != nil {
		return nil, err
	}

	if err := s.marshalAndSet(dagIns.Vars, &dagInsModel.Vars); err != nil {
		return nil, err
	}
	return &dagInsModel, nil
}

func (s *Store) fromDagInstanceModel(dagInsModel *model.DagInstance) (*entity.DagInstance, error) {
	dagIns := &entity.DagInstance{
		BaseInfo: entity.BaseInfo{
			ID:        dagInsModel.UID,
			CreatedAt: dagInsModel.CreatedAt.Unix(),
			UpdatedAt: dagInsModel.UpdatedAt.Unix(),
		},
		DagID:   dagInsModel.DagUID,
		Trigger: entity.Trigger(dagInsModel.Trigger),
		Worker:  dagInsModel.Worker,
		Vars:    nil, // filled in later
		Status:  entity.DagInstanceStatus(dagInsModel.Status),
		Reason:  utils.StringPtrToVal(dagInsModel.Reason, ""),
		Cmd:     nil, // filled in later
	}
	err := s.Unmarshal([]byte(utils.StringPtrToVal(dagInsModel.Vars, "")), &dagIns.Vars)
	if err != nil {
		return nil, err
	}
	err = s.Unmarshal([]byte(utils.StringPtrToVal(dagInsModel.Cmd, "")), &dagIns.Cmd)
	if err != nil {
		return nil, err
	}
	return dagIns, nil
}

func (s *Store) CreateDag(dag *entity.Dag) error {
	dagModel, err := s.toDagModel(dag)
	if err != nil {
		return err
	}

	taskModels := make([]*model.Task, 0, len(dag.Tasks))
	for _, task := range dag.Tasks {
		taskModel, err := s.toTaskModel(&task, dag.ID)
		if err != nil {
			return err
		}
		taskModels = append(taskModels, taskModel)
	}

	return s.query.Transaction(func(tx *query.Query) error {
		err := tx.Dag.WithContext(s.context).Create(dagModel)
		if err != nil {
			return err
		}
		return tx.Task.WithContext(s.context).Create(taskModels...)
	})

}

func (s *Store) CreateDagIns(dagIns *entity.DagInstance) error {
	dagInsModel, err := s.toDagInstanceModel(dagIns)
	if err != nil {
		return err
	}
	return s.query.DagInstance.WithContext(s.context).Create(dagInsModel)
}

func (s *Store) BatchCreatTaskIns(taskIns []*entity.TaskInstance) error {
	taskInsModels := make([]*model.TaskInstance, 0, len(taskIns))
	for _, taskIns := range taskIns {
		taskInsModel, err := s.toTaskInstanceModel(taskIns)
		if err != nil {
			return err
		}
		taskInsModels = append(taskInsModels, taskInsModel)
	}
	return s.query.TaskInstance.WithContext(s.context).Create(taskInsModels...)
}

// PatchTaskIns updates the task instance with the given ID (only non-zero fields will be updated).
func (s *Store) PatchTaskIns(taskIns *entity.TaskInstance) error {
	taskInsModel, err := s.toTaskInstanceModel(taskIns)
	if err != nil {
		return err
	}
	_, err = s.query.WithContext(s.context).TaskInstance.Where(s.query.TaskInstance.UID.Eq(taskIns.ID)).Updates(taskInsModel)
	if err != nil {
		return err
	}
	return nil
}

// PatchDagIns updates the dag instance with the given ID (only the specific fields will be updated).
func (s *Store) PatchDagIns(dagIns *entity.DagInstance, mustsPatchFields ...string) error {
	dagInsModel, err := s.toDagInstanceModel(dagIns)
	if err != nil {
		return err
	}

	updateFields := make([]field.Expr, 0, len(mustsPatchFields))
	for _, f := range mustsPatchFields {
		modelField, ok := s.query.DagInstance.GetFieldByName(f)
		if !ok {
			return fmt.Errorf("field %s not found", f)
		}
		updateFields = append(updateFields, modelField)
	}

	_, err = s.query.WithContext(s.context).DagInstance.Where(s.query.DagInstance.UID.Eq(dagIns.ID)).Select(
		updateFields...).Updates(dagInsModel)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDag updates the dag with the given ID (only the non-zero fields will be updated).
func (s *Store) UpdateDag(dag *entity.Dag) error {
	dagModel, err := s.toDagModel(dag)
	if err != nil {
		return err
	}

	_, err = s.query.WithContext(s.context).Dag.Where(s.query.Dag.UID.Eq(dag.ID)).Updates(dagModel)
	if err != nil {
		return err
	}

	return nil
}

// UpdateDagIns updates the dag instance with the given ID (only the non-zero fields will be updated).
func (s *Store) UpdateDagIns(dagIns *entity.DagInstance) error {
	dagInsModel, err := s.toDagInstanceModel(dagIns)
	if err != nil {
		return err
	}

	_, err = s.query.WithContext(s.context).DagInstance.Where(s.query.DagInstance.UID.Eq(dagIns.ID)).Updates(dagInsModel)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTaskIns updates the task instance with the given ID (only the non-zero fields will be updated).
func (s *Store) UpdateTaskIns(taskIns *entity.TaskInstance) error {
	taskInsModel, err := s.toTaskInstanceModel(taskIns)
	if err != nil {
		return err
	}

	_, err = s.query.WithContext(s.context).TaskInstance.Where(s.query.TaskInstance.UID.Eq(taskIns.ID)).Updates(taskInsModel)
	if err != nil {
		return err
	}

	return nil
}

// BatchUpdateDagIns batch updates the dag instance with the given IDs (only the non-zero fields will be updated).
func (s *Store) BatchUpdateDagIns(dagIns []*entity.DagInstance) error {
	dagInsModels := make([]*model.DagInstance, 0, len(dagIns))
	for _, dagIns := range dagIns {
		dagInsModel, err := s.toDagInstanceModel(dagIns)
		if err != nil {
			return err
		}
		dagInsModels = append(dagInsModels, dagInsModel)
	}

	return s.query.Transaction(func(tx *query.Query) error {
		for _, m := range dagInsModels {
			_, err := tx.WithContext(s.context).DagInstance.Updates(m)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// BatchUpdateTaskIns batch updates the task instance with the given IDs (only the non-zero fields will be updated).
func (s *Store) BatchUpdateTaskIns(taskIns []*entity.TaskInstance) error {
	taskInsModels := make([]*model.TaskInstance, 0, len(taskIns))
	for _, taskIns := range taskIns {
		taskInsModel, err := s.toTaskInstanceModel(taskIns)
		if err != nil {
			return err
		}
		taskInsModels = append(taskInsModels, taskInsModel)
	}

	return s.query.Transaction(func(tx *query.Query) error {
		for _, m := range taskInsModels {
			_, err := tx.WithContext(s.context).TaskInstance.Updates(m)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Store) GetTaskIns(taskInsID string) (*entity.TaskInstance, error) {
	m, err := s.query.WithContext(s.context).TaskInstance.Where(s.query.TaskInstance.UID.Eq(taskInsID)).First()
	if err != nil {
		return nil, err
	}

	taskInstance, err := s.fromTaskInstanceModel(m)
	if err != nil {
		return nil, err
	}
	return taskInstance, nil
}

func (s *Store) GetDag(dagId string) (*entity.Dag, error) {
	m, err := s.query.WithContext(s.context).Dag.Where(s.query.Dag.UID.Eq(dagId)).First()
	if err != nil {
		return nil, err
	}

	taskModels, err := s.query.WithContext(s.context).Task.Where(s.query.Task.DagUID.Eq(dagId)).Find()
	if err != nil {
		return nil, err
	}

	dag, err := s.fromDagModel(m)
	if err != nil {
		return nil, err
	}

	for _, taskModel := range taskModels {
		task, err := s.fromTaskModel(taskModel)
		if err != nil {
			return nil, err
		}
		dag.Tasks = append(dag.Tasks, *task)
	}
	return dag, nil
}

func (s *Store) GetDagInstance(dagInsId string) (*entity.DagInstance, error) {
	m, err := s.query.WithContext(s.context).DagInstance.Where(s.query.DagInstance.UID.Eq(dagInsId)).First()
	if err != nil {
		return nil, err
	}
	dagInstance, err := s.fromDagInstanceModel(m)
	if err != nil {
		return nil, err
	}
	dagInstance.ShareData = &entity.ShareData{} // TODO: load share data
	return dagInstance, nil
}

// ListDagInstance returns the dag instances filtered by the given conditions (only non-zero value will be used).
func (s *Store) ListDagInstance(input *mod.ListDagInstanceInput) ([]*entity.DagInstance, error) {
	query := &s.query.WithContext(s.context).DagInstance

	if input.DagID != "" {
		query = query.Where(s.query.DagInstance.DagUID.Eq(input.DagID))
	}
	if len(input.Status) != 0 {
		statusFilter := make([]string, 0, len(input.Status))
		for _, status := range input.Status {
			statusFilter = append(statusFilter, string(status))
		}
		query = query.Where(s.query.DagInstance.Status.In(statusFilter...))
	}
	if input.Worker != "" {
		query = query.Where(s.query.DagInstance.Worker.Eq(input.Worker))
	}
	if input.HasCmd {
		query = query.Where(s.query.DagInstance.Cmd.IsNotNull())
	}
	if input.UpdatedEnd != 0 {
		query = query.Where(s.query.DagInstance.UpdatedAt.Lte(time.Unix(input.UpdatedEnd, 0)))
	}

	dagInsModels, err := query.Order(s.query.DagInstance.CreatedAt.Desc()).Limit(int(input.Limit)).Offset(int(input.Offset)).Find()
	if err != nil {
		return nil, err
	}

	dagInsList := make([]*entity.DagInstance, 0, len(dagInsModels))
	for _, dagInsModel := range dagInsModels {
		dagIns, err := s.fromDagInstanceModel(dagInsModel)
		if err != nil {
			return nil, err
		}
		dagIns.ShareData = &entity.ShareData{} // TODO: load share data
		dagInsList = append(dagInsList, dagIns)
	}

	return dagInsList, nil
}

func (s *Store) ListTaskInstance(input *mod.ListTaskInstanceInput) ([]*entity.TaskInstance, error) {
	query := &s.query.WithContext(s.context).TaskInstance

	if len(input.IDs) != 0 {
		query = query.Where(s.query.TaskInstance.UID.In(input.IDs...))
	}
	if input.DagInsID != "" {
		query = query.Where(s.query.TaskInstance.DagInstanceUID.Eq(input.DagInsID))
	}
	if len(input.Status) != 0 {
		statusFilter := make([]string, 0, len(input.Status))
		for _, status := range input.Status {
			statusFilter = append(statusFilter, string(status))
		}
		query = query.Where(s.query.TaskInstance.Status.In(statusFilter...))
	}
	if input.Expired {
		query = query.Where(s.query.TaskInstance.UpdatedAt.LteCol(
			s.query.TaskInstance.TimeoutSecs.Sub(int32(time.Now().Unix() - 5)).Mul(int32(-1))))
	}

	taskInsModels, err := query.Find()
	if err != nil {
		return nil, err
	}

	taskInsList := make([]*entity.TaskInstance, 0, len(taskInsModels))
	for _, taskInsModel := range taskInsModels {
		taskIns, err := s.fromTaskInstanceModel(taskInsModel)
		if err != nil {
			return nil, err
		}
		taskInsList = append(taskInsList, taskIns)
	}

	return taskInsList, nil
}

func (s *Store) Marshal(obj interface{}) ([]byte, error) {
	if obj != nil {
		return json.Marshal(obj)
	}
	return []byte{}, nil
}

func (s *Store) Unmarshal(bytes []byte, ptr interface{}) error {
	if len(bytes) != 0 {
		return json.Unmarshal(bytes, ptr)
	}
	return nil
}

func (s *Store) Close() {
	return
}
