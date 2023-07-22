# Descripción
Crea un PDF a partir de un archivo con rutas de imágenes. Solo es necesario descargar el binario para el sistema operativo correspondiente y descargar, no es necesaria ninguna instalación adicional.\
Estos binarios pueden descargarse de la carpeta bin, donde solo habrá que seleccionar el sistema operativo y arquitectura. Actualmente se han compilado versiones para:
- GNU/Linux amd64
- Windows amd64

# Instrucciones
El programa puede ejercutarse por línea de comandos o a través de una interfaz de diálogos
## Interfaz de diálogos
Simplemente debe ejecutarse el binario. Aparecerá una ventana de información, un selector de archivos para seleccionar el archivo de origen con las URL (una dirección por línea), posteriormente seleccionar el archivo PDF resultante y esperar a que se termine el diálogo de progreso una vez finalizadas las descargas.\
Mensaje de inicio:\
![Diálogos](https://i.imgur.com/1x9a5gA.png)\
Selector de archivo de origen:\
![Diálogos](https://i.imgur.com/yzwyU3m.png.jpg)\
Selector para generar PDF:\
![Diálogos](https://i.imgur.com/3PuXkbT.png.jpg)\
Progreso:\
![Diálogos](https://i.imgur.com/BUZwg4m.png.jpg)\
Fin:\
![Diálogos](https://i.imgur.com/tJzJ40P.png.jpg)\

### CLI
La ejecución por línea de comandos permite modificar ajustes opcionales como la simultaneidad del número de descargas y el tiempo en segundos que debe esperarse entre cada lote de descargas simultáneas. Estos ajustes pueden incorporarse junto a los obligatorios: la ruta completa del archivo con las urls y la ruta completa del PDF que se quiere generar. 
Ejemplo de ejecución en GNU/Linux:
```bash
/home/usuario/descargas/imagenes_urls_pdf -ruta /home/usuario/descargas/urls.txt -destino /home/usuario/descargas/revista.pdf -simultaneidad 5 -espera 2
```
