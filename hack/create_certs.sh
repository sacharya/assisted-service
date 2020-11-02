#!/bin/bash
mkdir -p build/etc/assisted-service/nginx-certs
openssl req -x509 -nodes -days 365 -subj "/C=CA/ST=QC/O=Assisted Installer/CN=assisted-api.local.openshift.io" -addext "subjectAltName=DNS:assisted-api.local.openshift.io" -newkey rsa:2048 -keyout build/etc/assisted-service/nginx-certs/nginx-selfsigned.key -out build/etc/assisted-service/nginx-certs/nginx-selfsigned.crt;
chmod 644 build/etc/assisted-service/nginx-certs/nginx-selfsigned.key