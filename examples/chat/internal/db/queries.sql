-- name: CreateChatRoom :one
INSERT INTO chat_rooms (id, name, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, created_at, updated_at, last_message_sent_at;

-- name: InsertMessage :exec
INSERT INTO messages (chat_room_id, user_id, username, message, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, chat_room_id, user_id, username, message, created_at, updated_at;

-- name: UpdateChatRoomLastMessageSentAt :exec
UPDATE chat_rooms
SET last_message_sent_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: GetChatRoom :one
SELECT
    id,
    name,
    last_message_sent_at,
    created_at,
    updated_at
FROM chat_rooms
WHERE chat_rooms.id = ?;

-- name: CreateUser :one
INSERT INTO users (name, session_id, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, session_id, created_at, updated_at;

-- name: GetLastMessages :many
SELECT
    messages.id,
    messages.chat_room_id,
    messages.user_id,
    users.name AS user_name,
    messages.message,
    messages.created_at,
    messages.updated_at
FROM messages
         JOIN users ON messages.user_id = users.id
WHERE messages.chat_room_id = ?
ORDER BY messages.created_at
LIMIT ?;

-- name: GetUserBySessionId :one
SELECT * FROM users WHERE session_id = ?;
