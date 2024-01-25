-- Database: mgmt-auth
/*
DROP DATABASE IF EXISTS "mgmt-auth";

CREATE DATABASE "mgmt-auth"
    WITH
    OWNER = joakimengqvist
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
*/

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.users (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    privileges TEXT [],
    projects TEXT [],
    notes TEXT [] NOT NULL DEFAULT ARRAY[]::TEXT[],
    password character varying(60),
    user_active integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.users OWNER TO joakimengqvist;

ALTER TABLE ONLY public.users
ADD CONSTRAINT users_pkey PRIMARY KEY (id);

INSERT INTO "public"."users"("email","first_name","last_name", "privileges","projects","password","user_active","created_at","updated_at")
VALUES
(E'sudo',E'sudo',E'user', '{ "user_read", "user_write", "user_sudo", "project_read", "project_write", "project_sudo", "privilege_read", "note_read", "note_write", "note_sudo", "privilege_write", "privilege_sudo", "economics_read", "economics_write", "economics_sudo", "invoice_read", "invoice_write", "invoice_sudo", "product_read", "product_write", "product_sudo", "external_company_read", "external_company_write", "external_company_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'users',E'users',E'user', '{ "user_read", "user_write", "user_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'privileges',E'privileges',E'user', '{ "privilege_read", "privilege_write", "privilege_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'projects',E'projects',E'user', '{ "project_read", "project_write", "project_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'notes',E'notes',E'user', '{ "note_read", "note_write", "note_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'economics',E'economics',E'user', '{ "economics_read", "economics_write", "economics_sudo" }','{}',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public.user_settings (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
    user_id uuid,
	dark_theme boolean NOT NULL DEFAULT false,
	compact_ui boolean NOT NULL DEFAULT false,
    created_at timestamp without time zone,
    created_by uuid,
    updated_at timestamp without time zone,
    updated_by uuid
);

ALTER TABLE public.user_settings OWNER TO joakimengqvist;

ALTER TABLE ONLY public.user_settings
ADD CONSTRAINT user_settings_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.privileges (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    name character varying(255),
    description character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.privileges OWNER TO joakimengqvist;

ALTER TABLE ONLY public.privileges
ADD CONSTRAINT privileges_pkey PRIMARY KEY (id);

INSERT INTO "public"."privileges"("name", "description", "created_at","updated_at")
VALUES
(E'user_read',E'Allowing user to read users',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'user_write',E'Allowing user to write users',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'user_sudo',E'Allowing user to delete users',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'privilege_read',E'Allowing user to read privileges',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'privilege_write',E'Allowing user to write privileges',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'privilege_sudo',E'Allowing user to delete privileges',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'project_read',E'Allowing user to read projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'project_write',E'Allowing user to write projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'project_sudo',E'Allowing user to delete projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'note_read',E'Allowing user to read notes',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'sub_project_read',E'Allowing user to read projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'sub_project_write',E'Allowing user to write projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'sub_project_sudo',E'Allowing user to delete projects',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'note_read',E'Allowing user to read notes',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'note_write',E'Allowing user to write notes',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'note_sudo',E'Allowing user to delete notes',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'economics_read',E'Allowing user to read economics',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'economics_write',E'Allowing user to write economics',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'economics_sudo',E'Allowing user to delete economics',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');
(E'invoice_read',E'Allowing user to read invoices',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'invoice_write',E'Allowing user to write invoices',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'invoice_sudo',E'Allowing user to delete invoices',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');
(E'product_read',E'Allowing user to read products',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'product_write',E'Allowing user to write products',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'product_sudo',E'Allowing user to delete products',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');
(E'external_company_read',E'Allowing user to read external companies',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'external_company_write',E'Allowing user to write external companies',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00'),
(E'external_company_sudo',E'Allowing user to delete external companies',E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');