# Diagramas Estructurales UML

Los **diagramas estructurales** en UML representan la parte **estática** de un sistema, es decir, muestran cómo está organizado y cuáles son los elementos que lo componen (clases, objetos, paquetes, componentes, nodos, etc.) y las relaciones entre ellos.

Son muy útiles en la fase de **diseño del software**, ya que permiten visualizar la **arquitectura** del sistema como una “fotografía” de sus elementos principales.

---

## Principales Características de los Diagramas Estructurales
- 📌 Representan la **estructura estática** del sistema.  
- 📌 Muestran los **elementos que existen** (clases, objetos, paquetes, componentes) y cómo se relacionan.  
- 📌 Se enfocan en la **arquitectura del sistema**, no en su comportamiento en el tiempo.  
- 📌 Son útiles para **diseñar, documentar y mantener** el sistema.  
- 📌 Facilitan la **comprensión de la organización interna** y las dependencias entre partes.  
- 📌 Se consideran como el “**esqueleto**” del sistema de software.  

---

## Tipos de Diagramas Estructurales

1. **Diagrama de Clases**  
   Representa las clases del sistema, sus atributos, métodos y relaciones.
   - *Ejemplo:* Mostrar cómo se relacionan las clases `Usuario`, `Pedido` y `Producto` en un sistema de compras.

2. **Diagrama de Objetos**  
   Muestra instancias concretas de clases y sus relaciones en un momento determinado.
   - *Ejemplo:* Un objeto `usuario:Usuario` que tiene un `pedido:Pedido` con `producto:Producto`.

3. **Diagrama de Componentes**  
   Representa los módulos o componentes de software y sus dependencias.
   - *Ejemplo:* Un sistema dividido en módulos como `Base de Datos`, `API` y `Frontend`.

4. **Diagrama de Estructura Compuesta**  
   Detalla la estructura interna de una clase o componente, mostrando cómo se organiza.
   - *Ejemplo:* Una clase `Vehículo` que internamente contiene `Motor`, `Ruedas` y `Carrocería`.

5. **Diagrama de Paquetes**  
   Organiza clases y diagramas en grupos lógicos o paquetes.
   - *Ejemplo:* Paquetes como `Controladores`, `Modelos` y `Vistas` en un proyecto MVC.

6. **Diagrama de Despliegue**  
   Muestra cómo se distribuye el software en la infraestructura de hardware (servidores, dispositivos, etc.).
   - *Ejemplo:* Un sistema desplegado en un `Servidor Web`, un `Servidor de Base de Datos` y un `Cliente`.

---

## Ventajas de los Diagramas Estructurales
- 📌 Permiten **visualizar la arquitectura** general del sistema.  
- 📌 Facilitan la **organización y documentación** del proyecto.  
- 📌 Mejoran la **comunicación entre el equipo** de desarrollo.  
- 📌 Ayudan a **detectar dependencias y relaciones** críticas entre los componentes.  
- 📌 Útiles para el **mantenimiento** y **escalabilidad** del sistema.  

## Desventajas de los Diagramas Estructurales
- ⚠️ Requieren **tiempo y esfuerzo** para su elaboración y actualización.  
- ⚠️ Si no se mantienen actualizados, pueden volverse **obsoletos rápidamente**.  
- ⚠️ Pueden ser **complejos de entender** para personas sin conocimientos técnicos.  
- ⚠️ En proyectos pequeños, su elaboración puede ser vista como una **sobrecarga innecesaria**.  

---

## Conclusión
Los diagramas estructurales UML son fundamentales para entender la **composición y organización estática** de un sistema. Aunque requieren dedicación para su creación y mantenimiento, aportan claridad y mejoran la comunicación dentro del equipo de desarrollo.
