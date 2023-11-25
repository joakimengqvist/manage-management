import { ENDPOINTS } from "../endpoints";

/**
 * @returns Resolved promise returns the created product item ID
 */
export const createProduct = async (
  loggedInUserId : string, 
  name : string, 
  description : string,
  category : string,
  price : number,
  tax_percentage : number,
  ) : Promise<{
    error: boolean,
    message: string,
    data: string
  }> => {
    const payload = {
      name: name,
      description: description,
      category: category,
      price: price,
      tax_percentage: tax_percentage,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateProduct, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }