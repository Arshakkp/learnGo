package roles

import (
	"encoding/json"
	"net/http"

	"example.com/db"
	. "example.com/model"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {

	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	roles := []Roles{}
	db.Find(&roles)
	json, err := json.Marshal(roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(json)

}
