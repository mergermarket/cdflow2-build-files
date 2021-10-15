package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mergermarket/cdflow2-build-files/internal/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 2 && os.Args[1] == "requirements" {
		// requirements is a way for the release container to communciate its requirements to the
		// config container - this one doesn't have any
		if err := json.NewEncoder(os.Stdout).Encode(map[string]interface{}{}); err != nil {
			log.Panicln("error encoding requirements:", err)
		}
		return
	}

	params := map[string]interface{}{}
	if err := json.Unmarshal([]byte(os.Getenv("MANIFEST_PARAMS")), &params); err != nil {
		log.Fatalln("error loading MANIFEST_PARAMS:", err)
	}

	if params["path"] == nil {
		log.Fatalln("missing path")
	}
	sourcePath, ok := params["path"].(string)
	if !ok {
		log.Fatalln("incorrect path type")
	}

	path, err := app.SaveData("/build", sourcePath)

	data, err := json.Marshal(map[string]string{"path": path})
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("/release-metadata.json", data, 0644); err != nil {
		log.Fatalln("error writing release metadata:", err)
	}
}
