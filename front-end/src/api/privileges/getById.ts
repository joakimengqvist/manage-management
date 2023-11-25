import { Privilege } from "../../types";
import { ENDPOINTS } from "../endpoints";

export const getPrivilegeById = async (loggedInUserId : string, privilegeId : string) : Promise<{
    error: boolean,
    message: string,
    data: Privilege
}> => {
    const payload = {
      id: privilegeId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetPrivilegeById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }