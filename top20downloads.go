package main

import (
	"fmt"
	"os/exec"
)

func sumofdatausage() string {
	cmd, err := exec.Command("/bin/sh", "sumofdatausage.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("total size of artifacts in bytes")
	fmt.Println(string(cmd))
	return string(cmd)
}
func top20downloads() string {
	cmd, err := exec.Command("/bin/sh", "top20downloads.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("Noofdownloads repository path Sizeoftheartifact")
	fmt.Println(string(cmd))
	return string(cmd)
}
func top20uploads() string {
	cmd, err := exec.Command("/bin/sh", "top20uploads.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("Noofuploads repositorypath Sizeoftheartifact")
	fmt.Println(string(cmd))
	return string(cmd)
}
func main() {
	top20downloads()
	sumofdatausage()
	top20uploads()
}
