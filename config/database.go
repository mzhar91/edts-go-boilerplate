package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
	
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (dbConn *sql.DB) {
	var dsn string
	
	dbProvider := Cfg.Database.Provider
	dbHost := Cfg.Database.Host
	dbPort := &Cfg.Database.Port
	dbUser := Cfg.Database.User
	dbPassword := Cfg.Database.Password
	dbName := Cfg.Database.Name
	
	if dbProvider == "postgres" {
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName,
		)
	} else if dbProvider == "mysql" {
		connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
		val := url.Values{}
		val.Add("parseTime", "1")
		val.Add("charset", "utf8")
		dsn = fmt.Sprintf("%s?%s", connection, val.Encode())
	} else if dbProvider == "pgx" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d database=%s", dbHost, dbUser, dbPassword, dbPort, dbName)
		
		// Envar Get Env
		mustGetenv := func(k string) string {
			v := os.Getenv(k)
			if v == "" {
				log.Fatalf("Fatal Error: %s environment variable not set.", k)
			}
			return v
		}
		
		// connection is encrypted.
		if dbRootCert, ok := os.LookupEnv("DB_ROOT_CERT"); ok { // e.g., '/path/to/my/server-ca.pem'
			var (
				dbCert = mustGetenv("DB_CERT") // e.g. '/path/to/my/client-cert.pem'
				dbKey  = mustGetenv("DB_KEY")  // e.g. '/path/to/my/client-key.pem'
			)
			dsn += fmt.Sprintf(" sslmode=verify-ca sslrootcert=%s sslcert=%s sslkey=%s", dbRootCert, dbCert, dbKey)
		}
	}
	
	dbConn, err := sql.Open(dbProvider, dsn)
	if err != nil && Cfg.Debug {
		fmt.Println(err)
	} else {
		dbConn.SetMaxIdleConns(10)
		dbConn.SetMaxOpenConns(100)
		dbConn.SetConnMaxLifetime(time.Minute * 4)
	}
	
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	// defer dbConn.Close()
	
	return
}
