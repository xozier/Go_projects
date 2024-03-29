package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path  != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}


func formHandler( w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	}
	fmt.Fprintf(w , "POST request successful")
	name := r.FormValue("Name")
	age := r.FormValue("Age")
	fmt.Fprintf(w,"name: %s\n", name)
	fmt.Fprintf(w,"age:%s\n", age)
}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// it has to include "/form.html"
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListenAndServe(":8080", nil); err!= nil{
		log.Fatal(err)
	}
}