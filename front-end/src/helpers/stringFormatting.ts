export const replaceUnderscoreAndCapitalize = (inputString : string) => {
    const stringWithSpaces = inputString.replace(/_/g, ' ');
    const capitalizedString = stringWithSpaces.charAt(0).toUpperCase() + stringWithSpaces.slice(1);
    return capitalizedString;
  }

  export const replaceUnderscore = (inputString : string) => {
    const stringWithSpaces = inputString.replace(/_/g, ' ');
    return stringWithSpaces;
  }