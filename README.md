# razin

little implant using AES-over-TCP to communicate with C&C server

## Rationale

this is just a proof-of-concept to play around with networks, and things like that
as well as task execution in the background. In the future, I'd like to work on C2s,
and I feel like this kind of project is a good way to learn about some of the challenges
you get while doing this kind of stuff.

## Obligatory anti-skid disclaimer

Please don't use this software for evil things. Please use this software only on networks
you have been authorized to use this for. Please don't be dumb and use it for malware
campaigns or shit like that. The code is not even that good, you will get caught very
easily. I do not stand liable for your stupidity. I have written this to learn about
cybersecurity subjects and I am releasing it only to help other learn from my experience.


## Usage

### Compile and run the server

```bash

# go to server
cd server

# call `go build`
go build
```

Then, check usage prompt like so:
```bash
$ ./server --help
Command-and-control server for AES-over-TCP implants

Usage:
  server [flags]

Flags:
  -d, --debug        enable debug trace
  -h, --help         help for server
  -k, --key string   default AES key for communications (default "RAZINrazinRAZINrazinRAZINrazinRAZINraz")
  -p, --port int     default listening port (default 4444)
  -v, --version      version for server
```
