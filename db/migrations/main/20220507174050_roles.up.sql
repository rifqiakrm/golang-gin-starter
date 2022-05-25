BEGIN;

CREATE TABLE IF NOT EXISTS main.roles
(
    id         UUID         NOT NULL,
    name       VARCHAR(128) NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;