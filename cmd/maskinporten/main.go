package main

import (
	"github.com/tktip/maskinporten/pkg/maskinporten"
	"github.com/haraldfw/cfger"
	log "github.com/sirupsen/logrus"
)

func main() {
	// create handler, validation should happen in this method
	var mp maskinporten.Handler
	_, err := cfger.ReadStructuredCfg("env::CONFIG", &mp)
	if err != nil {
		log.Fatal(err)
	}

	if mp.Debug {
		log.SetLevel(log.DebugLevel)
	}

	//initialize the connector
	err = mp.Init()
	if err != nil {
		log.Fatal(err)
	}

	//grab an access token
	res, err := mp.CreateAccessToken()
	if err != nil {
		log.Error(err)
	}
	//print it!
	log.Info(res.AccessToken)
}
