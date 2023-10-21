import { ENDPOINTS } from "../../endpoints";

export const AddProjectsSubProjectConnection = async (userId : string, subProjectId : string, projectIds : Array<string>) => {
    const payload = {
      project_ids: projectIds,
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

    const response = await fetch(ENDPOINTS.AddProjectsSubProjectConnection, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }