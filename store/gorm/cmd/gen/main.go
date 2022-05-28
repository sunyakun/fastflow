package main

import (
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	mysqlConn string
	database  string
)

func init() {
	flag.StringVar(&mysqlConn, "mysql", "", "mysql connection string")
	flag.StringVar(&database, "db", "", "database name")
}

func validEmpty(key, val string) {
	if val == "" {
		panic(fmt.Sprintf("%s is empty", key))
	}
}

func main() {
	flag.Parse()

	validEmpty("mysql", mysqlConn)
	validEmpty("db", database)

	dbnameOpt := func(db *gorm.DB) string {
		return database
	}

	config := gen.Config{
		OutPath:          "./query",
		FieldWithTypeTag: true,
		FieldNullable:    true,
	}
	config.WithDbNameOpts(dbnameOpt)

	g := gen.NewGenerator(config)

	db, err := gorm.Open(mysql.Open(mysqlConn))
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
