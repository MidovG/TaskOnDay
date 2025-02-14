package entity

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tasks/internal/entity/userModel"
	"tasks/internal/repo"

	"github.com/gorilla/mux"
)

var database *sql.DB = repo.Connection()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("select * from base_crud_bd.users")

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	users := []userModel.User{}

	for rows.Next() {
		p := userModel.User{}
		err := rows.Scan(&p.Id, &p.Email, &p.UserName, &p.Password, &p.PhoneNumber, &p.DateOfBirth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	rows, err := database.Query("select * from base_crud_bd.users where id = ?;", params["id"])
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	users := []userModel.User{}

	for rows.Next() {
		p := userModel.User{}
		err := rows.Scan(&p.Id, &p.Email, &p.UserName, &p.Password, &p.PhoneNumber, &p.DateOfBirth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := userModel.User{
			Email:       r.FormValue("email"),
			UserName:    r.FormValue("username"),
			Password:    r.FormValue("password"),
			PhoneNumber: r.FormValue("phonenumber"),
			DateOfBirth: r.FormValue("dateofbirth"),
		}

		_, err = database.Exec("insert into base_crud_bd.users (email, user_name, password, phone_number, date_of_birth)  values (?, ?, ?, ?, ?)",
			user.Email,
			user.UserName,
			user.Password,
			user.PhoneNumber,
			user.DateOfBirth)

		if err != nil {
			log.Println(err)
		}

	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := database.Exec("delete from base_crud_bd.users where id = ?;", params["id"])
	if err != nil {
		log.Println(err)
	}
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.QueryRow("select * from base_crud_bd.users where id = ?", id)
	user := userModel.User{}
	err := row.Scan(&user.Id, &user.Email, &user.UserName, &user.Password, &user.PhoneNumber, &user.DateOfBirth)
	if err != nil {
		log.Println(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		params := mux.Vars(r)

		newPhoneNumber := r.FormValue("phonenumber")
		_, err = database.Exec("update base_crud_bd.users set phone_number = ? where id = ?;", newPhoneNumber, params["id"])

		if err != nil {
			log.Println(err)
		}

	}
}
