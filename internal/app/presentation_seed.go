package app

import (
	"context"
	"database/sql"
	_ "embed"
	"os"
	"time"
)

//go:embed presentation_seed.sql
var presentationSeed string

// SeedPresentation adds synthetic scenarios once, without replacing existing rows.
// It is restricted to the original demo installation and never runs in production mode.
func (a *App) SeedPresentation(ctx context.Context) error {
	if os.Getenv("DEMO_MODE") != "true" {
		return nil
	}
	if count(ctx, a.DB, `SELECT count(*) FROM usuario WHERE
	 (email='gerente@demo.local' AND papel='gerente' AND empresa_id=1) OR
	 (email='atendente@demo.local' AND papel='atendente') OR
	 (email='passageiro@demo.local' AND papel='passageiro' AND passageiro_id=1)`) != 3 {
		return nil
	}
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS demo_seed_version(version TEXT PRIMARY KEY,aplicado_em TEXT NOT NULL)`); err != nil {
		return err
	}
	var applied int
	if err = tx.QueryRowContext(ctx, `SELECT count(*) FROM demo_seed_version WHERE version='apresentacao_v1'`).Scan(&applied); err != nil {
		return err
	}
	if applied > 0 {
		return tx.Commit()
	}
	now := time.Now().In(a.Location)
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, a.Location).UTC().Format(time.RFC3339)
	if _, err = tx.ExecContext(ctx, presentationSeed, sql.Named("demo_day", tomorrow)); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, `INSERT INTO demo_seed_version VALUES('apresentacao_v1',?)`, a.now()); err != nil {
		return err
	}
	return tx.Commit()
}
