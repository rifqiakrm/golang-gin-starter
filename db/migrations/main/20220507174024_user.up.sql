BEGIN;

CREATE TABLE IF NOT EXISTS main.users
(
    id                    UUID         NOT NULL,
    name                  VARCHAR(128) NOT NULL,
    email                 VARCHAR(128) NOT NULL,
    password              VARCHAR(255) NOT NULL,
    phone_number          VARCHAR(50)  NOT NULL,
    otp                   VARCHAR(6),
    otp_password          VARCHAR(6),
    dob                   DATE,
    photo                 TEXT,
    status                VARCHAR(50),
    forgot_password_token TEXT,
    created_by            VARCHAR(128) NOT NULL,
    updated_by            VARCHAR(128) NOT NULL,
    deleted_by            VARCHAR(128),
    created_at            TIMESTAMPTZ  NOT NULL,
    updated_at            TIMESTAMPTZ  NOT NULL,
    deleted_at            TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;