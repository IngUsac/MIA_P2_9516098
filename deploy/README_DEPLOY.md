# Despliegue Proyecto 1 MIA

## Backend (EC2)

```bash
sudo apt update
sudo apt install golang-go -y

go version
```

Copiar proyecto.

```bash
go build
```

Crear servicio:

```bash
sudo cp deploy/ec2/mia.service /etc/systemd/system/

sudo systemctl daemon-reload

sudo systemctl enable mia

sudo systemctl start mia

sudo systemctl status mia
```

---

## Frontend (S3)

```bash
cd frontend

npm install

npm run build
```

Subir el contenido de:

```
frontend/build
```

Activar:

- Static Website Hosting

Index:

```
index.html
```

Error:

```
index.html
```

Aplicar:

```
deploy/s3/bucket-policy.json
```

---

## Variables de entorno

Producción:

```
REACT_APP_API_URL=http://IP_EC2:8080
```