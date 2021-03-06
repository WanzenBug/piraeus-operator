{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "operator.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "operator.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "operator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "operator.labels" -}}
app.kubernetes.io/name: {{ include "operator.name" . }}
helm.sh/chart: {{ include "operator.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Endpoint URL of LINSTOR controller
*/}}
{{- define "controller.endpoint" -}}
  {{- if empty .Values.linstorHttpsClientSecret -}}
    http://{{ template "operator.fullname" . }}-cs.{{ .Release.Namespace }}.svc:3370
  {{- else -}}
    https://{{ template "operator.fullname" . }}-cs.{{ .Release.Namespace }}.svc:3371
  {{- end -}}
{{- end -}}

{{/*
Environment to pass to stork
*/}}
{{- define "stork.config" -}}
- name: LS_CONTROLLERS
  value: {{ template "controller.endpoint" . }}
{{- if not (empty .Values.linstorHttpsClientSecret) }}
- name: LS_USER_CERTIFICATE
  valueFrom:
    secretKeyRef:
      name: {{ .Values.linstorHttpsClientSecret }}
      key: client.cert
- name: LS_USER_KEY
  valueFrom:
    secretKeyRef:
      name: {{ .Values.linstorHttpsClientSecret }}
      key: client.key
- name: LS_ROOT_CA
  valueFrom:
    secretKeyRef:
      name: {{ .Values.linstorHttpsClientSecret }}
      key: ca.pem
{{- end -}}
{{- end -}}
