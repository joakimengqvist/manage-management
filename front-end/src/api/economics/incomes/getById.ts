import { IncomeObject } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const getIncomeById = async (loggedInUserId : string, incomeId : string) : Promise<{
    error: boolean,
    message: string,
    data: IncomeObject
}> => {
    const payload = {
      id: incomeId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetIncomeById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }