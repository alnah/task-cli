-- name: CreateTask :one
INSERT INTO task (description)
VALUES ($1)
RETURNING *;

-- name: GetTaskForUpdate :one
SELECT * FROM task
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateTaskDescription :one
UPDATE task
SET description = $2
WHERE id = $1
RETURNING *;

-- name: UpdateTaskStatus :one
UPDATE task
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM task
WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM task
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTasksByStatus :many
SELECT * FROM task
WHERE status = $3
ORDER BY id
LIMIT $1
OFFSET $2;
