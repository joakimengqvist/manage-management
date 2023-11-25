import { ProjectNote } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

export const getAllProjectNotesByProjectId = async (loggedInUserId : string, projectId : string) : Promise<{
    error: boolean,
    message: string,
    data: Array<ProjectNote>
}> => {
    const payload = {
        id: projectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllProjectNotesByProjectId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }