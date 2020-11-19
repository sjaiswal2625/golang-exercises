package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"
	"os/exec"
	"time"
)

func main(){
	run:=1
	for run>0 {
	cmd:=exec.Command("cmd", "/c", "cls")
	cmd.Stdout=os.Stdout
	cmd.Run()
	fmt.Println("Hello, Welcome to user crud operations, choose:")
	fmt.Println("1 : to add user")
	fmt.Println("2 : to delete user")
	fmt.Println("3 : to update password")
	fmt.Println("4 : to login as existing user")
	fmt.Println("0 : to exit")
	m:=readfile("userdata.txt")
	//fmt.Println(m)
	var input string
	var name string
	var passwd string
	n, err:= fmt.Scanf("%s\n",&input)
	if err !=nil || n!= 1{
	fmt.Println(n,err)
	return
	}
	switch input{
		case "1" :
			fmt.Println("Enter name")
			n, err= fmt.Scanf("%s\n",&name)
			if err !=nil || n!= 1{
			fmt.Println(n,err)
			return
			}
			name=fmt.Sprintf("%-8s",name)
			if (validate(m, name)=="n"){
				fmt.Println("Enter password")
				n, err= fmt.Scanf("%s\n",&passwd)
				if err !=nil || n!= 1{
				fmt.Println(n,err)
				return
				}
				passwd=fmt.Sprintf("%-8s",passwd)
				useradd(m, name, passwd)
				fmt.Println("User added:",name)
			}else{
				fmt.Println("User already exist:",name)
			}
		case "2" :
			fmt.Println("Enter name")
			n, err= fmt.Scanf("%s\n",&name)
			if err !=nil || n!= 1{
			fmt.Println(n,err)
			return
			}
			name=fmt.Sprintf("%-8s",name)
			if (validate(m, name)=="y"){
				userdel(m, name)
				fmt.Println("User deleted:",name)
			}else{
				fmt.Println("Invalid username:",name)
			}
		case "3" :
			fmt.Println("Enter name")
			n, err= fmt.Scanf("%s\n",&name)
			if err !=nil || n!= 1{
			fmt.Println(n,err)
			return
			}
			name=fmt.Sprintf("%-8s",name)
			if (validate(m, name)=="y"){
				fmt.Println("Enter new password")
				n, err= fmt.Scanf("%s\n",&passwd)
				if err !=nil || n!= 1{
				fmt.Println(n,err)
				return
				}
				passwd=fmt.Sprintf("%-8s",passwd)
				useradd(m, name, passwd)
				fmt.Println("Password updated for",name)
			}else{
				fmt.Println("Invalid username:",name)
			}
		case "4" :
			fmt.Println("Enter name")
			n, err= fmt.Scanf("%s\n",&name)
			if err !=nil || n!= 1{
			fmt.Println(n,err)
			return
			}
			name=fmt.Sprintf("%-8s",name)
			if (validate(m, name)=="y"){
				fmt.Println("Enter password")
				n, err= fmt.Scanf("%s\n",&passwd)
				if err !=nil || n!= 1{
				fmt.Println(n,err)
				return
				}
				passwd=fmt.Sprintf("%-8s",passwd)
				if (m[name]==passwd){
					fmt.Println("Username and password matched: Access granted for",name)
				}else{
					fmt.Println("Password not matched: Wrong password for",name)
				}
			}else{
				fmt.Println("Username not valid: Access denied for",name)
			}
		case "0" :
			fmt.Println("Thanks for using this service. Exiting!!!!")
			run=0
		default :
			fmt.Println("Feature not yet added")
	}
	
	writefile(m, "userdata.txt")
	time.Sleep(2 * time.Second)
	
	}
}

func writefile(m map[string]string, fname string){
	
	keys := make([]string, 0, len(m))
	for k,_ := range m {
		keys = append(keys, k)
	}
	//fmt.Println(keys)
	sort.Strings(keys)
	//fmt.Println(keys)
	out:="Username"+"\t"+"Password"
	for _,k:=range(keys) {
		//fmt.Println(k, m[k])
		out=out+"\n"+fmt.Sprintf("%-8s",k)+"\t"+fmt.Sprintf("%-8s",m[k])
	}
	//fmt.Println(out)
	
	f,err:=os.Create(fname)
	if err != nil {
	log.Fatal(err)
	}
	
	defer f.Close()
	
	_, err2 := f.WriteString(out)
	
	if err2 != nil {
	log.Fatal(err2)
	}
	//fmt.Println("done")
	
}

func readfile(fname string) map[string]string {
	file, err := os.Open(fname)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	count:=0
	for scanner.Scan() {
		if count==0{
			count=count+1
		}else{
			lines = append(lines, scanner.Text())
		}
	}
 
	file.Close()
 
	datamap:=make(map[string]string)
	
	
	for _, eachline := range lines {
		s1:=strings.Split(eachline,"\t")
		datamap[s1[0]] = s1[1]
	}
	//fmt.Println(datamap)
	return datamap
	
}

func useradd(m map[string]string, name string, passwd string){
	m[name] = passwd
}
func userdel(m map[string]string, name string){
	delete(m, name)
}
func validate(m map[string]string, name string) string {
	out:="n"
	for k,_:=range(m){
		if k==name {
			out="y"
			break
		}
	}
	return out
}
