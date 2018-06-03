package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	tenable "github.com/mistsys/go-tenable"
)

func main() {
	accessKey := flag.String("accesskey", "", "AccessKey for Tenable.iO")
	secretKey := flag.String("secretkey", "", "SecretKey for Tenable.io")
	debug := flag.Bool("debug", false, "Debug mode")
	flag.Parse()

	client := tenable.NewClient(*accessKey, *secretKey)

	client.Debug = *debug

	deets, err := client.ScanDetail(context.Background(), "29")
	if err != nil {
		log.Println("Error getting scans detail", err)
	}
	fmt.Printf("%s", deets)
	return
	lst, err := client.ScansList(context.Background())
	if err != nil {
		log.Println("Error getting server scans list", err)
	}
	fmt.Printf("%q", lst)
	return

	//client.ScansList(context.Background())
	status, err := client.ServerStatus(context.Background())
	if err != nil {
		log.Println("Error getting server status. %s", err)
	}
	fmt.Printf("%v", status)
	props, err := client.ServerProperties(context.Background())
	if err != nil {
		log.Println("Error getting server properties. %s", err)
	}
	fmt.Printf("%v", props)
}
