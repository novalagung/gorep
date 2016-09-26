package main

import "path/filepath"
import "flag"
import "fmt"
import "strings"
import "io/ioutil"
import "os"

func main() {
	var flagPath string
	flag.StringVar(&flagPath, "path", "", "path files to replace")
	var flagFrom string
	flag.StringVar(&flagFrom, "from", "", "strings to replace")
	var flagTo string
	flag.StringVar(&flagTo, "to", "", "strings to replace")

	flag.Parse()

	if flagFrom == "" {
		return
	} else {
		flagFrom = fmt.Sprintf("%s%s", `"`, flagFrom)
	}

	if flagTo != "" {
		flagTo = fmt.Sprintf("%s%s", `"`, flagTo)
	}

	if flagPath == "." || flagPath == "" {
		flagPath, _ = os.Getwd()
	}

	err := filepath.Walk(flagPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			bts, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(bts)
			if strings.Contains(content, flagFrom) {
				content = strings.Replace(content, flagFrom, flagTo, -1)
				fmt.Println(path)
			}

			err = ioutil.WriteFile(path, []byte(content), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
}
