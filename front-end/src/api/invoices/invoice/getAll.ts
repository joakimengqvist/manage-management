import { Invoice } from "../../../interfaces";
import { ENDPOINTS } from "../../endpoints";

export const getAllInvoices = async (loggedInUserId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<Invoice>
}> => {

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'GET',
        headers: headers,
    };

    const response = await fetch(ENDPOINTS.GetAllInvoices, body)
        .then(response => {
            return response.json()})
        .catch(error => {
            return error
        });

      return response;
    }