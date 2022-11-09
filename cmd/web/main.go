package main

import (
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/pkg/models"
	"main/pkg/models/sqlite"
	"net/http"
	"os"
)

type App struct {
	DBModel  *sqlite.DBModel
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP")
	dsn := flag.String("dsn", "farms.sqlite", "sqlite3")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			errorLog.Fatal("DB Fatal error")
		}
	}(db)
	_, err = db.Exec(models.FarmsSchema)
	if err != nil {
		return
	}
	app := App{
		DBModel:  &sqlite.DBModel{DB: db},
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск веб-сервера на %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return nil, err
	}
	return db, nil
}

//func printPlants() {
//	db := ConnectDB()
//	pm := sqlite.DBModel{DB: db}
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//
//		}
//	}(db)
//	plant, err := pm.GetAllFarmPlants(1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//fmt.Println(plant)
//	jsonData, _ := json.MarshalIndent(plant, "", "  ")
//	fmt.Println(string(jsonData))
//}
