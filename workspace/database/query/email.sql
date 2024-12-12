-- name: CreateEmail :one
INSERT INTO email(title,body,status)
VALUES ($1,$2,$3)
RETURNING *;

-- name: ListEmails :many
SELECT * FROM email
        WHERE status = $1
        ORDER BY id
        LIMIT $2
        offset $3;


-- name: GetEmail :one
SELECT * FROM email
WHERE id = $1 LIMIT 1;


-- name: DeleteBook :exec
DELETE FROM email
WHERE id = $1;

-- name: UpdateEmail :one
UPDATE email
SET title = $2,
    body = $3,
    status = $4
WHERE id = $1
RETURNING *;