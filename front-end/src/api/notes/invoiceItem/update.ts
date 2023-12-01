import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";
import { InvoiceItemNote } from "../../../interfaces";

/**
 * @param noteId note id
 * @param user userId, name and email
 * @param title note title
 * @param note note text
 */
export const updateInvoiceItemNote = async (noteId : string, user : NoteAuthor, invoiceItemId : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: InvoiceItemNote
}> => {
    const payload = {
        id: noteId,
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        invoice_item_id: invoiceItemId,
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

    const response = await fetch(ENDPOINTS.UpdateInvoiceItemNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }