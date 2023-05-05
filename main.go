package main

import (
	"fmt"
	"go-gRPC/pb"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Takuma", // 修正: "name" を "Name" に変更
		Email:       "test@example.com",
		Occupation:  pb.Occupation_DEVELOPER,
		PhoneNumber: []string{"070-1234-5678", "080-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Takuma",
		},
		Birthdate: &pb.Date{
			Year:  1995,
			Month: 1,
			Day:   1,
		},
	}

	// binData, err := proto.Marshal(employee) // 修正: "binDate" を "binData" に変更
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil { // 修正: "filename:", "bindate", "perm:" を削除
	// 	log.Fatalln("Can't write", err) // 修正: "Can`s" を "Can't" に変更
	// }

	// in, err := ioutil.ReadFile("test.bin") // 修正: "filename:" を削除
	// if err != nil {
	// 	log.Fatalln("Can't read", err) // 修正: "Can`s" を "Can't" に変更
	// }

	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Failed to parse employee:", err)
	// }

	// fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Failed to encode employee:", err)
	}

	// fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Failed to parse employee:", err)
	}
	fmt.Println(readEmployee)
}
