import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";

/**
 * @param user userId, name and email
 * @param title note title
 * @param note note text
 * @returns Resolved promise returns the created invoice note ID
 */
export const createInvoiceNote = async (user : NoteAuthor, invoiceId : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: string
}> => {
    const payload = {
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        invoice_id: invoiceId,
        title: title,
        note: note
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", user.id);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.CreateInvoiceNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}