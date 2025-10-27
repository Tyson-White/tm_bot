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

	user, port, psw, db := pg.mustConnectionParams()

	connStr := fmt.Sprintf("user=%v port=%v password=%v dbname=%v sslmode=disable", user, port, psw, db)

	d, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	pg.DB = d

	return d
}

func (pg *PostgresProvider) mustConnectionParams() (string, string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Postgres env inititialize error", err)
	}

	user, uEx := os.LookupEnv("DB_USER")
	db, dbEx := os.LookupEnv("DB_NAME")
	psw, pswEx := os.LookupEnv("DB_PASSWORD")
	port, portEx := os.LookupEnv("DB_PORT")
	
	if !uEx || !dbEx || !pswEx {
		log.Fatalln("Missed env variables")
	}

	if user == "" || db == "" || psw == "" {
		log.Fatalln("Variables can't be empty")
	}

	if port == "" || !portEx {
		port = "5432"
	}

	return user, port, psw, db
}
