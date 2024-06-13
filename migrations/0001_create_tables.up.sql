CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "username" varchar,
  "full_name" varchar,
  "sex" int,
  "email" varchar,
  "password" varchar,
  "is_active" int,
  "is_login" int,
  "created_at" timestamp
);

CREATE TABLE "customer" (
  "id" uuid PRIMARY KEY,
  "full_name" varchar,
  "username" varchar,
  "password" varchar,
  "hp" varchar UNIQUE,
  "email" varchar,
  "sex" int,
  "is_active" int,
  "is_login" int,
  "created_by" uuid,
  "created_at" timestamp
);

CREATE TABLE "order" (
  "id" uuid PRIMARY KEY,
  "address_id" uuid,
  "customer_id" uuid,
  "product_name" varchar,
  "amount" int,
  "status" int,
  "created_by" uuid,
  "created_at" timestamp
);

CREATE TABLE "address" (
  "id" uuid PRIMARY KEY,
  "customer_id" uuid,
  "street_address" varchar,
  "province" varchar,
  "district" varchar,
  "sub_district" varchar,
  "zip_code" varchar,
  "created_by" uuid,
  "created_at" timestamp
);

ALTER TABLE "order" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");
ALTER TABLE "order" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");
ALTER TABLE "address" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");
ALTER TABLE "customer" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");
ALTER TABLE "order" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");
ALTER TABLE "address" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");
