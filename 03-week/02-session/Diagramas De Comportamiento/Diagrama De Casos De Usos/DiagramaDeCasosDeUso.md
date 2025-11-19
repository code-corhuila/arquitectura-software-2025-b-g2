# Diagrama de Casos de Uso (UML)

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

---

## Principales características

- Representa las funcionalidades vistas desde el usuario, no la implementación técnica.  
- Identifica quién usa el sistema (actores) y qué espera de él (casos de uso).  
- Es ideal para la comunicación con clientes y usuarios finales.  
- Puede detallar relaciones especiales entre casos de uso como inclusión, extensión o generalización.  
- Permite tener una visión general y sencilla del sistema.  

---

## Elementos principales

- **Actores** → Usuarios o sistemas externos que interactúan con el sistema.  
  - *Ejemplo:* Cliente, Administrador, Sistema de Pagos.  

- **Casos de uso** → Funcionalidades ofrecidas por el sistema.  
  - *Ejemplo:* Registrar Pedido, Procesar Pago, Generar Reporte.  

- **Sistema** → El rectángulo que encierra todos los casos de uso.  
  - *Ejemplo:* Sistema de Ventas.  

- **Relaciones:**
  - **Asociación (línea simple)** → Conecta actores con casos de uso que utilizan.  
  - **Inclusión («include»)** → Un caso de uso siempre incluye a otro.  
    - *Ejemplo:* Procesar Pedido incluye Verificar Stock.  
  - **Extensión («extend»)** → Un caso de uso opcional amplía a otro bajo ciertas condiciones.  
    - *Ejemplo:* Pagar puede extenderse con Aplicar Descuento.  
  - **Generalización** → Un actor o caso de uso hereda de otro.  
    - *Ejemplo:* Empleado hereda de Usuario.  

---

## Conexiones

### 1. Actor → Caso de uso
- Línea simple entre el actor y el caso de uso.  
- Representa que el actor utiliza esa funcionalidad.  

**Ejemplo:**  
Cliente -------- Realizar Pedido

---

### 2. Inclusión («include»)
- Línea con estereotipo <<include>>.  
- Indica que un caso de uso siempre invoca a otro.  

**Ejemplo:**  
Procesar Pedido → incluye → Verificar Stock

---

### 3. Extensión («extend»)
- Línea con estereotipo <<extend>>.  
- Un caso de uso se extiende opcionalmente con otro bajo ciertas condiciones.  

**Ejemplo:**  
Realizar Pago → extiende → Aplicar Descuento


---

### 4. Generalización
- Línea con triángulo vacío.  
- Representa herencia entre actores o casos de uso.  

**Ejemplo:**  
Empleado <|-- Administrador



---

## Ventajas

- Permite comunicar fácilmente requisitos con usuarios no técnicos.  
- Ofrece una visión clara y simple de lo que el sistema hace.  
- Ayuda a identificar actores y funcionalidades clave.  
- Sirve como base para crear documentación más detallada (casos de uso textuales, diagramas de secuencia, etc.).  
- Muy útil para planificación de pruebas de aceptación.  

---

## Desventajas

- No muestra el orden temporal de las interacciones (para eso se usa un diagrama de secuencia).  
- Puede volverse muy grande y complejo en sistemas extensos.  
- Requiere detalle adicional en descripciones textuales para ser completamente útil.  
- No refleja la implementación técnica, solo la visión funcional.  

---

## Conclusión

El diagrama de casos de uso es una herramienta esencial para capturar y comunicar requisitos funcionales, mostrando de forma sencilla qué hace el sistema y quién lo utiliza.  
Aunque no detalla la lógica interna, es clave en las fases iniciales del desarrollo para alinear la visión de usuarios, analistas y desarrolladores.