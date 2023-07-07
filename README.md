
# Protocol Buffers Example

This is an example project that demonstrates the usage of Protocol Buffers in Go. It includes a sample `.proto` file defining the data structures and a main Go program that showcases serialization, deserialization, and saving data to binary and JSON files.

## Prerequisites

Before running the code, make sure you have the following dependencies installed:

- Go (https://golang.org/doc/install)
- Protocol Buffers Compiler (`protoc`) - Make sure it is in your system's PATH (https://developers.google.com/protocol-buffers)

## Setup

1. Clone this repository or copy the `.proto` file and the main Go program to your project directory.

2. Generate Go code from the `.proto` file by running the following command:

   ```
   protoc --go_out=. --go-grpc_out=. employess.proto
   ```

   This will generate the necessary Go code for the defined data structures.

3. Update the import paths in the Go program to match your project's structure. Replace `import "proto_example/example/proto"` with the appropriate import path for your project.

## Usage

The main program (`main.go`) demonstrates the following:

1. Creating instances of the `Employee` and `Employees` data structures defined in the `.proto` file.

2. Serializing the `Employees` instance to binary data using the `gp.Marshal` function.

3. Deserializing the binary data back into a new `Employees` instance using the `gp.Unmarshal` function.

4. Saving the binary data to a file named "Binary".

5. Creating a JSON representation of the deserialized `Employees` object using `json.Marshal`.

6. Saving the JSON data to a file named "output.json".

To run the code, execute the following command in your terminal:

```
go run main.go
```

## Output

The program will output the following:

- The binary data representation of the serialized `Employees` object.

- The deserialized `Employees` object.

- The message "Data saved to file." indicating that the binary data has been saved to a file named "Binary".

- The message "JSON data saved to file." indicating that the JSON representation has been saved to a file named "output.json".

Make sure to inspect the generated binary and JSON files after running the program.

Feel free to modify the code or the `.proto` file to suit your needs and explore further possibilities with Protocol Buffers in Go.

