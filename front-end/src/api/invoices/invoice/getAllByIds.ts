import { ENDPOINTS } from "../../endpoints";

export const getInvoicesByIds = async (loggedInUserId : string, invoiceIds : Array<string>) => {
    const payload = {
      ids: invoiceIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.getAllInvoicesByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }