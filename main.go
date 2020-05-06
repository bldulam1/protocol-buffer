package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simplepb "github.com/bldulam1/grpc/4-go/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
}

func jsonDemo(sm *simplepb.SimpleMessage) {
	smAsString := toJSON(sm)

	sm2 := &simplepb.SimpleMessage{}

	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)

	if err != nil {
		log.Fatalln("Could not unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm *simplepb.SimpleMessage) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(toJSON(sm2))
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	err = proto.Unmarshal(data, pb)
	if err != nil {
		log.Fatalln("Can't unmarshall bytes", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SimpleList: []int32{1, 2, 3, 4},
	}

	return &sm
}
