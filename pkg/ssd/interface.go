package ssd

type Interface interface {
    ListDevices() ([]Device, error)
}

type Device struct {
    Name       string
    FsType     string
    MountPoint string
    Opts       []string
}
