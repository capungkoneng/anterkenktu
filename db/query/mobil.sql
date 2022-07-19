-- name: GetMobilJoinMany :many
select m.nama, k.nama_kategori, u.username, gm.url  
from mobil m
inner join kategori k on m.kategori_id  = k.id
inner join users u on u.username = m.user_id 
inner join gambar_mobil gm on gm.mobil_id  = m.id
where m.nama LIKE $1
ORDER BY m.id DESC
limit $2
Offset $3;

-- name: CreateMobil :one
INSERT INTO mobil (
  nama, 
  deskripsi,
  kategori_id,
  user_id,
  gambar,
  trf_6jam,
  trf_12jam,
  trf_24jam,
  seat,
  top_speed,
  max_power,
  pintu,
  gigi
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;