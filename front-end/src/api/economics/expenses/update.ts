import { ENDPOINTS } from "../../endpoints";
import { ExpenseObject } from "../../../interfaces";

export const updateExpense = async (
    id: string,
    project_id: string,
    expense_date: string,
	expense_category: string,
	vendor: string,
	description: string,
	amount: string,
	tax: string,
    status: string,
	currency: string,
	payment_method: string,
	created_by: string
) : Promise<{
    error: boolean,
    message: string,
    data: ExpenseObject
}> => {
    const payload = {
        project_id: project_id,
        id: id,
        expense_date: expense_date,
        expense_category: expense_category,
        vendor: vendor,
        description: description,
        amount: Number(amount),
        tax: Number(tax),
        status: status,
        currency: currency,
        payment_method: payment_method,
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

    const response = await fetch(ENDPOINTS.UpdateExpense, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}