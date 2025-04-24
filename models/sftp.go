package models

type SFTPConfiguration struct {
	Username string
	Password string
	Host     string
	Port     string
	BasePath string
}

type CSVRawFile struct {
	FileName         string
	RawFile          []byte
}

type Person struct {
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
    DateOfBirth string `json:"date_of_birth"` // ISO 8601 format: "YYYY-MM-DD"
    Address     string `json:"address"`
}