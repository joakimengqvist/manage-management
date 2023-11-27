import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../interfaces/notes";
import { ProjectNote } from "../../../interfaces";

/**
 * @param noteId note id
 * @param user userId, name and email
 * @param project project id
 * @param title note title
 * @param note note text
 */
export const updateProjectNote = async (noteId : string, user : NoteAuthor, project : string, title : string, note : string) : Promise<{
    error: boolean,
    message: string,
    data: ProjectNote
}> => {
    const payload = {
        id: noteId,
        author_id: user.id,
        author_name: user.name,
        author_email: user.email,
        project: project,
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

    const response = await fetch(ENDPOINTS.UpdateProjectNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }