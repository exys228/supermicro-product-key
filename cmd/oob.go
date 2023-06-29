package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/exys228/supermicro-product-key/pkg/oob"
)

func init() {
	rootCmd.AddCommand(oobCmd)
	oobCmd.AddCommand(oobEncodeCmd)
	oobCmd.AddCommand(oobBruteForceCmd)
}

var oobCmd = &cobra.Command{
	Use:   "oob",
	Short: "OOB product key operations",
}

var oobEncodeCmd = &cobra.Command{
	Use:   "encode MAC_ADDRESS",
	Short: "Encode an OOB product key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		macAddress := args[0]

		productKey, err := oob.EncodeOOBProductKey(macAddress)
		if err != nil {
			return err
		}
		fmt.Println(productKey)

		return nil
	},
}

var oobBruteForceCmd = &cobra.Command{
	Use:   "bruteforce PRODUCT_KEY",
	Short: "Find the MAC address associated with an OOB product key by brute force",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		productKey := args[0]

		fmt.Println("searching for mac address ...")

		mac, err := oob.BruteForceProductKeyMACAddress(productKey)
		if err != nil {
			return err
		}

		fmt.Printf("found match! mac = '%s'\n", mac)
		return nil
	},
}
