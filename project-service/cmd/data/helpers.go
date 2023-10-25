package data

import "strings"

func convertToPostgresArray(arrayElements []string) string {
	postgresArray := strings.Join(arrayElements, ",")
	postgresArray = "{" + postgresArray + "}"
	return postgresArray
}
