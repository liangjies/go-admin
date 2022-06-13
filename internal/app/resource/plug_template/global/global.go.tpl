package global

{{- if .HasGlobal }}

import "go-admin/internal/app/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}