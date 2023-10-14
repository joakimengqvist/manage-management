import { ENDPOINTS } from "../endpoints";
export const updatePrivilege = async (userId : string, id : string, name : string, description : string) => {
    const payload = {
      id: id,
      name: name,
      description: description,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdatePrivilege, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }