export const createExternalCompany = async (
    userId : string,
    company_name: string,
	company_registration_number: string,
	contact_person: string,
	contact_email: string,
	contact_phone: string,
	address: string,
	city: string,
	state_province: string,
	country: string,
	postal_code: string,
	payment_terms: string,
	billing_currency: string,
	bank_account_info: string,
	tax_identification_number: string,
	status: string,
	assigned_projects: Array<string>,
	invoice_pending: Array<string>,
	invoice_history: Array<string>,
	contractual_agreements: Array<string>,
) => {
    
    const payload = {
        company_name,
        company_registration_number,
        contact_person,
        contact_email,
        contact_phone,
        address,
        city,
        state_province,
        country,
        postal_code,
        payment_terms,
        billing_currency,
        bank_account_info,
        tax_identification_number,
        status,
        assigned_projects,
        invoice_pending,
        invoice_history,
        contractual_agreements
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/external-company/create-external-company", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }