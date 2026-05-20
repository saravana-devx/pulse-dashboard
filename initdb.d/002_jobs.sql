DO $$ BEGIN
    CREATE TYPE job_status AS ENUM (
        'pending',
        'running',
        'completed',
        'failed'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;  -- if type already exists, skip silently
END $$;

CREATE TABLE IF NOT EXISTS jobs (
    id           UUID        NOT NULL DEFAULT uuid_generate_v7(),  -- ← remove PRIMARY KEY here
    user_id      UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type         TEXT        NOT NULL,
    payload      JSONB       NOT NULL DEFAULT '{}',
    status       job_status  NOT NULL DEFAULT 'pending',
    priority     INT         NOT NULL DEFAULT 5,
    worker_id    TEXT,
    attempts     INT         NOT NULL DEFAULT 0,
    max_retries  INT         NOT NULL DEFAULT 3,
    error_msg    TEXT,
    scheduled_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_at   TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ,
    PRIMARY KEY (id, created_at)                    -- ← composite primary key
) PARTITION BY RANGE (created_at);

CREATE TABLE IF NOT EXISTS
    jobs_2026_05_01 PARTITION OF jobs FOR
VALUES
FROM
    ('2026-05-01') TO ('2026-05-02');

CREATE TABLE IF NOT EXISTS
    jobs_2026_05_02 PARTITION OF jobs FOR
VALUES
FROM
    ('2026-05-02') TO ('2026-05-03');

CREATE INDEX ON jobs (status, scheduled_at)
WHERE
    status = 'pending';

CREATE INDEX ON jobs (worker_id)
WHERE
    worker_id IS NOT NULL;
CREATE INDEX ON jobs (deleted_at)
WHERE deleted_at IS NULL;