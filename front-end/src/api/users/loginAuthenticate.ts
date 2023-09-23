export const loginAuthenticate = async (email : string, password : string) => {
    console.log(password)
    const payload = {
          email: email,
          password: "password"
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch("http://localhost:8080/auth/authenticate", body)
      .then(response => {
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }