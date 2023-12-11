CREATE TYPE "transation_result" AS ENUM (
  'createed',
  'processed',
  'successed',
  'failed'
);
ALTER TABLE "stock_transaction" ADD COLUMN  result transation_result NOT NULL DEFAULT 'createed';