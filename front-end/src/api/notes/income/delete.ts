import { ENDPOINTS } from "../../endpoints";

/**
 * @param loggedInUserId logged in user id
 */
export const deleteIncomeNote = async (loggedInUser : string, noteId : string) : Promise<{
    error: boolean,
    message: string,
    data: null
}> => {
    const payload = {
      id: noteId
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUser);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.DeleteIncomeNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }