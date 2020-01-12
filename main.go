package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)	

type Conf struct {
    code string
    arg  string
}

func main() {
	// Array of possible conf to pass to sonar-scanner cmd
	var params = []Conf{
		Conf{
			code: "sonar_host",
		 	arg: "sonar.host.url",
		},
		Conf{
			code: "sonar_token",
		 	arg: "sonar.login",
		},
		Conf{
			code: "project_key",
		 	arg: "sonar.projectKey",
		},
		Conf{
			code: "project_name",
		 	arg: "sonar.projectName",
		},
		Conf{
			code: "project_version",
		 	arg: "sonar.projectVersion",
		},
		Conf{
			code: "sources",
		 	arg: "sonar.sources",
		},
		Conf{
			code: "source_encoding",
		 	arg: "sonar.sourceEncoding",
		},
	}

	var args []string
	var i = 0
	for range params {
		var param Conf = params[i]
		i++
		// Get conf from env
		var value string = os.Getenv("PLUGIN_" + strings.ToUpper(param.code))
		if value != "" {
			// Add arg
			args = append(args, "-D" + param.arg +"="+ value)
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
