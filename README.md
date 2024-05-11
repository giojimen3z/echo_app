# Aplicación Ping-Pong en Go

## Descripción
Esta es una aplicación simple escrita en Go que utiliza el framework Echo para implementar servicios web. El propósito de esta aplicación es demostrar un endpoint simple que recibe peticiones POST y responde de acuerdo con la entrada especificada.

## Tecnologías
- **Lenguaje de Programación:** Go
- **Framework Web:** Echo v4

## Endpoint Disponible

- **POST /myapp/ping**
    - **Descripción:** Este endpoint recibe peticiones POST y simplemente responde con un "pong".
    - **Respuestas:**
        - `200 OK`: Devuelve "pong" si la solicitud es correcta.
        - `500 Internal Server Error`: Devuelve un error de servidor si ocurre cualquier otro caso.

## Cómo correr la aplicación localmente

### Prerrequisitos
Para ejecutar esta aplicación necesitas tener instalado Go en tu sistema. Puedes descargarlo e instalarlo desde [el sitio oficial de Go](https://golang.org/dl/).

### Instalación de dependencias
Antes de ejecutar la aplicación, necesitas instalar las dependencias del proyecto, lo cual puedes hacer ejecutando el siguiente comando en la raíz del proyecto:


### Ejecutar la aplicación
Para iniciar la aplicación, ejecuta el siguiente comando en la raíz del proyecto:


La aplicación estará disponible en `http://localhost:8080`.

### Pruebas
Puedes probar el endpoint utilizando `curl` con el siguiente comando:


Esperas recibir una respuesta de `pong`.

## Despliegue
Para instrucciones sobre cómo desplegar esta aplicación en un ambiente de producción, por favor revisa la sección sobre [despliegue en AWS](#) en nuestro documento de despliegue. (Asegúrate de agregar un enlace real o instrucciones de despliegue si es necesario).

## Contribuciones
Las contribuciones son bienvenidas. Si deseas contribuir al proyecto, por favor envía un pull request o abre un issue para discutir los cambios propuestos.

## Licencia
Este proyecto está licenciado bajo la Licencia MIT. Para más detalles, consulta el archivo `LICENSE` incluido en el repositorio.

