package xdb

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//Engine model sqlite
var Engine *xorm.Engine

//EngineTask for task db
var EngineTask *xorm.Engine

//EngineWarning for warning db
var EngineWarning *xorm.Engine

// Init xdb
func Init(showSQL bool) {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./model.db")
	if err != nil {
		panic(err)
	}

	EngineTask, err = xorm.NewEngine("sqlite3", "./task.db")

	EngineWarning, err = xorm.NewEngine("sqlite3", "./warning.db")

	Engine.ShowSQL(showSQL)
	EngineTask.ShowSQL(showSQL)
	EngineWarning.ShowSQL(showSQL)

}
