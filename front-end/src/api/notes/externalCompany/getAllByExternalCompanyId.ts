import { ENDPOINTS } from "../../endpoints";

export const getAllExternalCompanyNotesByExternalCompanyId = async (userId : string, externalCompanyId : string) => {
    const payload = {
        id: externalCompanyId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllExternalCompanyNotesByExternalCompanyId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }