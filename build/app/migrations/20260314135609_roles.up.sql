CREATE TABLE roles (
                       id BIGSERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT NOT NULL DEFAULT '',
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX ux_roles_name_active
    ON roles (name)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_roles_deleted_at ON roles (deleted_at);