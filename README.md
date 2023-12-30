# Snowflake ID Generator

The Snowflake ID Generator is a Go package that allows you to generate unique and distributed IDs based on a combination of timestamp, machine ID, and sequence number.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use the Snowflake ID Generator in your Go project, you can simply install it using the `go get` command:

```bash
go get -u github.com/naruepanart/snowflakeid
```

## Usage

To use the Snowflake ID Generator in your Go application, follow these steps:

1. Import the package:

   ```go
   import (
   	"fmt"
   	"github.com/naruepanart/snowflakeid"
   )
   ```

2. Create a Snowflake instance with a unique machine ID:

   ```go
   machineID := 1 // Set the machine ID according to the actual scenario
   sf := snowflakeid.NewSnowflake(int64(machineID))
   ```

3. Generate unique IDs using the `GenerateID` method:

   ```go
   for i := 0; i < 10; i++ {
       id := sf.GenerateID()
       fmt.Println(id)
   }
   ```

## Example

Here's a complete example of using the Snowflake ID Generator in a Go application:

```go
package main

import (
	"fmt"
	"github.com/naruepanart/snowflakeid"
)

func main() {
	machineID := 1 // Set the machine ID according to the actual scenario
	sf := snowflakeid.NewSnowflake(int64(machineID))

	// Generate and print 10 unique IDs
	for i := 0; i < 10; i++ {
		id := sf.GenerateID()
		fmt.Println(id)
	}
}

Adjust the machine ID based on your deployment environment.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.