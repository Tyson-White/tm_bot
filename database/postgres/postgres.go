package postgres

type PostgresDB struct{}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (pg *PostgresDB) Connect() {}
