package main

import (
    log "github.com/cihub/seelog"
    "html/template"
    "net/http"
)

const (
    PAGE_ERROR_404		= iota
    PAGE_ERROR_SYSTEM		= iota
    PAGE_SUCCESS_ADD_USER	= iota
    PAGE_ERROR_ILLEGAL_NAME	= iota
    PAGE_ERROR_USER_EXISTED	= iota
    PAGE_ERROR_ILLEGAL_PASS	= iota
    PAGE_ERROR_PASS_NOT_EQUAL	= iota
    PAGE_ERROR_NO_UID		= iota
)

var errorTemplate * template.Template
type ErrorInfo struct {
    Title string
    ErrorInfo string
} 

func loadTemplates(path string) bool {
    var err error
    errorTemplate, err = template.ParseFiles(path + "/" + "error.gtpl")
    if(err != nil) {
        log.Error("parse template file %s fail : %s", err.Error())
        return false;
    }

    return true;
}

func returnErrorPage(w http.ResponseWriter, errorCode int, title string, info string) {
    if(errorTemplate != nil) {
        ei := ErrorInfo { title, info }
        w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
        w.WriteHeader(errorCode)
        errorTemplate.Execute(w, ei)
    }
}

func showErrorPage(w http.ResponseWriter, err int) {
    switch err {
    case PAGE_ERROR_404:
        returnErrorPage(w, http.StatusNotFound, "Page Error", "Required Page is not existed!")
    case PAGE_ERROR_SYSTEM:
        returnErrorPage(w, http.StatusInternalServerError, "System Error", "System met error!")
    case PAGE_ERROR_ILLEGAL_NAME:
        returnErrorPage(w, http.StatusUnauthorized, "Register Fail", "User name are not legal!")
    case PAGE_ERROR_USER_EXISTED:
        returnErrorPage(w, http.StatusUnauthorized, "Register Fail", "User have existed!")
    case PAGE_ERROR_ILLEGAL_PASS:
        returnErrorPage(w, http.StatusUnauthorized, "Register Fail", "Password are not legal!")
    case PAGE_ERROR_PASS_NOT_EQUAL:
        returnErrorPage(w, http.StatusUnauthorized, "Register Fail", "Two password are not euqal!")
    case PAGE_ERROR_NO_UID:
        returnErrorPage(w, http.StatusNotFound, "Get User Fail", "No such user!")
    }
}


func showWebPage(w http.ResponseWriter, filename string) {
    t, err:= template.ParseFiles(filename)
    if err == nil { 
        t.Execute(w, nil)
    } else {
        log.Error("Parse %s fail : %s", filename, err.Error())
        showErrorPage(w, PAGE_ERROR_SYSTEM)
    }
}

func showLoginPage(w http.ResponseWriter) {
    showWebPage(w, "login.gtpl")
}

func showRegisterPage(w http.ResponseWriter) {
    showWebPage(w, "register.gtpl")
}

func showUserPage(w http.ResponseWriter, info* UserInfo) {
    t, err:= template.ParseFiles("UserInfo.gtpl")
    if err == nil { 
        t.Execute(w, info)
    } else {
        log.Error("Parse %s fail : %s", "UserInfo.gtpl", err.Error())
        showErrorPage(w, PAGE_ERROR_SYSTEM)
    }
}

