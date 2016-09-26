package main

import "path/filepath"
import "flag"
import "fmt"

import "strings"
import "io/ioutil"
import "regexp"
import "os"

func main() {
	var flagPath, flagFrom, flagTo string
	flag.StringVar(&flagPath, "path", "", "path files to replace")
	flag.StringVar(&flagFrom, "from", "", "strings to replace")
	flag.StringVar(&flagTo, "to", "", "strings to replace")
	flag.Parse()

	if flagFrom == "" {
		fmt.Println("ERROR", `argument -from="" is required`)
		return
	}

	if flagTo == "" {
		fmt.Println("ERROR", `argument -to="" is required`)
		return
	}

	if flagPath == "." || flagPath == "" {
		flagPath, _ = os.Getwd()
	}

	regexImport, err := regexp.Compile(`(?s)(import(.*?)\)|import.*$)`)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	regexImportedPackage, err := regexp.Compile(`"(.*?)"`)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	found := []string{}

	err = filepath.Walk(flagPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			bts, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(bts)
			matches := regexImport.FindAllString(content, -1)
			isExists := false

		isReplacable:
			for _, each := range matches {
				for _, eachLine := range strings.Split(each, "\n") {
					matchesInline := regexImportedPackage.FindAllString(eachLine, -1)
					if err != nil {
						return err
					}

					for _, eachSubline := range matchesInline {
						if strings.Contains(eachSubline, flagFrom) {
							isExists = true
							break isReplacable
						}
					}
				}
			}

			if isExists {
				content = strings.Replace(content, `"`+flagFrom+`"`, `"`+flagTo+`"`, -1)
				content = strings.Replace(content, `"`+flagFrom+`/`, `"`+flagTo+`/`, -1)
				found = append(found, path)
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

	for _, path := range found {
		fmt.Printf("found in %s\n", path)
	}

	if len(found) == 0 {
		fmt.Println("Nothing replaced")
	} else {
		fmt.Printf("Total %d file replaced\n", len(found))
	}

}
