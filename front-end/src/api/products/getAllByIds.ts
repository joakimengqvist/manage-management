import { Product } from "../../interfaces";
import { ENDPOINTS } from "../endpoints";

/**
 * @param productIds array of product ids
 */
export const getProductsByIds = async (loggedInUserId : string, productIds : Array<string>) : Promise<{
    error: boolean,
    message: string,
    data: Array<Product>
}> => {
    const payload = {
      ids: productIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.getProductsByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }