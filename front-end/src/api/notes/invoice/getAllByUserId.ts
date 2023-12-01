import { InvoiceNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

/**
 * @param loggedInUserId logged in user id
 */
export const getAllInvoiceNotesByUserId = async (loggedInUserId : string, authorId : string) : Promise<{
  error: boolean,
  message: string,
  data: Array<InvoiceNote>
}> => {
    const payload = {
      id: authorId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllInvoiceNotesByUserId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }