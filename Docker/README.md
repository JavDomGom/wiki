# Docker

## Concepto

Docker es un software que se utiliza principalmente para desplegar aplicaciones en espacios virtualizados. Estos espacios son conocidos como contenedores. Su uso está muy extendido y se utiliza mucho en equipos de desarrollo para desplegar de forma fácil y rápida las aplicaciones con las que trabajan sin interferir directamente en el resto de la computadora. Se podría decir que permite crear pequeñas computadoras efímeras, de usar y tirar.

En los siguientes puntos explicaré en detalle el funcionamiento y la utilización de Docker con algunos ejemplos.

## Máquina virtual Vs. Contenedores Docker

Hasta no hace mucho tiempo la manera de desplegar microservicios o pequeñas aplicaciones de forma independiente en una misma máquina era utilizando máquinas virtuales. Una máquina virtual es realmente una instalación de un sistema operativo dentro de otro sistema operativo, y para ello se utiliza un software llamado Hipervisor, que es quien nos permite poder tener múltiples sistemas en una misma máquina con independencia de recursos como pueden ser disco memoria, procesos, etc.

<p align="center"><img src="img/docker_00.png"></p>
<br>

Este tipo de instalaciones mediante un hipervisor son bastante pesadas y difíciles de mantener y administrar, pues se trata de la instalación de uno o más sistemas operativos huéspedes al completo dentro de un sistema operativo anfitrión. Este sistema operativo anfitrión deberá ser una máquina muy potente y con muchos recursos.

Docker propone un sistema en el que no se tenga que instalar un sistema operativo completo dentro de otro, sino aprovechar las características y recursos que ya están disponibles en la máquina anfitriona y compartirlos a través de una pieza fundamental llamada `Docker Engine`, es el motor de Docker, que no es más que un servicio corriendo en la máquina anfitriona en el que podremos desplegar los diferentes contenedores que solo compartirán la misma base del sistema operativo anfitrión.

Obviamente, dentro de cada contenedor encontraremos sus propias características, sus propias librerías, y una serie de recursos que podremos compartir para hacerlos visibles a otros contenedores o a la máquina anfitriona.


## Instalación de Docker

Para instalar Docker primero deberemos actualizar la información sobre paquetes que tenemos en nuestros repositorios locales mediante el siguiente comando:

```bash
~$ sudo apt-get update
Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
Get:2 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]                             
Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]                                     
Get:4 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]        
Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,150 kB]         
Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [923 kB]             
Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [435 kB]       
Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [798 kB]     
Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [721 kB]
Get:10 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [872 kB]
Get:11 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [665 kB]
Get:12 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [328 kB]
Fetched 7,332 kB in 2s (3,523 kB/s)                             
Reading package lists... Done
```

A continuación debemos instalar algunos paquetes adicionales que nos harán falta como dependencias para permitir que su sistema operativo acceda a los repositorios de Docker a través de HTTPS. Para ello ejecutaremos el siguiente comando:

```bash
~$ sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
Reading package lists... Done
Building dependency tree       
Reading state information... Done
ca-certificates is already the newest version (20170717~16.04.2).
ca-certificates set to manually installed.
curl is already the newest version (7.47.0-1ubuntu2.14).
software-properties-common is already the newest version (0.96.20.9).
software-properties-common set to manually installed.
The following NEW packages will be installed:
  apt-transport-https
0 upgraded, 1 newly installed, 0 to remove and 21 not upgraded.
Need to get 26.7 kB of archives.
After this operation, 216 kB of additional disk space will be used.
Do you want to continue? [Y/n] Y
Get:1 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 apt-transport-https amd64 1.2.32ubuntu0.1 [26.7 kB]
Fetched 26.7 kB in 0s (265 kB/s)              
Selecting previously unselected package apt-transport-https.
(Reading database ... 89418 files and directories currently installed.)
Preparing to unpack .../apt-transport-https_1.2.32ubuntu0.1_amd64.deb ...
Unpacking apt-transport-https (1.2.32ubuntu0.1) ...
Setting up apt-transport-https (1.2.32ubuntu0.1) ...
```

Aquí una breve descripción sobre cada uno de los paquetes que acabamos de instalar:

* **apt-transport-https**: Permite que el administrador de paquetes transfiera archivos y datos a través de HTTPS.
* **ca-certificates**: Permite que el sistema (y el navegador web) verifiquen los certificados de seguridad.
* **curl**: Esta es una herramienta muy popular para transferir datos.
* **software-properties-common**: Agrega scripts para administrar software.

Para asegurarnos de que el software que estamos instalando es auténtico verificaremos la integridad del mismo usando GPG mediante el siguiente comando:

```bash
~$ sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
OK
```

Ahora tenemos que añadir el repositorio oficial de Docker a nuestra lista de repositorios, de ese modo todo el software que exista en dichos repositorios estará disponible para ser instalado cuando lo necesitemos:

```bash
~$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu  $(lsb_release -cs)  stable"
```

Como hemos añadido un nuevo repositorio ahora debemos actualizar de nuevo la base de datos donde se incluye todo el software que hay en un repositorio. Para ello ejecutamos el siguiente comando:

```bash
~$ apt-get update
Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
Hit:2 http://archive.ubuntu.com/ubuntu xenial-updates InRelease         
Get:3 https://download.docker.com/linux/ubuntu xenial InRelease [66.2 kB]
Hit:4 http://archive.ubuntu.com/ubuntu xenial-backports InRelease                                    
Hit:5 http://security.ubuntu.com/ubuntu xenial-security InRelease       
Get:6 https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages [13.6 kB]
Fetched 79.8 kB in 0s (139 kB/s)                        
Reading package lists... Done
```

Ahora es el momento de instalar la última versión de Docker mediante el siguiente comando:

```bash
~$ sudo apt-get install docker-ce
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  aufs-tools cgroupfs-mount containerd.io docker-ce-cli git git-man iptables liberror-perl libgdbm3 libltdl7 libnfnetlink0 libperl5.22 patch perl perl-modules-5.22 pigz rename rsync
Suggested packages:
  mountall git-daemon-run | git-daemon-sysvinit git-doc git-el git-email git-gui gitk gitweb git-arch git-cvs git-mediawiki git-svn ed diffutils-doc perl-doc libterm-readline-gnu-perl
  | libterm-readline-perl-perl make
The following NEW packages will be installed:
  aufs-tools cgroupfs-mount containerd.io docker-ce docker-ce-cli git git-man iptables liberror-perl libgdbm3 libltdl7 libnfnetlink0 libperl5.22 patch perl perl-modules-5.22 pigz rename rsync
0 upgraded, 19 newly installed, 0 to remove and 21 not upgraded.
Need to get 95.7 MB of archives.
After this operation, 449 MB of additional disk space will be used.
Do you want to continue? [Y/n] Y
Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 libgdbm3 amd64 1.8.3-13.1 [16.9 kB]
Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libnfnetlink0 amd64 1.0.1-3 [13.3 kB]
Get:3 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 perl-modules-5.22 all 5.22.1-9ubuntu0.6 [2,629 kB]
…
…
…
Setting up libltdl7:amd64 (2.4.6-0.1) ...
Setting up patch (2.7.5-1ubuntu0.16.04.2) ...
Setting up rename (0.20-4) ...
update-alternatives: using /usr/bin/file-rename to provide /usr/bin/rename (rename) in auto mode
Processing triggers for libc-bin (2.23-0ubuntu11) ...
Processing triggers for systemd (229-4ubuntu21.27) ...
Processing triggers for ureadahead (0.100.0-19.1) ...
```

Para asegurarnos de que hemos instalado correctamente Docker podremos ejecutar el siguiente comando que nos devolverá la versión instalada:

```bash
~$ sudo docker --version
Docker version 19.03.9, build 9d988398e7
```

## Administrar Docker sin sudo
Primero hay que meter el usuario en el sudoers file. Para ello, como root se debe ejecutar el siguiente comando:

```bash
~# usermod -aG sudo alice
```

Luego añadirlo al grupo docker:

```bash
~$ sudo gpasswd -a alice docker
```

## ¿Qué es una imagen?

Una imagen Docker es una plantilla con una configuración específica que nos servirá para crear un tipo de objetos virtuales llamados contenedores, es decir un "molde" a partir del cual podremos obtener máquinas virtuales con la misma configuración, programas y características de base. Las imágenes las podemos crear o compilar nosotros mismo, pero también podemos utilizar cualquiera de las que han creado otras personas y empresas, y que están disponibles de forma pública en repositorios de imágenes como [DockerHub](https://hub.docker.com/), hay cientos de imágenes disponibles.

Por ejemplo, la imagen más típica para descargar y dar los primeros pasos con Docker es la imagen `hello-world`. Esta imagen la podremos descargar automáticamente (pull) de DockerHub y tener en nuestra lista de imágenes locales de la siguiente manera:

```bash
~$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
0e03bdcc26d7: Pull complete
Digest: sha256:6a65f928fb91fcfbc963f7aa6d57c8eeb426ad9a20c7ee045538ef34847f44f1
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

Ahora podemos obtener un listado de las imágenes que tenemos en local con el siguiente comando:

```bash
~$ docker images
REPOSITORY       TAG         IMAGE ID         CREATED           SIZE
hello-world      latest      bf756fb1ae65     4 months ago      13.3kB
```

Lo que hemos hecho realmente es instalar en nuestra máquina una imagen Docker ya compilada que está subida en DockerHub, una vez instalada la hemos creado un contenedor basado en esta imagen. Este contenedor tras mostrar el mensaje `Hello from Docker!` se termina y se apaga.

Si quisiéramos crear una imagen Docker nosotros mismos debemos compilarla en nuestra máquina local. Para ello es necesario crear un archivo llamado `Dockerfile` y escribir dentro de él una serie de instrucciones. Aquí tendríamos un ejemplo de cómo sería el contenido de un archivo `Dockerfile` bastante simple que contendría una imagen minimalista de un sistema Ubuntu:

```
FROM ubuntu:latest
```

Para compilar esta imagen en local tendríamos que ejecutar el siguiente comando :

```bash
~$ docker build .
```

Tendrá una salida de parecida a esta:

```bash
Sending build context to Docker daemon  2.048kB
Step 1/1 : FROM ubuntu:latest
latest: Pulling from library/ubuntu
d51af753c3d3: Pull complete
fc878cd0a91c: Pull complete
6154df8ff988: Pull complete
fee5db0ff82f: Pull complete
Digest: sha256:747d2dbbaaee995098c9792d99bd333c6783ce56150d1b11e333bbceed5c54d7
Status: Downloaded newer image for ubuntu:latest
 ---> 1d622ef86b13
Successfully built 1d622ef86b13
```

Al comando `docker` le pasamos la opción `build` para construir la imagen y finalmente un punto `.` que es la manera de decirle al comando `docker` que el archivo `Dockerfile` se encuentra en el mismo directorio en el que nos encontramos.

```bash
$ docker images
REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
ubuntu          latest  1d622ef86b13    4 weeks ago     73.9MB
hello-world     latest  bf756fb1ae65    4 months ago    13.3kB
```

A la hora de compilar una imagen usando un `Dockerfile` podemos usar muchas opciones, por ejemplo podemos usar la opción `-t` (tag) para etiquetar la imagen con el nombre `test_ubuntu`, luego dos puntos `:` para indicar la separación entre el nombre de la imagen y la versión, después la versión, por ejemplo `0.1.0`:

```bash
~$ docker build -t test_ubuntu:0.1.0 .
```

Obviamente son datos de ejemplo, a la imagen la podríamos haber nombrado o etiquetado con otro nombre, por ejemplo `mi_primera_imagen`, y la versión también podríamos haber puesto otra cosa como por ejemplo `v1` o `version_uno`, pero he elegido poner un nombre y una versión más acordes a las buenas prácticas. La diferencia de indicar la opción `-t` con una etiqueta con un nombre y una versión y no hacerlo es que con la etiqueta te creará una imagen Docker llamada `ubuntu` con versión `latest` y además creará una "copia" con el mismo identificador (ID) con el nombre que pusieras.

Para listar las imágenes que tenemos compiladas en local, bien si las hemos compilado a través de nuestro propio archivo `Dockerfile` o bien compilándolas desde `DockerHub`, debemos ejecutar el siguiente comando:

```bash
~$ docker images
REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
test_ubuntu     0.1.0   1d622ef86b13    4 weeks ago     73.9MB
ubuntu          latest  1d622ef86b13    4 weeks ago     73.9MB
hello-world     latest  bf756fb1ae65    4 months ago    13.3kB
```

Como podéis ver, al haber compilado la imagen indicando la etiqueta `test_ubuntu:0.1.0` se han creado aparentemente dos imágenes, pero en realidad es la misma, tienen el mismo ID `1d622ef86b13`. Si os fijáis bien se trata de una imagen minimalista muy muy reducida de un sistema Ubuntu, tiene un tamaño de tan solo `73.9MB`. Aunque pueda parecer muy poco tamaño, existen otras imágenes de sistemas GNU/Linux más reducidas todavía, como por ejemplo `Alpine`, veamos qué sucede si compilamos una imagen de un sistema `Alpine`:

1. Creamos el siguiente archivo `Dockerfile`:
    ```
    FROM apline:latest
    ```

2. Compilamos la imagen con `docker build`:
    ```bash
    ~$ docker build .
    ```

3. Listamos las imágenes una vez compiladas con `docker images`:
    ```bash
    ~$ docker images
    REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
    alpine          latest  f70734b6a266    4 weeks ago     5.61MB
    ```

Como se puede ver, la imagen de `Apline` es una imagen de tan solo `5.61MB` de tamaño, y nos puede servir como máquina GNU/Linux mínima de base para nuestros proyectos. Incluye todos los recursos mínimos para hacer funcionar un sistema operativo GNU/Linux. Si ocupa tan poco espacio es porque gracias al motor de Docker (Docker Engine) comparte algunos recursos con la máquina anfitriona, es decir nuestra computadora local donde tenemos instalado Docker.

Si quisiéramos eliminar una imagen previamente creada o descargada deberemos hacerlo de la siguiente forma:

1. Listamos las imágenes locales existentes:
    ```bash
    ~$ docker images
    REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
    ubuntu          latest  1d622ef86b13    4 weeks ago     73.9MB
    alpine          latest  f70734b6a266    4 weeks ago     5.61MB
    hello-world     latest  bf756fb1ae65    4 months ago    13.3kB
    ```

2. Si quisiéramos eliminar por ejemplo la imagen llamada `hello-world`, podríamos hacerlo de varias maneras, o bien indicando el nombre y versión de la imagen:
    ```bash
    ~$ docker rmi hello-world:latest
    ```
O bien indicando el `IMAGE ID` de la imagen, que en este ejemplo es `bf756fb1ae65`:
    ```bash
    ~$ docker rmi bf756fb1ae65
    ```
En ambos casos obtendremos el siguiente *output* por pantalla:
    ```bash
    Untagged: hello-world:latest
    Untagged: hello-world@sha256:6a65f928fb91fcfbc963f7aa6d57c8eeb426ad9a20c7ee045538ef34847f44f1
    Deleted: sha256:bf756fb1ae65adf866bd8c456593cd24beb6a0a061dedf42b26a993176745f6b
    Deleted: sha256:9c27e219663c25e0f28493790cc0b88bc973ba3b1686355f221c38a36978ac63
    ```

3. Si volvemos a listar las imágenes Docker locales existentes veremos que la imagen llamada `hello-world` ya no existe:
    ```bash
    ~$ docker images
    REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
    ubuntu          latest  1d622ef86b13    4 weeks ago     73.9MB
    alpine          latest  f70734b6a266    4 weeks ago     5.61MB
    ```

## Exportar e importar una imagen

Si quisiéramos exportar una imagen que ya tenemos compilada en local a un archivo TAR, por ejemplo para llevárnosla a otra máquina, podríamos hacerlo de la siguiente manera:

```bash
~$ docker save nombre_de_la_imagen:version > nombre_de_la_imagen_version.tar
```

Y para impiortarla desde el archivo TAR generado se puede hacer con el siguiente comando:

```bash
~$ docker load -i nombre_de_la_imagen_version.tar
```

## ¿Qué es un contenedor?

Un contenedor es una virtualización de un conjunto de programas, paquetes y librerías, instanciados a partir de una imagen Docker previamente compilada. para buscar un símil que nos suene familiar y que se pueda entender por todo el mundo, se podría decir que una imagen Docker es como un molde de galletas y los contenedores serían las diferentes galletas que puedes hacer a partir de un mismo molde. Tener muchas imágenes Docker diferentes es el equivalente a tener muchos moldes de galletas, por ejemplo en forma de estrella, de círculo, de árbol, etc. Siempre puedes hacer galletas de diferentes formas, solo tienes que utilizar el molde que mejor te venga.

Podemos iniciar o poner en funcionamiento contenedores Docker de varias maneras, bien porque tenemos la imagen ya compilada en local o bien porque Docker va a buscar la imagen al repositorio de DockerHub, donde hay centenares de imágenes ya compiladas y disponibles para usar.

Veamos un ejemplo en el que iniciamos un contenedor que no tenemos en local pero que se descargará automáticamente del repositorio de DockerHub:

```bash
~$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
0e03bdcc26d7: Pull complete
Digest: sha256:6a65f928fb91fcfbc963f7aa6d57c8eeb426ad9a20c7ee045538ef34847f44f1
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

Esto ya lo vimos en el punto anterior, como se puede ver, Docker primero va a buscar la imagen `hello-world` en local, y si no existe irá a buscarla al repositorio remoto DockerHub en internet.

Veamos ahora cómo iniciar un contenedor Docker de una imagen local, primero hay que listar las imágenes locales para ver cómo se llaman y qué versiones tenemos disponibles:

```bash
~$ docker images
REPOSITORY      TAG     IMAGE ID        CREATED         SIZE
ubuntu          latest  1d622ef86b13    4 weeks ago     73.9MB
alpine          latest  f70734b6a266    4 weeks ago     5.61MB
```

Vamos a iniciar un contenedor basándonos en la imagen de `ubuntu`, por lo tanto, podríamos hacerlo con el comando `docker run` indicando el nombre y la versión:

```bash
~$ docker run -it ubuntu:latest bash
```

O bien indicando el `IMAGE ID` que en este caso es `1d622ef86b13` de la siguiente forma:

```bash
~$ docker run -it 1d622ef86b13 bash
```

Los flags `-it` son `i` para indicar que queremos iniciar un contenedor de forma interactiva, es decir, que queremos iniciarlo y entrar dentro para trabajar en él, y `t` para indicar que queremos iniciar una sesión `tty` en consola. al final indicamos el tipo de intérprete o programa que queremos ejecutar al entrar, que en nuestro caso es `bash`. Al crear el contenedor accederemos a una shell parecida a esta:

```bash
root@5cbe59e47396:/#
```

En este caso `root` es el usuario con el que hemos entrado en este contenedor, `5cbe59e47396` es el ID del contenedor que se ha creado, cada vez será uno diferente, y `/` es el path o directorio raíz en el que nos encontramos nada más entrar.

Para ver qué contenedores que tenemos creados podremos usar el siguiente comando:

```bash
~$ docker ps -a
CONTAINER ID    IMAGE           COMMAND     CREATED         STATUS                      PORTS   NAMES
a19cedb69642    ubuntu:latest   "bash"      16 seconds ago  Exited (0) 13 seconds ago           laughing_banzai
```

El flag `-a` (all) es para mostrar no solo los contenedores que están ejecutándose, sino también los que están creados pero apagados.

En este caso tenemos un contenedor con un `STATUS` que dice `Exited (0) 13 seconds ago`. Este contenedor está creado pero en este momento no se está ejecutando. Podríamos levantarlo con el comando `docker start` del siguiente modo:

```bash
~$ docker start a19cedb69642
```

Donde `a19cedb69642` en este ejemplo es el `CONTAINER ID`, cada contenedor tiene uno diferente que se genera aleatoriamente cuando se crea el contenedor. Si ahora consultamos el estado de los contenedores como antes veremos el nuevo estado:

```bash
~$ docker ps -a
CONTAINER ID    IMAGE           COMMAND     CREATED         STATUS          PORTS   NAMES
a19cedb69642    ubuntu:latest   "bash"      3 minutes ago   Up 4 seconds            laughing_banzai
```

Para borrar o eliminar un contenedor podemos usar la instrucción `docker rm`, pero antes debe estar parado, así que primero lo paramos:

```bash
~$ docker stop a19cedb69642
a19cedb69642
```

Y una vez parado ejecutamos `docker rm` de la siguiente manera:

```bash
~$ docker rm a19cedb69642
a19cedb69642
```

## Un contenedor con una mase de datos MySQL

En este punto vamos a hacer un ejercicio un poco más complejo, vemos a crear una imagen Docker que contenga una base de datos MySQL y por otro lado haremos una imagen Docker con un sistema mínimo Alpine que se conecte a la base de datos de modo que pueda lanzar consultas (queries) y que la base de datos devuelva por pantalla los resultados.

1. Creamos la imagen Docker de MySQL directamente del servidor DockerHub de la siguiente manera:
    ```bash
    ~$ docker run -d --name mysql -e 'MYSQL_ROOT_PASSWORD=test' -d mysql:latest
    ```

2. Entramos dentro del container:
    ```bash
    ~$ docker exec -it mysql bash
    root@9961630bed6f:/#
    ```

3. Copiamos el archivo test/`setup.sql` dentro del container.

4. Ejecutamos el siguiente comando para crear la base de datos y la tabla:
    ```bash
    root@9961630bed6f:/# mysql -u root -p"test" < setup.sql
    ```
