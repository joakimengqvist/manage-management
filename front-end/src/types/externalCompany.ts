export type ExternalCompany = {
	id: string,
    company_name: string
	company_registration_number: string
	contact_person: string
	contact_email: string
	contact_phone: string
	address: string
	city: string
	state_province: string
	country: string
	postal_code: string
	payment_terms: string
	billing_currency: string
	bank_account_info: string,
	tax_identification_number: string
	status: string
	assigned_projects: Array<string>
	invoice_pending: Array<string>
	invoice_history: Array<string>
	contractual_agreements: Array<string>
	created_at: string
	created_by: string
	updated_at: string
	updated_by: string
}

