CREATE SEQUENCE scrapbooks_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."scrapbooks" (
    "id" bigint DEFAULT nextval('scrapbooks_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text,
    "user_id" bigint,
    CONSTRAINT "scrapbooks_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_scrapbooks_deleted_at" ON "public"."scrapbooks" USING btree ("deleted_at");

ALTER TABLE ONLY "public"."scrapbooks" ADD CONSTRAINT "fk_users_scrapbook" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE;
