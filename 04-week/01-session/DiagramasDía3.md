## Parte 3 del Foro **"DIAGRAMA DE COMPONENTES Y DE SECUENCIA"**
----------

## DIAGRAMA DE COMPONENTES
## ¿Qué es?

El diagrama de componentes es un diagrama estructural de UML que muestra cómo el sistema se divide en componentes físicos y reutilizables (como módulos, bibliotecas, archivos ejecutables o servicios), y cómo estos interactúan mediante interfaces y dependencias. se centra en la arquitectura física y modular del software.  
Es muy usado en la etapa de diseño e implementación porque muestra cómo se organizarán y desplegarán las partes del sistema.

---

## Estructura

Un diagrama de componentes está compuesto por:

- **Componentes** → Representados como rectángulos con pestañas en un costado (simbolizando un módulo o bloque de software).
- **Interfaces** → Puntos de acceso de los componentes (pueden ser proporcionadas o requeridas).
- **Dependencias** → Relaciones entre componentes que indican que uno necesita de otro para funcionar.
- **Artefactos** → Archivos físicos como `.jar`, `.dll`, `.exe`, bibliotecas, servicios, etc.
- **Paquetes** → Sirven para organizar y agrupar componentes relacionados.

---

# **DIAGRAMA DE COMPONENTES - GESTIÓN DE VUELOS**
# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
## SOLUCIÓN DIAGRAMA DE COMPONENTES
----

Este diagrama muestra cómo se organizan los servicios principales del sistema y las interfaces expuestas para su comunicación.

## Componentes principales

### GestiónVuelos
- Maneja la programación de vuelos, asignación de puertas, cambios de estado (programado, abordando, cerrado, despegado, cancelado).
- Provee servicios de consulta de estado y actualización en tiempo real.

### GestiónPasajeros
- Controla check-in, validación de documentación, emisión/validación de boarding pass y flujo de embarque.
- Se conecta con sistemas externos (migración, seguridad).

### IntegraciónDCS (Departure Control System)
- Permite interacción con sistemas de aerolíneas (DCS/PSS).
- Valida boarding pass, maneja excepciones como overbooking.

### FIDS (Flight Information Display System)
- Muestra la información del vuelo en pantallas de sala.
- Recibe actualizaciones en tiempo real desde GestiónVuelos.

## Relaciones
- **GestiónPasajeros ↔ IntegraciónDCS** → Intercambio de datos para validación de boarding pass y asignación de asiento.
- **GestiónVuelos ↔ FIDS** → El estado de los vuelos se refleja en pantallas de información.
- **GestiónPasajeros ↔ GestiónVuelos** → El embarque y el estado del pasajero afectan el estado del vuelo.
- **IntegraciónDCS ↔ GestiónVuelos** → Actualiza cambios operativos (ejemplo: overbooking, reubicaciones).
---

## Parte 3 del Foro **"DIAGRAMA DE COMPONENTES Y DE SECUENCIA"**
----------

## DIAGRAMA DE SECUENCIA
--
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

## SOLUCIÓN DIAGRAMA DE SECUENCIA
 - Embarque por Grupos

El diagrama de secuencia muestra cómo se coordina el embarque por grupos en el sistema.  

Primero, **GestiónVuelos** anuncia el inicio de embarque y envía la información a **FIDS**, que lo muestra en pantalla.  

Los pasajeros del grupo llamado presentan su boarding pass a **GestiónPasajeros**, quien lo valida con **IntegraciónDCS**.  

Si el pase es válido, el pasajero aborda y se registra en **GestiónVuelos**; si es inválido, se le niega el acceso.  

Finalmente, el estado del vuelo y de los embarques se actualiza en tiempo real en **FIDS** hasta completar el proceso. 

# Conexiones en el Diagrama de Secuencia

La conexión se da en los mensajes intercambiados entre los componentes:

- **GestiónVuelos ↔ FIDS** → conecta para anunciar y actualizar los grupos de embarque.  

- **Pasajero ↔ GestiónPasajeros** → conexión directa al presentar el boarding pass.  

- **GestiónPasajeros ↔ IntegraciónDCS** → conexión para validar el boarding pass en el sistema de la aerolínea.  

- **GestiónPasajeros ↔ GestiónVuelos** → conexión para registrar el estado de embarque en tiempo real.  

En otras palabras, cada flecha del diagrama representa una conexión lógica entre servicios.  
