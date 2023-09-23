export const createUser = async (
  userId : string,
  firstName : string,
  lastName : string,
  email : string,
  privileges : Array<string>, 
  projects : Array<string>,
  password : string
  ) => {
    const payload = {
      first_name: firstName,
      last_name: lastName,
      email: email,
      privileges: privileges,
      projects: projects,
      password: password
    };

    console.log('payload create user', payload)

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/auth/create-user", body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }