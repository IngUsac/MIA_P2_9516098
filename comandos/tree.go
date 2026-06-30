package comandos




import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)


type TreeNode struct {
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Children []TreeNode `json:"children,omitempty"`
}

func ReporteTREE(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo SuperBlock",
		)

		return
	}

	var dot strings.Builder

	dot.WriteString(
		"digraph TREE {\n",
	)

	dot.WriteString(
		"rankdir=LR;\n",
	)

	dot.WriteString(
		"node [shape=plaintext];\n",
	)

	visitados := make(
		map[int32]bool,
	)

	RecorrerInodoTREE(
		archivo,
		sb,
		0,
		&dot,
		visitados,
	)

	dot.WriteString(
		"}\n",
	)

	dotPath :=
		strings.TrimSuffix(
			path,
			filepath.Ext(path),
		) + ".dot"

	err = os.WriteFile(
		dotPath,
		[]byte(dot.String()),
		0644,
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo DOT",
		)

		return
	}

	err = exec.Command(
		"dot",
		"-Tpng",
		dotPath,
		"-o",
		path,
	).Run()

	if err != nil {

		fmt.Println(
			"ERROR ejecutando Graphviz",
		)

		return
	}

	fmt.Println(
		"Imagen generada:",
		path,
	)
}

func RecorrerInodoTREE(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroInodo int32,
	dot *strings.Builder,
	visitados map[int32]bool,
) {

	if visitados[numeroInodo] {
		return
	}

	visitados[numeroInodo] = true

	inodo, err := LeerInodoPorNumero(
		archivo,
		sb,
		numeroInodo,
	)

	if err != nil {
		return
	}

	GenerarNodoInodoTREE(
		dot,
		numeroInodo,
		inodo,
	)

	for i := 0; i < 12; i++ {

		if inodo.IBlock[i] == -1 {
			continue
		}

		numBloque :=
			inodo.IBlock[i]

		if inodo.IType == '0' {

			GenerarNodoFolderTREE(
				archivo,
				sb,
				numBloque,
				dot,
			)

			fmt.Fprintf(
				dot,
				"inode%d -> folder%d;\n",
				numeroInodo,
				numBloque,
			)

			folder, err :=
				LeerFolderPorNumero(
					archivo,
					sb,
					numBloque,
				)

			if err != nil {
				continue
			}

			for j := 0; j < 4; j++ {

				nombre :=
					utilidades.BytesAString(
						folder.BContent[j].BName[:],
					)

				if nombre == "." ||
					nombre == ".." ||
					nombre == "" {
					continue
				}

				hijo :=
					folder.BContent[j].BInodo

				if hijo < 0 {
					continue
				}

				fmt.Fprintf(
					dot,
					"folder%d -> inode%d;\n",
					numBloque,
					hijo,
				)

				RecorrerInodoTREE(
					archivo,
					sb,
					hijo,
					dot,
					visitados,
				)
			}

		} else {

			GenerarNodoArchivoTREE(
				archivo,
				sb,
				numBloque,
				dot,
			)

			fmt.Fprintf(
				dot,
				"inode%d -> file%d;\n",
				numeroInodo,
				numBloque,
			)
		}
	}
}

func GenerarNodoFolderTREE(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	dot *strings.Builder,
) {

	folder, err :=
		LeerFolderPorNumero(
			archivo,
			sb,
			numeroBloque,
		)

	if err != nil {
		return
	}

	fmt.Fprintf(
		dot,
		`folder%d [
shape=plaintext
label=<
<TABLE BORDER="1" CELLBORDER="1" CELLSPACING="0" BGCOLOR="salmon">
<TR><TD COLSPAN="2">Bloque Carpeta %d</TD></TR>`,
		numeroBloque,
		numeroBloque,
	)

	for i := 0; i < 4; i++ {

		nombre :=
			utilidades.BytesAString(
				folder.BContent[i].BName[:],
			)

		fmt.Fprintf(
			dot,
			"<TR><TD>%s</TD><TD>%d</TD></TR>",
			nombre,
			folder.BContent[i].BInodo,
		)
	}

	dot.WriteString(
		"</TABLE>>];\n",
	)
}



func GenerarNodoArchivoTREE(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	dot *strings.Builder,
) {

	fileBlock, err :=
		LeerFilePorNumero(
			archivo,
			sb,
			numeroBloque,
		)

	if err != nil {
		return
	}

	contenido :=
		utilidades.BytesAString(
			fileBlock.BContent[:],
		)

	contenido =
		strings.ReplaceAll(
			contenido,
			"\n",
			"<BR/>",
		)

	fmt.Fprintf(
		dot,
		`file%d [
shape=plaintext
label=<
<TABLE BORDER="1" CELLBORDER="1" CELLSPACING="0" BGCOLOR="khaki">
<TR><TD>Bloque Archivo %d</TD></TR>
<TR><TD>%s</TD></TR>
</TABLE>
>
];
`,
		numeroBloque,
		numeroBloque,
		contenido,
	)
}


func GenerarNodoInodoTREE(
	dot *strings.Builder,
	numero int32,
	inodo estructuras.Inode,
) {

	fmt.Fprintf(
		dot,
		`inode%d [
shape=plaintext
label=<
<TABLE BORDER="1" CELLBORDER="1" CELLSPACING="0" BGCOLOR="lightblue">
<TR>
<TD COLSPAN="2">
Inodo %d
</TD>
</TR>

<TR>
<TD>UID</TD>
<TD>%d</TD>
</TR>

<TR>
<TD>GID</TD>
<TD>%d</TD>
</TR>

<TR>
<TD>SIZE</TD>
<TD>%d</TD>
</TR>

<TR>
<TD>TYPE</TD>
<TD>%c</TD>
</TR>

<TR>
<TD>PERM</TD>
<TD>%d</TD>
</TR>
</TABLE>
>
];
`,
		numero,
		numero,
		inodo.IUid,
		inodo.IGid,
		inodo.ISize,
		inodo.IType,
		inodo.IPerm,
	)
}

func ConstruirArbol(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroInodo int32,
	nombre string,
) (TreeNode, error) {

	var nodo TreeNode

	nodo.Name = nombre

	posInodo := ObtenerPosicionInodo(
		sb,
		numeroInodo,
	)

	inode, err := LeerInodo(
		archivo,
		posInodo,
	)

	if err != nil {
		return TreeNode{}, err
	}

	// Determinar el tipo del inodo
	if inode.IType == '0' {

		nodo.Type = "folder"

	} else {

		nodo.Type = "file"

		// Los archivos son hojas del árbol
		return nodo, nil
	}

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for i := 0; i < 15; i++ {

		if inode.IBlock[i] == -1 {
			break
		}

		posBloque := sb.SBlockStart +
			(inode.IBlock[i] * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posBloque,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			if entrada.BInodo == -1 {
				continue
			}

			nombreHijo := utilidades.BytesAString(
				entrada.BName[:],
			)

			if nombreHijo == "." ||
				nombreHijo == ".." {
				continue
			}

			hijo, err := ConstruirArbol(
				archivo,
				sb,
				entrada.BInodo,
				nombreHijo,
			)

			if err != nil {
				continue
			}

			nodo.Children = append(
				nodo.Children,
				hijo,
			)
		}
	}

	return nodo, nil
}