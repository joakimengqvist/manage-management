export const getExternalCompanyById = async (userId : string, id : string) => {
    const payload = {
      id: id,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/external-company/get-external-company-by-id", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }