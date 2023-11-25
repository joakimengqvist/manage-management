import { ENDPOINTS } from "../endpoints";

export const deleteUser = async (loggedInUserId : string, userId : string) : Promise<{
    error: boolean,
    message: string
}> => {
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

    const response = await fetch(ENDPOINTS.DeleteUser, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }