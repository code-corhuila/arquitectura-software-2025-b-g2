## Parte 1 del Foro **"DIAGRAMA DE CASOS DE USO Y DE PAQUETES"**
----------

## DIAGRAMA CASO DE USO
----
## ¿Qué es?

El diagrama de casos de uso es un diagrama de comportamiento de UML que representa las funcionalidades principales que un sistema ofrece a los usuarios u otros sistemas (llamados actores).  
Muestra de manera gráfica qué puede hacer un sistema, sin entrar en detalle de cómo se implementa.

Es muy utilizado en la fase de análisis de requisitos, ya que comunica de forma clara qué espera el usuario del sistema.

---

## Estructura

Un diagrama de casos de uso está compuesto por:

- **Actores** → Representan usuarios u otros sistemas que interactúan con el sistema (se dibujan como un “muñeco” o ícono de persona).
- **Casos de uso** → Funcionalidades o servicios que el sistema proporciona al actor (se dibujan como óvalos).
- **Sistema** → Se representa con un rectángulo que contiene los casos de uso.
- **Relaciones** → Conexiones entre actores y casos de uso o entre casos de uso.
----
-----

# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
# SOLUCIÓN DIAGRAMA DE CASOS DE USO 
---
---
# Explicación del Diagrama de Casos de Uso - Sistema de Vuelos y Pasajeros ✈️

Este documento explica el diagrama UML de **casos de uso** que modela el **sistema de gestión de vuelos, abordaje y seguimiento de pasajeros** en un aeropuerto.

---

## Objetivo del Diagrama
Representar las **interacciones principales** entre actores externos (aerolínea, pasajeros, autoridades, etc.) y el sistema del aeropuerto.  
Se incluyen **funcionalidades críticas**, integraciones y manejo de **excepciones** como sobrecupo (*overbooking*), retrasos y cambios de puerta.

---

## 👥 Actores Principales

- **Pasajero** 🧳: Realiza check-in, presenta documentos, embarca.  
- **Aerolínea (DCS/PSS)** 🛫: Sistema externo que emite y valida pases de abordar.  
- **Autoridad Migratoria** 🛂: Controla documentos y permisos de salida/entrada.  
- **Seguridad Aeroportuaria** 👮: Realiza control y escaneo de seguridad.  
- **Pantallas FIDS** 📺: Sistema de información que muestra estado del vuelo y puerta.  
- **Operador del Aeropuerto** 🧑‍💼: Supervisa operaciones, maneja excepciones (cambio de puerta, retrasos, sobrecupo).  

---

## Casos de Uso Clave

- **Gestionar vuelo**: Programación, estado, asignación de puerta.  
- **Check-in de pasajero**: Registro y entrega de boarding pass.  
- **Validar boarding pass**: Confirmación con aerolínea.  
- **Control de seguridad**: Escaneo y validación de pasajero.  
- **Embarque de pasajero**: Acceso final a la aeronave.  
- **Seguimiento en tiempo real**: Registro de estado y alertas operativas.  

---

## Relaciones Importantes
## Puntos Claves
- Flechas ..>

En UML, la flecha con línea de puntos (..>) se usa para representar dependencias.
En casos de uso, las más comunes son:

<<include>>
- Es una relación obligatoria.
- Se usa cuando un caso de uso siempre necesita ejecutar otro para poder completarse.
---
<<extend>>
- Es una relación condicional.
- Se usa cuando un caso de uso puede extender opcionalmente a otro en una situación específica.
## Criterios de Calidad Cumplidos

- **Corrección UML** ✅: Uso de notación estándar UML 2.x, flechas de dependencia y estereotipos (`<<include>>`, `<<extend>>`).  
- **Consistencia** ✅: Casos de uso trazables a actores reales y escenarios del contexto.  
- **Claridad** ✅: Nombres significativos, límites del sistema bien definidos.  
- **Contexto** ✅: Cobertura de escenarios críticos (overbooking, retrasos, cambios de puerta).  
- **Presentación** ✅: Diagrama legible, agrupado por flujos lógicos y actores.  

---

## Conclusión

El diagrama refleja cómo el sistema de aeropuerto coordina **procesos de pasajeros, vuelos y seguridad**, integrándose con **aerolíneas y autoridades**.  
Además, incorpora el manejo de **eventos excepcionales** y garantiza **observabilidad** (registro de alertas y estados en tiempo real).  

-----
----

## Parte 1 del Foro **"DIAGRAMA DE CASOS DE USO Y DE PAQUETES"**
----------

## **DIAGRAMA CASO PAQUETES**
----

## ¿Qué es?

El diagrama de paquetes es un diagrama estructural de UML que organiza y agrupa los elementos de un sistema (clases, interfaces, componentes, casos de uso, etc.) en paquetes.  
Un paquete funciona como un contenedor que ayuda a dividir el sistema en módulos, facilitando la comprensión, el mantenimiento y la escalabilidad del software.

Este tipo de diagrama es muy usado para mostrar la arquitectura lógica y modular de un sistema, evitando la sobrecarga de diagramas demasiado grandes.

---

## Estructura

Un diagrama de paquetes está compuesto por:

- **Paquetes (Packages)** → Representan agrupaciones lógicas de elementos. Se dibujan como una carpeta (rectángulo con una solapa superior).
- **Elementos contenidos** → Pueden ser clases, interfaces, casos de uso, otros paquetes, etc.
- **Dependencias entre paquetes** → Representan cómo un paquete necesita de otro.
- **Relaciones jerárquicas** → Un paquete puede contener subpaquetes.
- **Accesibilidad de elementos** → Se puede definir qué elementos de un paquete son públicos (accesibles desde fuera) y cuáles privados (ocultos).
----
-----

# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
## SOLUCIÓN DIAGRAMA DE PAQUETES
---
# Explicación del código – Diagrama de Paquetes del Sistema de Vuelos

El diagrama organiza el sistema en capas de paquetes UML para representar la arquitectura:

---

## Presentación
Incluye interfaces de usuario como **Web/App/Kiosko** y **Pantallas FIDS** para mostrar información a pasajeros.

---

## Dominio
Contiene la lógica central: **Gestión de Vuelos, Pasajeros, Embarque, Seguimiento y Excepciones.**

Modela procesos como **Check-in, Reserva, Abordaje, Overbooking, Cambio de puerta, Retrasos y No show.**

---

## Integración
Conecta con **Sistemas externos (DCS/PSS, Migración, Seguridad)** y la **Base de Datos.**

---

## Relaciones
- Las interfaces de usuario se comunican con los procesos del dominio (Check-in, Tracking, Alertas).  
- Dentro del dominio, los paquetes se relacionan para reflejar dependencias reales (ej. *Check-in → Reserva*).  
- El dominio guarda y consulta información en la **Base de Datos.**  
- Se incluyen relaciones `<<extend>>` para excepciones como **Overbooking** o **Retrasos.**  

---

## En conclusión
El diagrama muestra una **arquitectura organizada por capas**, asegurando **claridad, trazabilidad y separación de responsabilidades** en el sistema de gestión de vuelos.
