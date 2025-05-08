package cmd

import (
    "fmt"
    "github.com/iltyty/devices-lister/pkg/ssd"
    "github.com/spf13/cobra"
    "os"
    "strings"
    "text/tabwriter"
)

var SSDCmd = &cobra.Command{
    Use:   "ssd",
    Short: "List all available SSD devices",
    Run: func(cmd *cobra.Command, args []string) {
        devices, err := ssd.NewUSBDeviceLister().ListDevices()
        if err != nil {
            panic(err)
        }

        w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
        _, _ = fmt.Fprintln(w, "SSD devices:")
        _, _ = fmt.Fprintln(w, "Name\tFilesystem Type\tMount Point\tOptions")

        for _, device := range devices {
            _, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
                device.Name,
                device.FsType,
                device.MountPoint,
                strings.Join(device.Opts, ","),
            )
        }
        _ = w.Flush()
    },
}
