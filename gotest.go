// Package main procedure in gotest.go
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/user/goerror"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotest"
	app.Usage = "./gotest [options]"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "error", Value: "main.go:120:ERROR:500:TestError", EnvVar: "ERROR"},
	}
	app.Action = func(c *cli.Context) {

		// The defined type is an object pointer derived from a go structure GoError
		// var gerr *goerror.GoError
		// This can be defined as
		// gerr := goerror.New("")
		// where gerr is a pointer to an empty GoError object

		// a new GoError object can be generated from a formatted string
		inpstr := c.String("error")
		fmt.Println("")
		fmt.Println("Error.Value = " + inpstr)
		fmt.Println("")
		fmt.Println("create New(Error.Value)..........")

		// goerror.New() method initialises GoError object from a formatted input string
		fmt.Println("gerr := goerror.New(Error.Value)")

		gerr := goerror.New(inpstr)

		fmt.Println("")

		// goerror.UserError() method returns the formatted user GoError message
		fmt.Println("test gerr.UserError().............")

		fmt.Println(gerr.UserError())

		fmt.Println("")

		// goerror.CodeError() method returns the formatted code GoError message
		fmt.Println("test gerr.CodeError().............")

		fmt.Println(gerr.CodeError())

		fmt.Println("")

		// goerror.PrintError() method prints the formatted user GoError message to stdout
		fmt.Println("test gerr.PrintError()............")

		gerr.PrintError()

		fmt.Println("")

		// goerror.PrintCodeError() method prints the formatted code GoError message to stdout
		fmt.Println("test gerr.PrintCodeError()........")

		gerr.PrintCodeError()

		fmt.Println("")

		// goerror.IsNil() method tests if a valid error exists
		fmt.Println("test gerr.IsNil().................")

		if gerr.IsNil() {
			fmt.Println("gerr.IsNil() == true")
		} else {
			fmt.Println("gerr.IsNil() == false")
		}

		fmt.Println("")

		// a new GoError object can also be generated from a standard error
		fmt.Println("create FromError(err)..........")
		fmt.Println("err := errors.New(\"A test error\")")
		fmt.Println("testerr := goerror.FromError(err)")

		err := errors.New("A test error")
		testerr := goerror.FromError(err)
		fmt.Println("")
		// and as is usual practice can be read as standard error
		fmt.Println("test fmt.Println(testerr)..........")
		fmt.Println(testerr)
	}
	app.Run(os.Args)
}
