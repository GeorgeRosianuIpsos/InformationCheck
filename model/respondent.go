package model

type Respondent struct {
	Id            []string `json:"id"`
	Name          []string `json:"name"`
	FirstName     []string `json:"first_name"`
	MiddleName    []string `json:"middle_name"`
	LastName      []string `json:"last_name"`
	Location      []string `json:"locations"`
	StreetAddress string   `json:"street_address"`
	Locality      string   `json:"locality"`
	Region        string   `json:"region"`
	Country       string   `json:"country"`
	PostalCode    []string `json:"postal_code"`
	Company       []string `json:"company"`
	School        []string `json:"school"`
	Phone         []string `json:"phone"`
	Email         []string `json:"email"`
	EmailHash     []string `json:"email_hash"`
	Profile       []string `json:"profile"`
	Lid           []string `json:"lid"`
	BirthDate     []string `json:"birth_date"`
}
