// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: mobil.sql

package db

import (
	"context"
	"database/sql"
)

const getMobilJoinMany = `-- name: GetMobilJoinMany :many
select m.nama, k.nama_kategori, u.username, gm.url  
from mobil m
inner join kategori k on m.kategori_id  = k.id
inner join users u on u.username = m.user_id 
inner join gambar_mobil gm on gm.mobil_id  = m.id
where m.nama LIKE $1
ORDER BY created_at DESC
limit $2
Offset $3
`

type GetMobilJoinManyParams struct {
	Nama   string `json:"nama"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetMobilJoinManyRow struct {
	Nama         string         `json:"nama"`
	NamaKategori string         `json:"nama_kategori"`
	Username     string         `json:"username"`
	Url          sql.NullString `json:"url"`
}

func (q *Queries) GetMobilJoinMany(ctx context.Context, arg GetMobilJoinManyParams) ([]GetMobilJoinManyRow, error) {
	rows, err := q.db.QueryContext(ctx, getMobilJoinMany, arg.Nama, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMobilJoinManyRow{}
	for rows.Next() {
		var i GetMobilJoinManyRow
		if err := rows.Scan(
			&i.Nama,
			&i.NamaKategori,
			&i.Username,
			&i.Url,
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
