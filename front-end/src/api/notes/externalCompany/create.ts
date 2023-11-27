import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";

/**
 * @param user userId, name and email
 * @param externalCompany external company id
 * @param title note title
 * @param note note text
 * @returns Resolved promise returns the created external company note ID
 */
export const createExternalCompanyNote = async (user : NoteAuthor, externalCompany : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: string
}> => {
    const payload = {
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        external_company: externalCompany,
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

    const response = await fetch(ENDPOINTS.CreateExternalCompanyNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}