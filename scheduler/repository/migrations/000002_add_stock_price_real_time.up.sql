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