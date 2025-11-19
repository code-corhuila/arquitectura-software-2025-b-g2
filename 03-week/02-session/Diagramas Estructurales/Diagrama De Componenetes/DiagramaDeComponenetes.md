# Diagrama de Componentes (UML)

## ¿Qué es?

El diagrama de componentes es un diagrama estructural de UML que muestra cómo el sistema se divide en componentes físicos y reutilizables (como módulos, bibliotecas, archivos ejecutables o servicios), y cómo estos interactúan mediante interfaces y dependencias.

Mientras que el diagrama de clases se enfoca en la lógica del sistema, el de componentes se centra en la arquitectura física y modular del software.  
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

## Principales características

- Representa la arquitectura física del sistema.
- Muestra componentes reutilizables y cómo interactúan.
- Se centra en la implementación, no en la lógica de negocio.
- Útil para planear el despliegue y empaquetado del software.
- Permite visualizar cómo se integran módulos desarrollados por diferentes equipos.

---

## Elementos principales

- **Componente** → Módulo de software (ejemplo: *GestorUsuarios*, *MotorPagos*).
- **Interfaz proporcionada** → Representa lo que un componente ofrece (se dibuja con un círculo).
- **Interfaz requerida** → Representa lo que un componente necesita (se dibuja con un semicírculo).
- **Dependencia** → Flecha punteada que indica que un componente depende de otro.
- **Artefactos** → Archivos o entregables físicos (ejemplo: *app.jar*, *auth.dll*).

---

## Conexiones

- Los componentes se conectan a través de interfaces (requeridas ↔ proporcionadas).
- Se pueden mostrar dependencias entre módulos (`--->` punteada).
- Puede haber jerarquías de componentes (subcomponentes dentro de un componente mayor).
- Se distinguen claramente las relaciones de implementación de las de uso.

---
# Relaciones en el Diagrama de Componentes (UML)

En lugar de los símbolos que se usan en **clases/objetos**, en componentes se emplean principalmente estas conexiones:

---

## Dependencia (punteada con flecha `--->`)
- Indica que un componente depende de otro para funcionar.  
- **Ejemplo:** el componente *CarritoCompras* depende de *GestiónProductos*.

---

## Interfaz proporcionada (círculo lleno ⚪)
- Muestra que un componente **ofrece una funcionalidad**.  
- **Ejemplo:** *GestiónUsuarios* proporciona la interfaz *ILogin*.

---

## Interfaz requerida (semicírculo abierto ◑)
- Muestra que un componente **necesita una funcionalidad**.  
- **Ejemplo:** *PortalWeb* requiere la interfaz *ILogin* para autenticar usuarios.

---

## Conexión de interfaces (⚪ ↔ ◑)
- Une una interfaz **proporcionada** con una **requerida**.  
- Es la forma “oficial” de representar que un componente usa la funcionalidad de otro.
---

**Ejemplo:**  
Un componente *CarritoCompras* requiere la interfaz de *GestiónProductos*, y a su vez proporciona la interfaz *CompraFinalizada*.

---

## Ventajas

- Permite ver la modularidad del sistema y cómo se organizan sus partes.
- Facilita la reutilización de software al identificar componentes independientes.
- Ayuda a la planificación de equipos de desarrollo (cada equipo puede trabajar en un componente).
- Es muy útil en proyectos grandes porque da una visión clara de la arquitectura física.
- Ayuda en la documentación para despliegue e integración continua.

---

## Desventajas

- Puede ser difícil de mantener actualizado si la arquitectura cambia constantemente.
- No muestra detalles internos del componente (solo su relación con otros).
- Puede volverse complejo en sistemas muy grandes con muchos módulos.
- Requiere una buena comprensión previa del sistema para que sea claro.

---

## Conclusión

El diagrama de componentes es una herramienta clave en la arquitectura de software, ya que muestra cómo se dividen, organizan y conectan los módulos físicos del sistema.  
Es fundamental para equipos de desarrollo e integración, porque permite entender la modularidad, dependencias y puntos de conexión del software.  
Su mayor valor está en visualizar la arquitectura real del sistema, facilitando la colaboración, la escalabilidad y el mantenimiento.
