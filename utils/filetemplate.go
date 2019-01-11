package utils

const defaultValues = `# Default values for %s.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
replicaCount: 1
image:
  repository: nginx
  tag: stable
  pullPolicy: IfNotPresent
nameOverride: ""
fullnameOverride: ""
service:
  type: ClusterIP
  port: 80
ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  paths: []
  hosts:
    - chart-example.local
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local
resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi
nodeSelector: {}
tolerations: []
affinity: {}
`

const defaultIgnore = `# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`

const defaultDeploymentVaule = `apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: %[1]s
  namespace: %[1]s
spec:
  replicas: 1
  selector:
    matchLabels:
      name: %[1]s
  strategy:
    rollingUpdate: 
      maxSurge: 100%%
      maxUnavailable: 30%%
    type: RollingUpdate
  template:
    metadata:
      annotations:
        quality: beta
        role: frontend-api
        sla: high
        tier: application
      labels:
        name:  %[1]s-%[1]s
        version: v1
    spec:
      containers:
      - env:
        - name: ENVIRONMENT
          value: beta
        - name: NODE_ENV
          value: develop
        - name: LOG_LEVEL
          value: debug
        image: mytest
        livenessProbe:
          httpGet:
            path: /health
            port: 3978
          initialDelaySeconds: 30
          timeoutSeconds: 3
        name: bot-ispd-gdpr
        ports:
        - containerPort: 8080
        readinessProbe:
          httpGet:
            path: /health
            port: 3978
          initialDelaySeconds: 30
          timeoutSeconds: 3
        resources:
          limits:
            cpu: 400m
            memory: 500Mi
          requests:
            cpu: 250m
            memory: 256Mi
`

const defaultServiceVaule = `apiVersion: v1
kind: Service
metadata:
  labels:
    name: %[1]s
  name: %[1]s
  namespace: %[1]s
  spec:
  ports:
  - name: http
    port: 3978
    protocol: TCP
    targetPort: 3978
  selector:
    name: %[1]s
  sessionAffinity: None
`

const defaultConfigmapValue = `apiVersion: v1
kind: ConfigMap
metadata:
  name: %[1]s
  namespace: %[1]s
`

const defaultDockerValue = `FROM alpine:latest

LABEL MAINTAINER xyz@gmail.com
`