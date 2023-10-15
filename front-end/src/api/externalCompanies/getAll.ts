import { ENDPOINTS } from "../endpoints";

export const getAllExternalCompanies = async (userId : string) => {

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'GET',
        headers: headers,
    };

    const response = await fetch(ENDPOINTS.GetAllExternalCompanies, body)
        .then(response => {
            return response.json()})
        .catch(error => {
            return error
        });

    return response;
}