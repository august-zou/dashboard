package web

import (  
    "log"
    "html/template"
    "net/http"
     "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/context"
    "github.com/august-zou/dashboard/utils"


)

func Home(w http.ResponseWriter, r *http.Request)  {
    
   db, _ := context.Get(r, utils.DBKey).(*sql.DB)

    // Execute the query
    rows, err := db.Query("SELECT * FROM user")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    utils.PrintRows(rows)
    
    t, err:= template.ParseFiles("view/template/index.html")
    if (err != nil) {
        log.Println(err)
    }

    t.Execute(w,nil) 
}
