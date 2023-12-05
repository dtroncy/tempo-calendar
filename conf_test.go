package tempocalendar

import (
	"testing"
)

func TestLoadConf(t *testing.T) {

	conf, err := loadConf()

	if err != nil {
		t.Fatalf(`Impossible de charger le fichier de conf`)
		return
	}

	if conf.TempoClientID != "tempo_clientID" || conf.TempoClientSecret != "tempo_clientSecret" {
		t.Fatalf(`Impossible de parser le fichier de conf`)
		return
	}
}
