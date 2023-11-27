import { SubProject } from "../../interfaces";
import { ENDPOINTS } from "../endpoints";

export const getSubProjectsByIds = async (loggedInUserId : string, subProjectIds : Array<string>) : Promise<{
  error: boolean,
  message: string,
  data: Array<SubProject>
}>=> {
    const payload = {
      ids: subProjectIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetSubProjectsByIds, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }