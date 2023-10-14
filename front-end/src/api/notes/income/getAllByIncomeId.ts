import { ENDPOINTS } from "../../endpoints";
export const getAllIncomeNotesByIncomeId = async (userId : string, incomeId : string) => {
    const payload = {
        id: incomeId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

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