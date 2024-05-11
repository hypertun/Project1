package main

import (
	"Project1/database"
	"Project1/handlers"
	"Project1/services"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbType = "postgres"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	dbConn, err := openDB()
	if err != nil {
		log.Fatalf("open repo: %v", err)
	}
	defer dbConn.Close()

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

	database := database.NewDatabase(dbConn)

	accountsService := services.NewAccountsService(database)
	accountsHandler := handlers.NewAccountsHandler(accountsService)

	transactionsService := services.NewTransactionsService(database)
	transactionsHandler := handlers.NewTransactionsHandler(transactionsService)

	version1Group := r.Group("/v1")
	{
		accountsGroup := version1Group.Group("/accounts")

		accountsGroup.POST("", accountsHandler.Post)
		accountsGroup.GET("/:account_id", accountsHandler.Get)

		transactionsGroup := version1Group.Group("/transactions")
		transactionsGroup.POST("", transactionsHandler.Post)
	}

	svr := &http.Server{
		Addr:              fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:           r,
		ReadHeaderTimeout: time.Second * 10,
	}

	// Start HTTP server.
	go func() {
		log.Printf("listening on port %s", os.Getenv("PORT"))

		if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	sig := <-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %v", err)
	}

	log.Printf("%s signal caught", sig)
	log.Print("server exited")
}

func openDB() (*sqlx.DB, error) {
	connstr := getConnectionString()

	db, err := sql.Open(dbType, connstr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	newConnstr := strings.Replace(connstr, dbType, "pgx", 1)

	m, err := migrate.New("file://migrations", newConnstr)
	if err != nil {
		return nil, err
	}

	defer m.Close()

	if err = m.Up(); err != nil &&
		!errors.Is(err, migrate.ErrNoChange) &&
		!errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	return sqlx.NewDb(db, dbType).Unsafe(), nil
}

func getConnectionString() string {
	connstr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?&connect_timeout=%d",
		dbType,
		os.Getenv("DB_USER"),
		url.QueryEscape(os.Getenv("DB_PASS")),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		2,
	)

	connstr = fmt.Sprintf("%s&sslmode=%s",
		connstr,
		"disable",
	)

	return connstr
}
