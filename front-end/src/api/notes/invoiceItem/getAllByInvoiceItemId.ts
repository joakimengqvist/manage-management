import { InvoiceItemNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllInvoiceItemNotesByInvoiceItemId = async (loggedInUserId : string, invoiceId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<InvoiceItemNote>
}> => {
    const payload = {
        id: invoiceId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllInvoiceItemNotesByInvoiceItemId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }