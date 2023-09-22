export const getAllUsers = async (userId : number) => {

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'GET',
        headers: headers,
    };

    const response = await fetch("http://localhost:8080/auth/get-all-users", body)
        .then(response => {
            return response.json()})
        .catch(error => {
            return error
        });

      return response;
    }