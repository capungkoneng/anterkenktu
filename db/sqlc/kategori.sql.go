// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: kategori.sql

package db

import (
	"context"
)

const createKategori = `-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori,
  deskripsi
) VALUES (
  $1, $2
) RETURNING id, nama_kategori, deskripsi, created_at
`

type CreateKategoriParams struct {
	NamaKategori string `json:"nama_kategori"`
	Deskripsi    string `json:"deskripsi"`
}

func (q *Queries) CreateKategori(ctx context.Context, arg CreateKategoriParams) (Kategori, error) {
	row := q.db.QueryRowContext(ctx, createKategori, arg.NamaKategori, arg.Deskripsi)
	var i Kategori
	err := row.Scan(
		&i.ID,
		&i.NamaKategori,
		&i.Deskripsi,
		&i.CreatedAt,
	)
	return i, err
}

const listKategori = `-- name: ListKategori :many
SELECT id, nama_kategori, deskripsi, created_at FROM kategori
ORDER BY id 
LIMIT $1
OFFSET $2
`

type ListKategoriParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListKategori(ctx context.Context, arg ListKategoriParams) ([]Kategori, error) {
	rows, err := q.db.QueryContext(ctx, listKategori, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Kategori{}
	for rows.Next() {
		var i Kategori
		if err := rows.Scan(
			&i.ID,
			&i.NamaKategori,
			&i.Deskripsi,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
