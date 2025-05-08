package main

import (
    "github.com/iltyty/devices-lister/cmd"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "dl",
    Short: "A CLI tool to list various system devices",
    Long:  `A CLI tool to list various system devices, including: USB devices, WiFi interfaces, Bluetooth devices, PCI devices, SSDs`,
}

func init() {
    rootCmd.AddCommand(cmd.USBCmd)
    rootCmd.AddCommand(cmd.SSDCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}
