CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT     NOT NULL UNIQUE,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    session_id TEXT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS chat_rooms
(
    id                   TEXT PRIMARY KEY,
    name                 TEXT     NOT NULL,
    last_message_sent_at TEXT,
    created_at           TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
) STRICT;

CREATE TABLE IF NOT EXISTS messages
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_room_id TEXT     NOT NULL,
    user_id      INTEGER  NOT NULL,
    username     TEXT     NOT NULL,
    message      TEXT     NOT NULL,
    created_at   TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
) STRICT;

CREATE INDEX IF NOT EXISTS idx_messages_chat_room_id ON messages (chat_room_id);
CREATE INDEX IF NOT EXISTS idx_messages_user_id ON messages (user_id);
