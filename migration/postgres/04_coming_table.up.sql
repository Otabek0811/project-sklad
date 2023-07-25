CREATE TABLE "coming" (
    "id" UUID PRIMARY KEY,
    "coming_id" VARCHAR(36) UNIQUE NOT NULL,  
    "filial_id" UUID REFERENCES "filial"("id"),
    "date_time" TIMESTAMP ,
    "status" VARCHAR(20) DEFAULT 'in process',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);