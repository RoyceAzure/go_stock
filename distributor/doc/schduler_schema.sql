-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-11-22T13:43:39.011Z

CREATE TABLE "client_register" (
  "client_uid" uuid NOT NULL,
  "stock_code" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "user_register" (
  "user_id" bigint NOT NULL,
  "stock_code" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "frontend_client" (
  "client_uid" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "ip" varchar NOT NULL,
  "region" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE UNIQUE INDEX ON "client_register" ("client_uid", "stock_code");

CREATE INDEX ON "client_register" ("client_uid");

CREATE UNIQUE INDEX ON "user_register" ("user_id", "stock_code");

CREATE INDEX ON "user_register" ("user_id");

CREATE INDEX ON "frontend_client" ("client_uid");

CREATE UNIQUE INDEX ON "frontend_client" ("ip");

ALTER TABLE "client_register" ADD FOREIGN KEY ("client_uid") REFERENCES "frontend_client" ("client_uid");
