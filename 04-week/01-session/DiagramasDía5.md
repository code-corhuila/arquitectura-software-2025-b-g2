## Parte 5 del Foro **"DIAGRAMA DE ESTRUCTURA COMPUESTA Y DE COMUNICACIÓN"**
----------

## DIAGRAMA DE ESTRUCTURA COMPUESTA 

### QUE ES?
es un tipo de diagrama estructural que muestra la estructura interna de una clase, componente o colaboración y cómo sus partes se relacionan para cumplir una funcionalidad.

### Elementos principales

**Partes (Parts):** Representan instancias internas que forman parte de la clase o componente.

**Puertos (Ports):** Puntos de interacción con el exterior.

**Conectores (Connectors):** Enlaces que muestran cómo las partes internas se comunican entre sí o con el exterior.

**Colaboraciones:** Roles o interacciones que ayudan a cumplir una función específica.

# **ACTIVIDAD**
----
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
# SOLUCIÓN DIAGRAMA DE ESTRUCTURA COMPUESTA 
---
Este diagrama muestra la **arquitectura interna** del módulo de Asignación de Puertas y Colas de Prioridad, así como sus **conexiones con otros servicios**.

### Componentes principales:

- **Asignación de Puertas (AP):**
  - **GestorPuertas (GP):** valida y asigna las puertas de embarque.  
  - **MotorPrioridad (MP):** solicita el siguiente grupo a embarcar.  
  - **ReglasNegocio (RN):** aplica excepciones (overbooking, cambio de puerta, retraso, no-show).  
  - **AdaptadorFIDS (AF):** envía actualizaciones a pantallas.  
  - **RepoVuelos (RV):** consulta y actualiza información del vuelo/gate.  
  - **EventBus/Logger (EV):** registra eventos, métricas y alertas.

- **Colas de Prioridad (CP):**
  - **GestorColas (GC):** administra la cola de pasajeros.  
  - **EstrategiaPrioridad (EP):** define el orden de abordaje (PMR, familias, business, etc.).  

### Conexiones externas (puertos):
- **IGestVuelos:** permite que AP consulte y actualice el estado del vuelo.  
- **IFIDS:** envía datos a pantallas de información.  
- **ISeguridad:** notifica los llamados de grupo a control de seguridad.  

### Notas importantes:
- Se incluyen reglas de negocio para manejar situaciones críticas (overbooking, retrasos, cambios).  
- El EventBus registra métricas y tiempos para **observabilidad y cumplimiento de SLA**.  

**En resumen:** este diagrama refleja la **organización interna y responsabilidades** de cada submódulo, y cómo se conectan con sistemas externos.

---

## Parte 5 del Foro **"DIAGRAMA DE ESTRUCTURA COMPUESTA Y DE COMUNICACIÓN"**
----------

## DIAGRAMA DE COMUNICACIÓN 
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

# **ACTIVIDAD**
----
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
# SOLUCIÓN DIAGRAMA DE COMUNICACIÓN 
---
# Diagrama de Comunicación – Proceso de Abordaje de Vuelo

Este diagrama muestra cómo se comunican los servicios entre sí durante el proceso de abordaje de un vuelo.  
Cada flecha numerada representa un mensaje o interacción concreta:

1. **Gestión de Vuelos (GV) → Asignación de Puertas (APS):** notifica estado del vuelo (ej. *ABORDANDO*).  
2. **Asignación de Puertas (APS) → Colas de Prioridad (CPS):** solicita el siguiente grupo de prioridad.  
3. **Colas de Prioridad (CPS) → Seguridad (SEG):** avisa que se debe llamar al grupo especial (ej. PMR, familias, business).  
4. **Seguridad (SEG) → DCS/PSS (aerolínea):** valida el *boarding pass* del pasajero.  
5. **DCS/PSS → Seguridad (SEG):** responde validación (*OK* o *REJECT*).  
6. **Seguridad (SEG) → Migración (MIG):** envía al pasajero a control migratorio.  
7. **Migración (MIG) → Seguridad (SEG):** devuelve respuesta de control.  
8. **Seguridad (SEG) → Gestión de Pasajeros (GP):** informa que el pasajero fue embarcado.  
9. **Gestión de Pasajeros (GP) → Gestión de Vuelos (GV):** actualiza el estado del pasajero en el vuelo.  
10. **Asignación de Puertas (APS) → FIDS:** actualiza información en pantallas (gate, grupos llamados).  
11. **Gestión de Vuelos (GV) → EventBus (EB):** publica eventos de estado del vuelo.  
12. **Seguridad (SEG) → EventBus (EB):** envía evento de control de seguridad (*scan*).  
13. **Colas de Prioridad (CPS) → EventBus (EB):** envía métricas de tiempo de espera en cola.

---

## Clasificación de servicios

- **Servicios internos:**  
  - Gestión de Vuelos (GV)  
  - Gestión de Pasajeros (GP)  
  - Asignación de Puertas (APS)  
  - Colas de Prioridad (CPS)  
  - Seguridad (SEG)  

- **Servicios externos:**  
  - DCS/PSS (aerolínea)  
  - Migración (MIG)  

- **Infraestructura:**  
  - FIDS  
  - EventBus (EB)  

---
**En resumen:** este diagrama refleja el flujo de mensajes y validaciones entre los módulos internos del aeropuerto (GV, GP, APS, CPS, SEG), los externos (DCS/PSS, Migración) y la infraestructura de soporte (FIDS, EventBus).

---