# byteflow

Small Go package for working with data sizes and transfer rates.

## Installation

```bash
go get github.com/Miklakapi/byteflow
```

## Size

```go
package main

import (
	"fmt"

	"github.com/Miklakapi/byteflow"
)

func main() {
	size := 1536 * byteflow.KiB

	fmt.Println(size)                // 1.5 MiB
	fmt.Println(size.BinaryString()) // 1.5 MiB
	fmt.Println(size.DecimalString()) // 1.57 MB
	fmt.Println(size.Bytes())        // 1572864
	fmt.Println(size.Mebibytes())    // 1.5
}
```

## Rate

```go
package main

import (
	"fmt"
	"time"

	"github.com/Miklakapi/byteflow"
)

func main() {
	size := 750 * byteflow.MiB
	duration := 3 * time.Second

	rate := byteflow.PerSecond(size, duration)

	fmt.Println(rate)                 // 250 MiB/s
	fmt.Println(rate.DecimalString()) // 262.14 MB/s
	fmt.Println(rate.BytesPerSecond()) // 262144000
}
```

## Parsing

```go
size, err := byteflow.ParseSize("1.5 MiB")
if err != nil {
	panic(err)
}

rate, err := byteflow.ParseRate("250 MiB/s")
if err != nil {
	panic(err)
}

fmt.Println(size) // 1.5 MiB
fmt.Println(rate) // 250 MiB/s
```

## Helpers

```go
size := byteflow.MustParseSize("10 MiB")
rate := byteflow.MustParseRate("25 MBps")

duration := byteflow.DurationFor(size, rate)
transferred := byteflow.TransferredIn(rate, duration)

fmt.Println(duration)    // 419.4304ms
fmt.Println(transferred) // 10 MiB
```

## Units

Binary units:

```go
byteflow.KiB
byteflow.MiB
byteflow.GiB
byteflow.TiB
```

Decimal units:

```go
byteflow.KB
byteflow.MB
byteflow.GB
byteflow.TB
```

Transfer rate units:

```go
byteflow.KiBps
byteflow.MiBps
byteflow.GiBps
byteflow.TiBps

byteflow.KBps
byteflow.MBps
byteflow.GBps
byteflow.TBps
```

## Notes

`String()` uses binary units by default.

```go
fmt.Println(1536 * byteflow.B) // 1.5 KiB
```

Use `DecimalString()` for decimal units.

```go
fmt.Println((1500 * byteflow.B).DecimalString()) // 1.5 KB
```
