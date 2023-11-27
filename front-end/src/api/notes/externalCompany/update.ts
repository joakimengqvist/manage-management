import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";
import { ExternalCompanyNote } from "../../../interfaces";

/**
 * @param noteId note id
 * @param user userId, name and email
 * @param external_company external company id
 * @param title note title
 * @param note note text
 */
export const updateExternalCompanyNote = async (noteId : string, user : NoteAuthor, external_company : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: ExternalCompanyNote
}> => {
    const payload = {
        id: noteId,
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        external_company: external_company,
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

    const response = await fetch(ENDPOINTS.UpdateExternalCompanyNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }