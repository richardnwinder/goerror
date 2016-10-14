# goerror
Error message module for golang. Create an error message with filename, lineno, status, arg, and message parameters.

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
		err := c.String("error")
		fmt.Println("Error.Value = " + err)
		fmt.Println("create New(Error.Value)..........")
    
    // goerror.New() method initialises GoError object from a formatted input string
		fmt.Println("err := goerror.New(Error.Value)")
    
		gerr := goerror.New(err)
    
    // goerror.UserError() method returns the formatted user GoError message
		fmt.Println("test err.UserError().............")
    
		fmt.Println(gerr.UserError())
    
    // goerror.CodeError() method returns the formatted code GoError message
		fmt.Println("test err.CodeError().............")
    
		fmt.Println(gerr.CodeError())
    
    // goerror.PrintError() method prints the formatted user GoError message to stdout
		fmt.Println("test err.PrintError()............")
    
		gerr.PrintError()
    
    // goerror.PrintCodeError() method prints the formatted code GoError message to stdout
		fmt.Println("test err.PrintCodeError()........")
    
		gerr.PrintCodeError()
    
    // goerror.IsNil() method tests if a valid error exists
		fmt.Println("test err.IsNil().................")
    
		if gerr.IsNil() {
			fmt.Println("gerr.IsNil() == true")
		} else {
			fmt.Println("gerr.IsNil() == false")
		}
    
    // a new GoError object can also be generated from a standard error
    
		err := errors.New("A test error")
		testerr := goerror.FromError(err)
    
    // and as is usual practice can be read as standard error
    
		fmt.Println(testerr)
		}
	app.Run(os.Args)
	}
