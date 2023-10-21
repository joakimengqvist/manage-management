import { ENDPOINTS } from "../endpoints";

export const updateSubProject = async (
  userId : string,
  id : string,
  name : string,
  status : string,
  description : string,
  priority : number,
  startDate : string,
  dueDate : string,
  estimatedDuration : number,
  notes : string[],
  invoices : string[],
  incomes : string[],
  expenses : string[],
) => {
    const payload = {
      id: id,
      name: name,
      status: status,
      description: description,
      priority: priority,
      start_date: startDate,
      due_date: dueDate,
      estimated_duration: estimatedDuration,
      notes: notes,
      invoices: invoices,
      incomes: incomes,
      expenses: expenses,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

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