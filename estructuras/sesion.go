package estructuras

type Sesion struct {
	Activa bool

	User string
    Pass string

	ID string

	UID int32
	GID int32
}

var SesionActual Sesion

