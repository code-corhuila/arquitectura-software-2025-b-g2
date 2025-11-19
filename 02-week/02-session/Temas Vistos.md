# Sesión 2 – Clase 2: Arquitectura de Software

## Tema central
En esta clase vimos cómo distintos **estilos y patrones arquitectónicos** definen no solo el diseño conceptual del software, sino también **cómo organizar los paquetes/carpetas** dentro de un proyecto.  

El profesor explicó que no es lo mismo **arquitectura de software** (visión global) que **patrones arquitectónicos** o **estructurales**, pero que la arquitectura sí guía la manera en que organizamos los módulos, repositorios y paquetes.  

---

## Arquitecturas y patrones estudiados

### 1. Arquitectura Monolítica
- **Definición:** toda la aplicación en un único artefacto o ejecutable.
- **Características:** acoplamiento alto, despliegue conjunto, escalabilidad vertical.  
- **Organización en paquetes:** se puede organizar en `frontend/`, `backend/`, `docs/`, `database/`.  
- **Ejemplo:** un sistema web donde backend y vistas están dentro del mismo proyecto.  
- **Nota:** si el frontend es separado (SPA), ya entra el modelo cliente-servidor.

---

### 2. Arquitectura en Capas
- **Definición:** organiza el software en niveles que se comunican de forma jerárquica.  
- **Modelos comunes:**  
  - 3 capas: **Presentación → Aplicación/Negocio → Datos**  
  - 4 capas (DDD/Clean): **Presentación → Aplicación → Dominio → Infraestructura**  
- **Organización en paquetes:** cada módulo/feature puede tener subpaquetes por capa:  
  `users/app/`, `users/domain/`, `users/infra/`.  
- **Evitar:** mezclar repositorios de BD dentro del dominio.

---

### 3. Patrón MVC (Model-View-Controller)
- **Definición:** separación de responsabilidades en **modelo, vista y controlador**.  
- **Modelo:** estado + lógica de negocio.  
- **Vista:** interfaz gráfica / UX.  
- **Controlador:** recibe solicitudes, invoca lógica, selecciona vista.  
- **Uso común:** aplicaciones web monolíticas o frameworks frontend (Angular, Vue, React).  
- **Ambigüedad a evitar:** el modelo no es solo “lógica de datos” ni solo “tablas de BD”.

---

### 4. Arquitectura Hexagonal (Puertos y Adaptadores)
- **Definición:** separar el núcleo de negocio de la infraestructura externa.  
- **Núcleo:** dominio (entidades) + casos de uso (aplicación).  
- **Puertos:** interfaces que definen entradas/salidas.  
- **Adaptadores:**  
  - Entrada: UI, API REST.  
  - Salida: bases de datos, servicios externos.  
- **Organización en paquetes:**  
- domain/ // entidades y lógica
- app/ // casos de uso
- ports/in/, ports/out/
- adapters/in/http, adapters/out/db
- infra/ // configuración

- **Nota:** el “front” aquí se entiende como adaptador de entrada.

---

### 5. Arquitectura Cliente-Servidor
- **Definición:** modelo de comunicación donde un cliente consume servicios de un servidor.  
- **Cliente:** frontend (navegador o app).  
- **Servidor:** backend (API, lógica de negocio).  
- **Organización:** proyectos separados, ej.:  
- `frontend/ (src, components, views, main.js)`  
- `backend/ (src, controllers, services, models, router, main.js)`  
- **Clave:** es un modelo de comunicación, no dicta cómo organizar internamente el backend (que puede ser monolítico, hexagonal, por capas, etc.).

---

### 6. Arquitectura Orientada a Servicios (SOA)
- **Definición:** aplicaciones compuestas por servicios autónomos con interfaces bien definidas.  
- **Relación:** base conceptual de los microservicios.  
- **Organización en repos:** cada servicio puede tener sus propias capas/hexagonal.

---

## Conexiones clave
- **Arquitectura de software** = decisión global (monolito, SOA, microservicios, cliente-servidor).  
- **Patrón arquitectónico** = solución probada para organizar internamente (MVC, hexagonal, capas).  
- **Patrón estructural** = cómo organizar clases/componentes (Adapter, Composite, también “capas”).  
- **Paquetes/carpetas** = reflejan los límites de la arquitectura; no son la arquitectura en sí.  

---

## Actividad en clase
 *“¿Se pueden aplicar capas, MVC, hexagonal y cliente-servidor dentro de un software monolítico?”*  

- **Respuesta:**  
- Sí, un monolito puede organizarse por capas o usar MVC internamente.  
- También puede aplicar hexagonal (núcleo y adaptadores) aunque siga siendo un solo artefacto.  
- Cliente-servidor solo aplica si hay separación de frontend y backend como artefactos distintos.  

---

## Conclusión de la sesión
- Un **monolito** no significa “desorden”, puede tener **capas, MVC o hexagonal** dentro.  
- **Cliente-servidor** no es una organización interna, sino cómo se comunican frontend y backend.  
- La **organización de paquetes** debe seguir las fronteras de la arquitectura:  
- No mezclar dominio con infraestructura.  
- Evitar acoplar frontend y backend en un mismo paquete.  
- La arquitectura elegida marca los límites; los paquetes los hacen visibles y respetables.