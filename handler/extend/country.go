package extend

import (
	"encoding/json"
	"fmt"
	"github.com/baiv84/personio/model"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CountryItem struct {
	Country_id  string
	Probability float32
}

type PersonNationality struct {
	Count   int32
	Name    string
	Country []CountryItem
}

func ExtractCountry(person *model.Person) {
	wrappedName := url.QueryEscape(person.FirstName)
	nationalityUrl := fmt.Sprintf("https://api.nationalize.io/?name=%s", wrappedName)

	nationalResp, err := http.Get(nationalityUrl)
	if err != nil {
		log.Fatal("error download name")
	}
	defer nationalResp.Body.Close()

	body, _ := io.ReadAll(nationalResp.Body)

	personNationality := PersonNationality{}
	json.Unmarshal(body, &personNationality)
	counts := personNationality.Country

	var probability float32
	var nationResult string

	for _, elem := range counts {
		if elem.Probability > probability {
			probability = elem.Probability
			nationResult = elem.Country_id
		}
	}
	person.Country = nationResult

}
