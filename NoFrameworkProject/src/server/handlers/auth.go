package handlers

import (
	"GoTask1/src/store"
	"fmt"
	"net/http"
)

// HandleBasicAuth code is required auth level
// 1 - any, 2 - waiter, 3 - admin
// if u pass 2 but user is admin in database func will return false
func HandleBasicAuth(r *http.Request, s *store.Store, code int) bool {
	if code < 1 || code > 3 {
		return false
	}
	log, pass, ok := r.BasicAuth()
	fmt.Println(log, pass, ok)
	if ok {
		u, err := s.User().Read(log)
		if err != nil {
			fmt.Println("1")
			return false
		}
		err, b := u.CompareWithHash(pass)
		if b == false || err != nil {
			fmt.Println(b, err, "2")
			return b
		}
		u.Sanitize()
		if code == 2 && u.Role != "waiter" {
			fmt.Println("3")
			return false
		} else if code == 3 && u.Role != "admin" {
			fmt.Println("4")
			return false
		}
		return b
	}
	return false
}
