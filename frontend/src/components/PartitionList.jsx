/*
PartitionList.jsx

Muestra las particiones del disco seleccionado.
*/

function PartitionList({ partitions }) {

    if (!partitions || partitions.length === 0) {

        return (
            <p>No hay particiones.</p>
        );

    }

    return (

        <div>

            <h2>Particiones</h2>

            <ul className="partition-list">

                {

                    partitions.map((partition, index) => (

                        <li
                            key={index}
                            className="partition-item"
                        >

                            💽 {partition.name}

                            {"  "}

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