/**
  @author: decision
  @date: 2023/9/6
  @note: Generator is used to generate a series of test params
**/

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	totpCmd := flag.NewFlagSet("totp", flag.ExitOnError)
	keypairCmd := flag.NewFlagSet("keypair", flag.ExitOnError)

	groupCmd := flag.NewFlagSet("group", flag.ExitOnError)
	groupSize := groupCmd.Int("size", 16, "Group size (default 16)")

	if len(os.Args) < 2 {
		fmt.Printf("expected 'totp', 'keypair or 'group' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "totp":
		totpCmd.Parse(os.Args[2:])
		GenerateTotpKey()
	case "keypair":
		keypairCmd.Parse(os.Args[2:])
		GenerateKeyPair()
	case "group":
		groupCmd.Parse(os.Args[2:])
		GenerateGroupKeys(*groupSize)
	}
}
