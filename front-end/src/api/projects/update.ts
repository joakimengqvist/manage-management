export const updateProject = async (userId : string, id : string, name : string, status : string) => {
    const payload = {
      id: id,
      name: name,
      status: status
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/project/update-project", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }