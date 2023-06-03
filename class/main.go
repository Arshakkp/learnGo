package class

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"example.com/db"
	"example.com/filereader"
	. "example.com/model"
	"github.com/gorilla/mux"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	classes := []Classes{}
	db.Order("id").Find(&classes)
	jsonData, err := json.Marshal(classes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
func GetClassAndRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["userId"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	roleClass := []Roleandclasses{}
	db.Order("id").Find(&roleClass, "user_id=?", id)
	jsonData, err := json.Marshal(roleClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
func AddClassRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	fmt.Println("hello")

	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	userClassRole := []Roleandclasses{}

	errmssg := json.NewDecoder(r.Body).Decode(&userClassRole)

	if errmssg != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	errDb := db.Create(&userClassRole)
	if errDb.Error != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func AddClassRoleThroughFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	fmt.Println("hello")
	vars := mux.Vars(r)
	idString := vars["userId"]
	idUint, _ := strconv.ParseUint(idString, 10, 64)

	id := uint(idUint)
	file, _ := filereader.FileReader(r)
	reader := csv.NewReader(file)
	flag := false
	userClassRole := []Roleandclasses{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if flag {
			roleString := record[0]
			classString := record[1]
			roleUint, _ := strconv.ParseUint(roleString, 10, 64)
			classUint, _ := strconv.ParseUint(classString, 10, 64)
			role := uint(roleUint)
			class := uint(classUint)
			parseData := Roleandclasses{
				UserId:  id,
				RoleId:  role,
				ClassId: class,
			}
			userClassRole = append(userClassRole, parseData)
		}
		flag = true

	}
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	errDb := db.Create(&userClassRole)
	if errDb.Error != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func GetClassAndRoleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["userId"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "roleandclass.csv"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	roleClasses := []Roleandclasses{}
	db.Order("id").Find(&roleClasses, "user_id=?", id)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header row
	header := []string{"Role_id", "Class_id"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}
	for _, roleClass := range roleClasses {
		roleIdString := strconv.FormatUint(uint64(roleClass.RoleId), 10)
		classIdString := strconv.FormatUint(uint64(roleClass.ClassId), 10)
		row := []string{roleIdString, classIdString}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing CSV:", err)
			return
		}
	}
	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write(jsonData)

}
