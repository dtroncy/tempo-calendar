package tempocalendar

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

func GetTempoCalendar(startDate string, endDate string) (*http.Response, error) {

	var response *http.Response

	conf, err := loadConf()

	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier de configuration : ", err)
		return response, err
	}

	config := &clientcredentials.Config{
		ClientID:     conf.TempoClientID,
		ClientSecret: conf.TempoClientSecret,
		TokenURL:     "https://digital.iservices.rte-france.com/token/oauth",
	}

	client := config.Client(context.Background())

	apiURL := "https://digital.iservices.rte-france.com/open_api/tempo_like_supply_contract/v1/tempo_like_calendars?start_date=" + startDate + "&end_date=" + endDate

	response, err = client.Get(apiURL)
	if err != nil {
		fmt.Printf("Erreur lors de la requête GET: %v\n", err)
		return response, err
	}
	defer response.Body.Close()

	return response, err
}
