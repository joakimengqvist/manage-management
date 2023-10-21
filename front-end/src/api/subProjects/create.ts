import { ENDPOINTS } from "../endpoints";

export const createSubProject = async (
  userId : string,
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
      name: name,
      description: description,
      status: status,
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

    const response = await fetch(ENDPOINTS.CreateSubProject, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }