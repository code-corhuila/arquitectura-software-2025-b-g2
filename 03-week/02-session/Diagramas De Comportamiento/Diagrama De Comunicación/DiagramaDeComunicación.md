# Diagrama de Comunicación (UML)

## ¿Qué es?

El diagrama de comunicación (antes llamado diagrama de colaboración) es un diagrama de comportamiento en UML que muestra la interacción entre objetos para cumplir una funcionalidad, enfocándose en las relaciones entre ellos más que en el orden temporal.

A diferencia del diagrama de secuencia, que destaca el tiempo, el de comunicación resalta la estructura de enlaces entre objetos y cómo fluye la comunicación a través de ellos.

---

## Estructura

Un diagrama de comunicación está compuesto por:

- **Objetos / Participantes** → Instancias que colaboran en el escenario (se dibujan como rectángulos).  
- **Enlaces (links)** → Líneas que conectan a los objetos para indicar que pueden comunicarse.  
- **Mensajes** → Se escriben sobre los enlaces, numerados para indicar el orden de envío.  
- **Secuencia de mensajes** → Se indica mediante numeración jerárquica (1, 1.1, 1.2, 2, etc.).  

---

## Principales características

- Representa cómo los objetos colaboran entre sí en un escenario.  
- Se centra en la estructura de comunicación, no en el tiempo exacto.  
- El orden de los mensajes se indica con números en lugar de flechas temporales.  
- Muy útil cuando se quiere destacar quién se comunica con quién en lugar de cuándo.  
- Es complementario al diagrama de secuencia: ambos muestran la misma interacción desde diferentes perspectivas.  

---

## Elementos principales

- **Objetos** → Los participantes de la interacción (ejemplo: `usuario:Cliente`, `sistema:App`).  
- **Enlaces** → Líneas que conectan los objetos para indicar una posible comunicación.  
- **Mensajes** → Texto escrito sobre el enlace con el número de orden y la operación.  
  - Ejemplo: `1: validarUsuario()`.  
- **Secuencia jerárquica** → Numeración que refleja la jerarquía de llamadas.  
  - Ejemplo:  
    ```
    1: login()
    1.1: verificarCredenciales()
    1.2: generarToken()
    ```

---

## Conexiones

### 1. Enlace (link)
- Línea sólida entre dos objetos.  
- Indica que pueden comunicarse.  
- **Ejemplo:**  
  `Cliente -------- Sistema`

### 2. Mensajes con orden
- Se escriben sobre el enlace con un número y acción.  
- Indican el flujo de llamadas.  
- **Ejemplo:**  
  `1: iniciarSesión(usuario, clave)`

### 3. Secuencia jerárquica
- Permite detallar llamadas internas dentro de una interacción.  
- **Ejemplo:**  

- 1: iniciarSesión()

  - 1.1: validarUsuario()

  - 1.2: consultarBaseDatos()

  - 1.3: devolverResultado()



### 4. Relación con casos de uso
- Se utiliza para detallar cómo se implementa un caso de uso mostrando las interacciones entre los objetos participantes.  

---

## Ventajas

- Muestra de manera clara qué objetos participan en la interacción.  
- Permite ver la estructura de comunicación dentro de un sistema.  
- Complementa al diagrama de secuencia, dando otra perspectiva del mismo escenario.  
- Útil para analizar responsabilidades entre objetos.  
- Sencillo para escenarios pequeños o medianos.  

---

## Desventajas

- Puede volverse difícil de leer si hay demasiados objetos o mensajes.  
- El orden temporal no es tan evidente como en el diagrama de secuencia.  
- Puede resultar redundante si ya existen diagramas de secuencia detallados.  
- No refleja la arquitectura interna, solo las interacciones externas.  

---

## Conclusión

El diagrama de comunicación es una herramienta útil para representar cómo colaboran los objetos entre sí en un escenario, destacando las relaciones y el flujo de mensajes numerados.  
Es complementario al diagrama de secuencia: mientras uno resalta el tiempo, el otro resalta la estructura de comunicación.  
Aunque puede complicarse en sistemas grandes, es valioso para comprender y documentar cómo los elementos trabajan juntos para cumplir un caso de uso.
