import { ENDPOINTS } from "../../endpoints";

export const deleteProjectNote = async (loggedInUser : string, noteId : string, authorId : string, projectId : string | undefined) : Promise<{
    error: boolean,
    message: string,
    data: null
}> => {
  
  if (!projectId) return {
    error: true,
    message: 'Project id is missing',
    data: null
  };

    const payload = {
      note_id: noteId,
      author_id: authorId,
      project_id: projectId
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUser);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.DeleteProjectNote, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }