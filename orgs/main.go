/*orgs give function for getting adding organisation to and from db */
package orgs

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"example.com/db"
	. "example.com/model"
	"github.com/gorilla/mux"
)

func GetOrg(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	org := Org{}
	intVar, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db.First(&org, intVar)
	jsonData, err := json.Marshal(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
func EditOrg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	vars := mux.Vars(r)
	idString := vars["id"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)

	org := Org{}
	org.ID = id

	errmssg := json.NewDecoder(r.Body).Decode(&org)
	if errmssg != nil {
		http.Error(w, errmssg.Error(), http.StatusInternalServerError)
	}
	db, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer sqlDB.Close()
	dbresult := db.Save(&org)
	if dbresult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func DeleteOrg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	vars := mux.Vars(r)
	idString := vars["id"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)

	org := Org{}
	user := User{}

	// errmssg := json.NewDecoder(r.Body).Decode(&org)
	// if errmssg != nil {
	// 	http.Error(w, errmssg.Error(), http.StatusInternalServerError)
	// }
	db, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer sqlDB.Close()
	dbresult := db.Where("id=?", id).Delete(&org)
	dbUserResult := db.Where("org_id=?", id).Delete(&user)
	if dbresult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if dbUserResult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func Orgs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageString := vars["page"]
	page, err := strconv.Atoi(pageString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	db, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	orgs := []Org{}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer sqlDB.Close()
	db.Find(&orgs)
	totalPage := int(math.Ceil(float64(len(orgs)) / 2))
	db.Order("id").Limit(2).Offset((page - 1) * 2).Find(&orgs)
	result := OrgPagination{}
	result.Data = orgs
	result.Page = page
	result.PageSize = 1
	result.TotalPage = totalPage

	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
func AddOrg(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	org := Org{}

	errmssg := json.NewDecoder(r.Body).Decode(&org)
	if errmssg != nil {
		http.Error(w, errmssg.Error(), http.StatusInternalServerError)
	}
	db, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer sqlDB.Close()
	dbresult := db.Create(&org)
	if dbresult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(jsonData)

}
