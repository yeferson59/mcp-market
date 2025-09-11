# Simple MCP Market Data

¡Bienvenido a **Simple MCP Market Data**!
Este proyecto es una implementación sencilla en Go para consultar información de acciones bursátiles utilizando el protocolo MCP y la API de Alpha Vantage.

---

## 🚀 Características

- Consulta datos detallados de acciones por símbolo.
- Obtén información financiera, sector, industria, capitalización de mercado, dividendos y mucho más.
- Integración con MCP para facilitar la interoperabilidad.
- Configuración sencilla mediante variables de entorno.

---

## 📦 Instalación

1. **Clona el repositorio:**

```bash
git clone https://github.com/tu-usuario/simple-mcp.git
cd simple-mcp
```

2. **Instala las dependencias:**

```bash
go mod tidy
```

3. **Configura tu API Key de Alpha Vantage:**
   Crea un archivo `.env` en la raíz con el siguiente contenido:
   ```
   API_URL=https://www.alphavantage.co
   API_KEY=tu_api_key
   ```

---

## 🛠️ Uso

Ejecuta el servidor MCP:

```bash
go run main.go
```

Luego, puedes realizar consultas de acciones enviando el símbolo correspondiente.

---

## 📁 Estructura del proyecto

- `main.go`: Código principal del servidor y lógica de consulta.
- `.env`: Variables de entorno para la configuración de la API.
- `go.mod` / `go.sum`: Dependencias del proyecto.

---

## 🤝 Contribuciones

¡Las contribuciones y sugerencias son bienvenidas!
Si deseas mejorar el proyecto, por favor abre un issue o envía un pull request.

---

## 📝 Créditos

Desarrollado por Yeferson Toloza Contreras.
Utiliza [Alpha Vantage](https://www.alphavantage.co/) y [Model Context Protocol](https://github.com/modelcontextprotocol/go-sdk).
