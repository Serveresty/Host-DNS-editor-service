package main

import (
	"YADROhostsDNS/pkg/api"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var setHostnameCmd = &cobra.Command{
		Use:   "set-hostname",
		Short: "Set hostname",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := api.NewHostnameServiceClient(conn)
			resp, err := client.SetHostname(context.Background(), &api.SetHostnameRequest{Hostname: args[0]})
			if err != nil {
				log.Fatalf("didn't set hostname: %v", err)
			}

			fmt.Println(resp.Message)
		},
	}

	var listDNSServersCmd = &cobra.Command{
		Use:   "list-dns",
		Short: "List DNS servers",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := api.NewHostnameServiceClient(conn)
			resp, err := client.ListDNSServers(context.Background(), &api.ListDNSServersRequest{})
			if err != nil {
				log.Fatalf("didn't set hostname: %v", err)
			}

			if len(resp.GetServers()) == 0 {
				fmt.Println("No DNS found")
			} else {
				fmt.Println("DNS servers: ")
				for _, server := range resp.GetServers() {
					fmt.Println(server)
				}
			}

		},
	}

	var addDNSServerCmd = &cobra.Command{
		Use:   "add-dns",
		Short: "Add DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := api.NewHostnameServiceClient(conn)
			resp, err := client.AddDNSServer(context.Background(), &api.AddDNSServerRequest{Server: args[0]})
			if err != nil {
				log.Fatalf("didn't added dns: %v", err)
			}
			fmt.Println(resp.Message)
		},
	}

	var removeDNSServerCmd = &cobra.Command{
		Use:   "remove-dns",
		Short: "Remove DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := api.NewHostnameServiceClient(conn)
			resp, err := client.RemoveDNSServer(context.Background(), &api.RemoveDNSServerRequest{Server: args[0]})
			if err != nil {
				log.Fatalf("didn't removed dns: %v", err)
			}
			fmt.Println(resp.Message)
		},
	}

	var rootCmd = &cobra.Command{Use: "hostnamecli"}
	setHostnameCmd.Flags().StringP("hostname", "H", "", "Set hostname")
	addDNSServerCmd.Flags().StringP("server", "S", "", "Add DNS server")
	removeDNSServerCmd.Flags().StringP("server", "S", "", "Remove DNS server")
	rootCmd.AddCommand(setHostnameCmd)
	rootCmd.AddCommand(listDNSServersCmd)
	rootCmd.AddCommand(addDNSServerCmd)
	rootCmd.AddCommand(removeDNSServerCmd)
	rootCmd.Execute()
}
