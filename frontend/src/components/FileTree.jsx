/*
FileTree.jsx

Visualiza el árbol del sistema de archivos.

Actualmente utiliza datos de prueba.
En el siguiente paso recibirá la información
desde el Backend REST.
*/

import { useState } from "react";

import "./FileTree.css";

/*
Nodo recursivo del árbol.
*/
function TreeNode({ node }) {

    const [expanded, setExpanded] = useState(true);

    const isFolder = node.type === "folder";

    return (

        <div className="tree-node">

            <div
                className="tree-item"
                onClick={() => {

                    if (isFolder) {

                        setExpanded(!expanded);

                    }

                }}
            >

                {

                    isFolder

                        ? (expanded ? "📂" : "📁")

                        : "📄"

                }

                {" "}

                {node.name}

            </div>

            {

                isFolder &&
                expanded &&
                node.children &&
                node.children.length > 0 &&

                (

                    <div className="tree-children">

                        {

                            node.children.map(

                                (child, index) => (

                                    <TreeNode

                                        key={index}

                                        node={child}

                                    />

                                )

                            )

                        }

                    </div>

                )

            }

        </div>

    );

}

/*
Componente principal.
*/
function FileTree() {

    /*
    Datos temporales.

    Serán reemplazados por la respuesta
    del endpoint GET /api/filesystem.
    */

    const tree = {

        name: "/",

        type: "folder",

        children: [

            {

                name: "users.txt",

                type: "file"

            },

            {

                name: "home",

                type: "folder",

                children: [

                    {

                        name: "user",

                        type: "folder",

                        children: [

                            {

                                name: "archivo.txt",

                                type: "file"

                            }

                        ]

                    }

                ]

            },

            {

                name: "etc",

                type: "folder",

                children: []

            }

        ]

    };

    return (

        <TreeNode

            node={tree}

        />

    );

}

export default FileTree;