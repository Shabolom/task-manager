CREATE TABLE statuses (
                          id BIGSERIAL PRIMARY KEY,
                          board_id BIGINT NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
                          name TEXT NOT NULL,
                          position INTEGER NOT NULL DEFAULT 0,
                          color TEXT NOT NULL DEFAULT '',
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          deleted_at TIMESTAMPTZ NULL

                          CONSTRAINT statuses_position_check CHECK (position >= 0)
);

CREATE UNIQUE INDEX ux_statuses_board_name_active
    ON statuses (board_id, name)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_statuses_board_id ON statuses (board_id);
CREATE INDEX idx_statuses_board_position ON statuses (board_id, position);
