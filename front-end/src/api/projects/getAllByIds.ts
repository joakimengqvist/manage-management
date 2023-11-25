import { ENDPOINTS } from "../endpoints";

export const getProjectsByIds = async (loggedInUserId : string, projectIds : Array<string>) => {
    const payload = {
      ids: projectIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetProjectsByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }