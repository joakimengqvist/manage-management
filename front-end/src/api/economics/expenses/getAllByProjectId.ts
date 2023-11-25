import { ENDPOINTS } from "../../endpoints";
import { ExpenseObject } from "../../../types";

export const getAllExpensesByProjectId = async (loggedInUserId : string, projectId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<ExpenseObject>
}> => {
    const payload = {
        id: projectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllExpensesByProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }