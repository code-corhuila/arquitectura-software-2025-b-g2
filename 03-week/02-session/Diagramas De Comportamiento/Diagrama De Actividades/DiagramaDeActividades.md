# Diagrama de Actividades (UML)

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

---

## Principales características

- Muestra procesos paso a paso dentro de un sistema.  
- Representa decisiones, paralelismo y concurrencia.  
- Permite visualizar flujos de negocio o algoritmos de manera gráfica.  
- Útil para analizar escenarios complejos con múltiples caminos posibles.  
- Similar a un diagrama de flujo, pero más potente en notación UML.  

---

## Elementos principales

- **Actividad** → Rectángulo redondeado que representa una acción.  
  - Ejemplo: *Validar usuario*.  

- **Nodo inicial** → Círculo sólido que marca el comienzo.  
  - Ejemplo: *inicio del proceso de compra*.  

- **Nodo final** → Círculo con borde y centro negro que marca el fin.  

- **Decisión / Fusión** → Rombo que permite bifurcar o unir flujos.  
  - Ejemplo: [usuario válido] / [usuario inválido].  

- **Fork / Join (paralelismo)** → Línea gruesa que divide o sincroniza múltiples actividades en paralelo.  

- **Swimlanes (carriles)** → Divisiones que asignan actividades a actores o subsistemas.  
  - Ejemplo: *carril Cliente y carril Sistema*.  

---

## Conexiones

### 1. Flujo secuencial
Flechas que conectan actividades en orden.  

**Ejemplo:**  
Iniciar Sesión → Validar Credenciales → Mostrar Menú


### 2. Decisión (if/else)
Rombo con condiciones en cada rama.  

**Ejemplo:**  
[válido] → Entrar al sistema

[inválido] → Mostrar error


### 3. Flujo paralelo (fork/join)
Línea horizontal que divide un proceso en varios simultáneos.  
Otra línea los une cuando deben sincronizarse.  

**Ejemplo:**  
Procesar Pago → en paralelo: Actualizar Inventario y Generar Factura


### 4. Swimlanes
Dividen el diagrama en columnas o filas según el responsable.  

**Ejemplo:**  

- **Carril Cliente:** Realizar Pedido  
- **Carril Sistema:** Validar Stock, Confirmar Compra  

---

## Ventajas

- Muestra claramente los procesos de negocio y flujos de trabajo.  
- Facilita la identificación de decisiones y condiciones.  
- Útil para modelar procesos paralelos y concurrentes.  
- Se entiende fácilmente, incluso por usuarios no técnicos.  
- Ideal para documentar escenarios de casos de uso.  

---

## Desventajas

- Puede volverse muy grande y difícil de seguir en procesos extensos.  
- Si se detalla demasiado, se parece a un pseudocódigo complejo.  
- Requiere esfuerzo para mantenerlo actualizado si cambian los procesos.  
- No muestra la estructura interna del sistema, solo el flujo de actividades.  

---

## Conclusión

El diagrama de actividades es fundamental para modelar procesos, algoritmos y flujos de trabajo dentro de un sistema.  
Ofrece una visión clara del orden, decisiones y paralelismos, facilitando tanto el análisis como la comunicación entre usuarios y desarrolladores.  

Aunque puede complicarse en sistemas grandes, es una herramienta clave en la fase de análisis y diseño funcional.