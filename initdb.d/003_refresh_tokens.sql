CREATE TABLE IF NOT EXISTS refresh_tokens (
    id              UUID         PRIMARY KEY DEFAULT uuid_generate_v7(),
    user_id         UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    family_id       UUID         NOT NULL DEFAULT uuid_generate_v7(),
    parent_id       UUID         REFERENCES refresh_tokens(id) ON DELETE SET NULL,
    token_hash      TEXT         NOT NULL UNIQUE,
    expires_at      TIMESTAMPTZ  NOT NULL,
    revoked_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_family    ON refresh_tokens (family_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user      ON refresh_tokens (user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_active    ON refresh_tokens (token_hash) WHERE revoked_at IS NULL;
