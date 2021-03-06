package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var url string
var image string
var registry string
var repository string
var user string
var password string

func init() {
	flag.StringVar(&url, "url", "", "url of file you want download")
	flag.StringVar(&image, "image", "", "full adress of your image  you want transfer")
	flag.StringVar(&registry, "registry", "docker.io", "your registry to save your image")
	flag.StringVar(&repository, "repository", "heishui", "your repository to save your image")
	flag.StringVar(&user, "user", "heishui", "user of your repository")
	flag.StringVar(&password, "password", "", "password of your repository")
}
func main() {
	// file := GetFile("https://golang.google.cn/dl/go1.15.windows-amd64.msi")
	// fmt.Println(file)
	flag.Parse()
	fmt.Println(runtime.GOOS)
	// fmt.Println(image)
	if image != "" {
		cmd := exec.Command("docker", "pull", image)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		myImage := GetImage(image)
		fmt.Println(myImage)
		cmd = exec.Command("docker", "tag", image, myImage)
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		cmd = exec.Command("docker", "login", "--username="+user, "--password="+password)
		fmt.Println(cmd.String())
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		cmd = exec.Command("docker", "push", myImage)
		fmt.Println(cmd.String())
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
	}
	if url != "" {
		cmd := exec.Command("wget", url)
		// cmd := exec.Command("ls")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		file := GetFile(url)
		fmt.Println(file)
		f, err := os.Create("Dockerfile")
		ct := "from scratch\ncopy " + file + " /"
		defer func() {
			if err = f.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		l, err := f.WriteString(ct)
		fmt.Println(l, " bytes written successfully!")
		newFullImage := registry + "/" + repository + "/" + strings.ToLower(file)
		cmd = exec.Command("docker", "build", "-t", newFullImage, ".")
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		cmd = exec.Command("docker", "login", "--username="+user, "--password="+password)
		fmt.Println(cmd.String())
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("docker", "push", newFullImage)
		stdoutStderr, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
		Inspect(newFullImage)
	}
}

// GetImage ... comment.
func GetImage(name string) string {
	imageArry := strings.Split(name, "/")
	newImage := imageArry[len(imageArry)-1]
	newFullImage := registry + "/" + repository + "/" + newImage
	return newFullImage
}

// GetFile ... fun.
func GetFile(url string) string {
	FileArry := strings.Split(url, "/")
	file := FileArry[len(FileArry)-1]
	// fmt.Println(file)
	return file
}

// Inspect comm.
func Inspect(item string) {
	CheckJq()
	fmt.Println(
		"docker pull", item,
		// "docker inspect", item, "| jq .[0].GraphDriver.Data.UpperDir",
	)
	fmt.Println(
		// "docker pull item",
		"docker inspect", item, "| jq .[0].GraphDriver.Data.UpperDir",
	)

}

// CheckJq command.
func CheckJq() {
	cmd := exec.Command("which", "jq")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		// InstallPackage("jq")
	}
	fmt.Printf("%s\n", stdoutStderr)
}

// InstallPackage in linux.
func InstallPackage(name string) {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	if path, err := exec.LookPath("yum"); err == nil {
		fmt.Println(path)
		cmd := exec.Command("yum", "install", "-y", name)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)

	} else if path, err := exec.LookPath("apt"); err == nil {
		fmt.Println(path)
		cmd := exec.Command("apt", "update")
		cmd.CombinedOutput()
		cmd = exec.Command("apt", "install", "-y", name)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
	} else {
		fmt.Println("not support operation system")
	}
}
