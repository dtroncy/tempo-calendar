package tempocalendar

import (
	"encoding/json"
	"os"
)

type Conf struct {
	TempoClientID     string `json:"tempo_clientID"`
	TempoClientSecret string `json:"tempo_clientSecret"`
}

func loadConf() (Conf, error) {
	var conf Conf

	file, err := os.Open("conf.json")
	if err != nil {
		return conf, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
