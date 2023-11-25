import { ENDPOINTS } from "../endpoints";

/**
 * @returns Resolved promise returns the created privilege ID
 */
export const createPrivilege = async (loggedInUserId : string, name : string, description : string) : Promise<{
    error: boolean,
    message: string,
    data: string
}> => {
    const payload = {
            name: name,
            description: description
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreatePrivilege, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }