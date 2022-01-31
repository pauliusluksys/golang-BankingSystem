package app

import (
	"bankingV2/domain"
	"bankingV2/errs"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

var routesWithoutAuth = []string{"Welcome", "InvestmentCreate", "InvestmentCompanyCreate", "InvestmentRiskLevelCreate", "InvestmentCategoryCreate", "CreateCustomerInvestment", "GetAllCustomerInvestments"}

func (a AuthMiddleware) authorizationHandler() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")
			if !isAuthorizationNeeded(currentRoute.GetName()) {
				next.ServeHTTP(w, r)
			} else {
				if authHeader != "" {
					token := getTokenFromHeader(authHeader)
					fmt.Println(token)
					isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

					if isAuthorized {
						next.ServeHTTP(w, r)
					} else {
						appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
						writeResponse(w, appError.Code, appError.AsMessage())
					}
				} else {
					writeResponse(w, http.StatusUnauthorized, "missing token")
				}
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	return header
}
func isAuthorizationNeeded(currentRouteName string) bool {
	for _, routeNameWA := range routesWithoutAuth {
		if routeNameWA == currentRouteName {
			return false
		}
	}
	return true
}
