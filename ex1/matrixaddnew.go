package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"sync"
	"log"
	"os"
)

func main() {
/////////////////////////////////////////////// FIRST MATRIX //////////////////////////////////////////////////////////////////////////////	
	data, err := ioutil.ReadFile("file1.txt")
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}
	var(
		rows int
		columns int
		total int
	)
	s:= string(data)
	s1:=strings.Split(s, " ")
	for i,v:=range(s1){
	if(strings.Contains(v, "\n")){
	columns = i
	break
	}
	}
	total = len(s1)
	rows = total/columns
	fmt.Printf("Matrix1 size is (%d x %d)\n",rows,columns)
	//fmt.Println(s1)
	for i:=1;i<total;i++{
	if(i%columns==0){
	s1[i]= s1[i][2:]
	}
	}
	count:=0
	var xi [][]int=make([][]int,rows,columns)
	for r:=0;r<rows;r++{
		for c:=0;c<columns;c++{
		a,_:=strconv.ParseInt(s1[count],10,64)
		xi[r]=append(xi[r],int(a))
		count=count+1
		}
	}
	a:=xi

///////////////////////////////////////////////////////////  SECOND MATRIX     //////////////////////////////////////////////////////////////
	data, err = ioutil.ReadFile("file2.txt")
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}
	var(
		rows2 int
		columns2 int
		total2 int
	)
	s= string(data)
	s1=strings.Split(s, " ")
	for i,v:=range(s1){
	if(strings.Contains(v, "\n")){
	columns2 = i
	break
	}
	}
	total2 = len(s1)
	rows2 = total2/columns2
	fmt.Printf("Matrix2 size is (%d x %d)\n",rows2,columns2)
	//fmt.Println(s1)
	for i:=1;i<total;i++{
	if(i%columns==0){
	s1[i]= s1[i][2:]
	}
	}
	count=0
	//var xi [][]int=make([][]int,rows,columns)
	for r:=0;r<rows;r++{
		for c:=0;c<columns;c++{
		a,_:=strconv.ParseInt(s1[count],10,64)
		xi[r]=append(xi[r],int(a))
		count=count+1
		}
	}
	b:=xi


///////////////////////////////////////////////////////////// ADDITION ////////////////////////////////////////////////////////////////
	
	var c [][]int = make([][]int,len(a),len(a))
	var wg sync.WaitGroup
	wg.Add(len(a))
	for i:=0;i<len(a);i++{
		I:=i
		go func(){
		c[I]=arrAdd(a[I], b[I])
		wg.Done()
		}()
		
	}
	wg.Wait()
	//fmt.Println(c)

////////////////////////////////////////////////////  WRITING OUTPUT IN FILE ///////////////////////////////////////////////////////////////////////////
	var str string
	//columns:=len(a[0])
	for r:=0;r<rows;r++{
		for co:=0;co<columns;co++{
			if(r==(rows-1))&&(co==(columns-1)){
			str=str+strconv.Itoa(c[r][co])
			}else{
				str=str+strconv.Itoa(c[r][co])+" "
				}
		}
		if(r!=(rows-1)){
		str=str+"\n"
		}
	}
	fmt.Println(str)
	
	//s:="1 2 3 4 5 \n6 7 8 9 10 \n11 12 13 14 15 \n16 17 18 19 20"
	f,err:=os.Create("out.txt")
	if err != nil {
	log.Fatal(err)
	}
	
	defer f.Close()
	
	_, err2 := f.WriteString(str)
	
	if err2 != nil {
	log.Fatal(err2)
	}
	fmt.Println("done")
	
	
}
//////////////////////////////////////////// FUNCTION TO ADD TWO SLICES ////////////////////////////////////////////////////////////////////////
func arrAdd(a,b []int) []int{
	var arrC []int = make([]int, len(a), len(a))
	ca:=make(chan int)
	cb:=make(chan int)
	cc:=make(chan int)
	go func(){
	for _,v:=range(a){
	ca<-v
	}
	}()
	go func(){
	for _,v:=range(b){
	cb<-v
	}
	}()
	go func(){
	for i:=0;i<len(a);i++{
	cc<-(<-ca + <-cb)
	}
	}()
	for i:=0;i<len(a);i++{
	arrC[i]=<-cc
	}
	close(ca)
	close(cb)
	close(cc)
	return arrC
}
