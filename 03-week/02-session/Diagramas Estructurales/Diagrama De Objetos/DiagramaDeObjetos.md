# Diagrama de Objetos (UML)

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

## Principales características

- Representa casos reales del sistema en un momento dado.  
- Complementa al diagrama de clases, mostrando su aplicación práctica.  
- Ayuda a validar el modelo de clases, porque se ven ejemplos concretos de cómo se relacionan.  
- Se usa para documentar escenarios o ejemplos de ejecución.  
- Útil en la fase de análisis y diseño, especialmente cuando se explican casos de uso con instancias reales.  

---

## Elementos principales

- **Objetos** → Instancias de clases con nombre y atributos.  
- **Atributos con valores** → Información concreta en el momento descrito.  
- **Líneas de relación** → Conectan los objetos según las asociaciones definidas en el diagrama de clases.  
- **Multiplicidad en la práctica** → Si una clase podía tener muchos elementos (ejemplo: un carrito puede tener muchos productos), en el diagrama de objetos se muestran las instancias reales que se estén usando (ejemplo: `carrito1` tiene exactamente 3 productos).  

---

## Conexiones

- Los objetos se conectan mediante asociaciones que ya están definidas en el diagrama de clases, pero aquí se muestran con instancias reales.  
- Puede haber enlaces directos (`--`) entre objetos que interactúan en ese escenario.  
- No se muestran generalizaciones ni herencias (eso se ve en el diagrama de clases).  

En el diagrama de objetos, lo que se representa son instancias reales de esas clases, por lo tanto no se muestran relaciones de herencia, implementación, agregación o composición con símbolos como <|--, *--, o--.

En su lugar:

✅ Se muestran asociaciones simples entre objetos, usando líneas (--).

✅ Estas asociaciones son las que derivan del diagrama de clases, pero con ejemplos concretos (instancias).

❌ No se usan herencia, dependencia, agregación o composición con sus símbolos, porque esas son relaciones de nivel de clase, no de objeto.

### Ejemplo:

En un diagrama de clases puedes tener:
Carrito *-- Producto (composición: un carrito no existe sin productos).

En un diagrama de objetos, verías algo así:

carrito1:Carrito conectado con tres instancias de productos → producto1:Producto, producto2:Producto, producto3:Producto.

### En resumen:
En el diagrama de objetos solo se muestran instancias y sus enlaces reales en ese momento, mientras que las relaciones avanzadas (<|--, *--, o--, -->) solo aparecen en el diagrama de clases.

---

## Ventajas

- Permite visualizar ejemplos concretos de cómo funciona el sistema.  
- Facilita la comprensión para quienes no dominan UML, porque muestra datos reales.  
- Ayuda a validar la lógica del modelo antes de pasar a la implementación.  
- Es una herramienta útil para explicar casos de uso con instancias específicas.  

---

## Desventajas

- No es práctico para sistemas grandes, ya que se llenaría de demasiados objetos.  
- Representa solo un estado temporal, por lo que no cubre todos los posibles escenarios.  
- Puede volverse redundante si ya se tienen muy claros los diagramas de clases.  
- No muestra la evolución del sistema (es estático en un momento dado).  

---

## Conclusión

El diagrama de objetos es una herramienta muy útil para **complementar el diagrama de clases**, ya que permite ver cómo los conceptos definidos en el modelo se convierten en **instancias concretas con datos reales** en la ejecución.  
Su mayor aporte está en la **claridad que brinda para validar el diseño** y para explicar a otras personas cómo interactúan los elementos del sistema en un caso particular.
