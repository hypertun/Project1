CREATE TABLE IF NOT EXISTS "accounts"
(
    "id" integer unique not null,
    "balance" numeric(16,5) default null,
    "created_at" timestamp default CURRENT_TIMESTAMP,
    "updated_at" timestamp default CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
CREATE INDEX accounts_ix ON "accounts" ("id");

CREATE TABLE IF NOT EXISTS "transactions"
(
    "source_account_id" integer not null,
    "destination_account_id" integer not null,
    "amount" numeric(16,5) default null,
    "created_at" timestamp default CURRENT_TIMESTAMP,
    FOREIGN KEY ("source_account_id") REFERENCES "accounts" ("id"),
    FOREIGN KEY ("destination_account_id") REFERENCES "accounts" ("id")
);
CREATE INDEX transactions_source_ix ON "transactions" ("source_account_id");
CREATE INDEX transactions_destination_ix ON "transactions" ("destination_account_id");
CREATE INDEX transactions_source_destination_ix ON "transactions" ("source_account_id", "destination_account_id");