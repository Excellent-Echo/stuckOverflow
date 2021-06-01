package jobsapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func JobsApi() entity.Data {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://remotive.io/api/remote-jobs?category=software-dev", nil)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)

	if err != nil || resp.StatusCode != 200 {
		panic(err)
	}

	byteData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var datas entity.Data

	json.Unmarshal(byteData, &datas)

	// for _, data := range datas.Jobs {
	// 	fmt.Println(data)
	// }
	// fmt.Println("Store data done")

	return datas
}

func connect() (*sql.DB, error) {
	err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func StoreToDB() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	datas := JobsApi()
	var counter int

	for _, data := range datas.Jobs {
		_, err := db.Exec("insert into jobs (title, url, company_name, job_type, publication_date) values (?,?,?,?,?)", data.Title, data.Url, data.CompanyName, data.JobType, data.PublicationDate)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		counter++
		fmt.Printf("%d insert success!", counter)
	}

}
