BEGIN;

CREATE TABLE IF NOT EXISTS main.role_permissions
(
    id            UUID         NOT NULL,
    role_id       UUID         NOT NULL,
    permission_id UUID         NOT NULL,
    created_by    VARCHAR(128) NOT NULL,
    updated_by    VARCHAR(128) NOT NULL,
    deleted_by    VARCHAR(128),
    created_at    TIMESTAMPTZ  NOT NULL,
    updated_at    TIMESTAMPTZ  NOT NULL,
    deleted_at    TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;