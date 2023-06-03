package filereader

import (
	"mime/multipart"
	"encoding/csv"
	"net/http"
)

func FileReader(r *http.Request) (multipart.File, error) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	return file, nil

}
func CsvReader( file multipart.File)([]string,  error){
	reader:=csv.NewReader(file)
	return reader.Read()
	
}