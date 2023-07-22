package main

import (
	"flag"
	"imagenes_urls_pdf/manejador"

	"github.com/ncruces/zenity"
)

func main() {
	var origen, destino string
	var simultaneidad, espera int

	flag.StringVar(&origen, "origen", "", "ruta completa del archivo con las url de las imágenes")
	flag.StringVar(&destino, "destino", "", "ruta completa del archivo pdf resultante")
	flag.IntVar(&simultaneidad, "simultaneidad", 1, "número de descargas simultáneas")
	flag.IntVar(&espera, "espera", 0, "segundos de espera entre el lote de descargas")
	flag.Parse()

	manejador := manejador.ConfigurarManejador(origen, destino, simultaneidad, espera)
	var progreso zenity.ProgressDialog
	if manejador.Dialogos {
		progreso, _ = zenity.Progress(
			zenity.Title("Descargando"),
			zenity.Pulsate(),
			zenity.NoCancel())
	}
	manejador.Descargar()
	manejador.Convertir()
	if manejador.Dialogos {
		progreso.Complete()
		progreso.Close()
	}
	manejador.Informar()
}
