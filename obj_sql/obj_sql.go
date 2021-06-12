package obj_sql

import (
	"SplatStatsGo/objects"
	"SplatStatsGo/secrets"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initClient() error {
	var err error
	db, err = sql.Open("mysql", "django_database_user:"+secrets.GetSecret("DJANGO_DATABASE_PASSWORD")+"@mysql://35.224.168.252/db")
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	return err
}

func ReadObjWithId(id int, objectTable string) interface{} {
	if db == nil {
		if err := initClient(); err != nil {
			panic(err)
		}
	}
	return db.QueryRow("SELECT * FROM " + objectTable + "where id = " + fmt.Sprint(id) + ";")
}

func WriteNewObj(object interface{}, objectTable string) {
	if db == nil {
		if err := initClient(); err != nil {
			panic(err)
		}
	}
	var newId int
	err := db.QueryRow("SELECT id FROM " + objectTable + " ORDER BY id DESC LIMIT 1;").Scan(&newId)
	if err != nil {
		panic(err)
	}
	newId += 1
	switch objectTable {
	case "two_battles_battle":
		WriteNewBattle(object.(objects.Battle))
	case "two_salmon_shift":
		WriteNewShift(object.(objects.Shift))
	}
}

func WriteNewBattle(battle objects.Battle) {

}

func WriteNewShift(shift objects.Shift) {

}
