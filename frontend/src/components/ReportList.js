function ReportList({

    reports,

    selected,

    onSelect,

}) {

    if (!reports || reports.length === 0) {

        return <p>No hay reportes.</p>;

    }

    return (

        <ul
            style={{
                listStyle:"none",
                padding:0,
                margin:0
            }}
        >

            {

                reports.map((report)=>(

                    <li
                        key={report.name}
                        onClick={()=>onSelect(report)}
                        style={{
                            padding:"8px",
                            cursor:"pointer",
                            background:
                                selected===report.name
                                ? "#dbeafe"
                                : "transparent"
                        }}
                    >

                        {report.name}

                    </li>

                ))

            }

        </ul>

    );

}

export default ReportList;