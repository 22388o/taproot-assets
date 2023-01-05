// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: macaroons.sql

package sqlc

import (
	"context"
)

const getRootKey = `-- name: GetRootKey :one
SELECT id, root_key FROM macaroons 
WHERE id = $1
`

func (q *Queries) GetRootKey(ctx context.Context, id []byte) (Macaroon, error) {
	row := q.db.QueryRowContext(ctx, getRootKey, id)
	var i Macaroon
	err := row.Scan(&i.ID, &i.RootKey)
	return i, err
}

const insertRootKey = `-- name: InsertRootKey :exec
INSERT INTO macaroons (id, root_key) VALUES ($1, $2)
`

type InsertRootKeyParams struct {
	ID      []byte
	RootKey []byte
}

func (q *Queries) InsertRootKey(ctx context.Context, arg InsertRootKeyParams) error {
	_, err := q.db.ExecContext(ctx, insertRootKey, arg.ID, arg.RootKey)
	return err
}
