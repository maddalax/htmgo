-- SQLite schema for User Management

-- User table
CREATE TABLE IF NOT EXISTS user
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    email      TEXT NOT NULL UNIQUE,
    password   TEXT NOT NULL,
    metadata   JSON DEFAULT '{}',
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now'))
);

-- Auth Token table
CREATE TABLE IF NOT EXISTS sessions
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER NOT NULL,
    session_id      TEXT    NOT NULL UNIQUE,
    created_at TEXT DEFAULT (datetime('now')),
    expires_at TEXT    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);

-- Indexes to improve query performance
CREATE INDEX IF NOT EXISTS idx_user_email ON user (email);
CREATE INDEX IF NOT EXISTS idx_session_id ON sessions (session_id);
CREATE INDEX IF NOT EXISTS idx_auth_sessions_user_id ON sessions (user_id);
