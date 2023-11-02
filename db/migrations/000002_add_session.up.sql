CREATE TABLE "session" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "cr_date" timestamptz NOT NULL DEFAULT (now()),
  "expired_date" timestamptz NOT NULL DEFAULT (now() + interval '3 days')
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");