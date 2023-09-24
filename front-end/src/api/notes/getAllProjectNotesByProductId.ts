// get-all-notes-by-project-id

export const getAllProjectNotesByProjectId = async (userId : string, projectId : string) => {
    const payload = {
        projectId: projectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/notes/get-all-notes-by-project-id", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }