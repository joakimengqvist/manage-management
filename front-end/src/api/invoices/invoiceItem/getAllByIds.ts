import { ENDPOINTS } from "../../endpoints";

export const getInvoiceItemsByIds = async (loggedInUserId : string, invoiceItemIds : Array<string>) => {
    const payload = {
      ids: invoiceItemIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.getAllInvoiceItemsByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }