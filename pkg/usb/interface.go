package usb

type Interface interface {
    ListDevices() ([]Device, error)
}

type Device struct {
    Bus                   string
    Name                  string
    Driver                string
    ProductID             string
    VendorID              string
    Version               string
    SerialNumber          string
    Speed                 string
    Manufacture           string
    LocationID            string
    CurrentAvailable      string
    CurrentRequired       string
    ExtraOperatingCurrent string
}
