package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "tie-explorer-go/routers"
	"tie-explorer-go/task"
)

func Initialize() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error("Failed to set the database engine, %s", err.Error())
		return
	}

	dbName := beego.AppConfig.String("postgres::db")
	dbHost := beego.AppConfig.String("postgres::host")
	dbUser := beego.AppConfig.String("postgres::user")
	dbPass := beego.AppConfig.String("postgres::pass")
	dbPort := beego.AppConfig.String("postgres::port")
	dbSsl := beego.AppConfig.String("postgres::ssl")

	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSsl)
	fmt.Println(dns)
	err = orm.RegisterDataBase("default", "postgres", dns)
	if err != nil {
		panic(fmt.Errorf("database connection failure, %s", err.Error()))
		return
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("Failed to synchronize database tables, %s", err.Error())
		return
	}
}

func main() {
	//orm.Debug = true

	// Initialize DB
	Initialize()

	// Scheduler
	task.StartTask()
	defer task.StopTask()

	// Api Server
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}
