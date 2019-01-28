package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

var (
	xdebug bool
	xncert int
	xfile  string
)

func init() {
        flag.BoolVar(&xdebug,"debug", false, "enable debugging")
        flag.IntVar(&xncert,"n", 1, "certificate number")
}


func main() {
	var ff *os.File

	flag.Parse()
	xfile:=flag.Arg(0)
	if xfile=="" {
		xfile="<STDIN>"
		ff=os.Stdin
	} else{
		var err error
		ff,err=os.Open(xfile)
		if err!=nil {
			fmt.Printf("Can't open file %s\n",xfile)
			os.Exit(1)
		}
		defer ff.Close()
	}
	if xdebug { fmt.Printf(":: Extracting certificate %d from file %s\n",xncert,xfile) }
	rr:=bufio.NewScanner(ff)
	
	var cert string
	var counter int=0
	var incert bool=false
	for  {
		ok:=rr.Scan()
		if !ok { 
			if xdebug { fmt.Printf(":: Exiting\n") } 
			break 
		} 
		s:=rr.Text()
		if xdebug { fmt.Printf(":: %s\n",s) }
		if s=="-----BEGIN CERTIFICATE-----" {
			if xdebug { fmt.Printf(":: in") }
			incert=true
			counter=counter+1
		}
		if incert {cert=cert+s+"\n"}
		if s=="-----END CERTIFICATE-----" {
			if xdebug { fmt.Printf(":: out") }
			incert=false
			if counter==xncert {
				fmt.Printf(cert)
			}
			cert=""
		}
	}
}
