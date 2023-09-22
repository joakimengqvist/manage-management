export const testBroker = async () => {
    const response = await fetch("http://localhost:8080", {method: 'POST'})
      .then(response => {
        return response.json();
    }).catch(error => {
          return error;
      });
      return response
    }