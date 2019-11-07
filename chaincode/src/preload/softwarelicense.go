
package main

import (
	"cycoreutils"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)

var SoftwareLicenseLogger = shim.NewLogger("softwarelicense")

// var selectorTemplate = `{ "selector": { "objectType": "{{.ObjectType}}",
//                         {{range $i, $e := .QueryFields}} "{{.FieldName}}": {{if .BoolFieldValue}} 
//                         { "$eq": {{.BoolFieldValue}} {{else}} { "{{.Operator}}": "{{.FieldValue}}" 
// 						{{end}} } {{if len $.QueryFields | sub 1 | eq $i | not}},{{end}} {{end}}}}`

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/// Record SoftwareLicense to the ledger
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func recordSoftwareLicense(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {
	var err error
	var collectionName string
	var Avalbytes []byte
	var SoftwareLicense = &softwarelicense{}

	// Convert the arg to a recordSoftwareLicense object
	SoftwareLicenseLogger.Info("recordSoftwareLicense() : Arguments for recordSoftwareLicense : ", args[0])
	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}

	SoftwareLicense.ObjectType = "SoftwareLicense"

	// Query and Retrieve the Full recordSoftwareLicense
	keys := []string{  SoftwareLicense.License  }
	SoftwareLicenseLogger.Debug("Keys for SoftwareLicense %s: ", keys)

	collectionName = ""

	Avalbytes, err = cycoreutils.QueryObject(stub, SoftwareLicense.ObjectType, keys, collectionName)

	if err != nil {
		return cycoreutils.ConstructResponse("SASTQRY003E", (errors.Wrapf(err, "Failed to query SoftwareLicense object")).Error(), nil)
	}

	if Avalbytes != nil {
		return cycoreutils.ConstructResponse("SASTQRY008E", fmt.Sprintf("SoftwareLicense already exists"), nil)
	}

	SoftwareLicenseBytes, _ := cycoreutils.ObjecttoJSON(SoftwareLicense)

	err = cycoreutils.UpdateObject(stub, SoftwareLicense.ObjectType, keys, SoftwareLicenseBytes,collectionName)
	
	if err != nil {
		SoftwareLicenseLogger.Errorf("recordSoftwareLicense() : Error inserting SoftwareLicense object into LedgerState %s", err)
		return cycoreutils.ConstructResponse("SASTUPD009E", (errors.Wrapf(err, "SoftwareLicense object update failed")).Error(), nil)
	}
	
	return cycoreutils.ConstructResponse("SASTREC011S", fmt.Sprintf("Successfully Recorded SoftwareLicense object"), SoftwareLicenseBytes)
}

//////////////////////////////////////////////////////////////
/// Query SoftwareLicense Info from the ledger
//////////////////////////////////////////////////////////////
func querySoftwareLicense(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {

	var err error
	var Avalbytes []byte
	var collectionName string
	var SoftwareLicense = &softwarelicense{}

	// In future , it should be > 1 and ,= mo_of_keys for object
	if len(args) != 1 {
		return cycoreutils.ConstructResponse("SUSRPARM001E", fmt.Sprintf("Expecting SoftwareLicense ID}. Received %d arguments", len(args)), nil)
	}

	SoftwareLicense.ObjectType = "SoftwareLicense"

	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrap(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}

	// Query and Retrieve the Full SoftwareLicense
	keys := []string{  SoftwareLicense.License  }
	SoftwareLicenseLogger.Info("Keys for SoftwareLicense : ", keys)

	collectionName = ""

	Avalbytes, err = cycoreutils.QueryObject(stub, SoftwareLicense.ObjectType, keys, collectionName)
		
	if err != nil {
		cycoreutils.ConstructResponse("SASTQRY003E", (errors.Wrapf(err, "Failed to query SoftwareLicense object")).Error(), nil)
	}

	if Avalbytes == nil {
		return cycoreutils.ConstructResponse("SASTDNE004E", fmt.Sprintf("SoftwareLicense not found"), nil)
	}

	err = cycoreutils.JSONtoObject(Avalbytes, SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert query result to SoftwareLicense object")).Error(), nil)
	}

	SoftwareLicenseLogger.Info("querySoftwareLicense() : Returning SoftwareLicense results")

	return cycoreutils.ConstructResponse("SASTQRY005S", fmt.Sprintf("Successfully Queried SoftwareLicense object"), Avalbytes)
}

//////////////////////////////////////////////////////////////
/// Query SoftwareLicense List from the ledger
//////////////////////////////////////////////////////////////
func querySoftwareLicenseList(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {

	var collectionName string
	collectionName = ""
    
	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"SoftwareLicense\"}}")
	SoftwareLicenseLogger.Info("Query List: queryString: ", queryString)

	queryResults, err := cycoreutils.GetQueryResultForQueryString(stub, queryString,collectionName)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTQRY006E", (errors.Wrapf(err, "Failed to query SoftwareLicense object list")).Error(), nil)
	}

	if queryResults == nil {
		return cycoreutils.ConstructResponse("SASTDNE004E", fmt.Sprintf("No SoftwareLicense record found"), nil)
	}

	return cycoreutils.ConstructResponse("SASTQRY007S", fmt.Sprintf("Successfully Retrieved the list of SoftwareLicense objects "), queryResults)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/// Update SoftwareLicense to the ledger by getting a copy of it, updating that copy, and overwriting the original to a newer version
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func updateSoftwareLicense(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {
    var err error
	var Avalbytes []byte
	var collectionName string
	
	var SoftwareLicense = &softwarelicense{}
	// Convert the arg to a updateSoftwareLicense object
	SoftwareLicenseLogger.Info("updateSoftwareLicense() : Arguments for Query: SoftwareLicense : ", args[0])
	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}
	SoftwareLicense.ObjectType = "SoftwareLicense"
	// Query and Retrieve the Full updateSoftwareLicense
	keys := []string{  SoftwareLicense.License  }
	SoftwareLicenseLogger.Info("Keys for SoftwareLicense : ", keys)

	collectionName = ""

	Avalbytes, err = cycoreutils.QueryObject(stub, SoftwareLicense.ObjectType, keys, collectionName)


	if err != nil {
		return cycoreutils.ConstructResponse("SASTQRY003E", (errors.Wrapf(err, "Failed to query SoftwareLicense object")).Error(), nil)
	}

	if Avalbytes == nil {
		return cycoreutils.ConstructResponse("SASTDNE004E", fmt.Sprintf("SoftwareLicense does not exist to update"), nil)
	}

	SoftwareLicenseBytes, _ := cycoreutils.ObjecttoJSON(SoftwareLicense)
	
	err = cycoreutils.UpdateObject(stub, SoftwareLicense.ObjectType, keys, SoftwareLicenseBytes, collectionName)
	
	if err != nil {
		SoftwareLicenseLogger.Errorf("updateSoftwareLicense() : Error updating SoftwareLicense object into LedgerState %s", err)
		return cycoreutils.ConstructResponse("SASTUPD009E", (errors.Wrapf(err, "SoftwareLicense object update failed")).Error(), nil)
	}

	return cycoreutils.ConstructResponse("SASTUPD010S", fmt.Sprintf("Successfully Updated SoftwareLicense object"), SoftwareLicenseBytes)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/// Delete SoftwareLicense from the ledger.
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func deleteSoftwareLicense(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {
    	
	var err error
	ObjectType := "SoftwareLicense"
	var collectionName string
	var SoftwareLicense = &softwarelicense{}

	collectionName = ""

	// In future , it should be > 1 and ,= mo_of_keys for object
	if len(args) != 1 {
		return cycoreutils.ConstructResponse("SUSRPARM001E", fmt.Sprintf("Expecting SoftwareLicense ID}. Received %d arguments", len(args)), nil)
	}
	// Convert the arg to a deleteSoftwareLicense object
	SoftwareLicenseLogger.Info("deleteSoftwareLicense() : Arguments for Query: SoftwareLicense : ", args[0])
	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}

	// Query and Retrieve the Full deleteSoftwareLicense
	keys := []string{  SoftwareLicense.License  }
	SoftwareLicenseLogger.Info("Keys for SoftwareLicense : ", keys)

	err = cycoreutils.DeleteObject(stub, ObjectType, keys,collectionName)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTDEL012E", (errors.Wrapf(err, "Failed to delete SoftwareLicense object")).Error(), nil)
	}

	return cycoreutils.ConstructResponse("SASTDEL013I", fmt.Sprintf("Successfully Deleted SoftwareLicense object"), nil)
}

//////////////////////////////////////////////////////////////
/// Query SoftwareLicense History from the ledger
//////////////////////////////////////////////////////////////
func querySoftwareLicenseHistory(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {
    	var err error
	var Avalbytes []byte
	var SoftwareLicense = &softwarelicense{}

	// In future , it should be > 1 and ,= mo_of_keys for object
	if len(args) != 1 {
		return cycoreutils.ConstructResponse("SUSRPARM001E", fmt.Sprintf("Expecting SoftwareLicense ID}. Received %d arguments", len(args)), nil)
	}

	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}

	// Query and Retrieve the Full SoftwareLicense
	keys := []string{  SoftwareLicense.License  }
	SoftwareLicenseLogger.Info("Keys for SoftwareLicense : ", keys)

    Avalbytes, err = cycoreutils.GetObjectHistory(stub, "SoftwareLicense", keys)
	if err != nil {
		return cycoreutils.ConstructResponse("SASTQRY003E", (errors.Wrapf(err, "Failed to query SoftwareLicense object history")).Error(), nil)
	}
	if Avalbytes == nil {
		return cycoreutils.ConstructResponse("SASTDNE004E", fmt.Sprintf("SoftwareLicense object not found"), nil)
	}

	return cycoreutils.ConstructResponse("SASTQRY014S", fmt.Sprintf("Successfully Queried SoftwareLicense object History"), Avalbytes)
}

func querySoftwareLicenseOwner(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response {
	var err error
	var collectionName string
	//var Avalbytes []byte
	var SoftwareLicense = &softwarelicense{}
						
	if len(args) != 1 {
		return cycoreutils.ConstructResponse("SUSRPARM001E", fmt.Sprintf("Expecting SoftwareLicense Owner}. Received %d arguments", len(args)), nil)
	}

	// Convert the arg to a SoftwareLicense object
	SoftwareLicenseLogger.Info("querySoftwareLicenseOwner() : Arguments for querySoftwareLicenseOwner : ", args[0])
	err = cycoreutils.JSONtoObject([]byte(args[0]), SoftwareLicense)

	if err != nil {
		return cycoreutils.ConstructResponse("SASTCONV002E", (errors.Wrapf(err, "Failed to convert arg[0] to SoftwareLicense object")).Error(), nil)
	}

	var owner = SoftwareLicense.Owner
	queryString := fmt.Sprintf("{\"selector\":{\"ObjectType\":\"SoftwareLicense\",\"Owner\":\"$owner\"}}")

	queryResults, err := GetQueryResultForQueryString(stub, queryString)
	
	if err != nil {
	cycoreutils.ConstructResponse("SASTQRY003E", (errors.Wrapf(err, "Failed to query SoftwareLicense object")).Error(), nil)
	}
						
	// if len(Avalbytes) < 3 {
	// return cycoreutils.ConstructResponse("SASTDNE004E", fmt.Sprintf("SoftwareLicense not found"), nil)
	// }
						
	// SoftwareLicenseLogger.Info("querySoftwareLicense() : Returning SoftwareLicense results")
	return cycoreutils.ConstructResponse("SASTQRY005S", fmt.Sprintf("Successfully Queried SoftwareLicense object"), queryResults)
}

// func invokeRichQuery(query richQuery, stub shim.ChaincodeStubInterface) ([]byte, error) {
// 	var selectorString bytes.Buffer
// 	var resultBytes []byte

// 	selectTempl, err := template.New("selector").Funcs(fns).Parse(selectorTemplate)
// 	if err != nil {
// 		logger.Infof("Error in the templates %s", err.Error())
// 		return nil, err
// 	}

// 	if err = selectTempl.Execute(&selectorString, query); err != nil {
// 		logger.Infof("Error executing the templates %s", err.Error())
// 		return nil, err
// 	}

// 	logger.Infof("++++InvokeRichQuery++++", selectorString.String())

// 	if resultBytes, err = cycoreutils.GetQueryResultForQueryString(stub, selectorString.String(), ""); err != nil {
// 		return nil, err
// 	}

// 	return resultBytes, nil
// }

// var fns = template.FuncMap{
// 	"eq": func(x, y interface{}) bool {
// 		return x == y
// 	},
// 	"sub": func(y, x int) int {
// 		return x - y
// 	},
// }