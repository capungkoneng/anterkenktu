-- name: ListMobilNew :many
SELECT m.id, m.nama, gm.url, gm.id  FROM Mobil m
inner join gambar_mobil gm on gm.mobil_id  = m.id
ORDER BY gm.id, m.id;

-- name: ListMobilNewB :many
SELECT * FROM Mobil m
ORDER BY id;