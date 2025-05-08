package ssd

import "github.com/shirou/gopsutil/v3/disk"

type macosDeviceLister struct{}

func NewUSBDeviceLister() Interface {
    return macosDeviceLister{}
}

func (m macosDeviceLister) ListDevices() (devices []Device, err error) {
    partitions, err := disk.Partitions(false)
    if err != nil {
        return
    }

    for _, p := range partitions {
        devices = append(devices, Device{
            Name:       p.Device,
            FsType:     p.Fstype,
            MountPoint: p.Mountpoint,
            Opts:       p.Opts,
        })
    }
    return devices, nil
}
