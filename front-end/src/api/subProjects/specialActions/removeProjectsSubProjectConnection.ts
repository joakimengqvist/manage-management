import { ENDPOINTS } from "../../endpoints";

/**
 * @param projectIds project ids in an array
 */
export const removeProjectsSubProjectConnection = async (
  loggedInUserId : string, 
  subProjectId : string, 
  projectIds : Array<string>
  ) : Promise<{
  error: boolean,
  message: string,
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

    const response = await fetch(ENDPOINTS.RemoveProjectsSubProjectConnection, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }