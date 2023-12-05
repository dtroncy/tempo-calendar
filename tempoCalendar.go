package tempocalendar

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/oauth2/clientcredentials"
)

func GetTempoCalendar(startDate string, endDate string) ([]byte, error) {

	var body []byte

	conf, err := loadConf()

	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier de configuration : ", err)
		return body, err
	}

	config := &clientcredentials.Config{
		ClientID:     conf.TempoClientID,
		ClientSecret: conf.TempoClientSecret,
		TokenURL:     "https://digital.iservices.rte-france.com/token/oauth",
	}

	client := config.Client(context.Background())

	apiURL := "https://digital.iservices.rte-france.com/open_api/tempo_like_supply_contract/v1/tempo_like_calendars?start_date=" + startDate + "&end_date=" + endDate

	response, err := client.Get(apiURL)
	if err != nil {
		fmt.Printf("Erreur lors de la requête GET: %v\n", err)
		return body, err
	}
	defer response.Body.Close()

	// Lecture du corps de la réponse
	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du corps de la réponse: %v\n", err)
		return body, err
	}

	return body, err
}
