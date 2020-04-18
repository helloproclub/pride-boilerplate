-- +migrate Up
CREATE TABLE IF NOT EXISTS "todos" (
    "id" UUID NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
);

-- +migrate Down
DROP TABLE IF EXISTS "todos";