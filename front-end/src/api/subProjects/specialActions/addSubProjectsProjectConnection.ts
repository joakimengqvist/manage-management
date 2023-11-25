import { ENDPOINTS } from "../../endpoints";

/**
 * @param subProjectIds sub project ids in an array
 */
export const addSubProjectsProjectConnection = async (loggedInUserId : string, subProjectIds : Array<string>, projectId : string) : Promise<{
  error: boolean,
  message: string,
}> => {
    const payload = {
      project_id: projectId,
      sub_project_ids: subProjectIds,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.AddSubProjectsProjectConnection, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }