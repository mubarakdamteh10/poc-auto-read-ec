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