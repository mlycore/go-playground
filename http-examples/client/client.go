package main

import (
	"io/ioutil"
	"github.com/mlycore/log"
	"net/http"
//	"time"
)

func main() {
	for {
		go func() {
			log.Infof("HTTP GET")
			client := &http.Client{
				//		Timeout: 1 * time.Second,
			}
			r, err := client.Get(`http://127.0.0.1:8080/`)
			if err != nil {
				log.Fatalf(err.Error())
			}
			defer r.Body.Close()
			bs, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatalf(err.Error())
			}
			log.Infof("HTTP Done.")
			log.Infof(string(bs))
		}()
	}

}
