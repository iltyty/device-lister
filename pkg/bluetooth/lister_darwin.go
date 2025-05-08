package bluetooth

import (
    "os/exec"
    "regexp"
    "strings"
)

var (
    addressPrefix   = "Address: "
    vendorIDPrefix  = "Vendor ID: "
    productIDPrefix = "Product ID: "
    minorTypePrefix = "Minor Type: "
    connectedRegex  = regexp.MustCompile(`Connected:\n([\s\S]*)Not Connected`)
)

type macosDeviceLister struct{}

func NewMacOSDeviceLister() Interface {
    return macosDeviceLister{}
}

func (m macosDeviceLister) ListDevices() (devices []Device, err error) {
    cmd := exec.Command("system_profiler", "SPBluetoothDataType")
    resp, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }

    matchRes := connectedRegex.FindStringSubmatch(string(resp))
    if matchRes == nil {
        return nil, nil
    }

    var device *Device
    lines := strings.Split(matchRes[1], "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        switch {
        case strings.HasSuffix(line, ":"):
            if device != nil {
                devices = append(devices, *device)
            }
            device = &Device{
                Name: strings.TrimSpace(line[:len(line)-1]),
            }
        case strings.HasPrefix(line, addressPrefix):
            device.Address = strings.TrimSpace(line[len(addressPrefix):])
        case strings.HasPrefix(line, vendorIDPrefix):
            device.VendorID = strings.TrimSpace(line[len(vendorIDPrefix):])
        case strings.HasPrefix(line, productIDPrefix):
            device.ProductID = strings.TrimSpace(line[len(productIDPrefix):])
        case strings.HasPrefix(line, minorTypePrefix):
            device.MinorType = strings.TrimSpace(line[len(minorTypePrefix):])
        }
    }

    if device != nil {
        devices = append(devices, *device)
    }
    return
}
