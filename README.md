# Proyecto 1 - MIA
## Fase 2 - Sistema de Archivos EXT2 con API REST

### Información

**Curso:** Manejo e Implementación de Archivos (MIA)

**Proyecto:** Proyecto 1 - Fase 2

**Lenguaje Backend:** Go

**Frontend:** React (Pendiente)

**Despliegue:** AWS EC2 + AWS S3 (Pendiente)

---

# Estado del Proyecto

## Paso 1 - Migración a API REST

### Objetivo

Convertir el Proyecto 1 en un Backend REST sin modificar la lógica del sistema de archivos.

### Estado

✅ En progreso

---

# Funcionalidades implementadas

## Proyecto 1

- [x] Analizador léxico
- [x] Parser
- [x] MKDISK
- [x] FDISK
- [x] MOUNT
- [x] MKFS
- [x] LOGIN
- [x] LOGOUT
- [x] MKGRP
- [x] RMGRP
- [x] MKUSR
- [x] RMUSR
- [x] CHGRP
- [x] MKDIR
- [x] MKFILE
- [x] CAT
- [x] REP

---

## Backend REST

- [x] Inicio en modo Consola
- [x] Inicio en modo Servidor
- [x] Endpoint principal
- [x] Endpoint `/api/status`
- [x] Endpoint `/api/execute`
- [x] GenericCommandHandler
- [x] APIResponse
- [x] Logger base

---

# Pendiente del Paso 1

- [ ] Capturar la salida del analizador para enviarla al Frontend.
- [ ] Devolver mensajes de ejecución en formato JSON.
- [ ] Manejo de errores HTTP.
- [ ] Middleware CORS.
- [ ] Manejo de sesiones.

---

# Paso 2

## Nuevos comandos

Pendientes de implementación:

- [ ] REMOVE
- [ ] EDIT
- [ ] RENAME
- [ ] COPY
- [ ] MOVE
- [ ] FDISK ADD
- [ ] FDISK DELETE

---

# Paso 3

## Frontend React

Pendiente.

Debe incluir:

- [ ] Login
- [ ] Home
- [ ] Consola
- [ ] Explorador de discos
- [ ] Explorador de particiones
- [ ] Navegador del sistema de archivos
- [ ] Visualizador de archivos
- [ ] Reportes

---

# Paso 4

## Integración

Pendiente.

- [ ] Backend + Frontend
- [ ] Consumo de API
- [ ] Navegación
- [ ] Reportes dinámicos

---

# Paso 5

## AWS

Pendiente.

- [ ] Instancia EC2
- [ ] Bucket S3
- [ ] Despliegue Backend
- [ ] Despliegue Frontend
- [ ] Configuración CORS
- [ ] Pruebas finales

---

# Historial de avances

## 2026-06-28

### Hito 1

Se reorganizó el Proyecto 1 para permitir dos modos de ejecución:

- Consola
- API REST

### Hito 2

Se implementó el servidor HTTP.

### Hito 3

Se implementó el endpoint `/api/status`.

### Hito 4

Se implementó `GenericCommandHandler`, eliminando la necesidad de crear un handler por cada comando.

---

# Estado general

Proyecto estable.

Compila correctamente en:

- ✅ Consola
- ✅ Servidor REST

Actualmente se trabaja en la captura de la salida del analizador para integrarla con el Frontend.