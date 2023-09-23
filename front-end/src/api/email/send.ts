export const sendEmail = async (userId : string, to : string, from : string, subject : string, message : string) => {
    const payload = {
      action: "mail",
      mail: {
          from,
          to,
          subject,
          message,
      }
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("X-user-id", userId.toString());

    const body = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(payload)
    };

    const respone = await fetch("http://localhost:8080/email/send", body)
      .then(response => {
  
        return response.json()
    }).catch(error => {
          return error;
      });

      return respone;
    }