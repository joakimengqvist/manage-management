-- Database: mgmt-project

-- DROP DATABASE IF EXISTS "mgmt-project";

/*
CREATE DATABASE "mgmt-project"
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

CREATE TABLE public.sub_projects (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    name character varying(255),
    description TEXT,
    status character varying(255),
    priority integer,
    start_date timestamp without time zone,
    due_date timestamp without time zone,
    estimated_duration integer,
    notes TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[],
    created_at timestamp without time zone,
    created_by uuid,
    updated_at timestamp without time zone,
    updated_by uuid,
    projects uuid[] NOT NULL DEFAULT ARRAY[]::uuid[],
    invoices uuid[] NOT NULL DEFAULT ARRAY[]::uuid[],
    incomes uuid[] NOT NULL DEFAULT ARRAY[]::uuid[],
    expenses uuid[] NOT NULL DEFAULT ARRAY[]::uuid[]
);

ALTER TABLE public.sub_projects OWNER TO joakimengqvist;

ALTER TABLE ONLY public.sub_projects
ADD CONSTRAINT sub_projects_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.projects (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
    name character varying(255),
    status character varying(255),
    notes TEXT [] NOT NULL DEFAULT ARRAY[]::TEXT[],
    sub_projects TEXT [] NOT NULL DEFAULT ARRAY[]::TEXT[],
    created_at timestamp without time zone,
    created_by uuid,
    updated_at timestamp without time zone,
    updated_by uuid
);

ALTER TABLE public.projects OWNER TO joakimengqvist;

ALTER TABLE ONLY public.projects
ADD CONSTRAINT projects_pkey PRIMARY KEY (id);

INSERT INTO "public"."projects"("name","status","created_at","created_by","updated_at","updated_by")
VALUES
(E'Big project',E'ongoing',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004'),
(E'Medium project',E'ongoing',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004'),
(E'Small project',E'ongoing',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004',E'2022-03-14 00:00:00',E'10048e68-5fc8-11ee-b6cd-0242ac120004');