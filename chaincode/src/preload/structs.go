package main

type softwarelicense struct {

ObjectType   string `json:"objectType"`

License string `json:"License"`

Owner string `json:"Owner"`

SoftwareName string `json:"SoftwareName"`
}

// type richQuery struct {
// 	ObjectType  string       `json:"objectType"`
// 	QueryFields []queryField `json:"queryField"`
// }
// type queryField struct {
// 	FieldName      string `json:"fieldName"`
// 	FieldValue     string `json:"fieldValue"`
// 	BoolFieldValue bool   `json:"boolFieldValue"`
// 	Operator       string `json:"operator"`
// }