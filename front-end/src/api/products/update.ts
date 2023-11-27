import { Product } from "../../interfaces";
import { ENDPOINTS } from "../endpoints";

export const updateProduct = async (
  loggedInUserId : string,
  id : string, 
  name : string, 
  description : string,
  category : string,
  price : number,
  tax : number
  ) : Promise<{
    error: boolean,
    message: string,
    data: Product
  }> => {
    const payload = {
      id: id,
      name: name,
      description: description,
      category: category,
      price: price,
      tax: tax,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateProduct, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }