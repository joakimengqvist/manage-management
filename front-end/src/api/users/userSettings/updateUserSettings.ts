import { ENDPOINTS } from "../../endpoints";

export const updateUserSettings = async (
    loggedInUserId: string,
    userId : string, 
    dark_theme : boolean,
    compact_ui : boolean,
    ) => {
    const payload = {
        user_id : userId, 
        dark_theme : dark_theme,
        compact_ui : compact_ui,
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", loggedInUserId);

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const response = await fetch(ENDPOINTS.UpdateUserSettings, body)
      .then(response => { 
        return response.json()
      })
      .catch(error => {
          return error
      });

      return response
    }