package manejador

import (
	"bytes"
	"io"
	"log"
	"os"
	"sort"

	"github.com/ncruces/zenity"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func (m *Manejador) Convertir() {
	archivoPDF, archivoPDFError := os.Create(m.Destino)
	if archivoPDFError != nil {
		if m.Dialogos {
			mensajeError := "no se ha podido crear el PDF " + archivoPDFError.Error()
			zenity.Info(mensajeError,
				zenity.Title("ERROR CRÍTICO"),
				zenity.Width(600),
				zenity.ErrorIcon)
		}
		log.Fatalln("no se ha podido crear el PDF:", archivoPDFError)
	}
	defer archivoPDF.Close()

	sort.Slice(m.Imagenes, func(i, j int) bool {
		return m.Imagenes[i].Orden < m.Imagenes[j].Orden
	})

	var imagenes []io.Reader
	for _, v := range m.Imagenes {
		r := bytes.NewReader(v.Bytes)
		imagenes = append(imagenes, r)
	}
	pdfError := api.ImportImages(nil, archivoPDF, imagenes, nil, nil)
	if pdfError != nil {
		if m.Dialogos {
			mensajeError := "no se ha podido crear el PDF a partir de las imágenes" + pdfError.Error()
			zenity.Info(mensajeError,
				zenity.Title("ERROR CRÍTICO"),
				zenity.Width(600),
				zenity.ErrorIcon)
		}
		log.Fatalln("no se ha podido crear el PDF a partir de las imágenes:", pdfError)
	}

}
