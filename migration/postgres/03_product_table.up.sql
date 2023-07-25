CREATE TABLE "product" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(36) NOT NULL,
    "price" NUMERIC NOT NULL,
    "barcode" VARCHAR NOT NULL,
    "category_id" UUID REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

