package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("##Flash Mount##\n\n")

	fmt.Println("Mount Name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("\n")

	fmt.Println("Full Network Address:")
	var address string
	fmt.Scan(&address)
	fmt.Println("\n")

	fmt.Println("Network User:")
	var user string
	fmt.Scan(&user)
	fmt.Println("\n")

	fmt.Println("Network Password:")
	var password string
	fmt.Scan(&password)
	fmt.Println("\n")

	os.Mkdir("/mnt/"+name, 0775)
	f, err := os.OpenFile("/etc/fstab", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(address + " /mnt/" + name + " cifs user,uid=500,rw,suid,username=" + user + ",password=" + password + " 0 0"); err != nil {
		fmt.Println(err)
	}

}
