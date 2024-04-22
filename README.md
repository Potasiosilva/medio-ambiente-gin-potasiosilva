# Proyecto de Reciclaje y Medio Ambiente

Este proyecto es una aplicación web simple que promueve la conciencia sobre el reciclaje y el cuidado del medio ambiente. Utiliza el framework web Gin para manejar las solicitudes HTTP y GORM para interactuar con una base de datos SQLite. El objetivo principal es permitir a los usuarios enviar sus ideas sobre cómo contribuir al reciclaje y al cuidado del medio ambiente.

## Características

- **Envío de Ideas**: Los usuarios pueden enviar sus ideas sobre cómo contribuir al reciclaje y al cuidado del medio ambiente.
- **Listado de Ideas**: Se muestra un listado de todas las ideas enviadas por los usuarios.
- **Validación de Usuarios**: Se verifica que cada usuario no envíe más de una idea para evitar duplicados.
- **Interfaz de Usuario**: Una interfaz de usuario simple y amigable para enviar ideas y ver las existentes.

## Requisitos

- Go 1.16 o superior
- Gin 1.7.4 o superior
- GORM 1.21.9 o superior

## Instalación

1. Clona el repositorio en tu máquina local.
2. Asegúrate de tener Go instalado y configurado correctamente.
3. Instala las dependencias ejecutando `go mod download`.
4. Ejecuta `go run .` para iniciar el servidor.

## Uso

- **Iniciar el servidor**: Ejecuta `go run .` en la terminal desde el directorio del proyecto.
- **Acceder a la aplicación**: Abre un navegador y navega a http://localhost:8080.
- **Enviar una idea**: Rellena el formulario en la página principal con tu nombre de usuario y tu idea, luego haz clic en "Enviar".
- **Ver ideas enviadas**: Navega a http://localhost:8080/ideas para ver todas las ideas enviadas por los usuarios.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request si tienes alguna sugerencia o deseas contribuir al proyecto.

## Licencia

Este proyecto está licenciado bajo la licencia MIT. Consulta el archivo LICENSE para más detalles.