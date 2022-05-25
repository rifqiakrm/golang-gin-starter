BEGIN;

CREATE TABLE IF NOT EXISTS main.notifications
(
    id          UUID         NOT NULL,
    user_id     UUID,
    title       TEXT         NOT NULL,
    description TEXT         NOT NULL,
    type        VARCHAR(50)  NOT NULL,
    extra       TEXT,
    is_read     BOOLEAN      NOT NULL,
    created_by  VARCHAR(128) NOT NULL,
    updated_by  VARCHAR(128) NOT NULL,
    deleted_by  VARCHAR(128),
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;