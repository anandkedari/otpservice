package main
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "strings"
)

type Message struct{
 SMSBody string `json:"smsbody"`
 SMSSenderNumber string `json:"smssender"`
 SMSRecepientNumber string `json:"smsrecepient"`
}

type ErrorMessage struct{
    ErrorCode string `json:"error_code"`
    ErrorDescription string `json:"error_desc"`
}

var Messages []Message
var errMessage ErrorMessage

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "OTP Service running")
}

func main() {
	Messages = []Message{
        // Message{SMSBody:"OTP1", SMSSenderNumber:"1234", SMSRecepientNumber:"9167471247"},
        // Message{SMSBody:"OTP2", SMSSenderNumber:"1111", SMSRecepientNumber:"916710000"},
	}
    handleRequests()
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", home)

    //Get SMS
    myRouter.HandleFunc("/sms", returnSMSBody).Methods("GET")
    myRouter.HandleFunc("/sms1", returnSMSBody).Methods("GET")
    myRouter.HandleFunc("/sms/{SMSRecepientNumber}", returnSingleSMSBody).Methods("GET")
    //Post new SMS
    myRouter.HandleFunc("/sms", addNewMessage).Methods("POST")
    //Post new SMS
    myRouter.HandleFunc("/sms/{SMSRecepientNumber}", deleteMessage).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func addNewMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // get the body of our POST request
    reqBody, _ := ioutil.ReadAll(r.Body)
    var message Message 
    json.Unmarshal(reqBody, &message) //unmarshal and save request in message variable

    removeWhiteSpaces(&message);
    if(validateMessage(message) && filteredMessage(message)){
        //deletes existing object from Messages array before adding new object for specified mobile number
        deleteExistingMessage(message)

        // Add new message to Messages array
        Messages = append(Messages, message)
        json.NewEncoder(w).Encode(message)
    } else {
        w.WriteHeader(400) // Return 400 Bad Request.
        json.NewEncoder(w).Encode(errMessage)
    }
}

func deleteExistingMessage(mes Message) {
    number := mes.SMSRecepientNumber
    fmt.Println("Endpoint Hit: deleteMessage from " + number)
    for index, message := range Messages {
        if message.SMSRecepientNumber == number {
            Messages = append(Messages[:index], Messages[index+1:]...)
        }
    }
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    number := vars["SMSRecepientNumber"]
    fmt.Println("Endpoint Hit: deleteMessage from " + number)
    for index, message := range Messages {
        if message.SMSRecepientNumber == number {
            Messages = append(Messages[:index], Messages[index+1:]...)
        }
    }
}

func returnSMSBody(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnSMSBody")
    	json.NewEncoder(w).Encode(Messages)
}

func returnSingleSMSBody(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: returnSingleSMSBody")
    vars := mux.Vars(r)
    key := vars["SMSRecepientNumber"]
    for _, currentmessage := range Messages {
        if currentmessage.SMSRecepientNumber == key {
            fmt.Println("SMS : " + currentmessage.SMSBody)
            json.NewEncoder(w).Encode(currentmessage)
        }
    }
}

func validateMessage(mes Message) bool{
    if ( (mes.SMSBody == "") || (mes.SMSSenderNumber == "")  || (mes.SMSRecepientNumber == "") ){
        errMessage.ErrorCode = "VALIDATION_FAILURE"
        errMessage.ErrorDescription = "Request body validation failed, check if all mandatory data is available"
        return false    
    }
    return true
}

func filteredMessage(mes Message) bool{
    if (strings.Contains(strings.ToLower(mes.SMSBody),strings.ToLower("OTP")) && strings.Contains(strings.ToLower(mes.SMSSenderNumber),strings.ToLower("IDFC"))){
        return true    
    }
    errMessage.ErrorCode = "FILTERATION_FAILURE"
    errMessage.ErrorDescription = "Request doesn't satisfy filter criteria(OTP & IDFC)"
    return false
}

func removeWhiteSpaces(mes *Message){
    mes.SMSBody = strings.TrimSpace(mes.SMSBody)
    mes.SMSSenderNumber = strings.TrimSpace(mes.SMSSenderNumber)
    mes.SMSRecepientNumber = strings.TrimSpace(mes.SMSRecepientNumber)
}