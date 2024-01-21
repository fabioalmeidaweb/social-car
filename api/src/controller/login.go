package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"social-car/src/auth"
	"social-car/src/database"
	"social-car/src/models"
	"social-car/src/repository"
	"social-car/src/responses"
	"social-car/src/security"
)

// Login is a function that handles the login functionality.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It does not return anything.
func Login(w http.ResponseWriter, r *http.Request) {
	regBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(regBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	savedUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Check(savedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(savedUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))

}
