CREATE TABLE "user" (
  "user_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "password" varchar,
  "email" varchar NOT NULL,
  "sso_identifer" varchar,
  "cr_date" timestamp DEFAULT (now()),
  "up_date" timestamp,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "fund" (
  "fund_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "balance" decimal NOT NULL,
  "currency_type" varchar NOT NULL,
  "cr_date" timestamp DEFAULT (now()),
  "up_date" timestamp,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stock" (
  "stock_id" BIGSERIAL PRIMARY KEY,
  "ticker_symbol" varchar NOT NULL,
  "comp_name" varchar NOT NULL,
  "current_price" decimal NOT NULL,
  "market_cap" bigint NOT NULL,
  "cr_date" timestamp DEFAULT (now()),
  "up_date" timestamp,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "user_stock" (
  "user_stock_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "stock_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "purchase_price_per_share" decimal NOT NULL,
  "purchased_date" timestamp NOT NULL,
  "cr_date" timestamp DEFAULT (now()),
  "up_date" timestamp,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE TABLE "stock_transaction" (
  "TransationId" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "stock_id" bigint NOT NULL,
  "transaction_type" varchar NOT NULL,
  "transaction_date" timestamp NOT NULL,
  "transation_amt" bigint NOT NULL,
  "transation_proce_per_share" decimal NOT NULL,
  "cr_date" timestamp DEFAULT (now()),
  "up_date" timestamp,
  "cr_user" varchar NOT NULL,
  "up_user" varchar
);

CREATE INDEX ON "user" ("user_id");

CREATE INDEX ON "fund" ("fund_id");

CREATE INDEX ON "fund" ("user_id");

CREATE INDEX ON "stock" ("stock_id");

CREATE INDEX ON "user_stock" ("user_stock_id");

CREATE INDEX ON "user_stock" ("user_id");

CREATE INDEX ON "user_stock" ("stock_id");

CREATE INDEX ON "stock_transaction" ("TransationId");

CREATE INDEX ON "stock_transaction" ("user_id", "stock_id");

ALTER TABLE "fund" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user_stock" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "stock_transaction" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user_stock" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("stock_id");

ALTER TABLE "stock_transaction" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("stock_id");