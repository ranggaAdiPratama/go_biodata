-- name: CheckAllHobby :many

SELECT * FROM hobbies WHERE user_id = $1 ORDER BY name;

-- name: CheckHobbyWithPage :many

SELECT *
FROM hobbies
WHERE user_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;

-- name: DeleteHobby :exec

DELETE FROM hobbies WHERE id = $1;

-- name: GetHobby :one

SELECT * FROM hobbies WHERE user_id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: CreateHobby :one

INSERT INTO hobbies (user_id, name) VALUES ($1, $2) RETURNING *;

-- name: UpdateHobby :one

UPDATE hobbies SET name = $2 WHERE id = $1 RETURNING *;