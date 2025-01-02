package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Pool *sql.DB

func InitMySQL(dsn string) error {
	db, err := sql.Open("msql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}
	err = initTables(db)
	if err != nil {
		return fmt.Errorf("error initializing tables: %v", err)
	}
	Pool = db
	return nil
}

func initTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Heroes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		description TEXT,
		type TINYINT,
		role TINYINT,
		action_list TEXT,
		hero_image TEXT,
		icon_image TEXT,
		banner_image TEXT,
		stats_res INT,
		P_Atk INT,
		A_Atk INT,
		W_Atk INT,
		P_Def INT,
		A_Def INT,
		W_Def INT,
		P_Boost INT,
		A_Boost INT,
		W_Boost INT,
		Speed INT,
		ActionDice INT
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating Heroes table: %v", err)
	}

	query = `CREATE TABLE IF NOT EXISTS Actions (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
	);`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating Actions table: %v", err)
	}
	return nil
}

func HandleTermination(db *sql.DB) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Terminating...")
		db.Close()
		os.Exit(0)
	}()
}
