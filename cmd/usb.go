package cmd

import (
    "fmt"
    "github.com/iltyty/devices-lister/pkg/usb"
    "github.com/spf13/cobra"
    "os"
    "text/tabwriter"
)

var USBCmd = &cobra.Command{
    Use:   "usb",
    Short: "List all available USB devices",
    Run: func(cmd *cobra.Command, args []string) {
        devices, err := usb.NewUSBDeviceLister().ListDevices()
        if err != nil {
            panic(err)
        }

        w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
        fmt.Fprintln(w, "USB devices:")
        fmt.Fprintln(w, "Bus\tName\tDriver\tManufacturer\tProduct ID\tVendor ID\tVersion\tSerial Number\tSpeed\tLocation ID")

        for _, device := range devices {
            fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
                device.Bus,
                device.Name,
                device.Driver,
                device.Manufacture,
                device.ProductID,
                device.VendorID,
                device.Version,
                device.SerialNumber,
                device.Speed,
                device.LocationID)
        }
        w.Flush()
    },
}
