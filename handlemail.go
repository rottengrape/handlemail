package main

import (
	"os"
	"bufio"
	"io"
	"bytes"
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
	bffw.Write(p[:len(p)-1])
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
	var nxt []byte
	for {
		l,err:=bffbyts.ReadBytes(byte('\n'))
		bs=append(bs,l...)
		if err==io.EOF{
			break
		}
		nxt=bffbyts.Next(3)
		if bytes.Equal(nxt,pattern)||bytes.Equal(nxt,patternAl){
			needSync=true
			break
		}else {
			bs=append(bs,nxt...)

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
	for _,arg:=range args{
		handleMail(arg)
	}

}
