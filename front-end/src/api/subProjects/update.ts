import { SubProject } from "../../interfaces";
import { ENDPOINTS } from "../endpoints";

export const updateSubProject = async (
  loggedInUserId : string,
  subProjectId : string,
  name : string,
  status : string,
  description : string,
  priority : number,
  startDate : string,
  dueDate : string,
  estimatedDuration : number,
) : Promise<{
  error: boolean,
  message: string,
  data: SubProject
}> => {
    const payload = {
      id: subProjectId,
      name: name,
      status: status,
      description: description,
      priority: priority,
      start_date: startDate,
      due_date: dueDate,
      estimated_duration: estimatedDuration,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateSubProject, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }