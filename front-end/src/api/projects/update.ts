import { ENDPOINTS } from "../endpoints";

export const updateProject = async (loggedInUserId : string, projectId : string, name : string, status : string) => {
    const payload = {
      id: projectId,
      name: name,
      status: status
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateProject, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }