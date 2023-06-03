/* server helps for setting and routing server*/
package server

import (
	"fmt"
	"log"
	"net/http"

	. "example.com/class"
	. "example.com/org"
	. "example.com/roles"
	. "example.com/users"
	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Starting Server.....")
	route()

}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Access-Control-Allow-Origin header to allow requests from any origin

		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func route() {
	router := mux.NewRouter()
	router.Handle("/org", corsMiddleware(http.HandlerFunc((GetOrg)))).Methods("GET")
	router.Handle("/classes", corsMiddleware(http.HandlerFunc((GetClasses)))).Methods("GET")
	router.Handle("/{userId}/classandrole", corsMiddleware(http.HandlerFunc((GetClassAndRole)))).Methods("GET")
	router.Handle("/{userId}/classandrolefile", corsMiddleware(http.HandlerFunc((GetClassAndRoleFile)))).Methods("GET")
	router.Handle("/{userId}/addclassandrole", corsMiddleware(http.HandlerFunc((AddClassRole)))).Methods("POST")
	router.Handle("/{userId}/addclassandrolefile", corsMiddleware(http.HandlerFunc((AddClassRoleThroughFile)))).Methods("POST")
	router.Handle("/roles", corsMiddleware(http.HandlerFunc((GetRoles)))).Methods("GET")
	router.Handle("/orgs/{page}", corsMiddleware(http.HandlerFunc((Orgs)))).Methods("GET")
	router.Handle("/add/org", corsMiddleware(http.HandlerFunc((AddOrg)))).Methods("POST")
	router.Handle("/edit/org/{id}", corsMiddleware(http.HandlerFunc((EditOrg)))).Methods("POST")
	router.Handle("/delete/org/{id}", corsMiddleware(http.HandlerFunc((DeleteOrg)))).Methods("GET")
	router.Handle("/org/{id}/users/{page}", corsMiddleware(http.HandlerFunc((GetUsers)))).Methods("GET")
	router.Handle("/org/user/{id}", corsMiddleware(http.HandlerFunc((GetUser)))).Methods("GET")
	router.Handle("/org/edit/user/{id}", corsMiddleware(http.HandlerFunc((EditUser)))).Methods("POST")
	router.Handle("/org/delete/{id}", corsMiddleware(http.HandlerFunc((DeleteUser)))).Methods("GET")
	router.Handle("/org/{id}/adduser", corsMiddleware(http.HandlerFunc((AddUser)))).Methods("POST")
	router.Handle("/org/{id}/adduserfile", corsMiddleware(http.HandlerFunc((AddUsersWithFile)))).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
