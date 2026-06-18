# Proyecto 1 - MIA

## Sistema de Archivos EXT2

### Bitácora de Avance

Fecha: Junio 2026

---

# RESUMEN GENERAL

Avance estimado del proyecto: **65%**

Se encuentra finalizada la infraestructura base del sistema:

* Administración de discos.
* Administración de particiones.
* Montaje de particiones.
* Formateo EXT2.
* Creación de estructuras iniciales.
* Manejo de sesiones.
* Administración de grupos.

Actualmente se detectó una limitación en el manejo de `users.txt`, debido a que únicamente se está utilizando un bloque de 64 bytes. Antes de implementar MKUSR y RMUSR se corregirá el soporte de múltiples bloques utilizando los apuntadores del inodo.

---

# COMANDOS IMPLEMENTADOS

## MKDISK

Estado: COMPLETADO

Porcentaje: 100%

Funcionalidades:

* Creación de discos.
* Escritura de MBR.
* Lectura y validación de MBR.
* Soporte para Fit.
* Creación física del archivo.

Pruebas realizadas:

* Creación de múltiples discos.
* Lectura correcta del MBR.

Resultado:

OK

---

## FDISK

Estado: COMPLETADO

Porcentaje: 100%

Funcionalidades:

* Creación de particiones primarias.
* Creación de particiones extendidas.
* Creación de particiones lógicas.
* Manejo de EBR.
* Lista enlazada de EBR.
* Validación de nombres repetidos.
* Lectura y actualización del MBR.

Pruebas realizadas:

* Particiones primarias.
* Particiones extendidas.
* Varias particiones lógicas.
* Persistencia después de reiniciar.

Resultado:

OK

---

## MOUNT

Estado: COMPLETADO

Porcentaje: 100%

Funcionalidades:

* Montaje de particiones.
* Generación de ID.
* Administración de particiones montadas.
* Búsqueda por ID.

Pruebas realizadas:

* Montaje de particiones lógicas.
* Recuperación de información mediante ID.

Resultado:

OK

---

## MKFS

Estado: COMPLETADO

Porcentaje: 95%

Funcionalidades:

* Creación de SuperBlock.
* Inicialización de bitmaps.
* Creación de inodo raíz.
* Creación de carpeta raíz.
* Creación de users.txt.
* Actualización de bitmaps.
* Actualización de contadores libres.

Pruebas realizadas:

* Formateo de particiones lógicas.
* Lectura del SuperBlock.
* Validación de bitmaps.
* Validación de users.txt.

Pendiente:

* Soporte completo de múltiples bloques para archivos.

Resultado:

FUNCIONAL

---

# ESTRUCTURAS EXT2

## SuperBlock

Estado: COMPLETADO

Porcentaje: 100%

Implementado:

* Cantidad de inodos.
* Cantidad de bloques.
* Inodos libres.
* Bloques libres.
* Primer inodo libre.
* Primer bloque libre.
* Direcciones de bitmaps.
* Direcciones de tablas.

Resultado:

OK

---

## Bitmap de Inodos

Estado: COMPLETADO

Porcentaje: 100%

Resultado:

OK

---

## Bitmap de Bloques

Estado: COMPLETADO

Porcentaje: 100%

Resultado:

OK

---

## Inodos

Estado: COMPLETADO

Porcentaje: 90%

Implementado:

* Inodo raíz.
* Inodo users.txt.
* Lectura y escritura.

Pendiente:

* Uso de múltiples bloques en archivos.

Resultado:

FUNCIONAL

---

## FolderBlock

Estado: COMPLETADO

Porcentaje: 100%

Implementado:

* Entrada .
* Entrada ..
* Entrada users.txt

Resultado:

OK

---

## FileBlock

Estado: COMPLETADO

Porcentaje: 90%

Implementado:

* Lectura.
* Escritura.

Pendiente:

* Encadenamiento de múltiples bloques para archivos grandes.

Resultado:

FUNCIONAL

---

# ADMINISTRACIÓN DE SESIONES

## LOGIN

Estado: COMPLETADO

Porcentaje: 100%

Implementado:

* Lectura de users.txt.
* Búsqueda de usuario.
* Validación de contraseña.
* Creación de sesión activa.
* Validación de sesión existente.

Pruebas realizadas:

* root / 123.
* Usuario inexistente.
* Contraseña incorrecta.
* Sesión duplicada.

Resultado:

OK

---

## LOGOUT

Estado: COMPLETADO

Porcentaje: 100%

Implementado:

* Cierre de sesión.
* Limpieza de SesionActual.

Resultado:

OK

---

# ADMINISTRACIÓN DE GRUPOS

## MKGRP

Estado: COMPLETADO

Porcentaje: 95%

Implementado:

* Validación de root.
* Lectura de users.txt.
* Validación de nombres repetidos.
* Generación automática de ID.
* Escritura en users.txt.
* Persistencia en disco.

Pruebas realizadas:

* Creación de grupos.
* Persistencia después de reiniciar.
* Recreación de grupos eliminados.

Pendiente:

* Adaptación a users.txt multibloque.

Resultado:

FUNCIONAL

---

## RMGRP

Estado: COMPLETADO

Porcentaje: 95%

Implementado:

* Validación de root.
* Lectura de users.txt.
* Eliminación lógica mediante ID = 0.
* Persistencia en disco.

Pruebas realizadas:

* Eliminación lógica.
* Validación de grupos eliminados.
* Persistencia después de reiniciar.

Pendiente:

* Adaptación a users.txt multibloque.

Resultado:

FUNCIONAL

---

# COMANDOS PENDIENTES

## MKUSR

Estado: NO INICIADO

Porcentaje: 0%

---

## RMUSR

Estado: NO INICIADO

Porcentaje: 0%

---

# PROBLEMA IDENTIFICADO

Durante las pruebas de MKGRP se detectó que users.txt actualmente utiliza únicamente un bloque de 64 bytes.

Evidencia:

* Hasta g6 funciona correctamente.
* A partir de g7 el contenido comienza a truncarse.

Causa:

* GuardarUsersTXT() y ObtenerContenidoUsersTXT() utilizan únicamente IBlock[0].

Corrección planificada:

* Implementar lectura multibloque.
* Implementar escritura multibloque.
* Utilizar los apuntadores IBlock[15] del inodo users.txt.

---

# SIGUIENTE OBJETIVO

1. Corregir users.txt para soportar múltiples bloques.
2. Actualizar bitmaps y SuperBlock al asignar nuevos bloques.
3. Implementar MKUSR.
4. Implementar RMUSR.
