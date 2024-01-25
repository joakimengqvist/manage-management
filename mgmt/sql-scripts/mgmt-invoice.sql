-- Database: mgmt-invoice

-- DROP DATABASE IF EXISTS "mgmt-invoice";
/*
CREATE DATABASE "mgmt-invoice"
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

CREATE TABLE invoices (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    company_id uuid,
    project_id uuid,
    sub_project_id uuid,
    income_id uuid,
    invoice_display_name VARCHAR(255),
    invoice_description VARCHAR(255),
    statistics_invoice bool DEFAULT false,
    invoice_items TEXT [],
    original_price DECIMAL(10, 2),
    actual_price DECIMAL(10, 2),    
    discount_percentage DECIMAL(10, 2),
    discount_amount DECIMAL(10, 2),
    original_tax DECIMAL(10, 2),
    actual_tax DECIMAL(10, 2),
    invoice_date timestamp without time zone,
    due_date timestamp without time zone,
    payment_date timestamp without time zone,
    paid BOOLEAN DEFAULT FALSE,
    status VARCHAR(255) NOT NULL,
    created_at timestamp without time zone,
    created_by VARCHAR(255) NOT NULL,
    updated_at timestamp without time zone,
    updated_by VARCHAR(255) NOT NULL
);

ALTER TABLE public.invoices OWNER TO joakimengqvist;

ALTER TABLE ONLY public.invoices
ADD CONSTRAINT invoices_pkey PRIMARY KEY (id);

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE invoice_items (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    product_id uuid,
    discount_percentage DECIMAL(10, 2),
    discount_amount DECIMAL(10, 2),
    original_price DECIMAL(10, 2),
    actual_price DECIMAL(10, 2),
    tax_percentage DECIMAL(10, 2),
    original_tax DECIMAL(10, 2),
    actual_tax DECIMAL(10, 2),
    quantity INT,
    created_at timestamp without time zone,
    created_by VARCHAR(255) NOT NULL,
    updated_at timestamp without time zone,
    updated_by VARCHAR(255) NOT NULL
);

ALTER TABLE public.invoice_items OWNER TO joakimengqvist;

ALTER TABLE ONLY public.invoice_items
ADD CONSTRAINT invoice_items_pkey PRIMARY KEY (id);