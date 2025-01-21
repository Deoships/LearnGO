package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "add tittle")
	flag.StringVar(&cf.Edit, "edit", "", "edit tittle")
	flag.IntVar(&cf.Del, "del", -1, "delete tittle")
	flag.IntVar(&cf.Toggle, "toggle", -1, "toggle tittle")
	flag.BoolVar(&cf.List, "list", false, "list all")
	

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute (todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add !="":
		todos.add(cf.Add)
	case cf.Edit !="":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) !=2 {
			fmt.Println("invalid format")
			os.Exit(1)
		}
		
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("invalid index")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("invalid command")		
	}
}