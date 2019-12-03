package main

import ("fmt"
		 "net/http")


func main(){
	fmt.Println("Server is now running on Port: 8000")

	http.HandleFunc("/", webPage) //creates
	http.ListenAndServe(":8000", nil) //server local host port 800: 
}


func webPage(w http.ResponseWriter, r * http.Request){
	fmt.Fprintln(w, "This is a go website thing.") //fmt means format 
	fmt.Fprintln(w, "Adding another line.") //fmt means format 

	} //* means reading through 