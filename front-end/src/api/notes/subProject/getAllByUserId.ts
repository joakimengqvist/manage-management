import { ENDPOINTS } from "../../endpoints";

export const getAllSubProjectNotesByUserId = async (loggedInUserId : string, userId : string) => {
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

    const response = await fetch(ENDPOINTS.GetAllSubProjectNotesByUserId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }