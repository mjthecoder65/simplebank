CREATE TABLE users (
  "username" VARCHAR PRIMARY KEY,
  "hashed_password" VARCHAR NOT NULL,
  "full_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE("owner", "currency");
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");