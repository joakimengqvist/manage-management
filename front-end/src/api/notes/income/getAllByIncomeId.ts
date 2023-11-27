import { IncomeNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllIncomeNotesByIncomeId = async (loggedInUserId : string, incomeId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<IncomeNote>
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

    const response = await fetch(ENDPOINTS.GetAllIncomeNotesByIncomeId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }