package main

import (
	"fmt"
	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	var err error
	if err = bindUser(); err != nil {
		fmt.Println("==========> %s <========")
		fmt.Printf("%s\n\n", err)
	}

	fmt.Println("==========> %v <========")
	fmt.Printf("%v\n\n", err)

	fmt.Println("==========> %-v <========")
	fmt.Printf("%-v\n\n", err)

	fmt.Println("==========> %+v <========")
	fmt.Printf("%+v\n\n", err)

	fmt.Println("==========> %#-v <========")
	fmt.Printf("%#-v\n\n", err)

	fmt.Println("==========> %#+v <========")
	fmt.Printf("%#+v\n\n", err)

	if errors.IsCode(err, code.ErrEncodingFailed) {
		fmt.Println("this is a ErrEncodingFailed error")
	}

	if errors.IsCode(err, code.ErrDatabase) {
		fmt.Println("this is a ErrDatabase error")
	}

	fmt.Println(errors.Cause(err))
}

func bindUser() error {
	if err := getUser(); err != nil {
		return errors.WrapC(err, code.ErrEncodingFailed, "encoding user")
	}

	return nil
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		return errors.Wrap(err, "get user failed.")
	}

	return nil
}

func queryDatabase() error {
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found")
}
