package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)	

type Conf struct {
    code string
    arg  string
}

func main() {
	var args []string
	for _, envVar := range os.Environ() {
		if (len(envVar) > 13 && string(envVar[0:13]) == "PLUGIN_SONAR.") {
			args = append(args, "-D" + toCamelCase(envVar[7:len(envVar)]))
		}
	}
	var debug string = os.Getenv("PLUGIN_DEBUG")
	// Print cmd in debug mof
	if (debug == "true") {
		fmt.Printf("DEBUG: run cmd: sonar-scanner %s\n", strings.Join(args, " "))
	}
	// Exec scan
	cmd := exec.Command("sonar-scanner", args...)
	output, err := cmd.CombinedOutput()

	// Print output
	if len(output) > 0 {
		fmt.Printf("INFO: Sonarqube code analysis result:\n %s\n", string(output))
	}

	// Error
	if err != nil {
		fmt.Printf("ERROR: Sonarqube code analysis error:\n %s\n", err)
	}
}

var link = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toCamelCase(str string) string {
	strValueArray := strings.Split(str , "=") 
  strArray := strings.Split(strings.ToLower(strValueArray[0]) , ".") 

  i :=0
  for _, part := range strArray  {
    strArray[i] = kebabToCamelCase(part)
    i++
  }

  return strings.Join(strArray ,".") + "=" + strValueArray[1]
}

 func kebabToCamelCase(kebabInput string) string {
	isToUpper := false
	camelCase := ""
	for _, char := range kebabInput  {
		
		if isToUpper && string(char) != "-"  {
			camelCase += strings.ToUpper(string(char))
			isToUpper = false
		} else {
			if string(char) == "-" {
				isToUpper = true
			} else {
				camelCase += string(char)
			}
		}
	}
	return camelCase

 }