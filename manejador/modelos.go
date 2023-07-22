package manejador

import (
	"net/http"
)

type InfoImagen struct {
	Bytes     []byte
	Error     error
	Extension string
	Orden     int
	Url       string
}

type Manejador struct {
	Origen        string
	Destino       string
	Espera        int
	Simultaneidad int
	Dialogos      bool
	Urls          []string
	Imagenes      []InfoImagen
	Errores       []string

	Cliente http.Client
}
