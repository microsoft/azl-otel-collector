// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package smartdata

// Machine ID constants
const (
	ProductUUIDPath         = "/sys/class/dmi/id/product_uuid"
	ProductManufacturerPath = "/sys/class/dmi/id/board_vendor"
	ProductModelNumberPath  = "/sys/class/dmi/id/product_name"
)

// MachineInfo holds some basic machine data.
type MachineInfo struct {
	UUID         string `json:"machine_uuid"`
	Manufacturer string `json:"manufacturer"`
	ModelNumber  string `json:"model_number"`
}
