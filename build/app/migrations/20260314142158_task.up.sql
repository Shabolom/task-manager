CREATE TABLE tasks (
                       id BIGSERIAL PRIMARY KEY,
                       board_id BIGINT NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
                       status_id BIGINT NOT NULL REFERENCES statuses(id) ON DELETE RESTRICT,
                       title TEXT NOT NULL,
                       description TEXT NOT NULL DEFAULT '',
                       priority TEXT NOT NULL DEFAULT 'medium',
                       assignee_id BIGINT NULL REFERENCES users(id) ON DELETE SET NULL,
                       creator_id BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
                       position INTEGER NOT NULL DEFAULT 0,
                       due_date TIMESTAMPTZ NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMPTZ NULL

                       CONSTRAINT tasks_priority_check CHECK (
                           priority IN ('low', 'medium', 'high')
                           ),
                       CONSTRAINT tasks_position_check CHECK (position >= 0)
);

CREATE INDEX idx_tasks_id ON tasks (id);