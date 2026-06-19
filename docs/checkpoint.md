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

Además:

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
CAT

Falta completo.

Página 23.

Debe:

leer uno o varios archivos
respetar permisos
concatenar resultados
MKDIR

Falta completo.

Página 36.

Debe:

crear carpetas
soportar -p
crear inodos carpeta
crear FolderBlock
actualizar bitmaps
MKFILE

Falta completo.

Página 35.

Debe:

crear archivos
soportar -r
soportar -size
soportar -cont
manejar bloques múltiples
actualizar bitmaps
Reportes pendientes

Según el enunciado corregido:

REP general

Falta implementar el comando REP.

Página 39.

Reporte MBR

Página 40.

Reporte DISK

Página 42.

Reporte INODE

Página 43.

Reporte BLOCK

Página 44.

Reporte BM_INODE

Página 45.

Reporte BM_BLOCK

Página 46.

Reporte SB

Página 47.

Reporte FILE

Página 48.

Reporte LS

Página 49.

Reporte TREE

Página 50.

Lo que NO aparece en este enunciado

Estos comandos no están en el PDF corregido:

❌ EDIT
❌ COPY
❌ MOVE
❌ REMOVE
❌ RENAME
❌ FIND
❌ CHMOD

Por lo tanto no debes invertir tiempo en ellos para este Proyecto 1.

Ruta óptima para terminar
CAT
MKDIR
MKFILE
REP MBR
REP DISK
REP SB
REP BM_INODE
REP BM_BLOCK
REP INODE
REP BLOCK
REP FILE
REP LS
REP TREE

Con esa ruta estarías cubriendo prácticamente el 100% de los requisitos del enunciado corregido.