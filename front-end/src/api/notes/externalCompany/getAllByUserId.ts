import { ENDPOINTS } from "../../endpoints";

export const getAllExternalCompanyNotesByUserId = async (loggedInUserId : string, userId : string) => {
    const payload = {
      id: userId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllExternalCompanyNotesByUserId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }