
CREATE USER mgmt_auth_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_auth_autobot;
GRANT azure_pg_admin TO mgmt_auth_autobot;

CREATE USER mgmt_project_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_project_autobot;
GRANT azure_pg_admin TO mgmt_project_autobot;

CREATE USER mgmt_company_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_company_autobot;
GRANT azure_pg_admin TO mgmt_company_autobot;

CREATE USER mgmt_economics_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_economics_autobot;
GRANT azure_pg_admin TO mgmt_economics_autobot;

CREATE USER mgmt_product_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_product_autobot;
GRANT azure_pg_admin TO mgmt_product_autobot;

CREATE USER mgmt_notes_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_notes_autobot;
GRANT azure_pg_admin TO mgmt_notes_autobot;

CREATE USER mgmt_invoice_autobot PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA public TO mgmt_invoice_autobot;
GRANT azure_pg_admin TO mgmt_invoice_autobot;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_auth_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_project_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_company_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_economics_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_product_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_notes_autobot;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mgmt_invoice_autobot;