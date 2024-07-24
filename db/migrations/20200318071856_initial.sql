-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE "public"."referential";
CREATE TABLE "public"."referential"
(
	"id"           UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"table_name"   VARCHAR(50)    NOT NULL,
	"field_name"   VARCHAR(50)    NOT NULL,
	"value"        VARCHAR(50)    NOT NULL,
	"sequence"     INT2           NOT NULL,
	"is_active"    INT2           NOT NULL DEFAULT 1,
	"created_by"   UUID                    DEFAULT NULL,
	"created_date" TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"   UUID                    DEFAULT NULL,
	"updated_date" TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "referential_unique_idx" ON "public"."referential" ("table_name", "field_name", "sequence");

DROP TABLE "public"."code_generator";
CREATE TABLE "public"."code_generator"
(
	"id"           UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"type"         VARCHAR(50)    NOT NULL,
	"prefix"       VARCHAR(50)    NOT NULL,
	"date"         TIMESTAMPTZ(6) NOT NULL DEFAULT NOW()::DATE,
	"sequence"     INT2           NOT NULL DEFAULT 1,
	"created_by"   UUID                    DEFAULT NULL,
	"created_date" TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"   UUID                    DEFAULT NULL,
	"updated_date" TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "code_generator_unique_prefix_idx" ON "public"."code_generator" ("prefix", "type", "date");

CREATE TABLE "public"."quote"
(
	"id"                   UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"company_id"           UUID           NOT NULL,
	"company_name"         VARCHAR(100)   NOT NULL,
	"account_manager_name" VARCHAR(50)    NOT NULL,
	"code"                 VARCHAR(20)    NOT NULL,
	"status"               UUID           NOT NULL,
	"closed_by"            UUID                    DEFAULT NULL,
	"closed_date"          TIMESTAMPTZ(6)          DEFAULT NULL,
	"created_by"           UUID           NOT NULL,
	"created_date"         TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"           UUID                    DEFAULT NULL,
	"updated_date"         TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "quote_unique_idx" ON "public"."quote" (LOWER("code"));


CREATE TABLE "public"."quote_original_item"
(
	"id"           UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_id"     UUID           NOT NULL,
	"name"         VARCHAR(255)   NOT NULL,
	"uom_label"    VARCHAR(20)             DEFAULT NULL,
	"uom"          VARCHAR(20)             DEFAULT NULL,
	"qty"          INT4           NOT NULL,
	"target_price" NUMERIC(30, 2)          DEFAULT NULL,
	"created_by"   UUID           NOT NULL,
	"created_date" TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"   UUID                    DEFAULT NULL,
	"updated_date" TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_item"
(
	"id"                    UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_original_id"     UUID           NOT NULL,
	"parent_id"             UUID                    DEFAULT NULL,
	"product_id"            UUID                    DEFAULT NULL,
	"product_sku"           VARCHAR(50)             DEFAULT NULL,
	"product_uom_label"     VARCHAR(20)             DEFAULT NULL,
	"product_uom"           VARCHAR(20)             DEFAULT NULL,
	"product_qty"           INT4                    DEFAULT NULL,
	"product_sales_price"   NUMERIC(30, 2)          DEFAULT NULL,
	"vendor_code"           VARCHAR(30)             DEFAULT NULL,
	"vendor_stock_status"   UUID                    DEFAULT NULL,
	"vendor_pricelist"      NUMERIC(30, 2)          DEFAULT NULL,
	"fulfillment_type"      UUID                    DEFAULT NULL,
	"contract_status"       UUID                    DEFAULT NULL,
	"contract_expired_date" TIMESTAMPTZ(6)          DEFAULT NULL,
	"contract_qty"          INT4                    DEFAULT NULL,
	"is_alternate_product"  INT2           NOT NULL DEFAULT 0,
	"status"                UUID           NOT NULL,
	"state"                 UUID           NOT NULL,
	"created_by"            UUID           NOT NULL,
	"created_date"          TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"            UUID                    DEFAULT NULL,
	"updated_date"          TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_item_history"
(
	"id"            UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_item_id" UUID           NOT NULL,
	"state"         UUID           NOT NULL,
	"created_by"    UUID           NOT NULL,
	"created_date"  TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"    UUID                    DEFAULT NULL,
	"updated_date"  TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_item_comment"
(
	"id"            UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_item_id" UUID           NOT NULL,
	"comment"       VARCHAR(255)   NOT NULL,
	"created_by"    UUID           NOT NULL,
	"created_date"  TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"    UUID                    DEFAULT NULL,
	"updated_date"  TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_release"
(
	"id"            UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_id"      UUID           NOT NULL,
	"code"          VARCHAR(20)    NOT NULL,
	"grand_total"   NUMERIC(30, 2) NOT NULL,
	"total_qty"     INT4           NOT NULL,
	"total_margin"  NUMERIC(30, 2) NOT NULL,
	"top"           INT2                    DEFAULT NULL,
	"expired_date"  TIMESTAMPTZ(6)          DEFAULT NULL,
	"shipping_cost" NUMERIC(30, 2) NOT NULL,
	"remarks"       VARCHAR(255)            DEFAULT NULL,
	"version"       INT2           NOT NULL,
	"status"        UUID           NOT NULL,
	"created_by"    UUID           NOT NULL,
	"created_date"  TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"    UUID                    DEFAULT NULL,
	"updated_date"  TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "quote_release_unique_idx" ON "public"."quote_release" (LOWER("code"));


CREATE TABLE "public"."quote_release_history"
(
	"id"               UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_release_id" UUID           NOT NULL,
	"quote_id"         UUID           NOT NULL,
	"code"             VARCHAR(20)    NOT NULL,
	"grand_total"      NUMERIC(30, 2) NOT NULL,
	"total_qty"        INT4           NOT NULL,
	"total_margin"     NUMERIC(30, 2) NOT NULL,
	"top"              INT2                    DEFAULT NULL,
	"expired_date"     TIMESTAMPTZ(6)          DEFAULT NULL,
	"shipping_cost"    NUMERIC(30, 2) NOT NULL,
	"remarks"          VARCHAR(255)            DEFAULT NULL,
	"version"          INT2           NOT NULL,
	"status"           UUID           NOT NULL,
	"created_by"       UUID           NOT NULL,
	"created_date"     TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"       UUID                    DEFAULT NULL,
	"updated_date"     TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_release_item"
(
	"id"               UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_release_id" UUID           NOT NULL,
	"quote_item_id"    UUID           NOT NULL,
	"status"           UUID           NOT NULL,
	"created_by"       UUID           NOT NULL,
	"created_date"     TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"       UUID                    DEFAULT NULL,
	"updated_date"     TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


CREATE TABLE "public"."quote_release_item_history"
(
	"id"                    UUID           NOT NULL DEFAULT "uuid_generate_v4"(),
	"quote_release_item_id" UUID           NOT NULL,
	"quote_release_id"      UUID           NOT NULL,
	"quote_item_id"         UUID           NOT NULL,
	"status"                UUID           NOT NULL,
	"created_by"            UUID           NOT NULL,
	"created_date"          TIMESTAMPTZ(6) NOT NULL DEFAULT NOW(),
	"updated_by"            UUID                    DEFAULT NULL,
	"updated_date"          TIMESTAMPTZ(6)          DEFAULT NULL,

	PRIMARY KEY ("id")
);


ALTER TABLE "public"."quote"
	ADD CONSTRAINT "quote_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_original_item"
	ADD CONSTRAINT "quote_original_item_quote_fkey" FOREIGN KEY ("quote_id")
		REFERENCES "public"."quote" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_item"
	ADD CONSTRAINT "quote_item_original_fkey" FOREIGN KEY ("quote_original_id")
		REFERENCES "public"."quote_original_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_item_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_item_state_fkey" FOREIGN KEY ("state")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_item_history"
	ADD CONSTRAINT "quote_item_history_item_fkey" FOREIGN KEY ("quote_item_id")
		REFERENCES "public"."quote_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_state_fkey" FOREIGN KEY ("state")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_item_comment"
	ADD CONSTRAINT "quote_item_comment_item_fkey" FOREIGN KEY ("quote_item_id")
		REFERENCES "public"."quote_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_release"
	ADD CONSTRAINT "quote_release_quote_fkey" FOREIGN KEY ("quote_id")
		REFERENCES "public"."quote" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_release_history"
	ADD CONSTRAINT "quote_release_history_quote_release_fkey" FOREIGN KEY ("quote_release_id")
		REFERENCES "public"."quote_release" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_history_quote_fkey" FOREIGN KEY ("quote_id")
		REFERENCES "public"."quote" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_history_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_release_item"
	ADD CONSTRAINT "quote_release_item_release_fkey" FOREIGN KEY ("quote_release_id")
		REFERENCES "public"."quote_release" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_item_item_fkey" FOREIGN KEY ("quote_item_id")
		REFERENCES "public"."quote_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_item_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "public"."quote_release_item_history"
	ADD CONSTRAINT "quote_release_item_history_release_fkey" FOREIGN KEY ("quote_release_id")
		REFERENCES "public"."quote_release" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_item_history_release_item_fkey" FOREIGN KEY ("quote_release_item_id")
		REFERENCES "public"."quote_release_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_item_history_item_fkey" FOREIGN KEY ("quote_item_id")
		REFERENCES "public"."quote_item" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED,
	ADD CONSTRAINT "quote_release_item_history_status_fkey" FOREIGN KEY ("status")
		REFERENCES "public"."referential" ("id") ON UPDATE CASCADE DEFERRABLE INITIALLY DEFERRED;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "public"."quote_release_item_history" CASCADE;
DROP TABLE "public"."quote_release_history" CASCADE;
DROP TABLE "public"."quote_release_item" CASCADE;
DROP TABLE "public"."quote_release" CASCADE;
DROP TABLE "public"."quote_item_history" CASCADE;
DROP TABLE "public"."quote_item_comment" CASCADE;
DROP TABLE "public"."quote_item" CASCADE;
DROP TABLE "public"."quote_original_item" CASCADE;
DROP TABLE "public"."quote" CASCADE;
DROP TABLE "public"."referential" CASCADE;
DROP TABLE "public"."code_generator" CASCADE;
