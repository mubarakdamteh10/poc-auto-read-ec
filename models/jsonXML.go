package models

import "encoding/xml"

// SoapEnvelope struct
type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Ns      string   `xml:"xmlns:ns,attr"`
	Header  *Header  `xml:"soapenv:Header"`
	Body    Body     `xml:"soapenv:Body"`
}

type Header struct {
	UserName string `xml:"ns:username"`
	Password string `xml:"ns:password"`
}

type Body struct {
	GetObjectByPath GetObjectByPath `xml:"ns:getObjectByPath"`
}

type GetObjectByPath struct {
	RepositoryId            string `xml:"ns:repositoryId"`
	Path                    string `xml:"ns:path"`
	Filter                  string `xml:"ns:filter,omitempty"`
	IncludeAllowableActions string `xml:"ns:includeAllowableActions,omitempty"`
	IncludeRelationships    string `xml:"ns:includeRelationships,omitempty"`
	RenditionFilter         string `xml:"ns:renditionFilter,omitempty"`
	IncludePolicyIds        string `xml:"ns:includePolicyIds,omitempty"`
	IncludeACL              string `xml:"ns:includeACL,omitempty"`
}

type XMLRawFile struct {
	FileName         string
	RawFile          []byte
}