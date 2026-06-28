/*
DiskList.jsx

Componente encargado de mostrar la lista de discos
disponibles y notificar cuál fue seleccionado.
*/

function DiskList({ disks, selectedDisk, onSelectDisk }) {

    if (disks.length === 0) {

        return (

            <p>No hay discos disponibles.</p>

        );

    }

    return (

        <div>

            <ul className="disk-list">

                {

                    disks.map((disk) => (

                        <li
                            key={disk.path}
                            className={
                                selectedDisk === disk.name
                                    ? "disk-item selected"
                                    : "disk-item"
                            }
                            onClick={() => onSelectDisk(disk.name)}
                        >

                            📀 {disk.name}

                        </li>

                    ))

                }

            </ul>

        </div>

    );

}

export default DiskList;