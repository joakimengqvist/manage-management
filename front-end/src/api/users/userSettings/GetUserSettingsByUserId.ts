import { ENDPOINTS } from "../../endpoints";

export const getUserSettingsByUserId = async (loggedInUserId : string, id : string) => {
    const payload = {
        id: id,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetUserSettingsByUserId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }