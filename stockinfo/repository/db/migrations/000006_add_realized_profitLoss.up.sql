CREATE TABLE "realized_profit_loss" (
  "id" BIGSERIAL PRIMARY KEY,
  "transation_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "product_name" varchar NOT NULL,
  "cost_per_price" decimal NOT NULL,
  "cost_total_price" decimal NOT NULL,
  "realized" decimal NOT NULL,
  "realized_precent" varchar NOT NULL
);