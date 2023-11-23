import { ENDPOINTS } from "../../endpoints";

export const createInvoiceItem = async (
    product_id: string,
    original_price: number,
    actual_price: number,
    tax_percentage: number,
    original_tax: number,
    actual_tax: number,
    discount_percentage: number,
    discount_amount: number,
    quantity: number,
    userId: string,
) => {
    const payload = {
        product_id: product_id,
        original_price: original_price,
        actual_price: actual_price,
        tax_percentage: tax_percentage,
        original_tax: original_tax,
        actual_tax: actual_tax,
        discount_percentage: discount_percentage,
        discount_amonut: discount_amount,
        quantity: quantity
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateInvoiceItem, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}