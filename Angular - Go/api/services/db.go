package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	DB struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Hostname string `yaml:"hostname"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"db"`
}

func getDbConnection() *sql.DB {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Hostname,
		config.DB.Dbname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
