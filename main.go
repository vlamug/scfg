package main

import (
	"fmt"
	"flag"
	"net/http"

	"github.com/vlamug/scfg/config"
	"github.com/vlamug/scfg/api"
	"github.com/vlamug/scfg/storage"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

type App struct {
	db *gorm.DB

	router *mux.Router
}

func (app *App) initDbConnection(configPath string) error {
	cfg, err := config.LoadDatabaseConfig(configPath)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DbName,
	)

	app.db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return nil
}

func (app *App) initRouter() {
	app.router = mux.NewRouter()
	app.router.HandleFunc("/get", api.GetHandler(storage.NewPostgresStorage(app.db))).Methods("POST")
}

func (app *App) listenAndServe() {
	server := http.Server{Addr:":9002", Handler: app.router}
	panic(server.ListenAndServe())
}

func (app *App) close() {
	app.db.Close()
}

func main() {
	dbConfigPath := flag.String("database.config-path", "etc/database.json", "The path to the database config")
	flag.Parse()

	app := App{}
	err := app.initDbConnection(*dbConfigPath)
	if err != nil {
		panic(err)
	}

	app.initRouter()
	app.listenAndServe()

	defer app.close()
}
