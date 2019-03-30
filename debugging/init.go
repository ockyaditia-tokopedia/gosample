package debugging

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ProbHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(body)
}
