package bluetooth

type Interface interface {
    ListDevices() ([]Device, error)
}

type Device struct {
    Name      string
    Address   string
    VendorID  string
    ProductID string
    MinorType string
}
