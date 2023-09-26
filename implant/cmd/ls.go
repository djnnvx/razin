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
    files, err := os.ReadDir(dirPath); if err != nil {
        return "", err
    }

    /* Iterate over files & directory */
    for _, f := range files {

        if f.IsDir() {

            /* retrieve file info & get output string */
            fInfo, err := f.Info()
            if err != nil {
                return "", err
            }

            /* retrieve file permissions */
            var owner *windows.SID
            var secDesc windows.Handle

            _ = acl.GetNamedSecurityInfo(
                f,
                acl.SE_FILE_OBJECT,
                acl.OWNER_SECURITY_INFORMATION,
                &owner,
                nil,
                nil,
                nil,
                &secDesc,
            );
            defer windows.LocalFree(secDesc)


            file_out := fmt.Sprintf("%v || Dir || %v || %v", fInfo.Size(), fInfo.ModTime(), f.Name())
            file_out := fmt.Sprintf("%s", )
            output = append(output, file_out)
        } else {



        }
    }

    return strings.Join(output, "\n"), nil
}
