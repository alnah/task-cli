CREATE TYPE "status" AS ENUM (
  'todo',
  'in-progress',
  'done'
);

CREATE TABLE "task" (
  "id" serial PRIMARY KEY,
  "description" varchar(100) NOT NULL,
  "status" status NOT NULL DEFAULT 'todo',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT null,
  "deleted_at" timestamptz DEFAULT null
);

CREATE INDEX ON "task" ("status");
