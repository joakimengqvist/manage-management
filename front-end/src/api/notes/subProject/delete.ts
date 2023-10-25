import { ENDPOINTS } from "../../endpoints";

export const deleteSubProjectNote = async (loggedInUser : string, noteId : string, authorId : string, subProjectId : string | undefined) => {
  
  if (!subProjectId) return;

    const payload = {
      note_id: noteId,
      author_id: authorId,
      sub_project_id: subProjectId
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUser);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.DeleteSubProjectNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }