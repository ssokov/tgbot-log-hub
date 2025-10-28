-- =============================================================================
-- Diagram Name: tgbot-log-hub
-- Created on: 29.10.2025 00:18:52
-- Diagram Version: 
-- =============================================================================

CREATE TABLE "admin_roles" (
	"id" int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"role_name" varchar(255) NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "admins" (
	"id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"login" varchar(255) NOT NULL,
	"email" varchar(255) NOT NULL,
	"password_hash" varchar(255) NOT NULL,
	"role_id" int4 NOT NULL,
	"status" int4 NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id"),
	CONSTRAINT "uk_admin_login" UNIQUE("login"),
	CONSTRAINT "uk_admin_email" UNIQUE("email")
);

CREATE TABLE "services_admins" (
	"service_id" int8 NOT NULL,
	"admin_id" int8 NOT NULL,
	"assigned_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("service_id","admin_id")
);

CREATE TABLE "services" (
	"id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" varchar(255) NOT NULL,
	"type_id" int4,
	"api_key" varchar(255) NOT NULL,
	"status" int4 NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "service_types" (
	"id" int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"type_name" varchar(255) NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "service_users" (
	"id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"tg_id" int8 NOT NULL DEFAULT -1,
	"nickname" varchar(255),
	"params" jsonb,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "service_logs" (
	"id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"type_id" int8 NOT NULL,
	"error_code" int4,
	"message" text,
	"service_id" int8 NOT NULL,
	"user_id" int8,
	"additional_data" jsonb,
	"created_at" timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY("id")
);

CREATE TABLE "log_types" (
	"id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"type_name" varchar(255) NOT NULL,
	PRIMARY KEY("id")
);


ALTER TABLE "admins" ADD CONSTRAINT "Ref_admin_to_admin_role" FOREIGN KEY ("role_id")
	REFERENCES "admin_roles"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "services_admins" ADD CONSTRAINT "Ref_service_admin_to_service" FOREIGN KEY ("service_id")
	REFERENCES "services"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "services_admins" ADD CONSTRAINT "Ref_service_admin_to_admin" FOREIGN KEY ("admin_id")
	REFERENCES "admins"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "services" ADD CONSTRAINT "Ref_service_to_service_type" FOREIGN KEY ("type_id")
	REFERENCES "service_types"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_logs" ADD CONSTRAINT "Ref_service_log_to_service_user" FOREIGN KEY ("user_id")
	REFERENCES "service_users"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_logs" ADD CONSTRAINT "Ref_service_log_to_service" FOREIGN KEY ("service_id")
	REFERENCES "services"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "service_logs" ADD CONSTRAINT "Ref_service_log_to_log_type" FOREIGN KEY ("type_id")
	REFERENCES "log_types"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;


