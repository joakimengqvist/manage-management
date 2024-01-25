-- Database: mgmt-note

-- DROP DATABASE IF EXISTS "mgmt-note";

/*
CREATE DATABASE "mgmt-note"
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

CREATE TABLE public.expense_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    expense_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.expense_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.expense_notes
ADD CONSTRAINT expense_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.external_company_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    external_company_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.external_company_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.external_company_notes
ADD CONSTRAINT external_company_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.income_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    income_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.income_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.income_notes
ADD CONSTRAINT income_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.invoice_item_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    invoice_item_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.invoice_item_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.invoice_item_notes
ADD CONSTRAINT invoice_item_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.invoice_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    invoice_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.invoice_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.invoice_notes
ADD CONSTRAINT invoice_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.product_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    product_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.product_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.product_notes
ADD CONSTRAINT product_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.project_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    project_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.project_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.project_notes
ADD CONSTRAINT project_notes_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.sub_project_notes (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
   	author_id character varying(255),
   	author_name character varying(255),
   	author_email character varying(255),
    sub_project_id uuid,
    title character varying(255),
    note TEXT,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.sub_project_notes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.sub_project_notes
ADD CONSTRAINT sub_project_notes_pkey PRIMARY KEY (id);