## **Sesión 2 - Semana 13: Jenkins y Docker**

En esta sesión el profesor **Jesús Ariel González Bonilla** nos explicó cómo utilizar **Jenkins** y **Docker** para automatizar el despliegue de aplicaciones.  
- El objetivo fue comprender el proceso completo para **construir una imagen de Jenkins**, configurarla dentro de un contenedor Docker** y **usarla para ejecutar pipelines de CI/CD**.

---

## **Contenido de la clase**

Durante la clase, el profesor nos enseñó los pasos necesarios para levantar una instancia de **Jenkins** utilizando **Docker**, desde la limpieza de instalaciones previas hasta el acceso a la interfaz web de Jenkins.  

También explicó cómo editar permisos, acceder como usuario root y configurar `sudo` dentro del contenedor para permitir que Jenkins pueda ejecutar comandos del sistema sin restricciones.

---

## **Pasos para montar la imagen de Jenkins en Docker**

```bash
# Paso 1: Limpiar instalación anterior (si existe)
docker stop jenkins
docker rm jenkins
docker volume rm jenkins_home
docker rmi jenkins

# Paso 2: Construir la imagen de Docker
docker build -t jenkins .

# Paso 3: Cambiar permisos en el volumen (opcional)
docker run --rm -v jenkins_home:/var/jenkins_home alpine chown -R 1000:1000 /var/jenkins_home

# Paso 4: Ejecutar el contenedor de Jenkins
docker run -d -p 8080:8080 -p 50000:50000 --name jenkins --privileged -v jenkins_home:/var/jenkins_home -v /var/run/docker.sock:/var/run/docker.sock jenkins

# Paso 5: Verificar que el contenedor está corriendo
docker ps -f name=jenkins

# Paso 6: Ver logs para confirmar el inicio
docker logs jenkins --tail 20

# Paso 7: Obtener la clave de administrador
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

## Con esta clave se puede acceder por primera vez al panel web de Jenkins:

- **URL:** [http://localhost:8080](http://localhost:8080)  
- **Usuario:** admin  
- **Contraseña inicial:** clave generada en el paso anterior  

---

## **Dockerfile utilizado**

```dockerfile
FROM jenkins/jenkins:lts-jdk17

USER root

RUN apt-get update && apt-get install -y --no-install-recommends       ca-certificates curl gnupg lsb-release sudo   && install -m 0755 -d /etc/apt/keyrings   && curl -fsSL https://download.docker.com/linux/debian/gpg      | gpg --dearmor -o /etc/apt/keyrings/docker.gpg   && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg]      https://download.docker.com/linux/debian $(lsb_release -cs) stable"      > /etc/apt/sources.list.d/docker.list   && apt-get update   && apt-get install -y --no-install-recommends docker-ce-cli docker-compose-plugin   && rm -rf /var/lib/apt/lists/*

RUN groupadd -f docker && usermod -aG docker jenkins   && echo "jenkins ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

USER jenkins

EXPOSE 8080 50000
```

---

## **Repositorios de referencia**

El profesor también compartió **tres repositorios en GitHub** para que los revisáramos y entendiéramos cómo se implementan los pipelines en Jenkins y la integración con proyectos reales:

1. [backend-security (Jenkinsfile)](https://github.com/code-corhuila/backend-security/blob/main/Jenkinsfile)  
2. [backend-security (carpeta principal)](https://github.com/code-corhuila/backend-security/tree/main)  
3. [backend-ubication](https://github.com/code-corhuila/backend-ubication)

---

## **Ejercicio práctico y tarea**

Además, nos compartió una **carpeta con un ejercicio** que podíamos usar para investigar y practicar cómo funciona el despliegue de proyectos mediante Jenkins y Docker.  

## Como **tarea**, nos indicó:
- Analizar el contenido de los repositorios enviados.  
- Documentar todo lo que se encuentre dentro de ellos (estructura, Jenkinsfile, scripts, etc.).  
- Preparar un informe o resumen con las observaciones y aprendizajes obtenidos.

---

## **Conclusión**

Esta sesión fue muy práctica y esencial para comprender cómo **automatizar procesos de integración continua** con Jenkins y Docker.  
Aprendimos a construir imágenes personalizadas, ejecutar contenedores y conectar Jenkins con el sistema anfitrión para manejar pipelines de despliegue de forma eficiente.

https://drive.google.com/drive/folders/15h1frurgsiC3p3Ay9moVSnruevGaOLGd?usp=sharing