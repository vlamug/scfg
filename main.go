package main

import (
	"net/http"
	"fmt"

	"github.com/vlamug/scfg/config"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"database/sql"
	"flag"
)

type App struct {
	db *sql.DB

	router *mux.Router
}

func (app *App) initDbConnection(configPath string) error {
	cfg, err := config.LoadDatabaseConfig(configPath)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DbName,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	app.db = conn

	return nil
}

func (app *App) initRouter() {
	app.router = mux.NewRouter()
	app.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("yes")
	})
}

func (app *App) listenAndServe() {
	server := http.Server{Addr:":9001", Handler: app.router}
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
