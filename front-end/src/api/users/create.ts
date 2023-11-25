import { ENDPOINTS } from "../endpoints";

/**
 * @returns Resolved promise returns the created product item ID
 */
export const createUser = async (
  loggedInUserId : string,
  firstName : string,
  lastName : string,
  email : string,
  privileges : Array<string>, 
  projects : Array<string>,
  password : string
  ) : Promise<{
    error: boolean,
    message: string,
    data: string
  }> => {
    const payload = {
      first_name: firstName,
      last_name: lastName,
      email: email,
      privileges: privileges,
      projects: projects,
      password: password
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateUser, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }