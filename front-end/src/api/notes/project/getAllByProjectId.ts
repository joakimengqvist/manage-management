import { ENDPOINTS } from "../../endpoints";

export const getAllProjectNotesByProjectId = async (userId : string, projectId : string) => {
    const payload = {
        id: projectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllProjectNotesByProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }