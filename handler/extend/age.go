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

type PersonAge struct {
	Count int
	Name  string
	Age   int
}

func ExtractAge(person *model.Person) {
	wrappedName := url.QueryEscape(person.FirstName)
	ageUrl := fmt.Sprintf("https://api.agify.io/?name=%s", wrappedName)

	ageResp, err := http.Get(ageUrl)
	if err != nil {
		log.Fatal("error download name")
	}
	defer ageResp.Body.Close()

	body, _ := io.ReadAll(ageResp.Body)
	personAge := PersonAge{}
	json.Unmarshal(body, &personAge)

	person.Age = personAge.Age

}
