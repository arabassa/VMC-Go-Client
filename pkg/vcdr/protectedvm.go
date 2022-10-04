package vcdr

import "encoding/json"

type GetProtectedVirtualMachinesResponse struct {
	Cursor string `json:"cursor"`
	Vms    []struct {
		ID       string      `json:"id"`
		Name     string      `json:"name"`
		Size     json.Number `json:"size"`
		VcdrVMID string      `json:"vcdr_vm_id"`
	} `json:"vms"`
}
