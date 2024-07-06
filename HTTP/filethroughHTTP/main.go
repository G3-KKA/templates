package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//	localhost:8080/
//	click on upload
//	choose a file
//	redirect to /immediate
//	file loaded to server
//	file returned to client

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="ie=edge" />
		<title>Document</title>
	  </head>
	  <body>
		<form
		  enctype="multipart/form-data"
		  action="http://localhost:8080/immediate"
		  method="post"
		>
		  <input type="file" name="myFile" />
		  <input type="submit" value="upload" />
		</form>
	  </body>
	</html>`)

}
func immediateReturn(w http.ResponseWriter, r *http.Request) {
	log.Println("immediateReturn")

	memlimit, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	if memlimit >= 8*1024*1024 /* series of partial scans to deal with big files and to not load 1G in memory */ {
		//	Logic for large files, it shouldnt be possible to store 1G in memory
	}

	//	Get file from request
	r.ParseMultipartForm(int64(memlimit))
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
	}

	//	Create a temporary file within our temp directory
	tempFile, err := os.CreateTemp("temp", "*.data")
	defer tempFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	//	Read file contents into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	//	Flush the temporary data to the disk
	flushed, err := tempFile.Write(fileBytes)
	if err != nil || flushed != len(fileBytes) {
		log.Println(err)
	}
	log.Println("Successfully Uploaded File: " + tempFile.Name())

	w.Header().Set("Content-Type", "application/json")

	//	tempFile.Name() actually contains the full path
	// Force a download with the content- disposition field
	w.Header().Set("Content-Disposition", "attachment; filename="+tempFile.Name())

	// Send back
	http.ServeFile(w, r, tempFile.Name())
}

// ========================================================
func setupHTTP() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/immediate", immediateReturn)
	err := os.Mkdir("temp", 0777)
	if err != nil {
		pe, ok := err.(*os.PathError)
		if !ok {
			log.Fatal(err)
		}
		if strings.Contains(pe.Error(), "file exists") {
			log.Println("temp directory already exists, continuing")
		} else {
			log.Fatal(err)
		}
	}

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		//	Probably port 8080 is already in use
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello World")
	setupHTTP()
}
