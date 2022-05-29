package main

import (
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	conn     string
	database string
)

func init() {
	flag.StringVar(&conn, "conn", "", "database connection string")
	flag.StringVar(&database, "db", "", "database name")
}

func validEmpty(key, val string) {
	if val == "" {
		panic(fmt.Sprintf("%s is empty", key))
	}
}

func main() {
	flag.Parse()

	validEmpty("conn", conn)
	validEmpty("db", database)

	config := gen.Config{
		OutPath:       "./query",
		FieldNullable: true,
	}
	config.WithDbNameOpts(func(db *gorm.DB) string {
		return database
	})

	g := gen.NewGenerator(config)

	db, err := gorm.Open(mysql.Open(conn))
	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModelAs("dag", "Dag"),
		g.GenerateModelAs("task", "Task"),
		g.GenerateModelAs("dag_instance", "DagInstance"),
		g.GenerateModelAs("task_instance", "TaskInstance"),
	)
	g.Execute()
}
