package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Response struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Create_at        int    `json:"create_at"`
	Create_ip_at     string `json:"create_ip_at"`
	Last_login_at    int    `json:"last_login_at"`
	Last_login_ip_at string `json:"last_login_ip_at"`
	Login_times      int    `json:"login_times"`
	Status           int    `json:"status"`
}

func RDSConnect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	var dbUser = os.Getenv("dbUser")
	var dbPass = os.Getenv("dbPass")
	var dbEndpoint = os.Getenv("dbEndpoint")
	var dbName = os.Getenv("dbName")
	connectStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbUser,
		dbPass,
		dbEndpoint,
		"3306",
		dbName,
		"utf8",
	)
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func RDSProcessing(db *sql.DB) (interface{}, error) {

	var id int
	var email string
	var phone string
	var username string
	var password string
	var create_at int
	var create_ip_at string
	var last_login_at int
	var last_login_ip_at string
	var login_times int
	var status int

	responses := []Response{}
	responseMap := Response{}

	getData, err := db.Query("SELECT * FROM account_user")
	defer getData.Close()
	if err != nil {
		return nil, err
	}

	for getData.Next() {
		if err := getData.Scan(
			&id,
			&email,
			&phone,
			&username,
			&password,
			&create_at,
			&create_ip_at,
			&last_login_at,
			&last_login_ip_at,
			&login_times,
			&status,
		); err != nil {
			return nil, err
		}
		fmt.Println(
			id,
			email,
			phone,
			username,
			password,
			create_at,
			create_ip_at,
			last_login_at,
			last_login_ip_at,
			login_times,
			status,
		)
		responseMap.Id = id
		responseMap.Email = email
		responseMap.Phone = phone
		responseMap.Username = username
		responseMap.Password = password
		responseMap.Create_at = create_at
		responseMap.Create_ip_at = create_ip_at
		responseMap.Last_login_at = last_login_at
		responseMap.Last_login_ip_at = last_login_ip_at
		responseMap.Login_times = login_times
		responseMap.Status = status

		responses = append(responses, responseMap)
	}

	params, _ := json.Marshal(responses)
	fmt.Println(string(params))

	defer db.Close()
	return string(params), nil
}

func run() (interface{}, error) {
	fmt.Println("Start to connect RDS")
	db, err := RDSConnect()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("RDS connected")
	fmt.Println("Start to process")
	response, err := RDSProcessing(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Process done!")
	return response, nil
}

func main() {
	lambda.Start(run)
}
