-- =============================================================================
-- Diagram Name: tgbot-log-hub
-- Created on: 18.10.2025 20:32:28
-- Diagram Version: 
-- =============================================================================

CREATE TABLE "admin_role" (
	"id" SERIAL NOT NULL,
	"role_name" varchar(255) NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "admin" (
	"id" BIGSERIAL NOT NULL,
	"login" varchar(255) NOT NULL,
	"email" varchar(255) NOT NULL,
	"password_hash" varchar(255) NOT NULL,
	"role_id" int4 NOT NULL,
	"status" int4 NOT NULL,
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id"),
	CONSTRAINT "uk_admin_login" UNIQUE("login"),
	CONSTRAINT "uk_admin_email" UNIQUE("email")
);

CREATE TABLE "service_admin" (
	"service_id" int8 NOT NULL,
	"admin_id" int8 NOT NULL,
	"assignedAt" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("service_id","admin_id")
);

CREATE TABLE "service" (
	"id" BIGSERIAL NOT NULL,
	"name" varchar(255) NOT NULL,
	"type_id" int4,
	"api_key" varchar(255) NOT NULL,
	"status" int4 NOT NULL,
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "service_type" (
	"id" SERIAL NOT NULL,
	"type_name" varchar(255) NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "service_user" (
	"id" BIGSERIAL NOT NULL,
	"tg_id" int8 NOT NULL DEFAULT -1,
	"nickname" varchar(255) NOT NULL,
	"params" jsonb,
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "service_log" (
	"id" BIGSERIAL NOT NULL,
	"type_id" int8 NOT NULL,
	"error_code" int4,
	"message" text,
	"service_id" int8 NOT NULL,
	"user_id" int8,
	"additional_data" jsonb,
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "log_type" (
	"id" BIGSERIAL NOT NULL,
	"type_name" varchar(255) NOT NULL,
	PRIMARY KEY("id")
);


ALTER TABLE "admin" ADD CONSTRAINT "Ref_admin_to_admin_role" FOREIGN KEY ("role_id")
	REFERENCES "admin_role"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_admin" ADD CONSTRAINT "Ref_service_admin_to_service" FOREIGN KEY ("service_id")
	REFERENCES "service"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_admin" ADD CONSTRAINT "Ref_service_admin_to_admin" FOREIGN KEY ("admin_id")
	REFERENCES "admin"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service" ADD CONSTRAINT "Ref_service_to_service_type" FOREIGN KEY ("type_id")
	REFERENCES "service_type"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_log" ADD CONSTRAINT "Ref_service_log_to_service_user" FOREIGN KEY ("user_id")
	REFERENCES "service_user"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_log" ADD CONSTRAINT "Ref_service_log_to_service" FOREIGN KEY ("service_id")
	REFERENCES "service"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_log" ADD CONSTRAINT "Ref_service_log_to_log_type" FOREIGN KEY ("type_id")
	REFERENCES "log_type"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;


