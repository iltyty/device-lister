package cmd

import (
    "fmt"
    "github.com/iltyty/devices-lister/pkg/bluetooth"
    "github.com/spf13/cobra"
    "os"
    "text/tabwriter"
)

var BluetoothCmd = &cobra.Command{
    Use:   "bluetooth",
    Short: "List all available bluetooth devices",
    Run: func(cmd *cobra.Command, args []string) {
        devices, err := bluetooth.NewMacOSDeviceLister().ListDevices()
        if err != nil {
            panic(err)
        }

        w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
        _, _ = fmt.Fprintln(w, "Bluetooth devices:")
        _, _ = fmt.Fprintln(w, "Name\tAddress\tProduct ID\tVendor ID\tMinor Type")

        for _, device := range devices {
            _, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
                device.Name,
                device.Address,
                device.ProductID,
                device.VendorID,
                device.MinorType,
            )
        }
        _ = w.Flush()
    },
}
