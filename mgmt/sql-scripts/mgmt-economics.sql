-- Database: mgmt-economics

-- DROP DATABASE IF EXISTS "mgmt-economics";
/*
CREATE DATABASE "mgmt-economics"
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

CREATE TABLE expenses (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    project_id VARCHAR(255) NOT NULL,
    expense_date TIMESTAMP NOT NULL,
    expense_category VARCHAR(255) NOT NULL,
    vendor VARCHAR(255) NOT NULL,
    description TEXT,
    amount DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2),
    status VARCHAR(255) NOT NULL DEFAULT 'idle',
    currency VARCHAR(3) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_by VARCHAR(255) NOT NULL,
    modified_at TIMESTAMP NOT NULL
);

ALTER TABLE public.expenses OWNER TO joakimengqvist;

ALTER TABLE ONLY public.expenses
ADD CONSTRAINT expenses_pkey PRIMARY KEY (id);
	
SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE incomes (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    invoice_id uuid,
    project_id VARCHAR(255) NOT NULL,
    income_date TIMESTAMP NOT NULL,
    income_category VARCHAR(255) NOT NULL,
    statistics_income bool DEFAULT false,
    vendor VARCHAR(255) NOT NULL,
    description TEXT,
    amount DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2),
    status VARCHAR(255) NOT NULL DEFAULT 'idle',
    currency VARCHAR(3) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE public.incomes OWNER TO joakimengqvist;

ALTER TABLE ONLY public.incomes
ADD CONSTRAINT incomes_pkey PRIMARY KEY (id);