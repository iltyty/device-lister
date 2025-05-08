package usb

import (
    "os/exec"
    "regexp"
    "strings"
)

var (
    busRegex              = regexp.MustCompile(`^\s*USB\s+(.*)\s+Bus:`)
    driverRegex           = regexp.MustCompile(`^\s*Host Controller Driver: (.*)`)
    productIDRegex        = regexp.MustCompile(`^\s*Product ID: (.*)`)
    vendorIDRegex         = regexp.MustCompile(`^\s*Vendor ID: (.*)`)
    versionRegex          = regexp.MustCompile(`^\s*Version: (.*)`)
    serialNumberRegex     = regexp.MustCompile(`^\s*Serial Number: (.*)`)
    speedRegex            = regexp.MustCompile(`^\s*Speed: (.*)`)
    manufacturerRegex     = regexp.MustCompile(`^\s*Manufacturer: (.*)`)
    locationIDRegex       = regexp.MustCompile(`^\s*Location ID: (.*)`)
    currentAvailableRegex = regexp.MustCompile(`^\s*Current Available (.*)`)
    currentRequiredRegex  = regexp.MustCompile(`^\s*Current Required (.*)`)
    extraCurrentRegex     = regexp.MustCompile(`^\s*Extra Operating Current (.*)`)
)

type macosDeviceLister struct {
}

func NewUSBDeviceLister() Interface {
    return macosDeviceLister{}
}

func (l macosDeviceLister) ListDevices() (devices []Device, err error) {
    cmd := exec.Command("system_profiler", "SPUSBDataType")
    resp, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }

    var device *Device
    lines := strings.Split(string(resp), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "USB:") || strings.TrimSpace(line) == "" {
            continue
        }

        switch {
        case busRegex.FindStringSubmatch(line) != nil:
            if device != nil && device.Name != "" {
                devices = append(devices, *device)
            }
            device = &Device{
                Bus: busRegex.FindStringSubmatch(line)[1],
            }
        case driverRegex.FindStringSubmatch(line) != nil:
            device.Driver = driverRegex.FindStringSubmatch(line)[1]
        case strings.HasSuffix(line, ":"):
            device.Name = strings.TrimSpace(line[:len(line)-1])
        case productIDRegex.FindStringSubmatch(line) != nil:
            device.ProductID = productIDRegex.FindStringSubmatch(line)[1]
        case vendorIDRegex.FindStringSubmatch(line) != nil:
            device.VendorID = vendorIDRegex.FindStringSubmatch(line)[1]
        case versionRegex.FindStringSubmatch(line) != nil:
            device.Version = versionRegex.FindStringSubmatch(line)[1]
        case serialNumberRegex.FindStringSubmatch(line) != nil:
            device.SerialNumber = serialNumberRegex.FindStringSubmatch(line)[1]
        case speedRegex.FindStringSubmatch(line) != nil:
            device.Speed = speedRegex.FindStringSubmatch(line)[1]
        case manufacturerRegex.FindStringSubmatch(line) != nil:
            device.Manufacture = manufacturerRegex.FindStringSubmatch(line)[1]
        case locationIDRegex.FindStringSubmatch(line) != nil:
            device.LocationID = locationIDRegex.FindStringSubmatch(line)[1]
        case currentAvailableRegex.FindStringSubmatch(line) != nil:
            device.CurrentAvailable = currentAvailableRegex.FindStringSubmatch(line)[1]
        case currentRequiredRegex.FindStringSubmatch(line) != nil:
            device.CurrentRequired = currentRequiredRegex.FindStringSubmatch(line)[1]
        case extraCurrentRegex.FindStringSubmatch(line) != nil:
            device.ExtraOperatingCurrent = extraCurrentRegex.FindStringSubmatch(line)[1]
        }
    }
    if device != nil && device.Name != "" {
        devices = append(devices, *device)
    }
    return
}
