

var n{{.Varname}} = []datas.{{.Varname}}{
{{- range  .Data}}
{ {{- range $key, $element := .}}
    {{- if eq $key "GVA_MODEL" }}
        GVA_MODEL:global.GVA_MODEL{ID:{{ $element}},CreatedAt: time.Now(), UpdatedAt: time.Now()},
    {{- else}}
        {{$key}}:{{ $element}},
    {{- end}}
{{- end}} },
{{- end }}
}



