import { ProductNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getProductNoteById = async (loggedInUserId : string, noteId : string) : Promise<{
    error: boolean,
    message: string,
    data: ProductNote
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

    const response = await fetch(ENDPOINTS.GetProductNoteById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }