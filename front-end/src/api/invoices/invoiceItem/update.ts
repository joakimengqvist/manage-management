import { InvoiceItem } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const updateInvoiceItem = async (
    id: string,
    product_id: string,
    discount_percentage: number,
    total_price: number,
    total_tax: number,
    quantity: number,
    loggedInUserId: string,
) : Promise<{
    error: boolean,
    message: string,
    data: InvoiceItem
}> => {
    const payload = {
        id: id,
        product_id: product_id,
        discount_percentage: discount_percentage,
        total_price: total_price,
        total_tax: total_tax,
        quantity: quantity,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateInvoiceItem, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}