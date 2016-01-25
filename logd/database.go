package main

import (
     _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "time"
    "fmt"
    log "github.com/cihub/seelog"
)

var myDB* sql.DB

type UserInfo struct {
    uid int
    name string
    pass string
    created string
}

func addNewUser(name string, pass string) int64 {
    stmt, err := myDB.Prepare("INSERT userinfo SET username=?,password=?,created=?")
    if err != nil {
        log.Error("addNewUser(), datbase prepare: ", err.Error())
        return -1;
    }

    res, err := stmt.Exec(name, createPassToken(name, pass), getTime())
    if err != nil {
        log.Error("addNewUser(), datbase Exec: ", err.Error())
        return -1;
    }

    id, err := res.LastInsertId()
    if err != nil {
        log.Error("addNewUser(), datbase LastInsertId: ", err.Error())
        return -1;
    }

    return id;
}

func getUserInfo(name string) (int, string, string) {
    rows, err := myDB.Query("SELECT * FROM userinfo where username = '%s'", name)

    if err == nil {
        for rows.Next() {
            var uid int
            var username string 
            var created string
            err = rows.Scan(&uid, &username, &created)
            if err == nil {
                return uid, username, created
            }
        }
    }

    return -1, "", ""
}

func getUserInfoById(id int) (*UserInfo) {
    rows, err := myDB.Query("SELECT * FROM userinfo where uid = '%d'", id)
    info := new(UserInfo)

    if err == nil {
        for rows.Next() {
            err = rows.Scan(&info.uid, &info.name, &info.pass, &info.created)
            if err == nil {       
                return info;
            }
        }
    }
    return nil
}


func connectDatabase(address string, database string, user string, pass string) bool {
    var err error
    dbConnect := fmt.Sprintf("%s:%s@/%s?charset=utf8", user, pass, database);
    myDB, err = sql.Open("mysql", dbConnect)
    
    if(err != nil) {
        log.Error("Failed to connect database %s with user %s for %s", address+"/"+database, user, err.Error());
        return false;
    }
    return true;
}

func waitAndConnectDatabase(count int, sleepSecond int) bool {
    for i:=0; i < count; i++ {
        if connectDatabase("", "test", "root", "root123") {
            return true;
        } else {
            time.Sleep(time.Duration(sleepSecond*1e9));
        }
    }
    log.Error("trying %d times to connect database, FAILED AND EXIT", count);
    return false;
}
