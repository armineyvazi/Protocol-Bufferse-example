package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gp "google.golang.org/protobuf/proto"

	"proto_example/example/proto"
)

func main() {
	hossein := &proto.Employee{
		Id:     1,
		Name:   "hossein",
		Salary: 3.2,
	}
	ahmad := &proto.Employee{
		Id:     1,
		Name:   "ahmad",
		Salary: 3.2,
	}
	rick := &proto.Employee{
		Id:     1,
		Name:   "rick",
		Salary: 3.2,
	}

	employees := proto.Employees{
		Employees: []*proto.Employee{
			ahmad,
			rick,
			hossein,
		},
	}

	// Serialize Employees to binary data
	data, err := gp.Marshal(&employees)
	if err != nil {
		log.Fatal("Failed to marshal employees:", err)
	}

	fmt.Println("Binary data:", data)

	// Deserialize binary data back into Employees
	deserializedEmployees := &proto.Employees{}
	err = gp.Unmarshal(data, deserializedEmployees)
	if err != nil {
		log.Fatal("Failed to unmarshal employees:", err)
	}

	//Print
	fmt.Println("Deserialized employees:", deserializedEmployees)

	//insert binary file
	creatBinaryFile(data)

	//insert json
	createJsonFile(deserializedEmployees)

}

func creatBinaryFile(data []byte) {
	err := ioutil.WriteFile("Binary", data, 0644) // Save data to a file named "output.txt"
	if err != nil {
		log.Fatal("Failed to write file:", err)
	}

	fmt.Println("Data saved to file.")
}

func createJsonFile(data any) {
	// Serialize deserializedEmployees to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("failed to marshal to json %v", err)
	}

	// Create or open the output file
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Printf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Failed to write to file:%v", err)
	}

	fmt.Println("JSON data saved to file.")
}
