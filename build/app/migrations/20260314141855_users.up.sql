CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE RESTRICT,
                       email TEXT NOT NULL,
                       password_hash TEXT NOT NULL,
                       full_name TEXT NOT NULL,
                       is_active BOOLEAN NOT NULL DEFAULT TRUE,
                       last_login_at TIMESTAMPTZ NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX ux_users_email_active
    ON users (email)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_users_role_id ON users (id);
CREATE INDEX idx_users_email ON users (email);