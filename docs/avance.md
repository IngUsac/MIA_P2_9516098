PROYECTO: MIA_P1_9516098 - Proyecto 1 Sistema de Archivos EXT2

HITO 01 - FINALIZACIÓN DE MKDISK Y FDISK

ESTADO GENERAL

* MKDISK finalizado y validado.
* FDISK finalizado y validado.
* Soporte para particiones Primarias, Extendidas y Lógicas.
* Manejo de EBR implementado.
* Código compilando correctamente.
* Limpieza de mensajes de depuración realizada.
* Commits realizados después de las validaciones finales.

ENTORNO

* Windows 11 + WSL2 Ubuntu 24.04
* Lenguaje: Go
* Proyecto: MIA_P1_9516098

FUNCIONALIDADES COMPLETADAS

MKDISK

* Creación de disco.
* Escritura de MBR.
* Lectura y validación de MBR.
* Persistencia de Fit del disco.

FDISK

* Creación de particiones Primarias.
* Creación de particiones Extendidas.
* Restricción de una única Extendida.
* Creación de EBR inicial.
* Creación de primera partición lógica reutilizando el EBR inicial.
* Creación de múltiples particiones lógicas.
* Lista enlazada de EBR funcional.
* Persistencia de PartFit.
* Validación de nombres duplicados.
* Validación de espacio disponible en disco.
* Validación de espacio disponible dentro de la Extendida.
* Obtención del último EBR.
* Enlace correcto mediante PartNext.

VALIDACIONES APROBADAS

* Creación de disco.
* Creación de Extendida.
* Creación de Logica1.
* Creación de Logica2.
* Creación de Logica3.
* Nombre duplicado.
* Segunda Extendida.
* Lógica demasiado grande.
* Salidas limpias sin mensajes temporales.

ESTADO DEL PROYECTO

* Avance estimado: 55% - 60%.

SIGUIENTE FASE
MOUNT

Antes de programar MOUNT se debe revisar nuevamente el enunciado para confirmar:

* Estructura de IDs de montaje.
* Formato requerido para los identificadores.
* Tabla de montajes en memoria.
* Restricciones de montaje.
* Comandos relacionados (MOUNT/UNMOUNT si aplica).

OBSERVACIÓN
Tomar este documento como punto de reanudación oficial del proyecto.
No modificar MKDISK ni FDISK salvo corrección de errores.
