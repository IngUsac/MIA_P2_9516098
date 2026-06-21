# Proyecto 1 - MIA
## Sistema de Archivos EXT2

### Estado actual del proyecto

#### Administración de discos y particiones

- ✓ MKDISK
- ✓ RMDISK
- ✓ FDISK
- ✓ MOUNT

#### Sistema de archivos

- ✓ MKFS (EXT2)
- ✓ LOGIN
- ✓ LOGOUT

#### Administración de grupos y usuarios

- ✓ MKGRP
- ✓ RMGRP
- ✓ MKUSR
- ✓ RMUSR
- ✓ CHGRP

#### Manejo de archivos y directorios

- ✓ MKDIR
- ✓ MKFILE
- ✓ CAT

#### Utilidades

- ✓ EXECUTE
- ✓ PAUSE

---

# Características implementadas

## EXT2

- ✓ SuperBlock
- ✓ Bitmap de Inodos
- ✓ Bitmap de Bloques
- ✓ Inodos
- ✓ FolderBlock
- ✓ FileBlock
- ✓ Creación de raíz (/)
- ✓ Creación automática de users.txt

## Manejo de sesiones

- ✓ Inicio de sesión
- ✓ Cierre de sesión
- ✓ Validación de permisos root
- ✓ Validación de sesión activa

## Directorios

- ✓ Creación de directorios
- ✓ Creación recursiva
- ✓ Navegación por rutas absolutas
- ✓ Validación de nombres máximos (12 caracteres)
- ✓ Detección de directorios existentes

## Archivos

- ✓ Creación de archivos vacíos
- ✓ Creación mediante tamaño (-size)
- ✓ Creación mediante contenido externo (-cont)
- ✓ Creación recursiva (-r)
- ✓ Lectura mediante CAT
- ✓ Archivos multinivel
- ✓ Archivos de múltiples bloques

## Gestión de users.txt

- ✓ Lectura de users.txt
- ✓ Escritura dinámica de users.txt
- ✓ Crecimiento automático de bloques
- ✓ Persistencia en disco
- ✓ Eliminación lógica de grupos
- ✓ Eliminación lógica de usuarios

## Sistema de bloques

### Directos

- ✓ IBlock[0] a IBlock[11]

### Indirecto Simple

- ✓ IBlock[12]
- ✓ Escritura
- ✓ Lectura
- ✓ Validación de límites

### Límites actualmente soportados

- ✓ 12 bloques directos
- ✓ 1 bloque indirecto simple
- ✓ 28 bloques de datos totales
- ✓ Hasta 1792 bytes por archivo

---

# Pruebas completadas

## Usuarios y grupos

- ✓ Crear grupo
- ✓ Grupo duplicado
- ✓ Eliminar grupo
- ✓ Crear usuario
- ✓ Usuario duplicado
- ✓ Eliminar usuario
- ✓ Cambiar grupo de usuario

## Directorios

- ✓ Directorio raíz
- ✓ Directorios anidados
- ✓ Creación recursiva
- ✓ Directorio existente
- ✓ Nombre inválido

## Archivos

- ✓ Archivo vacío
- ✓ Archivo con tamaño
- ✓ Archivo con contenido externo
- ✓ Archivo en rutas anidadas
- ✓ Archivos multibloque
- ✓ Archivos usando indirecto simple

## Sesiones

- ✓ Login correcto
- ✓ Login duplicado
- ✓ Logout
- ✓ Restricciones de root

---

# Pendiente según el enunciado

## Reportes

- ☐ REP MBR
- ☐ REP DISK
- ☐ REP SB
- ☐ REP BM_INODE
- ☐ REP BM_BLOCK
- ☐ REP INODE
- ☐ REP BLOCK
- ☐ REP FILE
- ☐ REP LS
- ☐ REP TREE

---

# Estado general

Avance estimado:

- Sistema de archivos: ~95%
- Administración de usuarios: 100%
- Administración de archivos y directorios: 100%
- Reportería: 0%

Pendiente principal: implementación completa de los comandos REP.