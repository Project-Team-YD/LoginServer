package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver       = "mysql"
	dbUser         = "root"
	dbPass         = "jaycci1@"
	dbMaxIdleConns = 10 // 최대 유휴 연결 수
	dbMaxOpenConns = 20 // 최대 열린 연결 수
)
const (
	loginDB   = "login_db"
	loginIp   = "13.125.254.231"
	loginPort = "3306"
	logDB     = "log_db"
	logIp     = "13.125.254.231"
	logPort   = "3306"

	maxIdleConnect = 10
	maxOpenConnect = 10
)

type GameDatabase struct {
	Database map[string]*sql.DB
}

var DBManager *GameDatabase

func StartDBConnection() {
	println("Start DB Connect!!")
	DBManager = &GameDatabase{
		Database: make(map[string]*sql.DB),
	}
	//-- Login DB Open
	login := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, loginIp, loginPort, loginDB)

	db, err := sql.Open(dbDriver, login)
	if err != nil {
		println("LoginDB::", err.Error())
	}
	defer db.Close()

	DBManager.Database["login"] = db

	//-- Log DB Open
	log := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, logIp, logPort, logDB)

	db, err = sql.Open(dbDriver, log)
	if err != nil {
		println("LogDB::", err.Error())
	}
	defer db.Close()

	DBManager.Database["log"] = db

	db.SetMaxIdleConns(maxIdleConnect)
	db.SetMaxOpenConns(maxOpenConnect)
	println("End DB Connect!!")
}
