import { ENDPOINTS } from "../../endpoints";

/**
 * @returns Resolved promise returns the created income ID
 */
export const createIncome = async (
    project_id: string,
    income_date: string,
	income_category: string,
    statistics_income: boolean,
	vendor: string,
	description: string,
	amount: string,
	tax: string,
    status: string,
	currency: string,
	created_by: string,
) : Promise<{
    error: boolean,
    message: string,
    data: string
}>=> {
    const payload = {
        project_id: project_id,
        income_date: income_date,
        income_category: income_category,
        statistics_income: statistics_income,
        vendor: vendor,
        description: description,
        amount: Number(amount),
        tax: Number(tax),
        status: status,
        currency: currency,
        created_by: created_by,
        updated_by: created_by,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", created_by);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateIncome, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}