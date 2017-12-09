package main

import (
	"fmt"
	"flag"
	"net/http"

	"github.com/vlamug/scfg/config"
	"github.com/vlamug/scfg/api"
	"github.com/vlamug/scfg/storage"
	"github.com/vlamug/scfg/loader"
	"github.com/vlamug/scfg/cache"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// DEFAULT_LISTEN_ADDR is default api address
	DEFAULT_LISTEN_ADDR = ":9002"
)

// App is main struct of api service
type App struct {
	db *gorm.DB
	loaderService *loader.Loader

	router *mux.Router
}

// initDbConnection initializes database connections
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

// initRouter initializes api endpoints
func (app *App) initRouter() {
	app.router = mux.NewRouter()
	app.router.HandleFunc("/get", api.GetHandler(app.loaderService)).Methods("POST")
	app.router.Handle("/metrics", promhttp.Handler())
}

// initLoaderService initializes service of loading config
func (app *App) initLoaderService() {
	app.loaderService = loader.NewLoader(storage.NewPostgresStorage(app.db), cache.NewMem())
}

// listenAndServe listens and serve requests
func (app *App) listenAndServe(listenAddr string) {
	server := http.Server{Addr:listenAddr, Handler: app.router}
	panic(server.ListenAndServe())
}

// close closes database connection
func (app *App) close() {
	app.db.Close()
}

func main() {
	listenAddr := flag.String("api.listen-addr", DEFAULT_LISTEN_ADDR, "the address of api")
	dbConfigPath := flag.String("database.config-path", "etc/database.json", "the path to the database config")
	flag.Parse()

	app := App{}
	err := app.initDbConnection(*dbConfigPath)
	if err != nil {
		panic(err)
	}

	app.initLoaderService()
	app.initRouter()
	app.listenAndServe(*listenAddr)

	defer app.close()
}
