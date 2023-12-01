import { ProductNote } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllProductNotesByProductId = async (loggedInUserId : string, productId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<ProductNote>
}> => {
    const payload = {
        id: productId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllProductNotesByProductId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }