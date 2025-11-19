# Diagrama de Tiempos (UML)

## ¿Qué es?

El **diagrama de tiempos** es un diagrama de comportamiento en UML que muestra la evolución de un objeto o varios objetos en función del tiempo.  
Se utiliza para representar cómo cambian los estados, valores o interacciones de un sistema a lo largo de un intervalo temporal.

Se parece a un diagrama de secuencia, pero en lugar de centrarse en los mensajes entre objetos, se enfoca en la duración y el orden cronológico de los estados o eventos.

---

## Estructura

Un diagrama de tiempos incluye:

- **Línea de vida** → Representa un objeto o participante, dispuesta en sentido vertical.  
- **Eje de tiempo** → Generalmente en el eje horizontal, indicando la progresión del tiempo.  
- **Estados o valores** → Condiciones del objeto representadas a lo largo de la línea de vida.  
- **Duraciones** → Intervalo de tiempo en el que un objeto permanece en un estado.  
- **Eventos** → Sucesos que provocan un cambio de estado o valor en un objeto.  
- **Restricciones temporales** → Reglas sobre duración o tiempos de transición (ejemplo: “máximo 5 segundos”).  
- **Sincronización** → Indica la coordinación entre varios objetos o procesos en el tiempo.  

---

## Principales características

- Representa el cambio de estado o valor en función del tiempo.  
- Útil para sistemas sensibles al tiempo (tiempo real, embebidos, telecomunicaciones, multimedia).  
- Permite visualizar restricciones temporales y tiempos máximos/mínimos.  
- Se puede usar para analizar rendimiento y sincronización de procesos.  

---

## Elementos principales

- **Eje de tiempo** → Marca la dirección del tiempo (horizontal o vertical según convenga).  
- **Objeto o participante** → Línea de vida que muestra la evolución de un elemento.  
- **Estado/Valor** → Situación en la que se encuentra un objeto en un instante dado.  
- **Evento** → Punto en el tiempo que provoca un cambio de estado.  
- **Duración** → Intervalo de tiempo que un objeto permanece en un estado.  
- **Restricción de tiempo** → Condiciones de permanencia (ejemplo: “≤ 3 seg”).  

---

## Conexiones

### 1. Línea de vida con estados
Un objeto cambia de estado en el tiempo.  

**Ejemplo:**
Un semáforo:  
**Verde (30s) → Amarillo (5s) → Rojo (40s).**

---

### 2. Sincronización entre objetos
Muestra cómo dos objetos deben coordinarse en el tiempo.  

**Ejemplo:** 
Un cliente espera respuesta del servidor en **máximo 2 segundos**.

---

### 3. Restricciones temporales
Se indican reglas sobre cuánto puede durar un estado o transición.  

**Ejemplo:**  
Un usuario debe autenticarse en **menos de 15 segundos** o la sesión expira.

---

## Ventajas

- Ideal para analizar el rendimiento temporal de un sistema.  
- Facilita el diseño de sistemas de tiempo real.  
- Útil para detectar cuellos de botella o retrasos.  
- Complementa a los diagramas de secuencia mostrando la variable tiempo.  

---

## Desventajas

- Menos usado que otros diagramas UML (más técnico).  
- Puede volverse complejo con muchos objetos y restricciones temporales.  
- Requiere conocimientos técnicos en sistemas de tiempo real para interpretarlo bien.  

---

## Conclusión

El **diagrama de tiempos** es clave en la ingeniería de software cuando el factor tiempo es determinante.  
Permite modelar **duraciones, sincronizaciones y restricciones temporales** de objetos en un sistema, siendo fundamental en áreas como:

- sistemas embebidos,  
- telecomunicaciones,  
- procesos industriales,  
- aplicaciones multimedia y en tiempo real.  

Aunque no es tan común en proyectos generales, es muy poderoso para verificar que los procesos cumplen con **límites de tiempo críticos**.