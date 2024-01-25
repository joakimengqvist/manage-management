-- Database: mgmt-company

-- DROP DATABASE IF EXISTS "mgmt-company";

/* 
CREATE DATABASE "mgmt-company"
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

CREATE TABLE external_companies (
    id uuid NOT NULL DEFAULT gen_random_uuid (),
    company_name VARCHAR(255) NOT NULL,
    company_registration_number VARCHAR(50),
    contact_person VARCHAR(100),
    contact_email VARCHAR(255),
    contact_phone VARCHAR(20),
    address VARCHAR(255),
    city VARCHAR(100),
    state_province VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    payment_terms VARCHAR(50),
    billing_currency VARCHAR(50),
    bank_account_info VARCHAR(255),
    tax_identification_number VARCHAR(50),
    created_at DATE,
    created_by DATE,
    updated_at DATE,
    updated_by DATE,
    status VARCHAR(20),
    assigned_projects TEXT[],
    invoices TEXT[],
    contractual_agreements TEXT[]
);

ALTER TABLE public.external_companies OWNER TO joakimengqvist;

ALTER TABLE ONLY public.external_companies
ADD CONSTRAINT external_companies_pkey PRIMARY KEY (id);