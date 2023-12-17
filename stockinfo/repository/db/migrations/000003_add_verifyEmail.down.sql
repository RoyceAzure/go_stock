DROP TABLE IF EXISTS "verify_email" CASCADE;

ALTER TABLE "user" DROP COLUMN "is_email_verified";