# Módulo 3 - Submódulo 3
## NLP convierte señales en significado

### Introducción
En este submódulo aprendí cómo las aplicaciones de **gestión de información no estructurada (UIMA)** descomponen preguntas humanas, generan posibles respuestas y seleccionan la más probable. Esto permite que los chatbots comprendan, razonen, aprendan e interactúen con los usuarios de manera efectiva.

### Estructura de un chatbot
Aprendí que los chatbots tienen dos componentes principales:

- **Frontend:** canal de mensajería que interactúa con el usuario, escuchando y respondiendo en tiempo real.  
- **Backend:** realiza la lógica y procesamiento de datos, recordando partes de la conversación y activando respuestas adecuadas.

El backend trabaja para interpretar preguntas complejas o mal redactadas mediante **clasificadores**, que relacionan múltiples formas de preguntar con un conjunto reducido de respuestas predefinidas.

### Intenciones, entidades y diálogos
El backend de un chatbot suele estar compuesto por tres elementos clave:

- **Intención:** propósito de la interacción del usuario, como consultar horarios, presentar quejas o hacer pedidos.  
- **Entidad:** sustantivo relevante dentro de la intención, como un lugar, persona u objeto.  
- **Diálogo:** estructura de flujo (tipo IF/THEN) que define cómo el chatbot responde a las preguntas y sigue la conversación.

Por ejemplo, para un chatbot de una cadena de restaurantes:

| Intención | Posibles entradas de usuario | Entidades |
|-----------|----------------------------|-----------|
| Abierto   | ¿Cuándo abren?             | Austin    |
|           | ¿Qué horario hacen?        | Planificación |
|           | ¿Ahora están abiertos?     | Hora      |
|           | ¿Hasta qué hora están abiertos? | Hora |

El diálogo permite que el chatbot responda correctamente aunque el usuario utilice frases mal redactadas o con errores tipográficos.

### Ejemplo de análisis NLP
Aprendí que el NLP, combinado con intenciones, entidades y diálogos, permite que la IA clasifique componentes de las preguntas y proporcione la respuesta más confiable. Por ejemplo, en el caso del **chatbot de Staples**, cuando un cliente dice: *“Quiero hacer otro pedido de bolígrafos negros”*, el sistema interpreta la intención, identifica la entidad, consulta el historial de compras y responde con la acción más adecuada en segundos.

### Reflexión personal
Este submódulo me enseñó cómo NLP convierte señales en significado y permite que los chatbots actúen de manera inteligente. Comprendí que los chatbots son útiles no solo por su disponibilidad constante, sino también por su capacidad para interpretar correctamente las preguntas humanas y ofrecer respuestas rápidas y confiables.
