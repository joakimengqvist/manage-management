import { Invoice } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const updateInvoice = async (
    id: string,
    company_id: string,
    project_id: string,
    sub_project_id: string,
    invoice_display_name: string,
    invoice_description: string,
    status: string,
    invoiceDate: string,
    dueDate: string,
    paymentDate: string,
    invoice_items: Array<never>,
    total_amount: number,
    discount_percentage: number,
    total_price: number,
    units_combined_price: number,
    total_tax: number,
    loggedInUserId: string,
) : Promise<{
    error: boolean,
    message: string,
    data: Invoice
}> => {
    const payload = {
        id: id,
        company_id: company_id,
        project_id: project_id,
        sub_project_id: sub_project_id,
        invoice_display_name: invoice_display_name,
        invoice_description: invoice_description,
        status: status,
        invoiceDate: invoiceDate,
        dueDate: dueDate,
        paymentDate: paymentDate,
        invoice_items: invoice_items,
        total_amount: total_amount,
        discount_percentage: discount_percentage,
        total_price: total_price,
        units_combined_price: units_combined_price,
        total_tax: total_tax,
        userId: loggedInUserId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateInvoice, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}