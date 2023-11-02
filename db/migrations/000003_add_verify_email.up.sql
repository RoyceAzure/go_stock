CREATE TABLE "verify_email" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "expired_date" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "verify_email" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user" ADD COLUMN "is_email_verify" bool NOT NULL DEFAULT false;