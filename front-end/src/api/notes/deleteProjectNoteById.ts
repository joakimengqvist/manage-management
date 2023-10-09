export const deleteProjectNoteById = async (loggedInUser : string, noteId : string, authorId : string, projectId : string) => {
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

    const response = await fetch("http://localhost:8080/notes/delete-project-note-by-id", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }