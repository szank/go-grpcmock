package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	testableMain()
}

func testableMain() {
	fmt.Println("running")

	//   var flags flag.FlagSet
	//   value := flags.Bool("param", false, "")

	opts := &protogen.Options{
		//   ParamFunc: flags.Set,
	}
	opts.Run(func(p *protogen.Plugin) error {
		return nil
	})
}
