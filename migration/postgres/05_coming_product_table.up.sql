CREATE TABLE "coming_products" (
    "id" UUID PRIMARY KEY,
    "coming_id" uuid REFERENCES "coming"("id"),
    "category_id" UUID REFERENCES "category"("id"),
    "name" VARCHAR NOT NULL,
    "barcode" VARCHAR NOT NULL,
    "amount"   int NOT NULL,
    "price"  NUMERIC NOT NULL,
    "total_price" NUMERIC NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);