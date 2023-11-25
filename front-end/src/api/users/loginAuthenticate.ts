import { ENDPOINTS } from "../endpoints";

/**
 * @param email email of user
 */
export const loginAuthenticate = async (email : string /* password : string */ ) : Promise<{
    error: boolean,
    message: string,
    data: null
}> => {

    const payload = {
          email: email,
          password: "password"
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.Authenticate, body)
      .then(response => {
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }