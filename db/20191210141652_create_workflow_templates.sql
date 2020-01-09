-- +goose Up
CREATE TABLE workflow_templates
(
    id          serial PRIMARY KEY,
    uid         varchar(36) UNIQUE NOT NULL CHECK(uid <> ''),
    name        text NOT NULL CHECK(name <> ''),
    namespace   varchar(36) NOT NULL,

    -- auditing info
    created_at  timestamp NOT NULL DEFAULT (NOW() at time zone 'utc'),
    modified_at timestamp,

    UNIQUE (name, namespace)
);

-- +goose Down
DROP TABLE workflow_templates;