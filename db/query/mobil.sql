-- name: GetMobilJoinMany :many
select m.nama, k.nama_kategori, u.username, gm.url  
from mobil m
inner join kategori k on m.kategori_id  = k.id
inner join users u on u.username = m.user_id 
inner join gambar_mobil gm on gm.mobil_id  = m.id
where m.nama LIKE $1
ORDER BY created_at DESC
limit $2
Offset $3;