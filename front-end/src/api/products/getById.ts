import { Product } from "../../types";
import { ENDPOINTS } from "../endpoints";

/**
 * @param productId product id
 */
export const getProductById = async (loggedInUserId : string, productId : string) : Promise<{
    error: boolean,
    message: string,
    data: Product
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

    const response = await fetch(ENDPOINTS.GetProductById, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }