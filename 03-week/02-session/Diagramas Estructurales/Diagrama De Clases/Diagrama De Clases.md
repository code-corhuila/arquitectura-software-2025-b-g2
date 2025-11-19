# Diagrama de Clases UML

El **Diagrama de Clases** es uno de los **diagramas estructurales** más importantes en UML.  
Su objetivo principal es **modelar la estructura estática del sistema** mostrando las **clases**, sus **atributos**, **métodos** y las **relaciones** que existen entre ellas.

Es considerado la **base de la mayoría de los modelos UML**, ya que permite representar el diseño conceptual y lógico de un sistema orientado a objetos.

---

## Principales Características
- 📌 Representa la **estructura estática** del sistema.  
- 📌 Describe **clases, atributos, métodos** y **relaciones**.  
- 📌 Permite identificar el **modelo conceptual** y el **modelo de diseño**.  
- 📌 Es el diagrama más utilizado en UML.  
- 📌 Sirve de **puente entre el análisis y la programación**.  
- 📌 Puede usarse tanto en la **fase de diseño** como en la **fase de documentación** del software.  

---

# Diagrama de Clases en UML  

## 1. Elementos del Diagrama de Clases  

### a) Clases  
Las clases son los bloques de construcción principales. Representan un conjunto de objetos que comparten las mismas características, relaciones y semántica.  

Se representan con un **rectángulo dividido en tres secciones**:  

- **Sección Superior:** Nombre de la clase.  
- **Sección Media:** Atributos (propiedades o campos).  
  - Sintaxis:  
    ```
    visibilidad nombre: tipo [multiplicidad] = valor_inicial {propiedad}
    ```
- **Sección Inferior:** Operaciones (métodos o funciones).  
  - Sintaxis:  
    ```
    visibilidad nombre(lista_de_parámetros): tipo_de_retorno {propiedad}
    ```

#### Visibilidad (Modificadores de Acceso)
- `+` Público: Accesible desde cualquier lugar.  
- `#` Protegido: Accesible solo por la clase y sus subclases.  
- `-` Privado: Accesible solo por la propia clase.  
- `~` Paquete: Accesible dentro del mismo paquete.  

---

## 2. Estructura de un Diagrama de Clases  
La estructura se compone de las clases y sus relaciones.  
Un diagrama bien diseñado debe ser **legible y coherente**, mostrando cómo interactúan los diferentes componentes del sistema.  
No es solo una lista de clases, sino una **red interconectada**.  

---

## 3. Relaciones entre Clases  

### a) Asociación  
- Representa una relación estructural entre clases.  
- Se dibuja con una línea recta entre ellas.  
- Puede tener:  
  - **Nombre:** Describe la naturaleza de la relación (ej. *trabaja en*).  
  - **Roles:** Cómo cada clase participa en la relación (ej. *empleado* – *departamento*).  
  - **Multiplicidad:** Número de objetos relacionados:  
    - `1` → Uno y solo uno.  
    - `0..1` → Cero o uno.  
    - `*` → Cero o más.  
    - `1..*` → Uno o más.  
    - `n` → Exactamente *n*.  
    - `n..m` → De *n* a *m*.  

### b) Agregación  
- Relación **todo-parte** débil.  
- La parte puede existir independientemente del todo.  
- Se representa con un **rombo vacío** en el lado del *todo*.  

### c) Composición  
- Relación **todo-parte** fuerte.  
- La parte **no puede existir** sin el todo.  
- Si el todo se destruye, las partes también.  
- Se representa con un **rombo sólido**.  

### d) Herencia (Generalización/Especialización)  
- Representa una relación **“es un”**.  
- Una subclase hereda atributos y operaciones de la superclase.  
- Se representa con una **flecha triangular hueca** hacia la superclase.  

### e) Dependencia  
- Una clase depende de otra (cambios en una pueden afectar a la otra).  
- Relación débil.  
- Se representa con una **línea discontinua con flecha**.  

---

## 4. Ventajas de Usar Diagramas de Clases  
- **Claridad y Comprensión:** Visualizan la estructura del sistema de manera clara.  
- **Comunicación Efectiva:** Herramienta común entre desarrolladores, analistas y clientes.  
- **Base para la Implementación:** Se pueden traducir directamente a código.  
- **Detección de Errores de Diseño:** Identificación de problemas antes de codificar.  
- **Mantenimiento y Refactorización:** Documentación útil para evolución del sistema.  

---

## 5. Desventajas de Usar Diagramas de Clases  
- **Complejidad:** Pueden volverse difíciles de leer en sistemas grandes.  
- **No Muestran Comportamiento Dinámico:** Solo la estructura estática.  
- **Curva de Aprendizaje:** Requieren conocimientos en notación UML.  
- **Sobrediseño:** Riesgo de invertir demasiado tiempo en el modelado.  

---

## 6. Conclusión  
El **diagrama de clases** es una herramienta indispensable en el desarrollo de software orientado a objetos.  
Funciona como un **plano arquitectónico** que ofrece una visión clara de la estructura, componentes y relaciones.  

Aunque no refleja el comportamiento dinámico, sus ventajas en términos de comunicación, documentación y detección de errores superan sus desventajas.  
Dominarlo es fundamental para cualquier profesional en el diseño y construcción de **sistemas de software complejos y bien estructurados**.  



# Ejemplos de Relaciones en Diagramas de Clases UML

---

## Ejemplo de Asociación

**Clases: Usuario – Compra**
 **Un Usuario realiza una o más Compras**

---

## Ejemplo de Agregación

**Clases: Universidad – Departamento**
 **Una Universidad tiene uno o más Departamentos. Un Departamento puede existir sin la Universidad.**

---

## Ejemplo de Composición

**Clases: Vehículo – Motor**
 **Un Vehículo tiene un Motor. Si el Vehículo se destruye, el Motor también deja de existir.**

---

## Ejemplo de Herencia

**Clases: Persona – Estudiante**
 **Un Estudiante es una Persona.**

---

## Ejemplo de Dependencia

**Clases: Reporte – BaseDeDatos**
 **La clase Reporte depende de la clase BaseDeDatos para generar un reporte.**

---------
# Diagrama de Clases - Carrito de Compras

Este diagrama de clases modela el funcionamiento de un sistema de **carrito de compras**.  
Permite representar los usuarios que interactúan con el sistema, los productos que se añaden al carrito, la gestión de pedidos y el procesamiento de pagos.

---

## Clases y atributos principales

### 1. Usuario
- **Atributos**: `idUsuario`, `nombre`, `email`, `contraseña`.  
- **Métodos**: 
  - `registrar()` → Permite registrar un nuevo usuario.  
  - `iniciarSesion()` → Inicia la sesión del usuario.  
  - `cerrarSesion()` → Cierra la sesión.  

### 2. Producto
- **Atributos**: `idProducto`, `nombre`, `precio`, `stock`.  
- **Métodos**:  
  - `actualizarStock(cantidad:int)` → Resta o suma stock.  
  - `obtenerInfo():string` → Devuelve información del producto.  

### 3. ItemCarrito
- **Atributos**: `cantidad`, `subtotal`.  
- **Métodos**:  
  - `calcularSubtotal():double` → Calcula el subtotal del ítem (cantidad * precio).  

### 4. Carrito
- **Atributos**: `idCarrito`, `fechaCreacion`, `total`.  
- **Métodos**:  
  - `agregarProducto(p:Producto, cantidad:int)` → Añade productos al carrito.  
  - `eliminarProducto(p:Producto)` → Elimina un producto del carrito.  
  - `calcularTotal():double` → Suma los subtotales de todos los ítems.  

### 5. Pedido
- **Atributos**: `idPedido`, `fecha`, `estado`.  
- **Métodos**:  
  - `confirmar()` → Confirma el pedido.  
  - `cancelar()` → Cancela el pedido.  

### 6. Pago
- **Atributos**: `idPago`, `monto`, `metodo`.  
- **Métodos**:  
  - `procesarPago():bool` → Procesa el pago y devuelve verdadero o falso.  
  - `generarRecibo():string` → Genera un recibo del pago realizado.  

---

## Relaciones entre clases

- **Usuario → Carrito**: Un usuario puede tener un carrito.  
- **Carrito *-- ItemCarrito**: Un carrito está compuesto por varios ítems.  
- **ItemCarrito → Producto**: Cada ítem corresponde a un producto específico.  
- **Carrito → Pedido**: Del carrito se genera un pedido.  
- **Pedido → Pago**: El pedido finaliza con un pago.  

---

## Conclusión

Este diagrama de clases modela el flujo básico de un **carrito de compras en línea**.  
Permite comprender cómo un usuario interactúa con productos, los agrega a un carrito, genera un pedido y finalmente realiza un pago.  

---

