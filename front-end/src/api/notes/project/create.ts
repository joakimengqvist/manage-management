import { ENDPOINTS } from "../../endpoints";
import { NoteAuthor } from "../../../types/state";

export const createProjectNote = async (user : NoteAuthor, project : string, title : string, note : string) => {
    const payload = {
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

    const response = await fetch(ENDPOINTS.CreateProjectNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
}