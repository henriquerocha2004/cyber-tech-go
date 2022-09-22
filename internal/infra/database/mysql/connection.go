package mysql

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type MysqlConnection struct {
	DbConnection *sqlx.DB
}

type MysqlConfig struct {
	User               string `json:"user"`
	Password           string `json:"password"`
	Port               string `json:"port"`
	Database           string `json:"db_name"`
	Host               string `json:"host"`
	AttemptsConnection int64  `json:"attempts_connection"`
}

var connectionAttempts int64

func NewMysqlConnection() *sqlx.DB {
	config := loadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	for {
		connection, err := connect(dsn)
		if err != nil {
			log.Println("Mysql is not ready")
			connectionAttempts++
		} else {
			log.Println("Connected to Mysql")
			return connection
		}

		if connectionAttempts > config.AttemptsConnection {
			log.Println(err)
			panic(err)
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}

}

func connect(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadConfig() *MysqlConfig {
	database, err := json.Marshal(viper.Get("database"))
	if err != nil {
		panic("failed to parse database config")
	}

	var mysqlConfig MysqlConfig
	err = json.Unmarshal(database, &mysqlConfig)
	if err != nil {
		panic("failed to parse database config")
	}

	return &mysqlConfig
}
