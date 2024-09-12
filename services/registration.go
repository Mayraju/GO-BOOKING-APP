package services

import (
	"fmt"

	"github.com/mayraju/go-booking-app/loaders"
	"github.com/mayraju/go-booking-app/models"
)

var (
	db, err = loaders.ConnectDB()
)

type Registration interface {
	RegistrationService(models.Registration)
	LoginService(email, password string)
}

type RegistrationServiceType struct{}

func (r RegistrationServiceType) RegistrationService(form models.Registration) string {

	if err != nil {
		fmt.Print("Error Occured===========")

	}
	if db == nil {
		fmt.Printf("Error to connecting to db")
		return ""
	}
	var createTable = `CREATE TABLE IF NOT EXISTS registration (
        id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
        firstName VARCHAR(50) NOT NULL,
        lastName VARCHAR(50) NOT NULL,
        email VARCHAR(50) NOT NULL,
        password VARCHAR(50) NOT NULL
    );`

	// Execute the SQL statement
	create, _ := db.Exec(createTable)

	fmt.Println("Table created successfully", create)
	stmt, err := db.Prepare("INSERT INTO registration(firstName,lastName,email,password) VALUES(?,?,?,?)")
	if err != nil {
		return "Error In Creating the table"
	}

	res, err := stmt.Exec(form.FirstName, form.LastName, form.Email, form.Password)
	if err != nil {
		return "Error In excuting the values"
	}

	fmt.Println(res, "result")

	return "Data and table inserted and created."

}

func (r RegistrationServiceType) LoginService(email, password string) ([]models.Registration, error) {
	fmt.Printf("email %s and password %s", email, password)
	var arr []models.Registration
	if db == nil {
		fmt.Println("Error in conneting the databas")
	}
	stmt, err := db.Query("SELECT * FROM registration")
	fmt.Print("after query=============")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	fmt.Print("after Close=============")
	for stmt.Next() {
		fmt.Println("IN next clasuee==========================")
		var result models.Registration
		error := stmt.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.Password)
		fmt.Println("After the Scan================", error)
		if error != nil {
			return nil, error
		}

		arr = append(arr, result)
		fmt.Println("after the append========================")
	}

	fmt.Print("after next =============")

	return arr, nil
}
