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
