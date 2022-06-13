package response

import "go-admin/internal/app/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
