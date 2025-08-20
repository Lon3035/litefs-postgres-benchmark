package db

import (
	"context"
	"database/sql"
	"regexp"
	"time"

	_ "github.com/lib/pq"
)

// Database is a generic interface for SQL databases.
type Database interface {
	Connect(connString string) error
	Close() error
	Query(query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
}

// PostgresDB implements Database for PostgreSQL.
type PostgresDB struct {
	DB *sql.DB
}

func (p *PostgresDB) Connect(connString string) error {
	var lastErr error
	baseWait := 5 * time.Second
	wait := baseWait
	for i := 0; i < 3; i++ {
		db, err := sql.Open("postgres", connString)
		if err == nil {
			// Teste die Verbindung
			if err = db.Ping(); err == nil {
				p.DB = db
				return nil
			}
			db.Close()
		}
		lastErr = err
		time.Sleep(wait)
		wait *= 2
	}
	return lastErr
}

func (p *PostgresDB) Close() error {
	if p.DB != nil {
		return p.DB.Close()
	}
	return nil
}

func (p *PostgresDB) Query(query string, args ...any) (*sql.Rows, error) {
	return p.DB.Query(query, args...)
}

func (p *PostgresDB) Exec(query string, args ...any) (sql.Result, error) {
	return p.DB.Exec(query, args...)
}

func (p *PostgresDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.DB.ExecContext(ctx, query, args...)
}

// SqliteDB implements Database for SQLite.
type SqliteDB struct {
	DB *sql.DB
}

func (s *SqliteDB) Connect(connString string) error {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *SqliteDB) Close() error {
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}

func (s *SqliteDB) Query(query string, args ...any) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SqliteDB) Exec(query string, args ...any) (sql.Result, error) {
	query = replacePostgresPlaceholders(query)
	return s.DB.Exec(query, args...)
}

func (s *SqliteDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	query = replacePostgresPlaceholders(query)
	return s.DB.ExecContext(ctx, query, args...)
}

var pgPlaceholderRegexp = regexp.MustCompile(`\$\d+`)

func replacePostgresPlaceholders(query string) string {
	return pgPlaceholderRegexp.ReplaceAllString(query, "?")
}
