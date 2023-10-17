package data

import "strings"

func parsePostgresArray(postgresArray string) []string {

	postgresArray = strings.Trim(postgresArray, "{}")

	if len(postgresArray) < 4 {
		return []string{}
	}

	arrayElements := strings.Split(postgresArray, ",")

	return arrayElements
}

func convertToPostgresArray(arrayElements []string) string {
	postgresArray := strings.Join(arrayElements, ",")
	postgresArray = "{" + postgresArray + "}"
	return postgresArray
}

func convertStringToPostgresArray(s string) string {
	returnedS := "{" + s + "}"
	return returnedS
}
