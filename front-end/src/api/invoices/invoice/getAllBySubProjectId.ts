import { Invoice } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const getAllInvoicesBySubProjectId = async (loggedInUserId : string, subProjectId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<Invoice>
}> => {
    const payload = {
        id: subProjectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllInvoicesBySubProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }