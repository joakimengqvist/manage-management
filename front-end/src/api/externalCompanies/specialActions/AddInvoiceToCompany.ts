import { ENDPOINTS } from "../../endpoints";

export const addInvoiceToCompany = async (
    loggedInUserId : string,
    companyId : string,
    invoiceId : string
    ) : Promise<{
    error: boolean,
    message: string,
    data: null
}> => {
    const payload = {
      company_id: companyId,
      invoice_id: invoiceId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.addInvoiceFromCompany, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }