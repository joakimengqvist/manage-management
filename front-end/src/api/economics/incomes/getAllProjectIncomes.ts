// /economics/get-all-project-expenses

export const getAllProjectIncomes = async (userId : string) => {

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'GET',
        headers: headers,
    };

    const response = await fetch("http://localhost:8080/economics/get-all-project-incomes", body)
        .then(response => {
            return response.json()})
        .catch(error => {
            return error
        });

      return response;
    }