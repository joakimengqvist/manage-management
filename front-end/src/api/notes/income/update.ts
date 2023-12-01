import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";
import { IncomeNote } from "../../../interfaces";

/**
 * @param noteId note id
 * @param user userId, name and email
 * @param title note title
 * @param note note text
 */
export const updateIncomeNote = async (noteId : string, user : NoteAuthor, incomeId : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: IncomeNote
}> => {
    const payload = {
        id: noteId,
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        income_id: incomeId,
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

    const response = await fetch(ENDPOINTS.UpdateIncomeNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }