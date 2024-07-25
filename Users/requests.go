package Users

import "github.com/ARTEM-SLV/Users/internal/service"

func Request(token string, userID int) string {
	if token != "token" {
		return "not authorized"
	}

	return service.GetUser(userID)
}
