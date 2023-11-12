CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "assigned_to" bigint NOT NULL,
  "created_by" bigint NOT NULL,
  "due_date" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tickets" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
ALTER TABLE "tickets" ADD FOREIGN KEY ("assigned_to") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "tickets" ("assigned_to");
