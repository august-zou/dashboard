package web

import (  
    "fmt"
    "html/template"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
    t, err:= template.ParseFiles("views/template/index.html")
    // t, err:= template.New("index").Parse("{{.Title}}")

    if (err != nil) {
        fmt.Println(err)
    }

    t.Execute(w,nil) 
}
