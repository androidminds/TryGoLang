package main

import (
     _ "github.com/go-sql-driver/mysql"
    _"database/sql"
    "crypto/md5"
    "time"
    "fmt"
     "regexp"
     "io"
)

const (
    PASS_LESS_LENGTH 	= 4
    PASS_MOST_LENGTH 	= 60
    NAME_LESS_LENGTH 	= 2
    NAME_MOST_LENGTH 	= 60
)

const (
    ERR_OK			= iota  
    ERR_NAME_LENGTH 		= iota  
    ERR_NAME_ILLEGAL_CHAR 	= iota  
    ERR_PASS_LENGTH		= iota
    ERR_PASS_NOT_EQUAL		= iota
    ERR_PASS_ILLEGAL_CHAR 	= iota  
    ERR_USER_EXISTED		= iota
)



func getTime() string {
    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    return tm.Format("2006-01-02 03:04:05 PM")
}

func islegalName(name string) bool {
    if len(name) < NAME_LESS_LENGTH || len(name) > NAME_MOST_LENGTH {
        return false
    }

    if m, _ := regexp.MatchString("^[a-z0-9A-Z\u4e00-\u9fa5]+$", name); !m {
        return false
    }

    return true
}

func islegalPassword(pass string) bool {
    if len(pass) < PASS_LESS_LENGTH || len(pass) > PASS_MOST_LENGTH {
        return false
    }  

    if m, _ := regexp.MatchString("^[a-z0-9A-Z]+$", pass); !m {
        return false
    }

    return true
}

func createPassToken(name string, pass string) string {
    h := md5.New()
    io.WriteString(h, pass)
    pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

    salt1 := "@2s#$%adss"
    salt2 := "s_^&*()sdf"

    io.WriteString(h, salt1)
    io.WriteString(h, name)
    io.WriteString(h, salt2)
    io.WriteString(h, pwmd5)

    return fmt.Sprintf("%x", h.Sum(nil))
}
