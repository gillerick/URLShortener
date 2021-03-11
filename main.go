package main


import (
	"URLShortener/models"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	base62 "URLShortener/utils"
	"log"
)


//DB stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

//Model the record struct
type Record struct {
	ID int `json:"id"`
	URL string `json:"url"`
}

//GetOriginalURL fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	//Fetch ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)

	//Handle response details
	if err != nil{
		w.Write([]byte(err.Error()))
	}else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}

}


//GetShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
	responseMap := map[string]interface{}{"encoded_string":base62.ToBase62(id)}
	if err != nil{
		w.Write([]byte(err.Error()))
	}else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}

}
func main()  {
	db, err = models.InitDB()



	//Test unit
	x := 100
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)

}
