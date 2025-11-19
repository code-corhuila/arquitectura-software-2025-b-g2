# PRIMERA PREGUNTA - Foro: ¿Qué metodología usarías en un sistema bancario?

## Introducción

La elección de una metodología de desarrollo afecta directamente el logro o fracaso de un sistema, más que todo cuando se trata de proyectos que manejan información sensible, como en este caso los sistemas bancarios. En estos entornos, la seguridad, la estabilidad y el control son factores que no se pueden dejar al azar. Por eso, es fundamental analizar qué metodología conviene aplicar, que permita entender cómo se pueden cumplir las obligaciones del sector financiero sin poner en riesgo los datos ni la confianza de los usuarios.

Teniendo en cuenta las preguntas futuras y la introducción al foro que el profesor nos proporcionó, haré principalmente una explicación de la metodología SCRUM, que ya conocía, y la metodología RUP, que nunca había escuchado, pero que creo fundamental definir antes de empezar con el desarrollo del foro.

## SCRUM

SCRUM es una metodología ágil, es decir, una forma de trabajar que busca entregar resultados rápidos y mejorar constantemente.  
Se usa mucho cuando el proyecto puede ir cambiando o ajustándose según las necesidades del cliente.  
El trabajo se divide en períodos cortos llamados “sprints”, que duran entre 2 y 4 semanas.  
En cada sprint, el equipo desarrolla una parte del sistema, la prueba y luego la mejora según los comentarios.  
No se enfoca tanto en la documentación, sino en la colaboración, la comunicación y la entrega continua.

## RUP (Rational Unified Process)

RUP es una metodología tradicional y estructurada, creada por IBM.  
Se basa en tener una planificación detallada, buena documentación y control en cada etapa del desarrollo del sistema.  
Divide el proyecto en cuatro fases principales:

- **Inicio:** se define qué se va a hacer y qué se necesita.  
- **Elaboración:** se diseña la arquitectura y se analizan los riesgos.  
- **Construcción:** se desarrolla el sistema siguiendo lo planeado.  
- **Transición:** se prueba y se entrega al usuario final.

RUP busca que todo esté documentado y probado antes de avanzar, lo que da mucha seguridad, aunque el proceso puede ser más lento que en metodologías ágiles.

## Escenario base

Un sistema bancario es una plataforma tecnológica que permite a los bancos realizar operaciones como transferencias, pagos, consultas de saldo y administración de cuentas. Para funcionar correctamente, debe ser seguro, confiable, estar siempre disponible y proteger la información de los clientes, además de cumplir con las normas legales y someterse a auditorías periódicas. Cualquier fallo puede generar pérdidas económicas o dañar la reputación de la entidad, por lo que su desarrollo requiere una metodología que garantice control, trazabilidad, planificación detallada y documentación en cada etapa.

## Pregunta 1

¿Consideras que SCRUM es adecuado para el desarrollo de un sistema bancario? ¿Por qué sí o por qué no?

Luego de investigar y leer un poco, desde mi punto de vista, SCRUM no es la metodología más adecuada para desarrollar un sistema bancario. Realmente me sorprende, ya que SCRUM es una de las metodologías más ágiles que he escuchado. Hoy en día se usan bastante, pero en este caso, este tipo de proyectos requiere una planeación estricta y un control muy riguroso en cada fase, porque maneja datos financieros que deben ser exactos y protegidos. SCRUM trabaja por ciclos cortos de entrega, donde los cambios son frecuentes y las decisiones se adaptan sobre la marcha; eso puede ser útil en otros tipos de software, pero en un sistema bancario podría generar riesgos si no se controla adecuadamente.

Mientras investigaba, encontré que la metodología que mejor se adapta a este escenario se llama RUP (Rational Unified Process), una metodología que jamás había escuchado, pero que es la más formal en este caso, porque se ajusta mejor, ya que permite definir los requisitos con detalle desde el principio, planificar bien las etapas y documentar todo el proceso. Esto ayuda a garantizar la seguridad, estabilidad y cumplimiento de las normas que exigen las instituciones financieras.

Aunque en lo que leí dice que SCRUM podría funcionar para desarrollar pequeñas partes del sistema, como interfaces o módulos de usuario, no sería adecuada para la estructura central y más importante que administra el dinero y los datos de los clientes, donde la precisión y el control son esenciales.

---

## Bibliografía

Kazman, R., Kruchten, P., Nord, R., & Tomayko, J. E. (2004, July 1). *Integrating software-architecture-centric methods into the Rational Unified Process.* CMU/SEI Technical Report CMU/SEI-2004-TR-011. Software Engineering Institute. https://doi.org/10.1184/R1/6574586.v1  

Letelay, K., Mola, S. A. S., & Go, R. (2023). *Challenges of agile software development in the banking sector: A systematic literature review.* JOIV: International Journal on Informatics Visualization, 9(1). https://doi.org/10.62527/joiv.9.1.2300  

Schwaber, K., & Sutherland, J. (2020, noviembre). *La Guía Scrum (traducida al español).* Scrum.org. https://agilenomadlife.com/wp-content/uploads/2021/07/Guia-Scrum-ORG-2020-Espanol.pdf  

Fernández, J. M., & Cadelli, S. (2014). *Convivencia de metodologías: Scrum y RUP en un proyecto de gran escala* [Tesis de grado, Universidad Nacional de La Plata]. Repositorio SEDICI. https://sedici.unlp.edu.ar/handle/10915/47082
