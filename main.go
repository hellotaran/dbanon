package main

import "bufio"
import "fmt"
import "github.com/mpchadwick/dbanon/src"
import "io/ioutil"
import "log"
import "os"

func main() {
	// sqlparser can be noisy
	// https://github.com/xwb1989/sqlparser/blob/120387863bf27d04bc07db8015110a6e96d0146c/ast.go#L52
	// We don't want to hear about it
	log.SetOutput(ioutil.Discard)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		fmt.Print(anomymizer.Anonymize(text))

		if err != nil {
			break
		}
	}
}