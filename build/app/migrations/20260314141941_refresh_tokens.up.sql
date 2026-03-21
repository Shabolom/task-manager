CREATE TABLE refresh_tokens (
                                id BIGSERIAL PRIMARY KEY,
                                user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                token TEXT NOT NULL UNIQUE,
                                expires_at TIMESTAMPTZ NOT NULL,
                                revoked_at TIMESTAMPTZ NULL,
                                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                deleted_at TIMESTAMPTZ NULL

                                CONSTRAINT refresh_tokens_token_check CHECK (char_length(trim(token)) > 0)
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);
