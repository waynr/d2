package e2etests

import (
	"fmt"
	"os"
	"path/filepath"
)

func printToFile(tcs []testCase) {
	dataPath := filepath.Join("output")
	err := os.MkdirAll(dataPath, 0755)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	for _, tc := range tcs {
		tc := tc
		path := filepath.Join(dataPath, tc.name+".d2")
		err := os.WriteFile(path, []byte(tc.script), 0644)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
	}
}
