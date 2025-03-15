-- name: GetAdminAuthByEmail :one
SELECT *
FROM admin_auth
WHERE admin_auth.email = $1;

-- name: CreateModeratorAuth :exec
INSERT INTO moderator_auth (email, password, type)
VALUES ($1, $2, $3);