import { ENDPOINTS } from "../endpoints";

/**
 * @returns Resolved promise returns the created project ID
 */
export const createProject = async (loggedInUserId : string, name : string, status : string) => {
    const payload = {
      name: name,
      status: status,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateProject, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }