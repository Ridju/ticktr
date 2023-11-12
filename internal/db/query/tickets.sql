-- name: CreateTicket :one
INSERT INTO tickets(
    title,
    description,
    assigned_to,
    created_by,
    due_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTicketByID :one
SELECT * FROM tickets 
WHERE id = $1 LIMIT 1;

-- name: GetTicketForUpdate :one
SELECT * FROM tickets 
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTickets :many
SELECT * FROM tickets 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTicketsForUser :many
SELECT * FROM tickets 
WHERE assigned_to = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTicketsByUser :many
SELECT * FROM tickets 
WHERE created_by = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTicket :one
UPDATE tickets 
SET title = $2,
    description = $3,
    assigned_to = $4,
    created_by = $5,
    due_date = $6
WHERE id = $1
RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets 
WHERE id = $1;