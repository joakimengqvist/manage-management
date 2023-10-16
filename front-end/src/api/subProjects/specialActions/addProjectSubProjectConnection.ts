import { ENDPOINTS } from "../../endpoints";

export const AddProjectSubProjectConnection = async (userId : string, subProjectId : string, projectId : string) => {
    const payload = {
      project_id: projectId,
      sub_project_id: subProjectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.AddProjectSubProjectConnection, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }