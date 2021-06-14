package shardsai

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/misbahulard/shards-ai/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func catIndicesRequest() func(*esapi.CatIndicesRequest) {
	return func(r *esapi.CatIndicesRequest) {
		r.S = []string{"index"}
		r.Bytes = "kb"
		r.Format = "json"
	}
}

func Run() {

	res, err := config.EsClient.Cat.Indices(catIndicesRequest())
	if err != nil {
		log.Error("Error when get elasticsearch indices")
		os.Exit(1)
	}
	defer res.Body.Close()

	var catIndices []CatIndex
	json.NewDecoder(res.Body).Decode(&catIndices)
	for _, i := range catIndices {
		log.Infof("%+v", i.Index)
	}

	log.Info("#####################################")

	var templates []Template
	viper.UnmarshalKey("templates", &templates)

	if len(templates) == 0 {
		log.Error("Please define at least one templates in the config file.")
		os.Exit(1)
	}

	log.Info("templates:")
	for _, t := range templates {
		log.Infof("- path: %s", t.Path)
		log.Infof("  indices: %s", t.Indices)
		log.Infof("  priority: %d", t.Priority)
	}

	for _, t := range templates {
		jsonFile, err := os.Open(t.Path)
		if err != nil {
			log.Errorf("Template file not found: %s, stacktrace: %s", t.Path, err.Error())
		}
		defer jsonFile.Close()

		jsonByte, _ := ioutil.ReadAll(jsonFile)

		var indexTemplate IndexTemplate
		json.Unmarshal(jsonByte, &indexTemplate)

		log.Infof("%+v", indexTemplate)
	}
}
