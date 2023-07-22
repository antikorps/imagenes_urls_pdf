package manejador

import (
	"log"
	"os"
	"strings"

	"github.com/ncruces/zenity"
)

func (m *Manejador) Informar() {
	if m.Dialogos {
		if len(m.Errores) == 0 {
			zenity.Info("PDF creado con éxito",
				zenity.Title("FIN"),
				zenity.InfoIcon)
			os.Exit(0)
		}
		advertenciaErrores := "ATENCIÓN: pdf creado, pero con errores\n" + strings.Join(m.Errores, "\n")
		zenity.Info(advertenciaErrores,
			zenity.Title("FIN"),
			zenity.Width(600),
			zenity.InfoIcon)
		os.Exit(0)
	}
	if len(m.Errores) > 0 {
		for _, v := range m.Errores {
			log.Println(v)
		}
		os.Exit(1)
	}
	os.Exit(0)
}
