package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/VikyArthya/go-auth/config"
	"github.com/VikyArthya/go-auth/entities"
	"github.com/VikyArthya/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["LoggedIn"] != true {
			http.Redirect(w, r, "login", http.StatusSeeOther)
		} else {
			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, nil)
		}
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, nil)

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "username", UserInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("Username atau Password salah")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("Username atau Password salah !")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["LoggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Username
			session.Values["nama_lengkap"] = user.NamaLengkap

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}
