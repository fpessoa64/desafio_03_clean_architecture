package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/fpessoa64/desafio03_clean_arch/internal/config"
	mysqlrepo "github.com/fpessoa64/desafio03_clean_arch/internal/repository/mysql"
	"github.com/fpessoa64/desafio03_clean_arch/internal/servers"
	"github.com/fpessoa64/desafio03_clean_arch/internal/usecase"
)

func main() {
	cfg := config.Load()
	if cfg.DBDSN == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	db := waitForDB(cfg.DBDSN, 30, 2*time.Second)
	defer db.Close()

	runMigrations(cfg.DBDSN)

	repo := mysqlrepo.NewOrderRepositoryMySQL(db)
	uc := usecase.NewOrderUsecase(repo)

	go func() {
		if err := servers.NewGrpc(cfg.GRPCPort).Start(uc); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	go func() {
		if err := servers.NewRest(cfg.RESTPort).Start(uc); err != nil {
			log.Fatalf("REST server error: %v", err)
		}
	}()

	if err := servers.NewGraphQL(cfg.GraphQLPort).Start(uc); err != nil {
		log.Fatalf("GraphQL server error: %v", err)
	}
}

func waitForDB(dsn string, maxAttempts int, delay time.Duration) *sql.DB {
	for i := 0; i < maxAttempts; i++ {
		db, err := sql.Open("mysql", dsn)
		if err == nil {
			if pingErr := db.Ping(); pingErr == nil {
				log.Println("connected to database")
				return db
			}
			db.Close()
		}
		log.Printf("waiting for database... attempt %d/%d", i+1, maxAttempts)
		time.Sleep(delay)
	}
	log.Fatal("could not connect to database after maximum attempts")
	return nil
}

func runMigrations(dsn string) {
	m, err := migrate.New("file://migrations", "mysql://"+dsn)
	if err != nil {
		log.Fatalf("failed to create migrator: %v", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("migrations applied successfully")
}
