-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO "public"."referential" ("sequence", "table_name", "field_name", "value")
VALUES (1, 'code_generator', 'type', 'RFQ'),
       (2, 'code_generator', 'type', 'SQ'),
       (1, 'quote', 'status', 'New'),
       (2, 'quote', 'status', 'In Progress'),
       (3, 'quote', 'status', 'Done'),
       (1, 'quote_item', 'status', 'New'),
       (2, 'quote_item', 'status', 'Ready For Release'),
       (3, 'quote_item', 'status', 'Released'),
       (1, 'quote_item', 'state', 'New'),
       (2, 'quote_item', 'state', 'MD Checking'),
       (3, 'quote_item', 'state', 'MD Release'),
       (4, 'quote_item', 'state', 'Sales Release'),
       (1, 'quote_release', 'status', 'New'),
       (2, 'quote_release', 'status', 'Done'),
       (1, 'quote_release_item', 'status', 'New'),
       (2, 'quote_release_item', 'status', 'Win'),
       (3, 'quote_release_item', 'status', 'Lose'),
       (4, 'quote_release_item', 'status', 'Purchased');


INSERT INTO "public"."code_generator" ("sequence", "type", "prefix")
VALUES (2, 'rfq', 'RFQ'),
       (1, 'sq', 'SQ');


INSERT INTO "public"."quote" ("code", "account_manager_name", "company_id", "company_name", "status", "created_by")
VALUES ('RFQ200318001', 'Lynet', '5615aa45-a85e-4f97-9a37-678705a47c87', 'Hermes 3',
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote'
	       AND "field_name" = 'status'
         LIMIT 1 OFFSET 0), '5278e3e1-6962-4ade-bf1e-2b428e8f53c4'),

       ('RFQ200318002', 'Lynet', '5615aa45-a85e-4f97-9a37-678705a47c87', 'Hermes 3',
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote'
	       AND "field_name" = 'status'
         LIMIT 1 OFFSET 0), '5278e3e1-6962-4ade-bf1e-2b428e8f53c4');


INSERT INTO "public"."quote_original_item" ("quote_id", "name", "qty", "uom", "uom_label", "target_price", "created_by")
VALUES ((SELECT "id"
         FROM "public"."quote"
         LIMIT 1 OFFSET 0), 'Keyboard', 20, 'pcs', 'Pieces', 500000, '5278e3e1-6962-4ade-bf1e-2b428e8f53c4'),

       ((SELECT "id"
         FROM "public"."quote"
         LIMIT 1 OFFSET 0), 'Mouse', 20, 'pcs', 'Pieces', 100000, '5278e3e1-6962-4ade-bf1e-2b428e8f53c4'),

       ((SELECT "id"
         FROM "public"."quote"
         LIMIT 1 OFFSET 1), 'Snowman Black Marker', 100, 'pcs', 'Pieces', 50000,
        '5278e3e1-6962-4ade-bf1e-2b428e8f53c4');


INSERT INTO "public"."quote_item" ("quote_original_id", "product_uom", "product_uom_label", "status", "state",
                                   "created_by")
VALUES ((SELECT "id"
         FROM "public"."quote_original_item"
         LIMIT 1 OFFSET 0), 'pcs', 'Pieces',
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'status'
         LIMIT 1 OFFSET 0),
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'state'
         LIMIT 1 OFFSET 0), '5278e3e1-6962-4ade-bf1e-2b428e8f53c4'),

       ((SELECT "id"
         FROM "public"."quote_original_item"
         LIMIT 1 OFFSET 1), 'pcs', 'Pieces',
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'status'
         LIMIT 1 OFFSET 0),
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'state'
         LIMIT 1 OFFSET 0), '5278e3e1-6962-4ade-bf1e-2b428e8f53c4'),

       ((SELECT "id"
         FROM "public"."quote_original_item"
         LIMIT 1 OFFSET 2), 'pcs', 'Pieces',
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'status'
         LIMIT 1 OFFSET 0),
        (SELECT "id"
         FROM "public"."referential"
         WHERE "table_name" = 'quote_item'
	       AND "field_name" = 'state'
         LIMIT 1 OFFSET 0), '5278e3e1-6962-4ade-bf1e-2b428e8f53c4');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
TRUNCATE TABLE "public"."quote_item" CASCADE;
TRUNCATE TABLE "public"."quote_original_item" CASCADE;
TRUNCATE TABLE "public"."quote" CASCADE;
TRUNCATE TABLE "public"."referential" CASCADE;
TRUNCATE TABLE "public"."code_generator" CASCADE;
