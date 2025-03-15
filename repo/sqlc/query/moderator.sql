-- name: GetAdminModeratorByEmail :one
SELECT *
FROM moderator_auth
WHERE moderator_auth.email = $1;