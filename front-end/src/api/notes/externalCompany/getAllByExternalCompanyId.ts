import { ExternalCompanyNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllExternalCompanyNotesByExternalCompanyId = async (loggedInUserId : string, externalCompanyId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<ExternalCompanyNote>
}> => {
    const payload = {
        id: externalCompanyId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllExternalCompanyNotesByExternalCompanyId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }