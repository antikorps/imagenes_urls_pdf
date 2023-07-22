package manejador

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ncruces/zenity"
)

func leerUrls(origen string) []string {
	var urls []string
	archivo, archivoError := os.Open(origen)
	if archivoError != nil {
		log.Fatalln("no se ha podido leer el archivo de origen", archivoError)
	}
	defer archivo.Close()
	escaner := bufio.NewScanner(archivo)
	escaner.Split(bufio.ScanLines)
	for escaner.Scan() {
		url := escaner.Text()
		if url == "" {
			continue
		}
		url = strings.TrimSpace(url)
		urls = append(urls, url)
	}

	return urls
}

func ConfigurarManejador(origen, destino string, simultaneidad, espera int) Manejador {
	var dialogos bool
	var err error

	if origen == "" || destino == "" {
		dialogos = true
		bienvenidaMensaje := `Crea un PDF de una lista de urls de imágenes.
Selecciona el archivo con las URLS y posteriormente otro archivo para el PDF.
Espera a que termine el diálogo de progreso...
... y disfruta del resultado :)`
		bienvenidaError := zenity.Info(bienvenidaMensaje,
			zenity.Title("Colección de imágenes a PDF"),
			zenity.Width(600),
			zenity.InfoIcon)
		if bienvenidaError != nil {
			log.Fatalln("no se ha podido lanzar el diálogo de bienvenida", bienvenidaError)
		}

		// Origen
		origen, err = zenity.SelectFile(
			zenity.Title("Selecciona el archivo con las imaǵenes"),
			zenity.FileFilters{
				{Name: "Archivos de texto", Patterns: []string{"*.txt", "*.csv"}, CaseFold: false},
			})
		if err != nil {
			if err == zenity.ErrCanceled {
				os.Exit(0)
			}
			log.Fatalln("no se ha podido lanzar el selector del archivo de origen", err)
		}
		// Destino
		destino, err = zenity.SelectFileSave(
			zenity.ConfirmOverwrite(),
			zenity.Title("Selecciona el pdf resultante"),
			zenity.FileFilters{
				{Name: "PDF", Patterns: []string{"*.pdf"}, CaseFold: false},
			})
		if err != nil {
			if err == zenity.ErrCanceled {
				os.Exit(0)
			}
			log.Fatalln("no se ha podido lanzar el selector del archivo para el pdf resultante", err)
		}
		if !strings.HasSuffix(destino, ".pdf") {
			destino += ".pdf"
		}
	}

	urls := leerUrls(origen)

	return Manejador{
		Origen:        origen,
		Destino:       destino,
		Simultaneidad: simultaneidad,
		Espera:        espera,
		Dialogos:      dialogos,
		Urls:          urls,
		Cliente: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
