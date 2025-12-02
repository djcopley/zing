{{/*
Expand the name of the chart.
*/}}
{{- define "zing.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Resolve the Redis master Service hostname for the application to connect to.
Order of precedence:
1) Explicit .Values.zing.redis.host
2) Bitnami Redis subchart fullnameOverride + "-master"
3) Default: <release-name>-redis-master
*/}}
{{- define "zing.redisHostname" -}}
{{- if .Values.zing.redis.host -}}
{{ .Values.zing.redis.host }}
{{- else if .Values.redis.fullnameOverride -}}
{{ printf "%s-master" .Values.redis.fullnameOverride }}
{{- else -}}
{{ printf "%s-redis-master" .Release.Name }}
{{- end -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "zing.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "zing.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "zing.labels" -}}
helm.sh/chart: {{ include "zing.chart" . }}
{{ include "zing.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "zing.selectorLabels" -}}
app.kubernetes.io/name: {{ include "zing.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "zing.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "zing.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
