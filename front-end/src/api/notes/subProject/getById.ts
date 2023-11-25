import { SubProjectNote } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const getSubProjectNoteById = async (loggedInUserId : string, noteId : string) : Promise<{
    error: boolean,
    message: string,
    data: SubProjectNote
}> => {
    const payload = {
      id: noteId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetSubProjectNoteById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }