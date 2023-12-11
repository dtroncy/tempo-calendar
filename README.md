# Tempo Calendar

Module go permettant de récupérer le calendrier des types de jours Tempo EDF à partir d'une date de début et d'une date de fin.

Importer le module :
```
import (
	tempo "github.com/dtroncy/tempo-calendar"
)
```
Ajouter un fichier de configuration ``conf.json`` à la racine de votre projet :
```
{
    "tempo_clientID": "tempo_clientID",
    "tempo_clientSecret": "tempo_clientSecret"
}
```
Pour obtenir ses informations, il faut s'inscrire sur le site suivant : https://data.rte-france.com/catalog/-/api/consumption/Tempo-Like-Supply-Contract/v1.1#


Utiliser le module :
```
tempo.GetTempoCalendar("Date de début", "Date de fin")
```
Exemple :
```
tempo.GetTempoCalendar("2023-12-02T00:00:00+01:00", "2023-12-05T00:00:00+01:00")
```

Le retour est le suivant :
```
{
    "tempo_like_calendars": {
        "start_date": "2023-12-02T00:00:00+01:00",
        "end_date": "2023-12-05T00:00:00+01:00",
        "values": [
            {
                "start_date": "2023-12-04T00:00:00+01:00",
                "end_date": "2023-12-05T00:00:00+01:00",
                "value": "BLUE",
                "updated_date": "2023-12-03T10:20:00+01:00"
            },
            {
                "start_date": "2023-12-03T00:00:00+01:00",
                "end_date": "2023-12-04T00:00:00+01:00",
                "value": "BLUE",
                "updated_date": "2023-12-02T10:20:00+01:00"
            },
            {
                "start_date": "2023-12-02T00:00:00+01:00",
                "end_date": "2023-12-03T00:00:00+01:00",
                "value": "WHITE",
                "updated_date": "2023-12-01T10:20:00+01:00"
            }
        ]
    }
}
```