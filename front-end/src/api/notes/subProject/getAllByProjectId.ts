import { ENDPOINTS } from "../../endpoints";

export const getAllSubProjectNotesBySubProjectId = async (userId : string, subProjectId : string) => {
    const payload = {
        id: subProjectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllSubProjectNotesBySubProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }