package core

import (
	"fmt"
	"os"
)

var blockChain BlockChain

func init() {
	blockChain = NewMemoryBlockChain()
}

type Client struct {
}

func (c *Client) Run() {
	c.validateArgs()
}

func (c *Client) validateArgs() {
	if len(os.Args) < 2 {
		c.help(-1)
		os.Exit(0)
	} else {
		arg1 := os.Args[1]
		switch arg1 {
		case "-l":
			println(argsContent[arg1])
		case "-h":
			c.help(-1)
		}
	}
}

var argsContent = map[string]string{
	"-h": "help",
	"-l": "list block chain",
}

func (c *Client) help(i int) {
	if i <= 0 {
		for k, v := range argsContent {
			fmt.Printf("%-4s  %s\n", k, v)
		}
	}
}
