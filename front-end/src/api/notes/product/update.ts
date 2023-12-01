import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";
import { ProductNote } from "../../../interfaces";

/**
 * @param noteId note id
 * @param user userId, name and email
 * @param title note title
 * @param note note text
 */
export const updateProductNote = async (noteId : string, user : NoteAuthor, productId : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: ProductNote
}> => {
    const payload = {
        id: noteId,
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        product_id: productId,
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

    const response = await fetch(ENDPOINTS.UpdateProductNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }