package main

import (
	"github.com/mlycore/log"
	"io"
	"net/http"
	"time"
)

func main() {
	//ClientReadTimeoutServer()
	ClientConnectTimeoutServer()
}

// ClientReadTimeoutServer will make a request read timeout event.
func ClientReadTimeoutServer()  {
	http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		log.Infof("wait a couple of seconds ...")
		time.Sleep(2 * time.Second)
		io.WriteString(w, `Hi`)
		log.Infof("Done.")
	})
	log.Infof(http.ListenAndServe(":8080", nil).Error())
}

// ClientConnectTimeoutServer will make a request connect timeout event.
func ClientConnectTimeoutServer()  {
	http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		log.Infof("wait a couple of seconds ...")
		time.Sleep(time.Hour)
		io.WriteString(w, `Hi`)
		log.Infof("Done.")
	})
	log.Infof(http.ListenAndServe(":8080", nil).Error())
}
