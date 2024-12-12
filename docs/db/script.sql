CREATE TABLE "email" (
                         "id" integer PRIMARY KEY,
                         "title" varchar,
                         "body" text,
                         "status" varchar,
                         "created_at" timestamp
);

COMMENT ON COLUMN "email"."body" IS 'Content of the post';
