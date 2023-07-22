package manejador

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func verificarExtension(extension string) bool {
	extensionesValidas := []string{"png", "jpg", "jpeg", "webp"}
	for _, v := range extensionesValidas {
		if v == extension {
			return true
		}
	}
	return false
}

func descarga(ruta string, orden int, cliente http.Client, canalDescargas chan InfoImagen) {
	defer wg.Done()
	peticion, peticionError := http.NewRequest("GET", ruta, nil)
	if peticionError != nil {
		canalDescargas <- InfoImagen{
			Orden: orden,
			Url:   ruta,
			Error: peticionError,
		}
		return
	}
	respuesta, respuestaError := cliente.Do(peticion)
	if respuestaError != nil {
		canalDescargas <- InfoImagen{
			Orden: orden,
			Url:   ruta,
			Error: respuestaError,
		}
		return
	}
	if respuesta.StatusCode != 200 {
		canalDescargas <- InfoImagen{
			Orden: orden,
			Url:   ruta,
			Error: errors.New("status code incorrecto" + respuesta.Status),
		}
		return
	}
	defer respuesta.Body.Close()

	var contentType string
	var extension string
	cabeceras := respuesta.Header
	contentType = cabeceras.Get("Content-Type")
	if contentType == "" {
		contentType = cabeceras.Get("content-type")
		if contentType == "" {
			for c, v := range cabeceras {
				if strings.Contains(strings.ToLower(c), "ontent") && strings.Contains(strings.ToLower(c), "ype") {
					contentType = v[0]
				}
			}
		}
	}
	infoContentType := strings.Split(contentType, "/")
	if len(infoContentType) > 0 {
		extension = strings.ToLower(infoContentType[1])
	}

	if !verificarExtension(extension) {
		canalDescargas <- InfoImagen{
			Orden: orden,
			Url:   ruta,
			Error: errors.New("no se ha podido verificar o no es válida la extensión de la respuesta: " + extension),
		}
		return
	}

	bytes, bytesError := ioutil.ReadAll(respuesta.Body)
	if bytesError != nil {
		canalDescargas <- InfoImagen{
			Orden: orden,
			Url:   ruta,
			Error: errors.New("no se ha podido obtenir los bytes del cuerpo de la respuesta: " + bytesError.Error()),
		}
		return
	}

	canalDescargas <- InfoImagen{
		Bytes:     bytes,
		Error:     nil,
		Extension: extension,
		Orden:     orden,
		Url:       ruta,
	}
}

func (m *Manejador) Descargar() {
	var infoImagen []InfoImagen
	imagenesDescargadas := 0
	continuar := true
	for continuar {
		canalDescargas := make(chan InfoImagen)
		if m.Simultaneidad > len(m.Urls) {
			m.Simultaneidad = len(m.Urls)
		}
		for i := 0; i < m.Simultaneidad; i++ {
			wg.Add(1)
			go descarga(m.Urls[imagenesDescargadas], imagenesDescargadas, m.Cliente, canalDescargas)
			imagenesDescargadas++
			if imagenesDescargadas == len(m.Urls) {
				continuar = false
				break
			}
		}
		go func() {
			wg.Wait()
			close(canalDescargas)
		}()
		time.Sleep(time.Duration(m.Espera) * time.Second)
		for c := range canalDescargas {
			if c.Error != nil {
				mensajeError := fmt.Sprintf("error en la url %v: %v. Línea %d en el archivo", c.Url, c.Error.Error(), c.Orden+1)
				m.Errores = append(m.Errores, mensajeError)
				continue
			}
			infoImagen = append(infoImagen, c)
		}
	}
	m.Imagenes = infoImagen
}
