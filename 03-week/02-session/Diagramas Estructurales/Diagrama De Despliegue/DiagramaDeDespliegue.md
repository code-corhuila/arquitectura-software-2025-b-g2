# Diagrama de Despliegue (UML)

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

---

## Principales características

- Describe la infraestructura física del sistema.  
- Muestra cómo los nodos se comunican y qué software se ejecuta en cada uno.  
- Es útil para documentar arquitecturas distribuidas, nubes, microservicios y sistemas cliente-servidor.  
- Puede complementar diagramas de componentes y casos de uso mostrando su ejecución real.  

---

## Elementos principales

- **Nodo (Node)** → Representa un recurso físico o lógico (ejemplo: `<<server>> Servidor Web`).  
- **Artefacto (Artifact)** → Archivo ejecutable o componente de software desplegado (ejemplo: `<<artifact>> app.war`).  
- **Asociación de comunicación** → Línea sólida que une dos nodos indicando un canal de comunicación (ejemplo: conexión HTTP, JDBC, TCP/IP).  
- **Dispositivos** → Nodos físicos como PC, móvil, impresora, servidor, router.  
- **Entornos de ejecución** → Nodos lógicos dentro de otros nodos (ejemplo: una JVM dentro de un servidor).  
- **Componentes** → El software desplegado dentro de los artefactos.  

---

## Conexiones

### 1. Asociación de comunicación (Communication Path)
- Línea sólida que conecta dos nodos.  
- Representa un canal físico o lógico de comunicación.  

**Ejemplo:**  

`Servidor Web -------- Servidor de Base de Datos`

---

### 2. Nodo con artefactos
- Dentro de un nodo (cubito) se colocan los artefactos desplegados.  
- Indica qué software corre en cada máquina o dispositivo.  

**Ejemplo:**  

`<<server>> Servidor Web` contiene `<<artifact>> app.war`

---

### 3. Conexión cliente-servidor
- Un cliente se conecta a un servidor a través de una línea de comunicación.  
- Permite representar arquitecturas cliente-servidor o de 3 capas.  

**Ejemplo:**  

`Cliente (PC) -------- Servidor Web -------- Servidor BD`

---

### 4. Nodos anidados
- Un nodo puede contener otros nodos lógicos, como máquinas virtuales, contenedores o entornos de ejecución.  

**Ejemplo:**  
<<server>> Servidor Aplicaciones

└─ <<executionEnvironment>> JVM

└─ <<artifact>> sistema.jar

---

### 5. Red de nodos
- Se pueden mostrar múltiples nodos interconectados formando una red distribuida.  
- Ideal para modelar arquitecturas de microservicios o sistemas cloud.  

**Ejemplo:**  
`Servidor API ↔ Servidor BD ↔ Servidor Cache ↔ Servidor Autenticación`

---

## Ventajas

- Muestra claramente dónde se ejecuta cada parte del software.  
- Ayuda a planificar la infraestructura física y lógica necesaria.  
- Útil para arquitecturas distribuidas, sistemas en la nube y despliegues de microservicios.  
- Complementa a los diagramas de componentes mostrando la implementación real.  

---

## Desventajas

- Puede ser demasiado detallado para sistemas pequeños.  
- Si la infraestructura cambia con frecuencia (ejemplo: entornos cloud), el diagrama puede quedar obsoleto rápido.  
- Puede volverse complejo en sistemas con muchos nodos distribuidos.  

---

## Conclusión

El diagrama de despliegue es fundamental para documentar la arquitectura física de un sistema, mostrando cómo los nodos de hardware se conectan entre sí y qué artefactos de software se ejecutan en cada uno.  
Es especialmente útil en el diseño de sistemas distribuidos, ya que permite visualizar la infraest