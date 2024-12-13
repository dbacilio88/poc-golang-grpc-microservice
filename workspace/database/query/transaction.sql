-- name: CreateTransaction :one
INSERT INTO transaction(email_id,status)
VALUES ($1,$2)
RETURNING *;

-- name: ListTransactions :many
SELECT * FROM transaction
        WHERE email_id = $1
        ORDER BY id
        LIMIT $2
        OFFSET $3;


-- name: GetTransaction :one
SELECT * FROM transaction
WHERE id = $1 LIMIT 1;