package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

func main() {

	// Callback which will invoke callto:<argument> to call a number.
	// Note that this works just on OS X (and probably only with
	// Skype) but it shows the idea.
	// opts.Call = func(num string) {
	// 	cmd := exec.Command("open", "callto:"+num)
	// 	cmd.Start()
	// 	cmd.Process.Release()
	// }

	// Make some fake arguments to parse.
	// args := []string{
	// 	"-vv",
	// 	"--offset=5",
	// 	"-n", "Me",
	// 	"-p", "3",
	// 	"-s", "hello",
	// 	"-s", "world",
	// 	"--ptrslice", "hello",
	// 	"--ptrslice", "world",
	// 	"--intmap", "a:1",
	// 	"--intmap", "b:5",
	// 	"--filename", "hello.go",
	// 	"id",
	// 	"10",
	// 	"remaining1",
	// 	"remaining2",
	// }

	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	// _, err := flags.ParseArgs(&opts, args)
	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("Offset: %d\n", opts.Offset)
	fmt.Printf("Name: %s\n", opts.Name)
	// fmt.Printf("Ptr: %d\n", *opts.Ptr)
	// fmt.Printf("StringSlice: %v\n", opts.StringSlice)
	// fmt.Printf("PtrSlice: [%v %v]\n", *opts.PtrSlice[0], *opts.PtrSlice[1])
	// fmt.Printf("IntMap: [a:%v b:%v]\n", opts.IntMap["a"], opts.IntMap["b"])
	// fmt.Printf("Filename: %v\n", opts.Filename)
	// fmt.Printf("Args.ID: %s\n", opts.Args.ID)
	// fmt.Printf("Args.Num: %d\n", opts.Args.Num)
	// fmt.Printf("Args.Rest: %v\n", opts.Args.Rest)
}

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// Example of automatic marshalling to desired type (uint)
	Offset uint `long:"offset" description:"Offset"`

	// Example of a callback, called each time the option is found.
	// Call func(string) `short:"c" description:"Call phone number"`

	// Example of a required flag
	Name string `short:"n" long:"name" description:"A name" required:"true" env:"ENV_NAME"`
}
