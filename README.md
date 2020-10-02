# guptp
Golang URI Path Template Parser

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
	uri :=`/api/v1/create/843/Sea%20Weed`
	template := `/api/v1/{Op}/{Id}/{Name}`

	//To map[string]string
	mapGot := guptp.ParseUriPathToMapStr(&uri, &template)
	fmt.Println(mapGot["Id"])
	//expected output: "843"

	fmt.Println(mapGot["Name"])
	//expected output: "Sea Weed"

	//To struct's fields
	structGot := struct{Op string; Id int; Name string}{}
	err := guptp.ParseUriPathToFields(&uri, &template, &structGot)
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
$ go run main.go
```

output:
```
843
Sea Weed
843
Sea Weed
```
