package utils

import (
	"net/http"
)

type ContexKey struct{}

func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(ContexKey{}).([]string)
	return fields[index]
}
