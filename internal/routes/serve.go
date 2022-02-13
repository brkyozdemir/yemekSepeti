package routes

import (
	"context"
	"fmt"
	"github.com/yemekSepeti/internal/utils"
	httpLog "github.com/yemekSepeti/internal/utils/log"
	"net/http"
	"strings"
)

func Serve(routes []route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var allow []string

		w.Header().Set("Content-Type", "application/json")
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}

				ctx := context.WithValue(r.Context(), utils.ContexKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))

				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)

			return
		}
		http.NotFound(w, r)
	}
}

func Start(routes []route) {
	fmt.Println("Route is starting...")
	http.ListenAndServe(":9000", httpLog.Logger(Serve(routes)))
}
