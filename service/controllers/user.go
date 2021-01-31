package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/EleMint/CONTd/service/contracts"
	"github.com/EleMint/CONTd/service/managers"
)

var manager *managers.UserManager

// InitUserController initializes the user controller
func InitUserController() {
	manager = managers.NewUserManager()
}

func parseUserGetFuncName(r *http.Request) string {
	path := strings.ToLower(r.URL.Path)

	fname := ""
	if path == "/user" ||
		path == "/user/" {
		fname = "get"
	} else if path == "/users" ||
		path == "/users/" {
		fname = "gets"
	}

	return fname
}

func parseUserPostFuncName(r *http.Request) string {
	path := strings.ToLower(r.URL.Path)

	fname := ""
	if path == "/user" ||
		path == "/user/" {
		fname = "create"
	}

	return fname
}

// UserHTTPHandleFunc http controller for accepting user requests
func UserHTTPHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.Proto, r.Method, r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		fname := parseUserGetFuncName(r)
		switch fname {
		case "get":
			handleHTTPUserGet(w, r)
		case "gets":
			handleHTTPUsersGet(w, r)
		default:
			http.Error(
				w,
				"request path not implemented",
				http.StatusNotImplemented,
			)
		}
	case http.MethodPost:
		fname := parseUserPostFuncName(r)
		switch fname {
		case "create":
			handleHTTPUserCreate(w, r)
		// case "updateAttributes":
		// 	handleHTTPUpdateUserAttributes(w, r)
		default:
			http.Error(
				w,
				"request path not implemented",
				http.StatusNotImplemented,
			)
		}
	}
}

func handleHTTPUserGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	email := r.URL.Query().Get("email")

	var rch chan *contracts.GetUserResponse

	if id != "" {
		rch = manager.GetUserByID(&contracts.GetUserByID{
			ID: id,
		})
	} else if email != "" {
		rch = manager.GetUserByEmail(&contracts.GetUserByEmail{
			Email: email,
		})
	}

	select {
	case resp := <-rch:
		marsh, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "failed to marshal response data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.Response.StatusCode)
		w.Write(marsh)
		return
		// TODO: add context.timeout case return timeout http error
	}
}

func handleHTTPUsersGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Users")
	emailStarts := r.URL.Query().Get("emailstarts")
	displayStarts := r.URL.Query().Get("displaystarts")
	l := r.URL.Query().Get("limit")
	if l == "" {
		l = "10"
	}
	limit, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		limit = 10
	}

	var rch chan *contracts.GetUsersResponse

	if emailStarts != "" {
		log.Println(emailStarts)
		rch = manager.GetUsersByStartEmail(&contracts.GetUsersByStartEmail{
			Substring: emailStarts,
			Limit:     limit,
		})
	} else if displayStarts != "" {
		rch = manager.GetUsersByStartDisplayName(&contracts.GetUsersByStartDisplayName{
			Substring: displayStarts,
			Limit:     limit,
		})
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	select {
	case resp := <-rch:
		marsh, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "failed to marshal response data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.Response.StatusCode)
		w.Write(marsh)
		return
		// TODO: add context.timeout case return timeout http error
	}
}

func handleHTTPUserGetStartEmail(w http.ResponseWriter, r *http.Request) {

}

func handleHTTPUserCreate(w http.ResponseWriter, r *http.Request) {
	cuContr := &contracts.CreateUser{}
	err := json.NewDecoder(r.Body).Decode(cuContr)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		return
	}

	valid := cuContr.Validate()
	if !valid {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	rch := manager.CreateUser(cuContr)
	select {
	case resp := <-rch:
		marsh, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "failed to marshal response data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.Response.StatusCode)
		w.Write(marsh)
		return
	}
}
