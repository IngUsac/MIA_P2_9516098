#!/bin/bash

echo "=========================================="
echo " PROYECTO 1 MIA - INSTALADOR EC2"
echo "=========================================="

echo
echo "[1/9] Actualizando Ubuntu..."
sudo apt update -y
sudo apt upgrade -y

echo
echo "[2/9] Instalando Go..."
sudo apt install golang-go -y

echo
echo "[3/9] Verificando Go..."
go version

echo
echo "[4/9] Creando carpetas..."
mkdir -p SALIDAS
mkdir -p SALIDAS/discos
mkdir -p SALIDAS/reportes

echo
echo "[5/9] Descargando dependencias..."
go mod tidy

echo
echo "[6/9] Compilando Backend..."
go build

if [ $? -ne 0 ]; then
    echo
    echo "ERROR: No fue posible compilar el proyecto."
    exit 1
fi

echo
echo "[7/9] Instalando servicio..."

sudo cp deploy/ec2/mia.service \
/etc/systemd/system/mia.service

sudo systemctl daemon-reload

sudo systemctl enable mia

echo
echo "[8/9] Iniciando servicio..."

sudo systemctl restart mia

echo
echo "[9/9] Estado del servicio"

sudo systemctl status mia --no-pager

echo
echo "=========================================="
echo " INSTALACION FINALIZADA"
echo "=========================================="

echo
echo "Prueba el backend con:"

echo
echo "http://IP_PUBLICA:8080/api/status"
echo