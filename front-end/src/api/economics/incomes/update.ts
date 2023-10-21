import { ENDPOINTS } from "../../endpoints";

export const updateProjectIncome = async (
    id: string,
    project_id: string,
    income_date: string,
	income_category: string,
	vendor: string,
	description: string,
	amount: string,
	tax: string,
    status: string,
	currency: string,
	payment_method: string,
	created_by: string,
) => {
    const payload = {
        project_id: project_id,
        id: id,
        income_date: income_date,
        income_category: income_category,
        vendor: vendor,
        description: description,
        amount: Number(amount),
        tax: Number(tax),
        status: status,
        currency: currency,
        payment_method: payment_method,
        created_by: created_by,
        modified_by: created_by,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", created_by);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateProjectIncome, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}