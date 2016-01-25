package main

import (
     _ "github.com/go-sql-driver/mysql"
    log "github.com/cihub/seelog"
    "net/http"
    "strconv"
    "fmt"
)


func checkUserName(w http.ResponseWriter, name string) bool {
    if !islegalName(name) {
        showErrorPage(w, PAGE_ERROR_ILLEGAL_NAME)
        return false
    } else if uid,_,_ := getUserInfo(name); uid >= 0 {
        showErrorPage(w, PAGE_ERROR_USER_EXISTED)
        return false
    } else {
        return true
    }
}

func checkPassword(w http.ResponseWriter, pass1 string, pass2 string) bool {
    if !islegalPassword(pass1) {
        showErrorPage(w, PAGE_ERROR_ILLEGAL_PASS)
        return false
    } else if pass1 != pass2 {
        showErrorPage(w, PAGE_ERROR_PASS_NOT_EQUAL)
        return false
    } else {
        return true
    }
}


func getUser(w http.ResponseWriter, r *http.Request) {
    if(r.Method == "GET") {
        fmt.Println("getUser %s", r.URL)
        return
    } else {
        r.ParseForm() 
        uid := r.Form["uid"][0]
    id, err := strconv.Atoi(uid);
    if err != nil {
        showErrorPage(w, PAGE_ERROR_NO_UID)
        return
    }

    if info := getUserInfoById(id); info != nil  {
	showUserPage(w, info)
    } else {
        showErrorPage(w, PAGE_ERROR_NO_UID)
    }
}
}

func createUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
	r.ParseForm() 
	name := r.Form["username"][0]
	pass1 := r.Form["password1"][0]
	pass2 := r.Form["password2"][0]
	if checkUserName(w, name) && checkPassword(w, pass1, pass2) {
	    if addNewUser(name, pass1) >= 0 {
	    // returnErrorPage(w, PAGE_SUCCESS_ADD_USER)
	    } else {
	        showErrorPage(w, PAGE_ERROR_SYSTEM)
	    }
	}   
    }
}

func register(w http.ResponseWriter, r *http.Request) {
    showRegisterPage(w)
}

func login(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        showLoginPage(w)
    } else if r.Method == "POST"{
        r.ParseForm() 
        name := r.Form["username"][0]
        pass := r.Form["password"][0]
        if islegalName(name) && islegalPassword(pass) {
            _, token, _ := getUserInfo(name)
            if token == createPassToken(name, pass) {
                log.Info("User %s login!", name)
                // todo gotoPage("user", uid)
                return
            }
        }
        // todo : returnResultPage(w, PAGE_ERROR_LOGIN);      
    }
}


func handleRoute() {
    http.HandleFunc("/", login)       
    http.HandleFunc("/login", login)        
    http.HandleFunc("/register", register) 
    http.HandleFunc("/createuser", createUser)
    http.HandleFunc("/user", getUser)
}

func main() {

    if !loadTemplates(".") {
        return
    }

    if !waitAndConnectDatabase(10/*times*/, 5/*seconds*/) {
        return;
    }

    handleRoute();

    port := "9090"
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Error("Failed to list on port %s : ", port, err.Error());
    }
}


