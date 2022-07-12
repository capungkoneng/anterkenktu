-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori,
  deskripsi
) VALUES (
  $1, $2
) RETURNING *;

-- name: ListKategori :many
SELECT * FROM kategori
ORDER BY id 
LIMIT $1
OFFSET $2;