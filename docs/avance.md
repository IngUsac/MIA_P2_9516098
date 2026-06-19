# Proyecto 1 - MIA

## Sistema de Archivos EXT2

### Bitácora de Avance

Fecha: Junio 2026

---

# RESUMEN GENERAL

Avance estimado del proyecto: **65%**

Se encuentra finalizada la infraestructura base del sistema:

---

# SIGUIENTE OBJETIVO

# MKDISK   100%
# FDISK    100%
# MOUNT    100%
## MKFS      95%  <------ 

# LOGIN    100%
# LOGOUT   100%

# MKGRP    100%
# RMGRP    100%

# MKUSR    100%
# RMUSR    100%

# EXECUTE  100%
# PAUSE    100%

## Usuarios y Grupos

### LOGIN
- [x] Login de usuario root
- [x] Validación de credenciales
- [x] Sesión activa global

### MKGRP
- [x] Crear grupos
- [x] Validar duplicados
- [x] Persistencia en users.txt

### RMGRP
- [x] Eliminación lógica de grupos
- [x] Validación de existencia

### MKUSR
- [x] Crear usuarios
- [x] Validar grupo existente
- [x] Validar usuario duplicado
- [x] Persistencia en users.txt

### RMUSR
- [x] Eliminación lógica de usuarios
- [x] Validación de existencia

### USERS.TXT
- [x] Lectura de múltiples bloques
- [x] Escritura de múltiples bloques
- [x] Persistencia validada