package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresProvider struct {
	DB *sqlx.DB
}

func New() PostgresProvider {
	pg := PostgresProvider{}

	return pg
}

func (pg *PostgresProvider) Connect() *sqlx.DB {

	user, psw, db := pg.mustConnectionParams()

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", user, psw, db)

	d, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	pg.DB = d

	return d
}

func (pg *PostgresProvider) mustConnectionParams() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Postgres env inititialize error", err)
	}

	user, uEx := os.LookupEnv("DB_USER")
	db, dbEx := os.LookupEnv("DB_NAME")
	psw, pswEx := os.LookupEnv("DB_PASSWORD")

	if !uEx || !dbEx || !pswEx {
		log.Fatalln("Missed env variables")
	}

	if user == "" || db == "" || psw == "" {
		log.Fatalln("Variables can't be empty")
	}

	return user, psw, db
}
