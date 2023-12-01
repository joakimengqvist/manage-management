import { Income } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllIncomesByProjectId = async (loggedInUserId : string, projectId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<Income>
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

    const response = await fetch(ENDPOINTS.GetAllIncomesByProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }