package cmd

import (
	"fmt"
	"os"
	"strings"

	acl "github.com/hectane/go-acl/api"
	"golang.org/x/sys/windows"
)

/*
Will emulate _ls_ command and return the output as a string
*/
func ExecLs(dirPath string) (string, error) {

	var output []string

	/* Read directory */
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return "", err
	}

	/* Iterate over files & directory */
	for _, f := range files {

		/* retrieve file info & get output string */
		fInfo, err := f.Info()
		if err != nil {
			return "", err
		}

		/* retrieve file permissions */
		var owner *windows.SID
		var secDesc windows.Handle

		_ = acl.GetNamedSecurityInfo(
			f.Name(),
			acl.SE_FILE_OBJECT,
			acl.OWNER_SECURITY_INFORMATION,
			&owner,
			nil,
			nil,
			nil,
			&secDesc,
		)
		defer windows.LocalFree(secDesc)

		sOwner := owner.String()

		if f.IsDir() {
			file_out := fmt.Sprintf("%s Dir  %v %v %s", sOwner, fInfo.Size(), fInfo.ModTime(), f.Name())
			output = append(output, file_out)

		} else {
			file_out := fmt.Sprintf("%s File  %v %v %s", sOwner, fInfo.Size(), fInfo.ModTime(), f.Name())
			output = append(output, file_out)
		}
	}

	return strings.Join(output, "\n"), nil
}
