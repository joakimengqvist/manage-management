import { ENDPOINTS } from "../../endpoints";
import { ExpenseObject } from "../../../interfaces";

export const getExpenseById = async (loggedInUserId : string, expenseId : string) : Promise<{
    error: boolean,
    message: string,
    data: ExpenseObject
}> => {
    const payload = {
      id: expenseId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetExpenseById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }