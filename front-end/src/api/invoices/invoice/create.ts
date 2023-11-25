import { ENDPOINTS } from "../../endpoints";

/**
 * @returns Resolved promise returns the created invoice ID
 */
export const createInvoice = async (
    company_id: string,
    project_id: string,
    sub_project_id: string,
    invoice_display_name: string,
    invoice_description: string,
    statistics_invoice: boolean,
    invoice_items: Array<never>,
    original_price: number,
    actual_price: number,
    discount_percentage: number,
    discount_amount: number,
    original_tax: number,
    actual_tax: number,
    invoiceDate: string,
    dueDate: string,
    paid: boolean,
    status: string,
    paymentDate: string,
    loggedInUserId: string
) : Promise<{
    error: boolean,
    message: string,
    data: string
}> => {
    const payload = {
        company_id: company_id,
        project_id: project_id,
        sub_project_id: sub_project_id,
        invoice_display_name: invoice_display_name,
        invoice_description: invoice_description,
        statistics_invoice: statistics_invoice,
        invoice_items: invoice_items,
        original_price: original_price,
        actual_price: actual_price,
        discount_percentage: discount_percentage,
        discount_amount: discount_amount,
        original_tax: original_tax,
        actual_tax: actual_tax,
        invoiceDate: invoiceDate,
        dueDate: dueDate,
        paid: paid,
        status: status,
        paymentDate: paymentDate,
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

    const response = await fetch(ENDPOINTS.CreateInvoice, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}