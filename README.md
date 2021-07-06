# guptp [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Golang URI Path Template Parser


## Supported field types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64

byte 

float32 float64

time.Time

```

## Adding dependency
```
import "github.com/seaweed843/guptp"
```

## Example:

init module:
```
$ go mod init main
```

main.go
```
package main

import "github.com/seaweed843/guptp"
import "fmt"

func main() {
	uriPath :=`/api/v1/create/843/Sea%20Weed`
	template := `/api/v1/{Op}/{Id}/{Name}`

	//To map[string]string
	mapGot := guptp.ParseUriPathToMapStr(&uriPath, &template)
	fmt.Println(mapGot["Id"])
	//expected output: "843"

	fmt.Println(mapGot["Name"])
	//expected output: "Sea Weed"

	//To struct's fields
	structGot := struct{Op string; Id int; Name string}{}
	err := guptp.ParseUriPathToFields(&uriPath, &template, &structGot)
	if err == nil {
		fmt.Println(structGot.Id)
		//expected output: 843

		fmt.Println(structGot.Name)
		//expected output: "Sea Weed"
	}

}
```

run:
```
$ go mod tidy
$ go run main.go
```

output:
```
843
Sea Weed
843
Sea Weed
```
