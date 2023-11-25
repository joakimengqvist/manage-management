import { ENDPOINTS } from "../../endpoints";

/**
 * @param projectIds project ids in an array
 */
export const addProjectsSubProjectConnection = async (
    loggedInUserId : string,
    subProjectId : string,
    projectIds : Array<string>
    ) : Promise<{
    error: boolean,
    message: string,
    data: null
}> => {
    const payload = {
      project_ids: projectIds,
      sub_project_id: subProjectId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

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