
package main

import (
	"fmt"
	"cycoreutils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

// ProjectKickerChaincode example - simple Chaincode implementation
type ProjectKickerChaincode struct {
	tableMap map[string]int
	funcMap  map[string]InvokeFunc
}

// Sample Asset struct
type Asset struct {
	ObjectType string `json:"objectType"`
	AssetID    string `json:"assetId"`
	AssetValue string `json:"assetValue"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Role struct {
	Role string `json:"role"`
	FunctionName string `json:"functionName"`
}

/////////////////////////////////////////////////////
// Constant for All function name that will be called from invoke
/////////////////////////////////////////////////////
const (

    QSoftwareLicense string = "querySoftwareLicense"
    RSoftwareLicense string = "recordSoftwareLicense"
    USoftwareLicense string = "updateSoftwareLicense"
    DSoftwareLicense string = "deleteSoftwareLicense"
    QSoftwareLicenseL string = "querySoftwareLicenseList"
	QSoftwareLicenseHistory string = "querySoftwareLicenseHistory"
	QSoftwareLicenseOwner string = "querySoftwareLicenseOwner"
    
)

var logger = shim.NewLogger("project-kicker-main")

type InvokeFunc func(stub shim.ChaincodeStubInterface, args []string) cycoreutils.Response

/////////////////////////////////////////////////////
// Map all the Functions here for Invoke
/////////////////////////////////////////////////////
func (t *ProjectKickerChaincode) initMaps() {
	t.tableMap = make(map[string]int)
	t.funcMap = make(map[string]InvokeFunc)

    t.funcMap[QSoftwareLicense] = querySoftwareLicense
    t.funcMap[RSoftwareLicense] = recordSoftwareLicense
    t.funcMap[USoftwareLicense] = updateSoftwareLicense
    t.funcMap[DSoftwareLicense] = deleteSoftwareLicense
    t.funcMap[QSoftwareLicenseL] = querySoftwareLicenseList
	t.funcMap[QSoftwareLicenseHistory] = querySoftwareLicenseHistory
	t.funcMap[QSoftwareLicenseOwner] = querySoftwareLicenseOwner
    
}

//////////////////////////////////////////////////////////////////////////////////
// Initialize Chaincode at Deploy Time
//////////////////////////////////////////////////////////////////////////////////
func (t *ProjectKickerChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Init ###########")
	t.initMaps()
	isInit = true
	logger.Info("ProjectKickerChaincode Init")
	return shim.Success((GetResponse("success", "Succesfully Initiated ProjectKickerChaincode")))
}

//////////////////////////////////////////////////////////////////////////////////
// Invoke Chaincode functions as requested by the Invoke Function
// In fabric 1.0 both Invoke and Query Requests are handled by Invoke
//////////////////////////////////////////////////////////////////////////////////
func (t *ProjectKickerChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//Temporay fix  if the initialization not done on the specific peer do it before Invoke a method
	if !isInit {
		t.initMaps()
		isInit = true
	}

	var returnValue cycoreutils.Response
	function, args := stub.GetFunctionAndParameters()
	f, ok := t.funcMap[function]
	var access bool
	var err error
	if ok {
		access, err = functionCheckAccess(stub, function)

	if err != nil {
		logger.Errorf("Error in fetching access. For function Name:%s ", function)
		return shim.Error(string("{\"function\":function,\"error\":err}"))
	}
	if !access {
		logger.Errorf("Invalid access for function:%s", function)
		return shim.Error(string("{\"status\":\"ERROR\",\"function\":"+function+",\"description\":\"Unauthorized\",\"detail\":\"Unauthorized\"}")) 
	}
		logger.Infof("########### Invoke/Query ###########: %s", function)

		rargs, err := stub.GetTransient()
		if err != nil {
			logger.Infof("arg[0] is not transient: %s", err.Error())
		} else {
			targs, ok1 := rargs["pargs"]
			if ok1 {
				args[0] = string(targs[:])
				logger.Info("arg[0] is transient")
			} else {
				logger.Error("arg[0] transient but could not read properly")
			}
		}

		returnValue = f(stub, args)
		returnBytes, err := cycoreutils.ObjecttoJSON(returnValue)
		if err != nil {
			logger.Errorf("Error converting the Result to JSON bytes: %s", err.Error())
			logger.Debugf("Returned Value: %+v", returnValue)
			return shim.Error(string("{\"status\":\"ERROR\",\"description\":\"Unexpected end of action. Please retry!\",\"detail\":\"" +
				(errors.Wrap(err, "Error converting the Result to JSON bytes")).Error() + "\"}"))
		}

		if returnValue.Status != "SUCCESS" {
			logger.Errorf("%s Error returned from chaincode %s", function, string(returnBytes))
			return shim.Error(string(returnBytes))
		}

		return shim.Success(returnBytes)

	}
	logger.Errorf("Invalid function name %s", function)
	return shim.Error(fmt.Sprintf("Invalid function %s", function))
}

func functionCheckAccess(stub shim.ChaincodeStubInterface, functionName string) (bool, error) {
	logger.Infof("functionCheckAccess() Entry--> ")
	var err error
	role, ok, err := cid.GetAttributeValue(stub, "role")

	if err != nil {
		return false, err
	}

	if !ok {
		return false, err
	}

	keyRole := Role{Role: role, FunctionName: functionName}

	keys,err := RoleToJSON(keyRole)

	logger.Infof("Keys for functionCheckAccess : ", keys)

	f := "queryFunctionAccessCheck"

	invokeArgs := toChaincodeArgs(f, string(keys))


	response := stub.InvokeChaincode("role_manager", invokeArgs, "defaultchannel")
	logger.Info("response status 1", response.Status)
	if response.Status != 200 {
		return false, nil
	}
	access := string(response.Payload[:])
	logger.Infof("Access recevied from AccessChecker chaincode is: %s", access)

	return true, nil
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

var isInit = false

func main() {
	logger.Info("ProjectKickerChaincode: main(): Init ")
	err := shim.Start(new(ProjectKickerChaincode))
	if err != nil {
		logger.Errorf("ProjectKickerChaincode: main(): Error starting Simple Chaincode : %s", err)
	}
}

//Prepare the response
func GetResponse(status string, message string) []byte {
	res := Response{Status: status, Message: message}
	logger.Info("GetResponse: Called For: ", res)
	response, err := cycoreutils.ObjecttoJSON(res)
	if err != nil {
		logger.Errorf(fmt.Sprintf("Invalid function %s", err))
	}
	return response
}

func init() {
	logger.SetLevel(shim.LogDebug)
}

// RoleToJSON : Converts Role Object to JSON
func RoleToJSON(role Role) ([]byte, error) {
	fmt.Println("RoleToJSON Init")
	rjson, err := json.Marshal(role)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Role Bytes : ", rjson)
	return rjson, nil
}
