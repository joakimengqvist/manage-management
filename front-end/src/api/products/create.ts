import { ENDPOINTS } from "../endpoints";

export const createProduct = async (
  userId : string, 
  name : string, 
  description : string,
  category : string,
  price : number,
  tax_percentage : number,
  ) => {
    const payload = {
      name: name,
      description: description,
      category: category,
      price: price,
      tax_percentage: tax_percentage,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

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