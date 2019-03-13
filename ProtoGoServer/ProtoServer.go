package main

import (
	"net/http"
	"strconv"

	proto "github.com/golang/protobuf/proto"
)

var personToSendOutProto []byte
var personToSendOutXML string
var personToSendOutJSON string

func main() {

	println("Started Server...")

	// THE PART OF THE PROTOCOL BUFFER
	var personToSendOut = &Person{
		Id:          500,
		Firstname:   "David",
		Lastname:    "Eilmsteiner",
		Email:       "test@test.com",
		City:        "Linz",
		Iban:        "BR55 1091 2392 0270 5316 9506 485Y H",
		Time:        "6:43 PM",
		Title:       "Mister",
		Shirtsize:   "M",
		PostalCode:  4240,
		MacAdress:   "97-C9-2C-F2-37-BF",
		Latitude:    float32(4.7524091),
		JobTitle:    "Student",
		Currency:    "Euro",
		Country:     "Austria",
		CountryCode: "AUT",
	}
	data, err := proto.Marshal(personToSendOut)
	if err != nil {
		println("ERROR: ", err)
	}
	personToSendOutProto = data
	println("Byte count of Proto: " + strconv.Itoa(len(personToSendOutProto)))

	// THE PART OF THE JSON
	personToSendOutJSON = `{"Person":{"id":1,"firstname":"David","lastname":"Eilmsteiner","email":"test@test.com","city":"Linz","iban":"BR55 1091 2392 0270 5316 9506 485Y H","time":"6:43 PM","title":"Mister","shirtsize":"M","postalCode":4240,"macAdress":"97-C9-2C-F2-37-BF","latitude":4.7524091,"jobTitle":"Student","currency":"Euro","country":"Austria","countryCode":"AUT"}}`
	println("Byte count of JSON: " + strconv.Itoa(len(personToSendOutJSON)))

	// THE PART OF THE XML
	personToSendOutXML = `<Person><id>1</id><firstname>David</firstname><lastname>Eilmsteiner</lastname><email>test@test.com</email><city>Linz</city><iban>BR55 1091 2392 0270 5316 9506 485Y H</iban><time>6:43 PM</time><title>Mister</title><shirtsize>M</shirtsize><postalCode>4240</postalCode><macAdress>97-C9-2C-F2-37-BF</macAdress><latitude>4.7524091</latitude><jobTitle>Student</jobTitle><currency>Euro</currency><country>Austria</country><countryCode>AUT</countryCode></Person>`
	println("Byte count of XML: " + strconv.Itoa(len(personToSendOutXML)))

	http.HandleFunc("/sendProto", sendProto)
	http.HandleFunc("/sendXml", sendXML)
	http.HandleFunc("/sendJson", sendJSON)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sendProto(w http.ResponseWriter, r *http.Request) {

	println("Received request for Protocol Buffer...")

	w.Write(personToSendOutProto)
}

func sendXML(w http.ResponseWriter, r *http.Request) {

	println("Received request for XML...")

	w.Write([]byte(personToSendOutXML))
}

func sendJSON(w http.ResponseWriter, r *http.Request) {

	println("Received request for JSON...")

	w.Write([]byte(personToSendOutJSON))
}
