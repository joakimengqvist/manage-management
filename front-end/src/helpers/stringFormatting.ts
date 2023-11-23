export const replaceUnderscoreAndCapitalize = (inputString : string) => {
    const stringWithSpaces = inputString.replace(/_/g, ' ');
    const capitalizedString = stringWithSpaces.charAt(0).toUpperCase() + stringWithSpaces.slice(1);
    return capitalizedString;
  }

  export const replaceUnderscore = (inputString : string) => {
    const stringWithSpaces = inputString.replace(/_/g, ' ');
    return stringWithSpaces;
  }

export const formatNumberWithSpaces = (input : string | number) => {
  if (typeof input === 'number') {
      return input.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
  } 
  
  if (typeof input === 'string') {
      const number = parseFloat(input);
      if (!isNaN(number)) {
          return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
      } 
    }
    
  return input;
    
}