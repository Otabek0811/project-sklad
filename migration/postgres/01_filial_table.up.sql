CREATE TABLE "filial"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "address" VARCHAR,
    "phone_number" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
