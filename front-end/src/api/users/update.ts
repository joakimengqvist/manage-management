import { User } from "../../interfaces";
import { ENDPOINTS } from "../endpoints";

export const updateUser = async (
  loggedInUserId : string, 
  id : string, 
  firstName : string, 
  lastName : string, 
  email : string, 
  privileges : Array<string>,
  projects : Array<string>
  ) : Promise<{
    error: boolean,
    message: string,
    data: User
  }> => {
    const payload = {
        id: id,
        first_name: firstName,
        last_name: lastName,
        email: email,
        projects: projects,
        privileges: privileges
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateUser, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }