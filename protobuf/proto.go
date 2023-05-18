package protobuffers

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func TestProtoColBuffer() {

	elliot := &Person{
		Name: "Eliot",
		Age:  32,
	}
	data, err := proto.Marshal(elliot)
	fmt.Println(data)

	if err != nil {
		fmt.Println("Error In Protocal Marshal")
	}
	newEll := &Person{}
	err = proto.Unmarshal(data, newEll)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newEll.GetAge())
	fmt.Println(newEll.GetName())
}
