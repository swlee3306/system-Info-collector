package disk

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
	"system-Info-collector/pkg/config"
)

// GetDiskInfo executes 'df -h' command and parses the output into DiskInfo struct.
func GetDiskInfo() ([]*config.DiskInfo, error) {
	cmd := exec.Command("df", "-h")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var disks []*config.DiskInfo

	for _, line := range lines[1:] { // Skip header
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		diskInfo := &config.DiskInfo{
			Filesystems: []struct {
				Filesystem string `json:"filesystem"`
				Size       string `json:"size"`
				Used       string `json:"used"`
				Avail      string `json:"avail"`
				Use        string `json:"use%"`
				MountedOn  string `json:"mounted_on"`
			}{
				{
					Filesystem: fields[0],
					Size:       fields[1],
					Used:       fields[2],
					Avail:      fields[3],
					Use:        fields[4],
					MountedOn:  fields[5],
				},
			},
		}
		disks = append(disks, diskInfo)
	}

	return disks, nil
}

// GetDiskUsage retrieves disk usage statistics
func GetDiskUsage(path string) (*config.DiskUsage, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs(path, &stat)
	if err != nil {
		return nil, err
	}

	total := float64(stat.Blocks * uint64(stat.Bsize))
	free := float64(stat.Bfree * uint64(stat.Bsize))
	used := total - free
	usage := (used / total) * 100

	return &config.DiskUsage{
		Total: total,
		Used:  used,
		Free:  free,
		Usage: usage,
	}, nil
}
