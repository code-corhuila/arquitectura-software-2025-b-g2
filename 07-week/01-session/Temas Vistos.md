# Sesión 1 - Semana 7  

En la **sesión 1 de la semana 7**, el profesor abordó el tema de la **arquitectura en N-Capas**, haciendo énfasis en dos enfoques principales para organizar un proyecto:  

- **N-Capas by Module**  
- **N-Capas All Project**  

## Contenido de la clase  

1. **Explicación teórica**  
   - Se describió qué significa trabajar en una arquitectura **N-Capas** y cómo esta busca separar responsabilidades en diferentes capas, tales como *controller*, *service*, *repository* y *model/entity*.  
   - El profesor explicó cómo estas capas se comunican entre sí, siguiendo un flujo organizado y reduciendo el acoplamiento entre componentes.  

2. **Escenario N-Capas by Module**  
   - Se mostró cómo en este enfoque cada **módulo funcional** (ejemplo: *Producto*, *Usuario*, *Pedido*) contiene internamente todas sus capas: `controller`, `service`, `repository`, `entity`.  
   - El profesor utilizó **árboles de directorios** y ejemplos prácticos para ilustrar cómo se organiza el código en módulos independientes pero consistentes.  

3. **Escenario N-Capas All Project**  
   - Se presentó el enfoque tradicional en el que la aplicación se organiza **por capas globales**.  
   - El proyecto tiene carpetas únicas para controladores, servicios, repositorios y entidades, y dentro de ellas se agrupan todas las funcionalidades.  
   - Con ejemplos y diagramas de estructura, se mostró cómo los módulos comparten estas capas de manera centralizada.  

## Conclusión de la sesión  

La clase brindó una visión completa de cómo organizar proyectos en **N-Capas**, entendiendo no solo la teoría, sino también con ejemplos prácticos y representaciones visuales de las estructuras de carpetas.  
Gracias a esto, se pudo comprender mejor la diferencia entre **estructuras centralizadas (All Project)** y **estructuras modulares (By Module)** dentro de un proyecto de software.  
