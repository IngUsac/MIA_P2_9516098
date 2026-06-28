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

Convertir el Proyecto 1 en un Backend REST sin modificar la lógica del sistema de archivos existente.

### Estado

✅ **Completado**

---

# Funcionalidades implementadas

## Proyecto 1

* [x] Analizador léxico
* [x] Parser
* [x] MKDISK
* [x] FDISK
* [x] MOUNT
* [x] MKFS
* [x] LOGIN
* [x] LOGOUT
* [x] MKGRP
* [x] RMGRP
* [x] MKUSR
* [x] RMUSR
* [x] CHGRP
* [x] MKDIR
* [x] MKFILE
* [x] CAT
* [x] REP

---

## Backend REST

* [x] Inicio en modo Consola
* [x] Inicio en modo Servidor
* [x] Endpoint principal (`/`)
* [x] Endpoint `/api/status`
* [x] Endpoint `/api/execute`
* [x] GenericCommandHandler
* [x] Respuestas JSON (`APIResponse`)
* [x] Captura de la salida del analizador
* [x] Ejecución de comandos mediante API REST
* [x] Compatibilidad con el Proyecto 1 sin modificar el analizador
* [x] Logger base

---

# Paso 2 - Nuevos comandos

## Estado

🟡 Pendiente

### Comandos a implementar

* [ ] FDISK ADD
* [ ] FDISK DELETE
* [ ] REMOVE
* [ ] EDIT
* [ ] RENAME
* [ ] COPY
* [ ] MOVE

---

# Paso 3 - Frontend React

## Estado

🟡 Pendiente

### Funcionalidades

* [ ] Login
* [ ] Home
* [ ] Consola Web
* [ ] Explorador de discos
* [ ] Explorador de particiones
* [ ] Navegador del sistema de archivos
* [ ] Visualizador de archivos
* [ ] Visualizador de reportes
* [ ] Consumo de la API REST

---

# Paso 4 - Integración

## Estado

🟡 Pendiente

### Funcionalidades

* [ ] Integración Backend + Frontend
* [ ] Manejo de sesiones
* [ ] Navegación del sistema de archivos
* [ ] Reportes dinámicos
* [ ] Manejo de errores HTTP
* [ ] Middleware CORS

---

# Paso 5 - Despliegue AWS

## Estado

🟡 Pendiente

### Infraestructura

* [ ] Crear instancia EC2
* [ ] Configurar Ubuntu
* [ ] Desplegar Backend
* [ ] Crear Bucket S3
* [ ] Desplegar Frontend
* [ ] Configurar CORS
* [ ] Pruebas de integración
* [ ] Documentación de despliegue

---

# Historial de avances

## 2026-06-28

### Hito 1

Reestructuración del proyecto para soportar dos modos de ejecución:

* Consola
* API REST

### Hito 2

Implementación del servidor HTTP.

### Hito 3

Implementación del endpoint `/api/status`.

### Hito 4

Implementación de `GenericCommandHandler`.

### Hito 5

Implementación del endpoint `/api/execute`.

### Hito 6

Captura automática de la salida del analizador y envío al cliente mediante respuestas JSON.

---

# Pruebas realizadas

## Consola

* [x] El proyecto compila correctamente.
* [x] Todos los comandos existentes continúan funcionando.

## API REST

* [x] El servidor inicia correctamente.
* [x] `/api/status` responde correctamente.
* [x] `/api/execute` ejecuta comandos del Proyecto 1.
* [x] La salida del analizador es devuelta en formato JSON.
* [x] Se verificó la ejecución de `MKDISK` mediante la API.

---

# Estado general

🟢 **Proyecto estable**

El Proyecto 1 continúa funcionando tanto en modo consola como mediante API REST.

La migración del backend se considera finalizada y validada.

El siguiente objetivo es implementar los nuevos comandos requeridos por la **Fase 2** antes de comenzar el desarrollo del Frontend React.
