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

CREATE INDEX ON "stock_day_avg_all" ("code");

CREATE INDEX ON "stock_day_avg_all" ("stock_name");