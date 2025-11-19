## Parte 6 del Foro **"DIAGRAMA DE OBJETOS Y TIMING"**
----------

## DIAGRAMA DE OBJETOS

## ¿Qué es?

El diagrama de objetos es un diagrama estructural de UML que muestra una instancia específica de las clases y sus relaciones en un momento determinado del tiempo.  
Mientras que el diagrama de clases representa el modelo estático (la definición de clases y sus relaciones), el diagrama de objetos es como una **fotografía del sistema en ejecución**, con objetos reales creados a partir de esas clases.

---

## Estructura

Un diagrama de objetos está compuesto por:

- **Objetos (instancias de clases)** → Representados por rectángulos con el nombre del objeto y su clase (ejemplo: `carrito1:Carrito`).  
- **Atributos con valores específicos** → A diferencia del diagrama de clases, aquí los atributos muestran datos concretos (ejemplo: `total = 20000`).  
- **Relaciones entre objetos** → Muestran cómo se conectan en la práctica (asociaciones, dependencias, agregación, composición, etc.).  
- **Nombres de objetos** → Suelen escribirse en minúsculas para diferenciarlos de las clases. 
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
# SOLUCIÓN DIAGRAMA DE OBJETOS 
---
## 1. Diagrama de Objetos – Instancias del Vuelo AV123

El diagrama de objetos es una **fotografía estática** del sistema en un instante puntual.  
No muestra clases genéricas, sino **objetos concretos** con valores reales en sus atributos y cómo se relacionan entre sí.

### ¿Qué representa?
- **Instancia de Vuelo (AV123 : Vuelo)**  
  Representa el vuelo específico AV123 que va de Bogotá (BOG) a Miami (MIA).  
  Tiene un **estado actual** (`ABORDANDO`), una **hora programada** (`14:30`) y una **puerta asignada** (`A5`).

- **Instancia de Puerta (Gate A5 : Puerta)**  
  Es la puerta física del aeropuerto donde se realiza el embarque.  
  Su estado es `DISPONIBLE` y está vinculada directamente al vuelo AV123.

- **Instancia de Pasajero (Pax #987654 : Pasajero)**  
  Ejemplo de un pasajero real: Sofía A., con prioridad **PMR** (Personas con Movilidad Reducida).  
  Su estado de embarque es `EN_COLA`, indicando que está esperando para ingresar al avión.

- **Instancia de Boarding Pass (BP : BoardingPass)**  
  Documento digital/físico que valida al pasajero para el vuelo.  
  Contiene datos concretos: asiento `12C`, zona `1`, estado `válido`.  
  Está **asociado al pasajero**.

- **Instancia de Cola (Cola Prioridad : Cola)**  
  Representa la fila de embarque donde los pasajeros esperan por grupo.  
  Muestra qué grupos están configurados (`PMR, Familias, Business, General`) y cuál es el **siguiente en abordar** (`PMR`).

### Relaciones claves en el diagrama
- **Vuelo ↔ Puerta (asignadoA)**  
  El vuelo AV123 está asignado a la puerta A5.  
  Esto refleja la infraestructura: un vuelo no puede embarcar sin una puerta disponible.

- **Pasajero ↔ BoardingPass (porta)**  
  El pasajero #987654 porta un boarding pass válido.  
  Sin esta relación, el pasajero no podría pasar los controles de seguridad ni embarcar.

- **Pasajero ↔ Cola (esperaEn)**  
  El pasajero está esperando en la **Cola de Prioridad**.  
  Esto refleja que, en ese momento del proceso, no ha sido llamado a embarcar.

### ¿Qué aporta este diagrama?
- Permite ver el **estado exacto** de todos los elementos del sistema en un momento dado.  
- Sirve como complemento a los diagramas de clases y de secuencia, mostrando cómo las entidades abstractas (Vuelo, Pasajero, Puerta…) se instancian en **objetos reales con datos concretos**.  
- Es útil para pruebas, casos de uso y simulación de escenarios reales (ejemplo: “¿qué pasa si un pasajero no tiene boarding pass válido?”).

En resumen: este diagrama responde a la pregunta **“¿Cómo están organizados los objetos en este momento?”**, mientras que el de tiempos responde a **“¿Cómo evolucionan en el tiempo?”**.

---
## Parte 6 del Foro **"DIAGRAMA DE OBJETOS Y TIMING"**
----

## DIAGRAMA DE TIMING
## ¿Qué es?

El **diagrama de tiempos** es un diagrama de comportamiento en UML que muestra la evolución de un objeto o varios objetos en función del tiempo.  
Se utiliza para representar cómo cambian los estados, valores o interacciones de un sistema a lo largo de un intervalo temporal.

Se parece a un diagrama de secuencia, pero en lugar de centrarse en los mensajes entre objetos, se enfoca en la duración y el orden cronológico de los estados o eventos.

## Estructura

Un diagrama de tiempos incluye:

- **Línea de vida** → Representa un objeto o participante, dispuesta en sentido vertical.  
- **Eje de tiempo** → Generalmente en el eje horizontal, indicando la progresión del tiempo.  
- **Estados o valores** → Condiciones del objeto representadas a lo largo de la línea de vida.  
- **Duraciones** → Intervalo de tiempo en el que un objeto permanece en un estado.  
- **Eventos** → Sucesos que provocan un cambio de estado o valor en un objeto.  
- **Restricciones temporales** → Reglas sobre duración o tiempos de transición (ejemplo: “máximo 5 segundos”).  
- **Sincronización** → Indica la coordinación entre varios objetos o procesos en el tiempo.  

# **ACTIVIDAD**

**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------

# SOLUCIÓN DIAGRAMA DE TIMING
---
Este diagrama es **dinámico** y muestra cómo cambian los estados de distintos actores/sistemas a lo largo del tiempo en el proceso de embarque.

### Entidades monitoreadas
- **Pax #987654** → Estado del pasajero en las fases de check-in, cola, seguridad, etc.
- **Cola Prioridad (QC)** → Apertura progresiva de grupos de embarque.
- **Seguridad (SEC)** → Punto de control/escaneo de boarding pass.
- **Gate A5 (GATE)** → Preparación de la puerta y abordaje físico.
- **Gestión Vuelos (GV)** → Estado administrativo del vuelo.
- **FIDS** → Sistema de pantallas que anuncia llamadas de abordaje.

### Línea de tiempo (minutos relativos a la salida)
- **@0** → Estado inicial:
  - Pasajero en `CheckIn`.
  - Cola cerrada.
  - Seguridad libre.
  - Puerta en `Preparando`.
  - Gestión de vuelos en `Programado`.
  - Pantallas en modo estático.

- **@-45 minutos (antes de salida)**:
  - GV cambia estado a `AbordajeProx`.
  - FIDS muestra anuncio previo → aviso a pasajeros.

- **@-30 minutos**:
  - Se abre la cola para grupo `PMR`.
  - El pasajero #987654 entra a `EnColaPMR`.
  - FIDS muestra `LlamadoPMR`.

- **@-25 minutos**:
  - Seguridad inicia `Escaneo`.
  - Pasajero pasa a estado `Seguridad`.

- **@-20 minutos**:
  - Se abre la cola para `Familias`.
  - FIDS muestra `LlamadoFamilias`.

- **@-15 minutos**:
  - Fin de los eventos modelados en este ejemplo (el proceso continúa con Business y General, no dibujados).

### Conexiones clave
- El **Gestión de Vuelos (GV)** orquesta los tiempos → dispara cambios de estado en cola y pantallas.
- El **FIDS** actúa como canal visible para pasajeros.
- La **Cola Prioridad (QC)** se abre progresivamente por grupos.
- La **Seguridad (SEC)** valida al pasajero antes de marcarlo como embarcado.
- La **Puerta (GATE)** está sincronizada con el flujo de pasajeros.

Este diagrama muestra **la evolución temporal** del embarque, a diferencia del diagrama de objetos que muestra un estado puntual.

---