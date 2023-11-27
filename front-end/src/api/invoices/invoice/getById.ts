import { Invoice } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getInvoiceById = async (loggedInUserId : string, invoiceId : string) : Promise<{
    error: boolean,
    message: string,
    data: Invoice
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

    const response = await fetch(ENDPOINTS.GetInvoiceById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }