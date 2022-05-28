package main

import (
	"fmt"
	"time"

	"github.com/shiningrush/fastflow"
	mongoKeeper "github.com/shiningrush/fastflow/keeper/mongo"
	"github.com/shiningrush/fastflow/pkg/entity"
	"github.com/shiningrush/fastflow/pkg/entity/run"
	"github.com/shiningrush/fastflow/pkg/log"
	"github.com/shiningrush/fastflow/pkg/mod"
	mongoStore "github.com/shiningrush/fastflow/store/mongo"
)

type CustomAction struct {
}

func (a *CustomAction) Name() string {
	return "CustomAction"
}

func (a *CustomAction) Run(ctx run.ExecuteContext, params interface{}) error {
	fmt.Println("action start: ", time.Now())
	fmt.Println(params)
	return nil
}

func (a *CustomAction) ParameterNew() interface{} {
	return &struct{ Name, Address string }{}
}

func main() {
	// Register action
	fastflow.RegisterAction([]run.Action{
		&CustomAction{},
	})

	// init keeper, it used to e
	keeper := mongoKeeper.NewKeeper(&mongoKeeper.KeeperOption{
		Key: "worker-1",
		// if your mongo does not set user/pwd, you should remove it
		ConnStr:  "mongodb://admin:pwd@localhost:27017/fastflow?authSource=admin",
		Database: "fastflow",
		Prefix:   "test",
	})
	if err := keeper.Init(); err != nil {
		log.Fatal("init keeper failed: %w", err)
	}

	// init store
	st := mongoStore.NewStore(&mongoStore.StoreOption{
		// if your mongo does not set user/pwd, you should remove it
		ConnStr:  "mongodb://admin:pwd@localhost:27017/?authSource=admin",
		Database: "fastflow",
		Prefix:   "test",
	})
	if err := st.Init(); err != nil {
		log.Fatal("init store failed: %w", err)
	}

	go func() {
		dag := entity.Dag{
			BaseInfo: entity.BaseInfo{
				ID: "dag-1",
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
		if err := mod.GetStore().UpdateDag(&dag); err != nil {
			err = mod.GetStore().CreateDag(&dag)
			if err != nil {
				log.Fatal("create dag failed: %w", err)
			}
		}

		dagins, err := dag.Run(entity.TriggerManually, map[string]string{
			"name": "shining",
		})
		if err != nil {
			log.Fatal("run dag failed: %w", err)
		}

		if err := mod.GetStore().CreateDagIns(dagins); err != nil {
			log.Fatal("create dag ins failed: %w", err)
		}
	}()

	// start fastflow
	if err := fastflow.Start(&fastflow.InitialOption{
		Keeper: keeper,
		Store:  st,
	}); err != nil {
		panic(fmt.Sprintf("init fastflow failed: %s", err))
	}
}
