package main

import "fmt"
type application struct {}
type config struct {}

func main() {
	flag.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	fmt.Println("Hello, World!")
}
