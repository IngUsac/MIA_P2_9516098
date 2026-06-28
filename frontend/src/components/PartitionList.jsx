/*
PartitionList.jsx

Muestra las particiones del disco seleccionado.
Permite seleccionar una partición.
*/

function PartitionList({

    partitions,

    selectedPartition,

    onSelectPartition

}) {

    if (!partitions || partitions.length === 0) {

        return (

            <p>No hay particiones.</p>

        );

    }

    return (

        <div>

            <ul className="partition-list">

                {

                    partitions.map((partition, index) => (

                        <li

                            key={index}

                            className={

                                selectedPartition === partition.name

                                    ? "partition-item selected"

                                    : "partition-item"

                            }

                            onClick={() =>

                                onSelectPartition(partition)

                            }

                        >

                            💽 {partition.name}

                            {" "}

                            <small>

                                ({partition.typeName})

                            </small>

                        </li>

                    ))

                }

            </ul>

        </div>

    );

}

export default PartitionList;