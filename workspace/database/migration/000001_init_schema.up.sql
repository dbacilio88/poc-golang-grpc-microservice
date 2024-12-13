CREATE  TABLE if not exists"email" (
    "id" bigserial PRIMARY KEY,
    "title" varchar(100) NOT NULL,
    "body" varchar(100) NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamptz not null default (now())
);

COMMENT ON COLUMN "email"."body" IS 'Content of the post';

CREATE  TABLE if not exists "transaction"(
    "id" bigserial PRIMARY KEY,
    "email_id" bigint NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamptz not null default (now())
);

ALTER TABLE "transaction" ADD foreign key ("email_id") references "email"("id");