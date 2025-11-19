## Parte 2 del Foro **"DIAGRAMA DE CLASES DE USO Y DE ACTIVIDADES"**
----------
## DIAGRAMA DE CLASES

**¿QUÉ ES?**

El **Diagrama de Clases** es uno de los **diagramas estructurales** más importantes en UML.  
Su objetivo principal es **modelar la estructura estática del sistema** mostrando las **clases**, sus **atributos**, **métodos** y las **relaciones** que existen entre ellas.

Es considerado la **base de la mayoría de los modelos UML**, ya que permite representar el diseño conceptual y lógico de un sistema orientado a objetos.

---

## 1. Elementos básicos
- **Clase** → Representa una entidad.
+Atributos
+Métodos()
- **Enumeración** → Conjunto de valores posibles.
- **Interfaces** → Definen métodos que deben implementarse.

---

## 2. Relaciones

### 🔹 Asociación (—)
- Relación general entre clases.
- Ejemplo: `ClaseA -- ClaseB`

### 🔹 Agregación (◦—)
- Relación débil, una clase contiene a otra pero ambas pueden existir de forma independiente.
- Ejemplo: `ClaseA o-- ClaseB`

### 🔹 Composición (◆—)
- Relación fuerte, una clase depende totalmente de la otra (si la principal muere, la contenida también).
- Ejemplo: `ClaseA *-- ClaseB`

### 🔹 Herencia (△—)
- Una clase hereda atributos y métodos de otra.
- Ejemplo: `ClaseHija <|-- ClasePadre`

### 🔹 Implementación
- Una clase implementa una interfaz.
- Ejemplo: `ClaseA ..|> Interfaz`

---

# **DIAGRAMA DE CLASES - GESTIÓN DE VUELOS**
# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
## SOLUCIÓN DIAGRAMA DE CLASES

# Diagrama de Clases - Gestión de Vuelos

Este diagrama modela las entidades y relaciones principales en un sistema de gestión de vuelos.

## Clases
- **Vuelo**: Contiene datos del vuelo (código, origen, destino, hora, estado).
- **Pasajero**: Representa a cada pasajero con su documento y nivel de prioridad.
- **BoardingPass**: Tarjeta de embarque con código QR, clase de vuelo y estado.
- **Puerta**: Puerta de embarque con número y estado.
- **IntegracionServicio**: Sistemas externos conectados (DCS, PSS, Migración, FIDS).
- **Evento**: Registro de sucesos asociados a un vuelo.

## Relaciones
- **Vuelo *-- Pasajero** → Un vuelo está compuesto por muchos pasajeros.
- **Pasajero *-- BoardingPass** → Cada pasajero tiene al menos un boarding pass.
- **Vuelo o-- BoardingPass** → Un vuelo puede generar múltiples boarding pass (agregación).
- **Vuelo -- Puerta** → Cada vuelo se asigna a una puerta.
- **Vuelo o-- Evento** → Los vuelos registran eventos relacionados.
- **IntegracionServicio -- Vuelo** → Los servicios externos se asocian a los vuelos.

---
 En resumen: El **Vuelo** es la entidad central, relacionada con pasajeros, boarding passes, puerta de embarque, eventos y servicios externos.

-------------
----------
## Parte 2 del Foro **"DIAGRAMA DE CLASES DE USO Y DE ACTIVIDADES"**
----------
## ¿Qué es?

El diagrama de actividades es un diagrama de comportamiento en UML que describe los flujos de trabajo y procesos dentro de un sistema.  
Representa las actividades que se realizan, el orden en que ocurren y las decisiones que pueden alterar el flujo.  

Se usa mucho para modelar procesos de negocio, algoritmos o escenarios de casos de uso, mostrando tanto la secuencia como la concurrencia de actividades.  

---

## Estructura

Un diagrama de actividades está compuesto por:

- **Actividades** → Acciones o pasos de un proceso (se dibujan como rectángulos redondeados).  
- **Flujos (edges)** → Flechas que conectan las actividades y marcan el orden.  
- **Nodos iniciales** → Punto de inicio del proceso (círculo negro sólido).  
- **Nodos finales** → Punto de terminación (círculo con borde y un punto negro dentro).  
- **Decisiones** → Rombos que representan bifurcaciones en el flujo.  
- **Uniones y bifurcaciones (fork/join)** → Líneas que dividen o sincronizan actividades paralelas.  
- **Swimlanes (carriles)** → Divisiones verticales/horizontales que indican responsabilidades (quién ejecuta cada actividad).  

--------
# **DIAGRAMA DE ACTIVIDADES - GESTIÓN DE VUELOS**
# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
## SOLUCIÓN DIAGRAMA DE ACTIVIDADES
El diagrama de actividad representa el proceso que sigue un pasajero desde que llega al aeropuerto hasta el embarque:

1. **Check-in:**  
   - El pasajero presenta sus documentos.  
   - Si son válidos, se genera el *boarding pass*.  
   - Si no, el proceso se detiene.  

2. **Control de seguridad:**  
   - El pasajero pasa por revisión de seguridad.  
   - Si aprueba, accede a la zona de embarque.  
   - Si no, es retenido y el flujo termina.  

3. **Embarque:**  
   - El pasajero espera la llamada de abordaje.  
   - Presenta su *boarding pass* en la puerta.  
   - Si es válido, sube al avión.  
   - Si no, se le niega el acceso.  

En conclusión, el diagrama muestra de forma clara las decisiones críticas del proceso (validación de documentos, seguridad y boarding pass) y cómo cada una determina si el pasajero continúa o no hasta abordar el vuelo.