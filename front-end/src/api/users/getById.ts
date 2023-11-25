import { User } from "../../types";
import { ENDPOINTS } from "../endpoints";

export const getUserById = async (loggedInUserId : string, userId : string) : Promise<{
    error: boolean,
    message: string,
    data: User
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

    const response = await fetch(ENDPOINTS.GetUserById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }