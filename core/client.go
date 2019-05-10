package core

import (
	"fmt"
	"os"
)

type Client struct {
	blockChain *BlockChain
}

func NewClient(address string) *Client {
	return &Client{NewDbBlockChain(address, "", "")}
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
		case "list":
			println(argsContent[arg1])
		case "help":
			c.help(-1)
		case "add":
			if len(os.Args) < 3 {
				println(argsContent[arg1])
			} else {
				data := os.Args[2]
				c.blockChain.AddBlock([]byte(data))
				c.blockChain.LastBlock().Print()
			}
		case "send":
			println(argsContent[arg1])
		}
	}
}

var argsContent = map[string]string{
	"help": "help",
	"list": "list block chain",
	"add":  "[data] , add a new block with data",
	"send": "-amount [float] -from [from address] -to [to address]",
}

func (c *Client) help(i int) {
	if i <= 0 {
		for k, v := range argsContent {
			fmt.Printf("%-6s  %s\n", k, v)
		}
	}
}
