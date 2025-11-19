# Diagrama de Estados (UML)

## ¿Qué es?

El diagrama de estados es un diagrama de comportamiento en UML que muestra los diferentes estados por los que pasa un objeto a lo largo de su ciclo de vida, así como los eventos o transiciones que provocan esos cambios.

Se utiliza principalmente para modelar objetos dinámicos, es decir, aquellos cuyo comportamiento depende de eventos externos o internos (ejemplo: un pedido, una sesión de usuario, una máquina expendedora).

---

## Estructura

Un diagrama de estados incluye:

- **Estado** → Situación o condición en la que se encuentra un objeto (se representa como un rectángulo redondeado).
- **Transición** → Flecha que conecta estados y que ocurre cuando sucede un evento.
- **Evento** → Acción que dispara una transición (ejemplo: “clic en botón”, “tiempo expirado”).
- **Acción** → Respuesta que ocurre como parte de la transición o dentro de un estado.
- **Estado inicial** → Punto de inicio del ciclo de vida (círculo sólido).
- **Estado final** → Punto donde termina el ciclo de vida del objeto (círculo con borde y un punto negro dentro).
- **Estados compuestos** → Estados que contienen subestados internos.
- **Historial** → Indica que al volver a un estado compuesto, el objeto retoma el último subestado en el que estaba.

---

## Principales características

- Representa el ciclo de vida de un objeto.
- Permite visualizar eventos, transiciones y respuestas del sistema.
- Se centra en un único objeto o entidad, no en todo el proceso general.
- Modela estados internos, no flujos de trabajo (a diferencia del diagrama de actividades).
- Útil para sistemas reactivos (que responden a eventos externos).

---

## Elementos principales

- **Estado** → Representa una condición estable del objeto.  
  Ejemplo: *Pendiente, Pagado, Enviado.*

- **Estado inicial** → Círculo sólido que marca el comienzo del ciclo.  
  Ejemplo: *al crear un pedido.*

- **Estado final** → Círculo con borde y centro negro que marca el fin del ciclo.

- **Transición** → Flecha que conecta estados, etiquetada con:  
  `Evento [condición] / acción`.

- **Estado compuesto** → Un estado que se divide en subestados más detallados.

- **Historial (H)** → Marca para retomar el último estado previo dentro de un estado compuesto.

---

## Conexiones

### 1. Flujo básico
El objeto pasa de un estado a otro por un evento.  
**Ejemplo:**  
Pendiente --(Pago realizado)--> Pagado


### 2. Transición con condición y acción
Incluye condición y acción en la flecha.  
**Ejemplo:**  
Activo --(Tiempo > 30min) / Cerrar sesión--> Inactivo


### 3. Estado compuesto
Un estado general se divide en subestados.  
**Ejemplo:**  
Procesando Pedido contiene: Verificando stock, Confirmando pago


### 4. Historial
El sistema recuerda en qué subestado estaba antes de salir.  
**Ejemplo:**  
Un reproductor de música vuelve al punto donde se detuvo (H en el estado Reproducción).

---

## Ventajas

- Muestra claramente el ciclo de vida de un objeto.
- Útil para sistemas donde el comportamiento depende de eventos.
- Facilita el diseño de máquinas de estados finitos.
- Ayuda a detectar estados redundantes o inexistentes.
- Ideal para modelar procesos controlados por eventos.

---

## Desventajas

- Se centra en un solo objeto, no en todo el sistema.
- Puede volverse muy complejo si el objeto tiene muchos estados.
- No muestra responsabilidades de actores externos (para eso sirven otros diagramas).
- Menos intuitivo para usuarios no técnicos comparado con un diagrama de actividades.

---

## Conclusión

El diagrama de estados es clave para modelar el comportamiento dinámico de un objeto dentro de un sistema.  
Permite ver cómo cambia un objeto con el tiempo según los eventos que ocurren, y ayuda a diseñar procesos claros, evitando estados muertos o transiciones incorrectas.

Aunque puede ser complejo en sistemas grandes, es fundamental en el diseño de sistemas reactivos y orientados a eventos.
