package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/shiningrush/fastflow/pkg/entity"
	"github.com/shiningrush/fastflow/pkg/mod"
	"github.com/shiningrush/fastflow/store/gorm/model"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Dag{}, &model.DagInstance{}, &model.Task{}, &model.TaskInstance{})
	assert.NoError(t, err)

	return db
}

func CreateTestDag(_ *testing.T, dagID string) *entity.Dag {
	dag := entity.Dag{
		BaseInfo: entity.BaseInfo{
			ID: dagID,
		},
		Status: entity.DagStatusNormal,
		Name:   "custom-dag",
		Desc:   "custom",
		Tasks: []entity.Task{
			{ID: "task-1", ActionName: "CustomAction", Params: map[string]interface{}{
				"name": "{{name}}",
			}},
			{ID: "task-2", ActionName: "CustomAction", DependOn: []string{"task-1"}, Params: map[string]interface{}{
				"name": "{{name}}",
			}},
			{ID: "task-3", ActionName: "CustomAction", DependOn: []string{"task-2"}},
		},
		Vars: entity.DagVars{
			"name": entity.DagVar{
				Desc:         "name",
				DefaultValue: "unknown",
			},
		},
	}
	return &dag
}

func TestStore(t *testing.T) {
	var err error
	db := SetupTestDB(t)
	store := NewStore(db)

	dagID := "dag-1"
	dag := CreateTestDag(t, "dag-1")

	// create and get dag
	err = store.CreateDag(dag)
	assert.NoError(t, err)
	dagFromDB, err := store.GetDag(dagID)
	assert.NoError(t, err)
	assert.True(t, assert.ObjectsAreEqualValues(dag, dagFromDB))

	// create and get dag instance
	dagIns, err := dag.Run(entity.TriggerManually, map[string]string{})
	assert.NoError(t, err)
	dagIns.ID = "dagins-1"
	assert.NoError(t, store.CreateDagIns(dagIns))
	dagInsFromDB, err := store.GetDagInstance("dagins-1")
	assert.NoError(t, err)
	assert.True(t, assert.ObjectsAreEqualValues(dagIns, dagInsFromDB))
	// list dag instance
	dagInstanceList, err := store.ListDagInstance(&mod.ListDagInstanceInput{
		DagID: dag.ID,
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dagInstanceList))
	assert.True(t, assert.ObjectsAreEqualValues(dagIns, dagInstanceList[0]))

	// create and get task instance
	taskIns := &entity.TaskInstance{
		BaseInfo: entity.BaseInfo{
			ID: "taskins-1",
		},
		TaskID:      dag.Tasks[0].ID,
		DagInsID:    dagIns.ID,
		Name:        dag.Tasks[0].Name,
		DependOn:    dag.Tasks[0].DependOn,
		ActionName:  dag.Tasks[0].ActionName,
		TimeoutSecs: dag.Tasks[0].TimeoutSecs,
		Params:      dag.Tasks[0].Params,
		Traces:      nil,
		Status:      entity.TaskInstanceStatusInit,
		Reason:      "reason",
		PreChecks:   dag.Tasks[0].PreChecks,
		Patch:       nil,
		Context:     nil,
	}
	assert.NoError(t, store.BatchCreatTaskIns([]*entity.TaskInstance{taskIns}))
	taskInsList, err := store.ListTaskInstance(&mod.ListTaskInstanceInput{
		IDs: []string{"taskins-1"},
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(taskInsList))
	assert.True(t, assert.ObjectsAreEqualValues(taskIns, taskInsList[0]))
}

func TestStore_ListDagInstance(t *testing.T) {

}

func TestStore_ListTaskInstance(t *testing.T) {

}

func TestStore_UpdateAndPatch(t *testing.T) {

}
