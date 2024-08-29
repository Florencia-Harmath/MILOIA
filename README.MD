# MiloIA - Chat en Tiempo Real

MiloIA es una aplicación de chat en tiempo real que permite a los usuarios conectarse y comunicarse a través de WebSockets. Utiliza Go para la lógica del servidor, PostgreSQL para el almacenamiento de datos y Redis para la gestión de mensajes y la conexión en tiempo real.

## Tabla de Contenidos

- [Características](#características)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Uso](#uso)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

## Características

- Conexión en tiempo real mediante WebSockets.
- Autenticación basada en JWT.
- Almacenamiento de mensajes en PostgreSQL.
- Gestión de sesiones y usuarios con Redis.
- Migración automática de la base de datos.

## Requisitos

- Go 1.18 o superior
- PostgreSQL
- Redis
- [Viper](https://github.com/spf13/viper) para la configuración

## Instalación

Sigue estos pasos para instalar y ejecutar la aplicación en tu máquina local:

1. **Clona el repositorio:**

    ```bash
    git clone https://github.com/tuusuario/miloia.git
    cd miloia
    ```

2. **Instala las dependencias:**

    Asegúrate de tener Go instalado. Luego, instala las dependencias del proyecto:

    ```bash
    go mod tidy
    ```

3. **Configura la base de datos y Redis:**

    Asegúrate de tener PostgreSQL y Redis instalados y funcionando. Crea una base de datos en PostgreSQL y asegúrate de que Redis esté activo.

4. **Crea el archivo `.env`:**

    Crea un archivo `.env` en la raíz del proyecto con la siguiente configuración:

    ```env
    DB_HOST
    DB_PORT
    DB_NAME
    REDIS_ADDR
    JWT_SECRET
    ```

5. **Compila y ejecuta la aplicación:**

    ```bash
    go run main.go
    ```

## Configuración

La configuración se carga desde un archivo `.env` utilizando la biblioteca [Viper](https://github.com/spf13/viper). Asegúrate de definir las siguientes variables de entorno:

- `DB_HOST`: Dirección del servidor PostgreSQL.
- `DB_PORT`: Puerto del servidor PostgreSQL.
- `DB_NAME`: Nombre de la base de datos PostgreSQL.
- `REDIS_ADDR`: Dirección del servidor Redis.
- `JWT_SECRET`: Clave secreta para la firma de tokens JWT.

## Uso

Una vez que la aplicación esté en ejecución, el servidor estará disponible en `http://localhost:8080`. Puedes conectar a la aplicación usando WebSockets. Asegúrate de enviar el token JWT en el encabezado `Authorization` para autenticarte.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura:


- **`main.go`**: Archivo principal que inicia la aplicación.
- **`internal/config/config.go`**: Maneja la carga de la configuración.
- **`internal/database/database.go`**: Conecta a PostgreSQL y Redis y realiza migraciones.
- **`internal/chat/hub.go`**: Implementa la lógica del hub de WebSocket.
- **`internal/handlers/chat_handlers.go`**: Maneja las conexiones WebSocket.
- **`pkg/auth/auth.go`**: Proporciona funciones para la autenticación JWT.
- **`internal/models/user.go`**: Define el modelo de usuario.
- **`internal/models/message.go`**: Define el modelo de mensaje.

## Contribuciones

¡Las contribuciones son bienvenidas! Si quieres contribuir al proyecto, por favor sigue estos pasos:

1. Forkea el repositorio.
2. Crea una nueva rama para tu feature o bugfix (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y asegúrate de que el código esté bien probado.
4. Envía un pull request detallando los cambios que realizaste.

## Licencia

Este proyecto está licenciado bajo la [Licencia MIT](LICENSE).

