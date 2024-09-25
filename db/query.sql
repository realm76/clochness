-- name: GetEntry :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;

-- name: ListEntries :many
SELECT *
FROM entries
ORDER BY created_at DESC;

-- name: CreateEntry :one
INSERT INTO entries (
                     user_id,
                     project_id,
                     description,
                     start_date,
                     end_date
)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
       )
RETURNING *;

-- name: UpdateEntry :exec
UPDATE entries
set user_id = $2,
    project_id  = $3,
    description = $4,
    start_date = $5,
    end_date = $6
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE
FROM entries
WHERE id = $1;
