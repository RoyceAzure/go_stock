-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-11-02T09:57:41.602Z

CREATE TABLE "user" (
  "user_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "sso_identifer" varchar,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "fund" (
  "fund_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "balance" decimal NOT NULL,
  "currency_type" varchar NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stock" (
  "stock_id" BIGSERIAL PRIMARY KEY,
  "ticker_symbol" varchar NOT NULL,
  "comp_name" varchar NOT NULL,
  "current_price" decimal NOT NULL,
  "market_cap" bigint NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "user_stock" (
  "user_stock_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "stock_id" bigint NOT NULL,
  "quantity" int NOT NULL DEFAULT 1,
  "purchase_price_per_share" decimal NOT NULL,
  "purchased_date" timestamp NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stock_transaction" (
  "transation_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "stock_id" bigint NOT NULL,
  "fund_id" bigint NOT NULL,
  "transaction_type" varchar NOT NULL,
  "transaction_date" timestamp NOT NULL DEFAULT (now()),
  "transation_amt" int NOT NULL,
  "transation_price_per_share" decimal NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE INDEX ON "user" ("user_id");

CREATE INDEX ON "fund" ("fund_id");

CREATE INDEX ON "fund" ("user_id");

CREATE UNIQUE INDEX ON "fund" ("user_id", "currency_type");

CREATE INDEX ON "stock" ("stock_id");

CREATE INDEX ON "user_stock" ("user_stock_id");

CREATE INDEX ON "user_stock" ("user_id");

CREATE INDEX ON "user_stock" ("stock_id");

CREATE INDEX ON "stock_transaction" ("transation_id");

CREATE INDEX ON "stock_transaction" ("user_id", "stock_id");

ALTER TABLE "fund" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user_stock" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "stock_transaction" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user_stock" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("stock_id");

ALTER TABLE "stock_transaction" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("stock_id");

ALTER TABLE "stock_transaction" ADD FOREIGN KEY ("fund_id") REFERENCES "fund" ("fund_id");

