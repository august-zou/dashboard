package main

import (  
    "log"
    "net/http"
    "database/sql"
    
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/context"
    "github.com/codegangsta/negroni"
    
    "github.com/august-zou/dashboard/router/web"
    "github.com/august-zou/dashboard/utils"
)


var db *sql.DB = setupDB()

func setupDB() *sql.DB {
    db, err := sql.Open("mysql", "zx:123456@/dashboard")
    if err != nil {
		log.Println(err)
		panic(err)
	}
	return db  
}

func ContextMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("======")
    context.Set(r, utils.DBKey, db)
    next(rw, r)
  // do some stuff after
}

func main()  {
    defer db.Close()
    
    r := mux.NewRouter()
    
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("./view/static"))))
	
    // ***** START: web home *****
    r.HandleFunc("/", web.Home)

    // r.HandleFunc("/home/index", web.Home)
    // ***** END: web home *****

    n := negroni.Classic()
    n.Use(negroni.HandlerFunc(ContextMiddleware))
    
    n.UseHandler(r)
    log.Fatal(http.ListenAndServe(":8080", n))

}