ALTER TABLE "order" DROP CONSTRAINT "order_customer_id_fkey";
ALTER TABLE "order" DROP CONSTRAINT "order_address_id_fkey";
ALTER TABLE "address" DROP CONSTRAINT "address_customer_id_fkey";
ALTER TABLE "customer" DROP CONSTRAINT "customer_created_by_fkey";
ALTER TABLE "order" DROP CONSTRAINT "order_created_by_fkey";
ALTER TABLE "address" DROP CONSTRAINT "address_created_by_fkey";

DROP TABLE "user";
DROP TABLE "customer";
DROP TABLE "order";
DROP TABLE "address";
