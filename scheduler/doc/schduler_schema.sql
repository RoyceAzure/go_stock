-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-11-16T09:29:17.166Z

CREATE TABLE "stock_day_avg_all" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "stock_name" varchar NOT NULL,
  "close_price" decimal NOT NULL,
  "monthly_avg_price" decimal NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL DEFAULT 'SYSTEM',
  "up_user" varchar
);

CREATE TABLE "stock_price_realtime" (
  "code" varchar NOT NULL,
  "stock_name" varchar NOT NULL,
  "trade_volume" decimal NOT NULL,
  "trade_value" decimal NOT NULL,
  "opening_price" decimal NOT NULL,
  "highest_price" decimal NOT NULL,
  "lowest_price" decimal NOT NULL,
  "closing_price" decimal NOT NULL,
  "change" decimal NOT NULL,
  "transaction" decimal NOT NULL,
  "trans_time" timestamptz NOT NULL
);

CREATE INDEX ON "stock_day_avg_all" ("code");

CREATE INDEX ON "stock_day_avg_all" ("stock_name");
