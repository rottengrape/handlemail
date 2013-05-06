package main

import (
	"os"
	"bufio"
	"io"
	"bytes"
	"strings"
)

var (
	pattern = []byte{'-','-','\n'}
	patternAl = []byte{'-','-',' '}

)

func errExit(err error){
	if err!=nil{
		panic(err)
	}
}

func reCreate(arg string,byts *[]byte){
	f,err:=os.Create(arg)
	defer f.Close()
	errExit(err)
	bffw:=bufio.NewWriter(f)
	p:=*byts
	bffw.Write(p[:len(p)])
	bffw.Flush()

}

func readMail(arg string)(bs []byte,needSync bool){
	f,err:=os.Open(arg)
	defer f.Close()
	errExit(err)
	bufr:=bufio.NewReader(f)
	bffbyts:=bytes.Buffer{}
	bffbyts.ReadFrom(bufr)
	bs=make([]byte,0)
	var prefix []byte
	for {
		l,err:=bffbyts.ReadBytes(byte('\n'))
		if err==io.EOF{
			break
		}
		if len(l)>3{
			prefix=l[:3]
			if bytes.Equal(prefix,pattern)||bytes.Equal(prefix,patternAl){
			needSync=true
			break

			}else{
				bs=append(bs,l...)
			}

		}else{
			bs=append(bs,l...)
		}

	}


	return

}

func handleMail(arg string){
	bs,needSync:=readMail(arg)
	if needSync{
		reCreate(arg,&bs)
	}

}

func main() {
	args := os.Args
	args = args[1:]
	pwd:=os.Getenv("PWD")+"/"
	for _,arg:=range args{
		if !strings.HasPrefix(arg,"/"){
			arg=pwd+arg
		}
		handleMail(arg)
	}

}
