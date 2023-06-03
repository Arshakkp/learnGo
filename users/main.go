/*users give functions for adding and getting users from db*/
package users

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"

	"net/http"
	"strconv"

	"example.com/db"
	"example.com/filereader"
	. "example.com/model"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pageString := vars["page"]
	page, err := strconv.Atoi(pageString)
	limit := 2
	if err != nil {
		page = 1
		limit = -1
	}
	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	defer sqlDB.Close()
	users := []UserWithRoleAndClass{}
	db.Order("id").Limit(limit).Offset((page-1)*limit).Raw("SELECT users.* ,classes.class ,roles.role FROM users LEFT JOIN roleandclasses ON users.id = roleandclasses.user_id LEFT JOIN classes ON roleandclasses.class_id =classes.id LEFT JOIN roles ON roleandclasses.role_id=roles.id WHERE users.org_id=?  ORDER BY users.id ", id).Scan(&users)
	totalPage := int(math.Ceil(float64(len(users)) / 2))

	db.Order("id").Limit(limit).Offset((page-1)*limit).Raw("SELECT users.* ,classes.class ,roles.role FROM users LEFT JOIN roleandclasses ON users.id = roleandclasses.user_id LEFT JOIN classes ON roleandclasses.class_id =classes.id LEFT JOIN roles ON roleandclasses.role_id=roles.id WHERE users.org_id=?  ORDER BY users.id LIMIT ? OFFSET ?", id, limit, (page-1)*limit).Scan(&users)
	// db.Order("id").Limit(limit).Offset((page - 1) * limit).Joins("JOIN roleandclasses ON users.id=roleandclasses.user_id").Joins("JOIN roles ON roleandclasses.role_id=roles.id").Joins("JOIN classes ON roleandclasses.role_id=classes.id").Find(&users, "org_id=?", id)
	// if limit != -1 {
	// 	db.Order("id").Limit(limit).Offset((page-1)*limit).Find(&users, "org_id=?", id)
	// }
	result := UserPagination{}
	result.Data = users
	result.Page = page
	result.PageSize = limit
	result.TotalPage = totalPage

	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write(jsonData)

}
func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := db.Connect()
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	defer sqlDB.Close()
	user := User{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	errmssg := json.NewDecoder(r.Body).Decode(&user)
	user.OrgId = idInt
	if errmssg != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	errDb := db.Create(&user)
	if errDb.Error != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write(jsonData)

}
func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	vars := mux.Vars(r)
	idString := vars["id"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)

	user := User{}
	user.ID = id

	errmssg := json.NewDecoder(r.Body).Decode(&user)
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
	dbresult := db.Save(&user)
	if dbresult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	vars := mux.Vars(r)
	idString := vars["id"]
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := uint(idUint)
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

	dbUserResult := db.Where("id=?", id).Delete(&user)

	if dbUserResult.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Write(jsonData)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user := User{}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.First(&user, id)
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func AddUsersWithFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	vars := mux.Vars(r)
	idString := vars["id"]
	idiNT64, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	id := int(idiNT64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	file, _ := filereader.FileReader(r)
	defer file.Close()

	reader := csv.NewReader(file)
	flag := false
	users := []User{}
	rolesAndclasses := []Roleandclasses{}
	nameIndex := -1
	addressIndex := -1
	ageindex := -1
	emailIndex := -1
	classIndex := -1
	roleIndex := -1
	userId := []uint{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if flag {
			if record[emailIndex] == "" {
				break
			}
			fmt.Println(id)
			user := User{}
			user.Name = record[nameIndex]
			user.Address = record[addressIndex]
			user.Age, _ = strconv.Atoi(record[ageindex])
			user.Email = record[emailIndex]
			user.OrgId = id
			if roleIndex != -1 && classIndex != -1 {
				roleAndClass := Roleandclasses{}
				roleId, _ := strconv.ParseUint(record[roleIndex], 10, 64)
				classId, _ := strconv.ParseUint(record[classIndex], 10, 64)
				roleAndClass.RoleId = uint(roleId)
				roleAndClass.ClassId = uint(classId)
				rolesAndclasses = append(rolesAndclasses, roleAndClass)
			}

			users = append(users, user)

		} else {

			for i, val := range record {
				switch val {
				case "name":
					nameIndex = i
				case "address":
					addressIndex = i
				case "age":
					ageindex = i
				case "email":
					emailIndex = i
				case "role id":
					roleIndex = i
				case "class id":
					classIndex = i
				}
			}
			fmt.Println(ageindex, addressIndex, nameIndex, emailIndex)
		}
		flag = true
	}
	db, _ := db.Connect()
	for _, value := range users {
		dbVal := User{}
		result := db.Where("email = ?", value.Email).First(&dbVal)
		if result.Error != nil {
			errDb := db.Create(&value)
			userId = append(userId, value.ID)
			if errDb.Error != nil {
				log.Fatal(errDb)
			}

		} else {
			value.ID = dbVal.ID
			userId = append(userId, dbVal.ID)
			errDb := db.Save(&value)
			if errDb.Error != nil {
				log.Fatal(errDb)
			}
		}

	}
	for i, value := range rolesAndclasses {
		if value.ClassId == 0 && value.RoleId == 0 {
			continue
		}
		if userId[i] != uint(0) {
			value.UserId = userId[i]
		}

		dbVal := Roleandclasses{}
		result := db.Where("user_id = ? AND class_id=?", value.UserId, value.ClassId).First(&dbVal)
		if result.Error != nil {
			fmt.Print("here problem")
			errDb := db.Create(&value)
			if errDb.Error != nil {
				log.Fatal(errDb)
			}

		} else {
			value.ID = dbVal.ID
			errDb := db.Save(&value)
			if errDb.Error != nil {
				log.Fatal(errDb)
			}
		}
	}

}
