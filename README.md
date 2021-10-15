# utils
### common tools for golang 


```go



package main

import (
	"github.com/daqnext/utils/bytes_util"
	"github.com/daqnext/utils/color_util"
	"github.com/daqnext/utils/hash_util"
	"github.com/daqnext/utils/ip_util"
	"github.com/daqnext/utils/path_util"
	"github.com/daqnext/utils/rand_util"
	"github.com/daqnext/utils/runtime_util"
	"github.com/daqnext/utils/time_util"
)

func main() {

	color_util.ColorPrintln(color_util.Yellow, time_util.GetUTCDate())
	color_util.ColorPrintln(color_util.Yellow, time_util.GetUTCDateTime())
	color_util.ColorPrintln(color_util.Red, bytes_util.Format(12312341344))
	color_util.ColorPrintln(color_util.Green, rand_util.GenRandStr(80))
	color_util.ColorPrintln(color_util.Red, runtime_util.StackString(100))
	color_util.ColorPrintln(color_util.Green, hash_util.MD5Hash_String("1234214"))
	color_util.ColorPrintln(color_util.Green, hash_util.MD5Hash_StringArray([]string{"123", "1234"}))
	color_util.ColorPrintln(color_util.Green, hash_util.MD5Hash_StringArray([]string{}))
	color_util.ColorPrintln(color_util.Green, hash_util.MD5Hash_String(""))
	color_util.ColorPrintln(color_util.Green, path_util.GetAbsPath("/subfolder/subsubfolder"))
	color_util.ColorPrintln(color_util.Green, path_util.GetAbsPath("/subfolder/xxx.json"))
	path_util.ExEPathPrintln()

	bytesnum, err := bytes_util.Parse("11.47TB") //case insensitive
	if err != nil {
		color_util.ColorPrintln(color_util.Red, err.Error())
	} else {
		color_util.ColorPrintln(color_util.Green, bytesnum)
	}

	ipstr, iperr := ip_util.GetPubIp(false)
	if iperr != nil {
		color_util.ColorPrintln(color_util.Red, iperr)
	} else {
		color_util.ColorPrintln(color_util.Blue, ipstr)
	}

}




```

