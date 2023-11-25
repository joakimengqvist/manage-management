import { ExpenseNote } from "../../../types";
import { ENDPOINTS } from "../../endpoints";

/**
 * @param authorId author id
 */
export const getAllExpenseNotesByUserId = async (loggedInUserId : string, authorId : string) : Promise<{
  error: boolean,
  message: string,
  data: Array<ExpenseNote>
}> => {
    const payload = {
      id: authorId,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.GetAllExpenseNotesByUserId, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }