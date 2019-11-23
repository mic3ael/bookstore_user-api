package users

import (
	"fmt"

	"github.com/mic3ael/bookstore_user-api/datasources/psql/bookstore_db"
	"github.com/mic3ael/bookstore_user-api/utils/date_utils"
	"github.com/mic3ael/bookstore_user-api/utils/psql_utils"

	"github.com/mic3ael/bookstore_user-api/utils/errors"
)

const (
	quetyInsertUser = "INSERT INTO users (first_name, last_name, email, password, created_on, updated_on) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;"
	quetyGetUser    = "SELECT id, first_name, last_name, email, created_on, updated_on FROM users WHERE id=$1 AND deleted=false;"
	queryUpdateUser = "UPDATE users SET first_name=$2, last_name=$3, email=$4, updated_on=$5 WHERE id=$1;"
	queryDeleteUser = "UPDATE users SET deleted=true, updated_on=$2 WHERE id=$1;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := bookstoredb.Client.Prepare(quetyGetUser)
	if err != nil {
		fmt.Printf("err %v", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedOn, &user.UpdatedOn); getErr != nil {
		return psqlutils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := bookstoredb.Client.Prepare(quetyInsertUser)
	if err != nil {
		fmt.Printf("err %v", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	date := dateutils.GetNowUint()
	user.CreatedOn = date
	user.Password = "123"
	user.Deleted = false
	user.UpdatedOn = date
	var userID uint64

	saveErr := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Password, user.CreatedOn, user.UpdatedOn).Scan(&userID)

	if saveErr != nil {
		fmt.Println(saveErr)
		return psqlutils.ParseError(saveErr)
	}

	user.ID = uint64(userID)
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := bookstoredb.Client.Prepare(queryUpdateUser)
	if err != nil {
		fmt.Printf("err %v", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.UpdatedOn = dateutils.GetNowUint()

	_, updateErr := stmt.Exec(user.ID, user.FirstName, user.LastName, user.Email, user.UpdatedOn)

	if updateErr != nil {
		return psqlutils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := bookstoredb.Client.Prepare(queryDeleteUser)
	if err != nil {
		fmt.Printf("err %v", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.UpdatedOn = dateutils.GetNowUint()

	_, deleteErr := stmt.Exec(user.ID, user.UpdatedOn)

	if deleteErr != nil {
		return psqlutils.ParseError(deleteErr)
	}

	return nil
}
