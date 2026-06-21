# CHECKPOINT PROYECTO 1 MIA

## Estado actual

### Comandos implementados y probados

* MKDISK
* RMDISK
* FDISK
* MOUNT
* MKFS
* LOGIN
* LOGOUT
* MKGRP
* RMGRP
* MKUSR
* RMUSR
* CHGRP
* EXECUTE
* PAUSE
* Comentarios (#)

### Mejoras implementadas

* Extensión obligatoria .dsk
* Ejecución de scripts mediante EXECUTE
* Soporte para PAUSE
* Lectura y escritura multibloque de users.txt
* Persistencia validada mediante logout/login
* Separación de responsabilidades:

  * filesystem.go -> acceso al sistema de archivos
  * users.go -> usuarios y grupos

### Archivos importantes

#### filesystem.go

Contiene:

* LeerInodo()
* LeerFolderBlock()
* LeerFileBlock()
* BuscarEntradaEnFolder()
* ObtenerInodoPorRuta()

#### users.go

Contiene:

* ObtenerContenidoUsersTXT()
* GuardarUsersTXT()
* BuscarUsuario()
* ExisteUsuarioActivo()
* ExisteGrupoActivo()

### Estructuras

Content:

* BName [12]byte
* BInodo int32

FolderBlock:

* BContent [4]Content

FileBlock:

* BContent [64]byte

Inode:

* IUid
* IGid
* ISize
* IBlock[15]
* IType
* IPerm

### Último avance validado

CHGRP funciona correctamente.

Prueba realizada:

chgrp -user=u1 -grp=root

Resultado:

2,U,g1,u1,123

→

2,U,root,u1,123

Persistencia validada después de logout/login.

### Próximo objetivo

Implementar CAT.

Ruta acordada:

1. LeerArchivo()
2. CAT
3. MKDIR
4. MKFILE
5. REP

### Observaciones

Actualmente el sistema soporta:

/
└── users.txt

Todavía no existen:

* MKDIR
* MKFILE

Por lo tanto la primera prueba de CAT será:

cat -file1="/users.txt"


Avance estimado: 60% – 65%

La razón es que ya completaste prácticamente toda la primera mitad del proyecto:

Comandos completados

✅ MKDISK
✅ RMDISK (si ya lo tienes implementado)
✅ FDISK (primarias, extendida y lógicas)
✅ MOUNT
✅ MKFS (EXT2 básico)
✅ LOGIN
✅ LOGOUT
✅ MKGRP
✅ RMGRP
✅ MKUSR
✅ RMUSR
✅ CHGRP
✅ EXECUTE (extra)
✅ PAUSE (extra)
✅ Comentarios en scripts (extra)
✅ CAT 

Además de las funcionalidades para :

✅ MBR
✅ Particiones
✅ EBR
✅ SuperBlock
✅ Inodos
✅ FolderBlock
✅ FileBlock
✅ Bitmap Inodos
✅ Bitmap Bloques
✅ users.txt multibloque
✅ Sesiones
✅ Persistencia en disco

Todo eso corresponde a la mayor parte de las páginas 1–33 del enunciado.

Comandos que todavía faltan 

CAT en proceso, especificado en Página 23.

Debe:

leer uno o varios archivos
respetar permisos
concatenar resultados

MKDIR Falta completo, especificado en Página 36.

Debe:

crear carpetas
soportar -p
crear inodos carpeta
crear FolderBlock
actualizar bitmaps

MKFILE Falta completo, especificado  Página 35.

Debe:

crear archivos
soportar -r
soportar -size
soportar -cont
manejar bloques múltiples
actualizar bitmaps
Reportes pendientes

Según el enunciado corregido falta implementar:

Los Reportes en genaral con el comando REP especificado en la Página 39. que hacen referencia a que reporte realizar segun el parametro -name 

ejemplo: rep -id=A118 -path=/reportes/nombre_reporte.jpg -name=MBR

el parametro -name puede tomar los siguientes valores:

MBR especificados en Página 40 

EBR especificados en Página 41

DISK especificado en Página 42.

INODE especificado en Página 43.

BLOCK especificado en Página 44.

BM_INODE especificado en Página 45.

BM_BLOC especificado en Página 46.

SB especificado en Página 47.

FILE especificado en Página 48.

LS especificado en Página 49.

TREE especificado en Página 50.

