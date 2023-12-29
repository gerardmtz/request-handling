package main

import (

    "fmt"
    "log"
    "net/http"
)


func hello (w http.ResponseWriter, r *http.Request) {
   if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    // Request handler (like an improvised multiplexer)
    switch r.Method {
        case "GET":
            // response to the request
            // with the named file
            // or directory
            http.ServeFile(w, r, "form.html")

        case "POST":
            // Call ParseForm() to parse the raw query and
            // update r.PostForm and r.Form
            err := r.ParseForm()


            // Every time we use function Fprintf()
            // we are writting into the w io.writer
            // the following string format.

            if err != nil {
                fmt.Fprintf(w, "ParseForm() err: %v", err)
                return
            }

            fmt.Fprintf(w, "Post from Website! r.PostForm = %v\n",
                        r.PostForm)
            
            name := r.FormValue("name")
            address := r.FormValue("address")
            fmt.Fprintf(w, "Name = %s\n", name)
            fmt.Fprintf(w, "Address = %s\n", address)

        default:
            fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }

}

func main() {

    // register the handler 
    // function for the given pattern
    // in the DefaultServeMux

    // Here "hello" is the function
    // thar is used by the handler

    http.HandleFunc("/", hello)

    fmt.Printf("Startin the server for testing HTTPP POST...\n")
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        log.Fatal(err)
    }
}
