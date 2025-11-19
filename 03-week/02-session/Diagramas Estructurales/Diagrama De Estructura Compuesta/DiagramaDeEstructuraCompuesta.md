# Diagrama de Estructura Compuesta (UML)

## ¿Qué es?

El diagrama de estructura compuesta es un diagrama estructural de UML que representa la estructura interna de una clase, componente o colaboración, mostrando las partes que lo conforman y las interacciones entre ellas.
A diferencia de otros diagramas más generales, este se centra en el detalle interno, ilustrando cómo se organizan las partes de un elemento y cómo se comunican entre sí para cumplir una función específica.

Se usa mucho para describir la arquitectura interna de un componente o subsistema.

---

## Estructura

Un diagrama de estructura compuesta está compuesto por:

- **Partes (parts)** → Son instancias de roles dentro de una clase o componente (se dibujan como rectángulos dentro de otro rectángulo).
- **Conectores (connectors)** → Representan las conexiones entre partes (se dibujan con líneas).
- **Puertos (ports)** → Puntos de interacción de un elemento con su entorno (se dibujan como cuadrados pequeños en el borde de un rectángulo).
- **Interfaces** → Asociadas a los puertos, muestran los servicios ofrecidos o requeridos.
- **Colaboraciones** → Grupos de roles y relaciones que trabajan juntos para un propósito específico.

---

## Principales características

- Muestra la estructura interna de un elemento (clase, componente o colaboración).
- Representa las partes y roles internos que conforman un elemento complejo.
- Describe cómo colaboran internamente los elementos para lograr un comportamiento.
- Es útil para entender la arquitectura interna y las responsabilidades de cada parte.
- Permite detallar la implementación interna de un componente o subsistema.

---

## Elementos principales

- **Clases / Componentes** → El contenedor principal que se descompone.
- **Partes** → Subunidades internas de un elemento (ejemplo: dentro de *Vehículo*, una parte *motor:Motor*).
- **Puertos** → Puntos de acceso de un elemento hacia el exterior.
- **Conectores** → Representan enlaces de comunicación entre partes internas o con el exterior.
- **Interfaces** → Definen lo que se ofrece o requiere a través de los puertos.
- **Colaboraciones** → Definen cómo las partes trabajan en conjunto para un propósito.

---

## Conexiones
---

## 1. Conectores (Connectors)

- Son líneas sólidas que unen dos partes dentro de un mismo elemento.  
- Representan la **comunicación directa** entre esas partes.

**Ejemplo:** dentro de un *Vehículo*, el conector une el motor con el sistema de frenos:

[motor:Motor] -------- [sistemaFrenos:Frenos]

## 2. Puertos (Ports)

- Son pequeños cuadrados en el borde del rectángulo que representa un elemento (clase, componente o subsistema).  
- Indican los **puntos de acceso o interacción hacia el exterior**.

**Ejemplo:** en un *Componente de Autenticación*, el puerto `authPort` expone el servicio *ILogin*:

[Componente Autenticación]
└─ ◼ authPort

---

## 3. Interfaces proporcionadas (⚪)

- Se dibujan como un círculo conectado a un puerto o parte.  
- Muestran qué funcionalidad **ofrece un elemento**.

**Ejemplo:** el *Motor* proporciona la interfaz *IGenerarPotencia*:


[motor:Motor] ⚪ IGenerarPotencia

---

## 4. Interfaces requeridas (◑)

- Se dibujan como un semicírculo conectado a un puerto o parte.  
- Muestran qué funcionalidad **necesita un elemento**.

**Ejemplo:** el *Sistema de Frenos* requiere la interfaz *ISuministroHidráulico*:


[sistemaFrenos:Frenos] ◑ ISuministroHidráulico

---

## 5. Conexión de interfaces (⚪ ↔ ◑)

- Es una línea que une un círculo y un semicírculo.  
- Muestra que un **servicio ofrecido es utilizado por otro elemento**.

**Ejemplo:** el *Motor* proporciona potencia y el *Sistema de Frenos* la requiere:


[motor:Motor] ⚪ IGenerarPotencia -------- ◑ ISuministroHidráulico [sistemaFrenos:Frenos]


---

## 6. Conectores delegados

- Son líneas desde un puerto del contenedor hacia una parte interna.  
- Indican que el **puerto externo pasa la interacción** a un elemento interno.

**Ejemplo:** un *ServidorWeb* tiene un puerto `httpPort`, pero delega la petición al *ManejadorDePeticiones* interno:


[ServidorWeb]
◼ httpPort -------- [ManejadorDePeticiones]

---

## Ventajas

- Permite comprender cómo está compuesto internamente un elemento complejo.
- Detalla la colaboración entre subpartes, lo que ayuda en el diseño detallado.
- Facilita la documentación de la arquitectura interna de un sistema o subsistema.
- Ayuda a la planificación de equipos de desarrollo (cada parte puede corresponder a un módulo distinto).
- Es útil para modelar sistemas embebidos o arquitecturas basadas en componentes.

---

## Desventajas

- Puede ser difícil de interpretar si el sistema tiene demasiadas partes internas.
- Requiere un alto nivel de detalle que puede no ser necesario en todas las fases del desarrollo.
- Puede volverse redundante si ya existen diagramas de clases y de componentes bien definidos.
- No es adecuado para sistemas pequeños, porque se vuelve innecesario.

---

## Conclusión

El diagrama de estructura compuesta es una herramienta poderosa para detallar la organización interna de clases, componentes o colaboraciones, mostrando cómo se relacionan y trabajan juntas sus partes.  
Es especialmente útil en la fase de diseño detallado y en arquitecturas complejas, donde se necesita entender no solo qué hace un elemento, sino también cómo está compuesto internamente y cómo interactúa con el entorno a través de puertos e interfaces.
