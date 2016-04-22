package main

import (  
    "net/http"
    "github.com/gorilla/mux"

    "github.com/august-zou/dashboard/routers/web"
)

func main()  {
    
    r := mux.NewRouter()
    
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("./views/static"))))
	
    // ***** START: web home *****
    r.HandleFunc("/", web.Home)

    // r.HandleFunc("/home/index", web.Home)
    // ***** END: web home *****

    http.ListenAndServe(":8080", r)

}