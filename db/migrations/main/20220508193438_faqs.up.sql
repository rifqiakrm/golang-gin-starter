BEGIN;

CREATE TABLE IF NOT EXISTS main.faqs
(
    id         UUID         NOT NULL,
    question   TEXT         NOT NULL,
    answer     TEXT         NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id)
);

COMMIT;