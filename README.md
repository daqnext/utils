# utils
### common tools for golang 


```go
 package main

import (
	"github.com/daqnext/utils/bytes"
	"github.com/daqnext/utils/color"
	"github.com/daqnext/utils/hash"
	"github.com/daqnext/utils/path"
	"github.com/daqnext/utils/rand"
	"github.com/daqnext/utils/runtime"
	"github.com/daqnext/utils/time"
)

func main() {

	color.ColorPrintln(color.Yellow, time.GetUTCDate())
	color.ColorPrintln(color.Yellow, time.GetUTCDateTime())
	color.ColorPrintln(color.Red, bytes.Format(12312341344))
	color.ColorPrintln(color.Green, rand.GenRandStr(80))
	color.ColorPrintln(color.Red, runtime.StackString(100))
	color.ColorPrintln(color.Green, hash.MD5Hash_String("1234214"))
	color.ColorPrintln(color.Green, hash.MD5Hash_StringArray([]string{"123", "1234"}))
	color.ColorPrintln(color.Green, hash.MD5Hash_StringArray([]string{}))
	color.ColorPrintln(color.Green, hash.MD5Hash_String(""))
	path.ExEPathPrintln()
	color.ColorPrintln(color.Green, path.GetAbsPath("/subfolder/subsubfolder"))
	bytesnum, err := bytes.Parse("11.47TB") //case insensitive
	if err != nil {
		color.ColorPrintln(color.Red, err.Error())
	} else {
		color.ColorPrintln(color.Green, bytesnum)
	}
}

```

