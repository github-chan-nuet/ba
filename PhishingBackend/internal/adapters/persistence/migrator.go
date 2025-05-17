package persistence

import (
	"context"
	"embed"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
	"io/fs"
)

//go:embed postgresSchemaMigration/*.sql
var postgresSchemaFS embed.FS

func migrateDatabaseSchema() error {
	config := newDbConfig()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.getPostgresConnString())
	if err != nil {
		return err
	}
	m, err := migrate.NewMigrator(ctx, conn, "schema_version")
	if err != nil {
		return err
	}
	// migrate.FindMigrations in LoadMigrations scans dir non-recursively
	schema, err := fs.Sub(postgresSchemaFS, "postgresSchemaMigration")
	if err != nil {
		return err
	}
	if err := m.LoadMigrations(schema); err != nil {
		return err
	}
	return m.Migrate(ctx)
}
