package main

import (
	"encoding/json"
    "net/http"
	"bytes"
	"log"
	"io/ioutil"
	"net/url"
    
)

type image struct {
	Image string
}

type containerCreateResp struct {
	ID string `json:"Id"`
	Warnings string `json:" Warnings"`
}

func main() {
    //Pull the image
	img := image{"alpine:latest"}
	data := url.Values{}
	data.Set("fromImage", img.Image)
	response, err := http.PostForm("http://localhost:5555/images/create", data)
	if err != nil {
        log.Fatalf("An Error Occured downloading the image %v", err)
    }

	//Encode the data
	postBody, err := json.Marshal(img)
	if err != nil {
		log.Fatalf("Error encoding the data as a Json. Error:%v", err)
	}
	responseBody := bytes.NewBuffer(postBody)
	response, err = http.Post("http://localhost:5555/containers/create","application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }

	 defer response.Body.Close()

	 body, err := ioutil.ReadAll(response.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
   log.Printf(sb)


}