export const updateUserCall = async (
  userId : string, 
  id : string, 
  firstName : string, 
  lastName : string, 
  email : string, 
  privileges : Array<string>,
  projects : Array<string>
  ) => {
    const payload = {
        id: id,
        first_name: firstName,
        last_name: lastName,
        email: email,
        projects: projects,
        privileges: privileges
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/auth/update-user", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }