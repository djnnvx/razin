package cmd

import (
	"fmt"
	"os"
	"strings"
)

func ExecLs(dirPath string) (string, error) {
	var output []string

	/* Read directory */
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err.Error() + "\n", err
	}

	/* Iterate over files & directory */
	for _, f := range files {

		/* retrieve file info & get output string */
		fInfo, err := f.Info()
		if err != nil {
			return "", err
		}

		file_out := fmt.Sprintf(
			"%v %v (%v) %s", fInfo.Mode(), fInfo.ModTime(), fInfo.Size(), f.Name(),
		)
		output = append(output, file_out)
	}

	return strings.Join(output, "\n") + "\n", nil
}
