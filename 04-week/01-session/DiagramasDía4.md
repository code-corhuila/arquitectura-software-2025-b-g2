## Parte 4 del Foro **"DIAGRAMA DE DESPLIEGUE Y DE ESTADOS"**
----------

## DIAGRAMA DE DESPLIEGUE

## ¿Qué es?

El diagrama de despliegue es un diagrama estructural de UML que muestra la arquitectura física de un sistema, es decir, cómo el software se ejecuta sobre la infraestructura de hardware.  
Representa los nodos de hardware (servidores, dispositivos, computadoras, móviles, etc.) y cómo se conectan entre sí, además de indicar qué artefactos de software se despliegan en esos nodos.

Se usa mucho en sistemas distribuidos, aplicaciones cliente-servidor y arquitecturas en red.

---

## Estructura

Un diagrama de despliegue está compuesto por:

- **Nodos (Nodes)** → Elementos físicos (servidores, PCs, móviles, routers). Se dibujan como cubos tridimensionales.
- **Artefactos (Artifacts)** → Archivos o componentes de software que se ejecutan en un nodo (se dibujan como rectángulos con la palabra clave *artifact*).
- **Asociaciones de comunicación** → Líneas que unen nodos para representar la red o los canales de comunicación.
- **Componentes / aplicaciones** → El software desplegado dentro de un artefacto.
- **Estereotipos** → Se usan para identificar tipos de nodos o artefactos, como `<<device>>`, `<<database>>`, `<<server>>`, `<<executable>>`.

# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
# SOLUCIÓN DIAGRAMA DE DESPLIEGUE
---
---
# 📌 Diagrama de Despliegue (Cloud + On-Prem)

Este tipo de diagrama no modela lógica de negocio ni flujos, sino la **infraestructura física y lógica** donde se despliega el sistema: qué nodos existen, qué roles cumplen y cómo se organiza la arquitectura.

---

## 🔹 Estructura General

Está dividido en **2 entornos principales**:

- **Cloud – Proveedor Nube** → Servicios modernos desplegados en la nube.  
- **On-Prem – Data Center del Aeropuerto** → Infraestructura local, dentro del aeropuerto.  

---

## Cloud – Proveedor Nube

### API Gateway (WAF, RateLimit)
- Primer punto de entrada.  
- Aplica seguridad (**Web Application Firewall**), control de tráfico y límites de peticiones.  

### Balanceador L7 (HTTPS)
- Balanceador de capa 7 (**HTTP/HTTPS**).  
- Distribuye solicitudes hacia los servicios dentro del clúster Kubernetes.  

### Cluster Kubernetes (AKS/EKS/GKE)
- **Orquestador de contenedores**.  
- Allí se despliegan los microservicios y adaptadores:  
  - **FlightSvc** → Servicio de gestión de vuelos.  
  - **PaxSvc** → Servicio de pasajeros.  
  - **BoardingSvc** → Servicio de embarque.  
  - **FIDS-Adapter** → Integra con pantallas de información de vuelos (FIDS).  
  - **Security-Adapter** → Conexión con sistemas de seguridad.  
  - **Migracion-Adapter** → Conexión con sistemas de migración/migración de pasajeros.  
  - **EventBus Client** → Cliente de mensajería/eventos (**Kafka, RabbitMQ, etc.**).  

### DMZ (Demilitarized Zone)
- Zona de seguridad intermedia, protege los sistemas internos.  
- Incluye:  
  - **Reverse Proxy** → Encaminamiento de tráfico interno.  
  - **Firewall** → Control de accesos, reglas de red.  

### Observabilidad
- Módulo de monitoreo centralizado.  
- Incluye:  
  - **Tracing (OpenTelemetry)** → Seguimiento de llamadas distribuidas.  
  - **Logs (ELK/CloudLogs)** → Centralización de logs.  
  - **Métricas (Prom/Grafana)** → Métricas y dashboards.  

---

## On-Prem – Data Center Aeropuerto
- Nodo **Data Center Aeropuerto**.  
- Representa la infraestructura propia del aeropuerto que se conecta con la nube.  
- Allí podrían estar sistemas legados (**check-in, migración, seguridad física, etc.**) que interactúan con los adaptadores en la nube.  

---

## Resumen de lo que modela
Este diagrama refleja:  
- **Arquitectura híbrida** → parte en la nube, parte en el aeropuerto.  
- **Seguridad por capas** → API Gateway → Balanceador → Proxy/Firewall (DMZ).  
- **Despliegue en contenedores (Kubernetes)** → microservicios desacoplados.  
- **Observabilidad completa** → logs, métricas y trazas.  
- **Interoperabilidad con sistemas on-premise** mediante adaptadores.  

En palabras simples: este diagrama muestra cómo se despliega el sistema de gestión de vuelos/pasajeros en un entorno híbrido, con seguridad, escalabilidad y monitoreo integrados.

----

## Parte 4 del Foro **"DIAGRAMA DE DESPLIEGUE Y DE ESTADOS"**
----------
## DIAGRAMA DE ESTADOS

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

# **ACTIVIDAD**
**Contexto**. Se requiere modelar el núcleo de un sistema para un aeropuerto que gestione vuelos, pasajeros, control de check-in, seguridad, asignación de puertas, embarque y seguimiento en tiempo real (estado del vuelo, ubicación del pasajero en proceso, alertas). El sistema se integra con servicios de aerolíneas (emisión/validación de boarding pass), autoridad migratoria, control de seguridad y pantallas de información en sala. Debe soportar picos de demanda, trazabilidad y auditoría, y manejar excepciones (overbooking, cambio de puerta, retrasos, no show).

**Alcance mínimo**
- **Gestión de vuelos:** programación, estado (programado, abordando, cerrado, despegado, cancelado), asignación de puerta.
- **Gestión de pasajeros:** check-in, control de documentación, validación de boarding pass, embarque.
- **Flujo operativo:** colas por prioridad (PMR, familias, business), excepciones y reubicaciones.
- **Integraciones:** aerolínea (DCS/PSS), autoridad migratoria, seguridad aeroportuaria, pantallas (FIDS).
- **Observabilidad:** registro de eventos, tiempos de proceso y alertas operativas.
-------
# SOLUCIÓN DIAGRAMA DE ESTADOS 
---
---
Este diagrama de estados unificado describe el ciclo de vida de un Vuelo y, en paralelo, cómo evoluciona el BoardingPass de cada pasajero

**El diagrama une los dos mundos:**

- **Vuelo = control operativo.**

- **BoardingPass = control individual de pasajeros.**

Cada transición del BoardingPass depende de un estado del vuelo. Es decir:

- **No puedes hacer check-in si el vuelo no está en AbiertoCheckIn.**

- **No puedes embarcar si el vuelo no está en Embarcando.**

- **No puedes estar Abordado si el vuelo aún no cerró embarque o despegó.**

- **1. Inicio**

El diagrama empieza en [*] → Programado.

Significa que un Vuelo inicia en estado Programado (ya existe en el sistema pero aún no ha abierto procesos).

En ese mismo momento se pueden emitir BoardingPass para los pasajeros (Emitido).

- **2. AbiertoCheckIn**

El Vuelo cambia a AbiertoCheckIn cuando se habilita el proceso de check-in.

En este estado, los BoardingPass emitidos pueden cambiar de Emitido a CheckIn cuando el pasajero hace su registro en línea o en mostrador.

- **3. CerradoCheckIn**

Cuando llega la hora límite, el vuelo cambia a CerradoCheckIn.

Ya no se permiten nuevos check-ins.

Solo avanzan quienes tengan BoardingPass en estado CheckIn.

- **4. Embarcando**

El vuelo entra en Embarcando al abrirse la puerta de abordaje.

Aquí, el BoardingPass pasa por varias fases:

Seguridad → el pasajero pasa por control.

Embarque → el pase se valida en la puerta.

Abordado → el pasajero ya está dentro del avión.

- **5. CerradoEmbarque**

Una vez que se cierran las puertas, el vuelo entra a CerradoEmbarque.

Ya no se permiten más pases de abordar → cualquier BoardingPass no validado queda inválido.

- *6. EnVuelo**

El estado EnVuelo inicia en el despegue.

En este punto, todos los BoardingPass válidos deben estar en estado Abordado.

- **7. Aterrizado → Finalizado**

Después del aterrizaje, el vuelo pasa a Aterrizado y finalmente a Finalizado.

Con esto se da por terminado el ciclo del vuelo y del BoardingPass.
