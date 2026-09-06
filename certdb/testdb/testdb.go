package testdb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	mysqlTruncateTables = `
TRUNCATE certificates;
TRUNCATE ocsp_responses;
`

	pgTruncateTables = `
CREATE OR REPLACE FUNCTION truncate_tables() RETURNS void AS $$
DECLARE
    statements CURSOR FOR
        SELECT tablename FROM pg_tables
        WHERE tablename != 'goose_db_version'
          AND tableowner = session_user
          AND schemaname = 'public';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT truncate_tables();
`

	sqliteTruncateTables = `
DELETE FROM certificates;
DELETE FROM ocsp_responses;
`
)

func MySQLDB() *sqlx.DB { _ = "STUB: not implemented"; return nil }

func PostgreSQLDB() *sqlx.DB { _ = "STUB: not implemented"; return nil }

func SQLiteDB(dbpath string) *sqlx.DB { _ = "STUB: not implemented"; return nil }

func Truncate(db *sqlx.DB) { _ = "STUB: not implemented"; return }
