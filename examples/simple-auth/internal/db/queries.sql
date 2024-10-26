-- Queries for User Management

-- name: CreateUser :one
INSERT INTO user (email, password, metadata)
VALUES (?, ?, ?)
RETURNING id;

-- name: CreateSession :exec
INSERT INTO sessions (user_id, session_id, expires_at)
VALUES (?, ?, ?);

-- name: GetUserByToken :one
SELECT u.*
FROM user u
         JOIN sessions t ON u.id = t.user_id
WHERE t.session_id = ?
  AND t.expires_at > datetime('now');

-- name: GetUserByID :one
SELECT *
FROM user
WHERE id = ?;


-- name: GetUserByEmail :one
SELECT *
FROM user
WHERE email = ?;

-- name: UpdateUserMetadata :exec
UPDATE user SET metadata = json_patch(COALESCE(metadata, '{}'), ?) WHERE id = ?;
