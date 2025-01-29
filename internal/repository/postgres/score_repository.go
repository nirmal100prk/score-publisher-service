package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXDatabase struct {
	Db *pgxpool.Pool
}

type PgxRepository interface {
	InsertScore(ctx context.Context, point int64) error
}

func NewPgxRepository(PgServices *PGXDatabase) PgxRepository {
	return &PGXDatabase{
		Db: PgServices.Db,
	}
}

func (r *PGXDatabase) InsertScore(ctx context.Context, point int64) error {

	query := `INSERT INTO scores (score)  VALUES (  $1)`
	res, err := r.Db.Exec(ctx, query, point)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("query failed %v", "0 rows affected")
	}

	return nil
}
