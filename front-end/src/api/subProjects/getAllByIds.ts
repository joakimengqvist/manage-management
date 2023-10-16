import { ENDPOINTS } from "../endpoints";

export const getSubProjectsByIds = async (userId : string, ids : Array<string>) => {
    const payload = {
      ids: ids,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetSubProjectsByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }