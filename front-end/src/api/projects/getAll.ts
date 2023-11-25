import { ENDPOINTS } from "../endpoints";

export const getAllProjects = async (loggedInUserId : string) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'GET',
        headers: headers,
    };

    const response = await fetch(ENDPOINTS.GetAllProjects, body)
        .then(response => {
            return response.json()})
        .catch(error => {
            return error
        });

      return response;
    }