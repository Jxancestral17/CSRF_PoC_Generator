package main

import (
	"os"

	"github.com/Jxancestral17/CSRF_PoC_Generator/poc"
)

var (
	meth int
	enc  int
	url  string
	data string
)

func main() {

	if len(os.Args) > 1 {

		if os.Args[1] == "-h" {
			poc.Help()

		} else {
			if os.Args[1] >= "0" || os.Args[1] <= "2" {
				switch os.Args[1] {
				case "0":
					meth = 0
				case "1":
					meth = 1
				case "2":
					meth = 2
				}
			}
			if os.Args[2] >= "0" || os.Args[2] <= "2" {
				switch os.Args[2] {
				case "0":
					enc = 0
				case "1":
					enc = 1
				case "2":
					enc = 2
				}
			}
			if os.Args[3] != "" {
				url = os.Args[3]
			}
			if os.Args[4] != "" {
				data = os.Args[4]
			}

			poc.ReturnHTML(poc.GetMethod()[meth], poc.GetEncodigHTML()[enc], data, url)

		}
	} else {
		poc.Help()
	}

}
