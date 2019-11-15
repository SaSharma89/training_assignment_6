package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"hash/fnv"
)


func converter( inURL string ) (outURL string) {
	h := fnv.New64a()
	h.Write([]byte(inURL))
	return fmt.Sprintf("%s/%d", OutURLBase, h.Sum64())
}

func convert(w http.ResponseWriter, req *http.Request){
	fmt.Println("Input URL :: ", req.FormValue("test"))

	inURL := req.FormValue("test")

	outURL := converter(inURL)

	fmt.Println("outURL after hash ", outURL)

	masterTmpl, err := template.New("master2").Parse(Secondtmpl)
	if err != nil {
		log.Fatal(err)
	}

	data := Datastruct{ PageTitle : "URL Conversion",
						InURL : inURL,
						OutURL: outURL,
					 }

	masterTmpl.Execute(w, data)

	reg := URLregister{ InURL : inURL,
				OutURL: outURL, }

	URLregisterList = append(URLregisterList, reg)
}

func displayList(w http.ResponseWriter, req *http.Request){

	masterTmpl, err := template.New("master3").Parse(Thirdtmpl)
	if err != nil {
		log.Fatal(err)
	}

	masterTmpl.Execute(w, URLregisterList)
}

func handleHello(w http.ResponseWriter, req *http.Request){
	masterTmpl, err := template.New("master").Parse(Firsttmpl)
	if err != nil {
		log.Fatal(err)
	}

	data := Datastruct{ PageTitle : "URL Conversion", }
	masterTmpl.Execute(w, data)
}

func main() {
	fmt.Println("Go training final assingment. Starting server...")
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/convert", convert)
	http.HandleFunc("/list", displayList)
	fmt.Println(http.ListenAndServe("localhost:8080", nil))

}
