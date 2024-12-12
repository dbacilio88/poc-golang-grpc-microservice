CREATE TABLE "email" (
    "id" bigserial PRIMARY KEY,
    "title" varchar(100) NOT NULL,
    "body" varchar(100) NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamptz not null default (now())
);

COMMENT ON COLUMN "email"."body" IS 'Content of the post';
