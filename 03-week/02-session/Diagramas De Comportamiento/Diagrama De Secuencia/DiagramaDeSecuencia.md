# Diagrama de Secuencia (UML)

## ¿Qué es?

El diagrama de secuencia es un diagrama de comportamiento en UML que muestra la interacción entre objetos o componentes en el tiempo, mediante el envío de mensajes.  
Se centra en el orden cronológico de los mensajes, reflejando cómo colaboran los objetos para cumplir una funcionalidad específica.

Es uno de los diagramas más usados en el diseño, ya que une casos de uso con la lógica de implementación.

---

## Estructura

Un diagrama de secuencia está compuesto por:

- **Objetos / Participantes** → Elementos que interactúan (dibujados como rectángulos en la parte superior).  
- **Líneas de vida (lifelines)** → Líneas verticales bajo los objetos, que representan la existencia de cada uno a lo largo del tiempo.  
- **Mensajes** → Flechas horizontales que indican la comunicación entre los objetos.  
- **Activaciones** → Barras delgadas sobre las líneas de vida que muestran cuándo un objeto está ejecutando una acción.  
- **Condiciones y bucles** → Notaciones especiales que indican decisiones o repeticiones.  

---

## Principales características

- Representa interacciones dinámicas en orden temporal.  
- Muestra qué mensajes se envían, en qué orden y entre qué objetos.  
- Se lee de arriba hacia abajo (el tiempo fluye verticalmente).  
- Es muy útil para detallar escenarios de un caso de uso.  
- Permite identificar responsabilidades entre clases/objetos.  

---

## Elementos principales

- **Actor / Objeto / Componente** → Los participantes del escenario.  
  Ejemplo: Usuario, Sistema, Base de Datos.  

- **Línea de vida (lifeline)** → Línea punteada que baja desde cada participante.  

- **Mensaje** → Flecha horizontal que representa una invocación, respuesta o comunicación.  
  - **Síncrono (→):** el emisor espera la respuesta.  
  - **Asíncrono (↦):** el emisor no espera respuesta.  
  - **Retorno (↩):** respuesta devuelta.  

- **Bloques de control:**  
  - **alt** → Alternativa (condicional).  
  - **opt** → Opción (ejecución opcional).  
  - **loop** → Bucle o repetición.  
  - **par** → Procesos paralelos.  

---

## Conexiones

### 1. Mensaje síncrono (→)
- Línea sólida con flecha llena.  
- El emisor espera respuesta.  
- **Ejemplo:**  
  `Usuario → Sistema : iniciarSesión()`

### 2. Mensaje asíncrono (↦)
- Línea sólida con flecha abierta.  
- El emisor no espera respuesta inmediata.  
- **Ejemplo:**  
  `Sistema ↦ ServicioCorreo : enviarNotificación()`

### 3. Mensaje de retorno (↩)
- Línea punteada con flecha abierta.  
- Representa la respuesta a una invocación.  
- **Ejemplo:**  
  `BaseDatos ↩ Sistema : resultadoConsulta`

### 4. Bloque condicional (alt)
- Dos o más fragmentos, separados por condiciones.  
- Muestra un if / else.  
- **Ejemplo:**  
  - `alt [login válido] → mostrar menú`  
  - `alt [login inválido] → mostrar error`

### 5. Bucle (loop)
- Fragmento con la etiqueta loop.  
- Indica repetición.  
- **Ejemplo:**  
  `loop [mientras haya ítems] → procesar ítem`

### 6. Opción (opt)
- Fragmento opcional que se ejecuta solo si la condición se cumple.  
- **Ejemplo:**  
  `opt [cliente es VIP] → aplicar descuento`

---

## Ventajas

- Representa de manera clara el flujo temporal de interacciones.  
- Ayuda a validar escenarios de casos de uso.  
- Facilita la traducción a código, ya que muestra el orden de invocaciones.  
- Identifica responsabilidades entre objetos.  
- Útil para detallar protocolos de comunicación entre sistemas.  

---

## Desventajas

- Puede volverse muy grande y difícil de leer en procesos complejos.  
- Se enfoca solo en un escenario concreto (hay que hacer varios para cubrir todos los casos).  
- Requiere tiempo de elaboración para ser detallado y entendible.  
- No refleja la estructura interna, solo la interacción.  

---

## Conclusión

El diagrama de secuencia es una herramienta poderosa para describir cómo colaboran los objetos en un escenario, mostrando el orden exacto de los mensajes en el tiempo.  
Es fundamental para conectar requisitos funcionales (casos de uso) con el diseño técnico (clases, objetos y métodos).  
Aunque puede complicarse en sistemas grandes, su uso es clave para documentar, validar y comunicar la lógica dinámica del sistema.
