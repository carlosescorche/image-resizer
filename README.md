## Redimensionador y Convertidor de Imágenes

Este programa te permite redimensionar imágenes a una lista de anchos especificados y convertirlas a formatos JPEG o WEBP. Es útil para preparar imágenes para diferentes tamaños de pantalla u optimizarlas para su uso en la web.

## Requisitos Previos
- **Go**: Es necesario tener instalado Go en tu máquina. Puedes descargarlo e instalarlo desde [golang.org](https://golang.org/dl/).
- **Librerías Externas**: El programa hace uso de bibliotecas externas para el manejo de imágenes. Asegúrate de tener instalados los siguientes paquetes de Go:
  - `github.com/chai2010/webp`
  - `golang.org/x/image/draw`

## Instalación
1. **Clonar el Repositorio**
    Clona el repositorio o descarga los archivos directamente en tu máquina local:
    ```
    git clone https://github.com/carlosescorche/image-resizer.git
    cd image-resizer
    ```

2. **Instalar dependencias**
    Ejecuta el siguiente comando para obtener las librerías necesarias de Go:
    ```
    go get -u github.com/chai2010/webp
    go get -u golang.org/x/image/draw
    ```

## Uso
Para usar este programa, sigue los pasos que se indican a continuación:

1. **Preparar tus imágenes** 
- Asegúrate de que las imágenes que deseas redimensionar sean accesibles y estén soportadas por el programa (JPEG, PNG y WEBP son soportados para la entrada).

2. **Ejecutar el programa** 
- Utiliza el siguiente comando para ejecutar el programa. Reemplaza los marcadores de posición con valores reales:

    ```
    go run main.go -source=ruta/a/imagen-origen.jpg -target=ruta/a/destino -widths=1080,720,320 -format=webp
    ```
    - **source:** Ruta al archivo de imagen fuente.
    - **target:** Ruta base para los archivos de salida. Los nombres de archivo se complementarán con el ancho y el formato.
    - **widths:** Lista separada por comas de los anchos a los cuales quieres redimensionar la imagen.
    - **format:** Formato de salida de la imagen (ya sea 'webp' o 'jpg').

3. **Verificar la salida**
Las imágenes redimensionadas se guardarán en el directorio de destino especificado con nombres que indican su tamaño y formato.

## Ejemplo
```
go run main.go -source=./img/muestra.png -target=./salida/muestra -widths=800,500,300 -format=jpg
```

Este comando redimensionará la imagen muestra.png a anchos de 800, 500 y 300 píxeles y las guardará en formato JPEG en el directorio ./salida.