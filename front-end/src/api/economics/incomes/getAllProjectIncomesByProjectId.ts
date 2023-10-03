export const getAllProjectIncomesByProjectId = async (userId : string, projectId : string) => {
    const payload = {
        project_id: projectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/economics/get-all-project-incomes-by-project-id", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }