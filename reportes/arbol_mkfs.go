package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Generando gráfica MKFS...")

	err := GenerarGraficaMKFS()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Gráfica generada correctamente.")
}

func GenerarGraficaMKFS() error {

	fechaHora := time.Now().Format("02/01/2006 15:04:05")

	dot := fmt.Sprintf(`
digraph MKFS {

    graph [
        pad="0.5",
        nodesep="0.6",
        ranksep="1.0"
    ];

    rankdir=TB;

    node [
        shape=record,
        style=filled,
        fillcolor=lightyellow
    ];

    titulo [
        fillcolor=lightblue,
        label="{REPORTE MKFS|Fecha: %s}"
    ];

    superblock [
        fillcolor=lightgreen,
        label="{SUPERBLOCK|Filesystem EXT2}"
    ];

    inode0 [
        label="{INODO 0|Directorio Raíz /}"
    ];

    block0 [
        fillcolor=white,
        label="{BLOQUE 0|.|..|users.txt}"
    ];

    inode1 [
        label="{INODO 1|Archivo users.txt}"
    ];

    block1 [
        fillcolor=white,
        label="{BLOQUE 1|1,G,root\\n1,U,root,root,123}"
    ];

    titulo -> superblock;
    superblock -> inode0;
    superblock -> inode1;

    inode0 -> block0;
    block0 -> inode1;

    inode1 -> block1;
}
`, fechaHora)

	dotFile := "arbol_mkfs.dot"

	err := os.WriteFile(
		dotFile,
		[]byte(dot),
		0644,
	)

	if err != nil {
		return err
	}

	pngFile := "grafica.png"

	cmd := exec.Command(
		"dot",
		"-Tpng",
		dotFile,
		"-o",
		pngFile,
	)

	err = cmd.Run()

	if err != nil {
		return fmt.Errorf(
			"no se pudo ejecutar Graphviz. Verifica que 'dot' esté instalado y en el PATH: %v",
			err,
		)
	}

	AbrirArchivo(pngFile)

	return nil
}

func AbrirArchivo(ruta string) {

	switch runtime.GOOS {

	case "windows":

		exec.Command(
			"rundll32",
			"url.dll,FileProtocolHandler",
			ruta,
		).Start()

	case "linux":

		exec.Command(
			"xdg-open",
			ruta,
		).Start()

	case "darwin":

		exec.Command(
			"open",
			ruta,
		).Start()
	}
}