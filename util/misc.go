package util

import (
	"fmt"

	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
)

func CountUserDbStructs(collection []db.User) int {
	count := 0

	for _, item := range collection {
		count++

		fmt.Println(item)
	}

	return count
}
