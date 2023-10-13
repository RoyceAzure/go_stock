CREATE TABLE "users" (
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

CREATE TABLE "funds" (
  "fund_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "balance" decimal NOT NULL,
  "currency_type" varchar NOT NULL,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stocks" (
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

CREATE TABLE "user_stocks" (
  "user_stock_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint,
  "stock_id" bigint,
  "quantity" int DEFAULT 1,
  "purchase_price_per_share" decimal,
  "purchased_date" timestamp,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stock_transactions" (
  "TransationId" BIGSERIAL PRIMARY KEY,
  "user_id" bigint,
  "stock_id" bigint,
  "transaction_type" varchar,
  "transaction_date" timestamp DEFAULT (now()),
  "transation_amt" decimal,
  "transation_proce_per_share" decimal,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "up_date" timestamptz,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE INDEX ON "users" ("user_id");

CREATE INDEX ON "funds" ("fund_id");

CREATE INDEX ON "funds" ("user_id");

CREATE UNIQUE INDEX ON "funds" ("user_id", "currency_type");

CREATE INDEX ON "stocks" ("stock_id");

CREATE INDEX ON "user_stocks" ("user_stock_id");

CREATE INDEX ON "user_stocks" ("user_id");

CREATE INDEX ON "user_stocks" ("stock_id");

CREATE INDEX ON "stock_transactions" ("TransationId");

CREATE INDEX ON "stock_transactions" ("user_id", "stock_id");

ALTER TABLE "funds" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_stocks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "stock_transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_stocks" ADD FOREIGN KEY ("stock_id") REFERENCES "stocks" ("stock_id");

ALTER TABLE "stock_transactions" ADD FOREIGN KEY ("stock_id") REFERENCES "stocks" ("stock_id");
