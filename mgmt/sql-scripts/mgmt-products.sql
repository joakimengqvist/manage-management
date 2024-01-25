-- Database: mgmt-product

-- DROP DATABASE IF EXISTS "mgmt-product";

/* 
CREATE DATABASE "mgmt-product"
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

CREATE TABLE products (
   	id uuid NOT NULL DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(50),
    price DECIMAL(10, 2) NOT NULL,
    tax_percentage DECIMAL(5, 2) DEFAULT 0.00,
    created_at timestamp without time zone,
    created_by uuid,
    updated_at timestamp without time zone,
    updated_by uuid
);

ALTER TABLE public.products OWNER TO joakimengqvist;

ALTER TABLE ONLY public.products
ADD CONSTRAINT products_pkey PRIMARY KEY (id);
