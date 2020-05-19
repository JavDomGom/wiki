# Sistema GNU/Linux

## Breve introducción al sistema operativo

GNU, también conocido como GNU/Linux o Linux sin más, es un sistema operativo de tipo Unix, así como una gran colección de programas informáticos que componen al sistema, desarrollado por Richard Stallman para el Proyecto GNU y auspiciado por la Free Software Foundation. Está formado por software libre, mayoritariamente bajo términos de copyleft. GNU es el acrónimo recursivo de "GNU's Not Unix" (en español: GNU no es Unix), nombre elegido debido a que GNU sigue un diseño tipo Unix y se mantiene compatible con éste, pero se distingue de Unix por ser software libre y por no contener código de Unix.

El sobrenombre Linux es realmente el nombre del kernel o núcleo del sistema, es el principal responsable de facilitar a los distintos programas acceso seguro al hardware de la computadora o en forma básica, es el encargado de gestionar recursos, a través de servicios de llamada al sistema. Muchas personas se refieren al sistema operativo simplemente con el nombre “Linux”, y aunque es aceptado realmente es inexacto e injusto, pues el nombre del sistema es realmente GNU, pero se acepta igualmente llamarlo GNU/Linux, ambos forman un sistema operativo completo y funcional.

## Acceso al sistema

### Protocolo SSH

Una de las grandes ventajas de utilizar la terminal es que podemos acceder a terminales en otros servidores muy fácilmente. El protocolo más utilizado para acceder a terminales de forma remota es SSH (Secure SHell). El protocolo SSH tiene un gran número de posibilidades, pero el uso más habitual es utilizarlo para abrir terminales en servidores remotos que tienen un servicio SSH. Este protocolo SSH es seguro porque cifra las comunicaciones entre el cliente y el servidor. Se diseñó como una alternativa segura a telnet. No debemos usar el protocolo Telnet porque las comunicaciones en Telnet, incluidas las claves de acceso, no están cifradas y cualquiera puede tener acceso a ellas.
Para acceder a una computadora que implemente el protocolo SSH podemos usar el programa ssh, pero previamente tenemos que tener una cuenta en esa computadora. Imaginemos que *Alice* tiene una cuenta en un servidor que tiene un servicio SSH. Para conectarse puede hacer:

```bash
~$ ssh alice@servidor
```

En este ejemplo podemos cambiar servidor por la dirección IP o el nombre de la computadora a la que nos queremos conectar. Si el nombre de la cuenta de usuario (alice) en el servidor cliente y en el servidor es el mismo podemos obviar el nombre de usuario, de modo que se podría hacer la conexión SSH de la siguiente manera:

```bash
~$ ssh servidor
```
A continuación el servidor nos pedirá la clave correspondiente al usuario con el que nos queremos conectar.

Podemos incorporar algunas opciones adicionales como `ServerAliveInterval` para indicar un número de segundos que el cliente esperará antes de enviar un paquete nulo al servidor con el objetivo de mantener viva la conexión, esto puede venirnos muy bien en aquellas conexiones a servidores en los que si no realiza ninguna acción durante un periodo de tiempo se pierde la conexión. Un ejemplo de uso sería:

```bash
~$ ssh -o ServerAliveInterval=10 alice@servidor
```

En este ejemplo se enviará un paquete nulo cada 10 segundos, para que no se pierda la conexión automáticamente tras un tiempo. Es una manera de indicarle al servidor que seguimos ahí.

Existen clientes SSH para Windows con los que nos podemos conectar. Uno muy común es [PuTTY](https://www.putty.org/) y otro un poco más completo es [MobaXterm](https://mobaxterm.mobatek.net/).

## Manejo del sistema a través de la Shell o terminal

La Shell (o terminal) es un intérprete de comandos. Es simplemente un modo alternativo de controlar un servidor basado en una interfaz de texto. La terminal nos permite ejecutar software escribiendo el nombre del programa que queremos ejecutar en la terminal. Podemos pedirle al servidor que ejecute un programa mediante el ratón haciendo click en distintos lugares del escritorio o podemos escribir una orden para conseguir el mismo objetivo. Ninguna de las dos formas de comunicarse con el servidor es mejor que la otra aunque en ciertas ocasiones puede resultar más conveniente utilizar una u otra. Las ventajas de la línea de comandos son:

* **Necesidad**: Existe mucho software que está sólo disponible en la terminal. Esto es especialmente cierto en el área de la administración de sistemas y ciberseguridad.
* **Flexibilidad**: Los programas gráficos suelen ser muy adecuados para realizar la tarea para la que han sido creados, pero son difíciles de adaptar para otras tareas. Los programas diseñados para ser usados en la línea de comandos suelen ser muy versátiles.
* **Reproducibilidad**: Documentar y repetir el proceso seguido para realizar un análisis con un programa gráfico es muy costoso puesto que es difícil describir la secuencia de clicks y doble clicks que hemos realizado. Por el contrario, los procesos realizados mediante la línea de comandos son muy fáciles de documentar puesto que tan sólo debemos guardar el texto que hemos introducido en la pantalla.
* **Fiabilidad**: Los programas básicos de Unix fueron creados en los años 70 y han sido probados por innumerables usuarios por lo que se han convertido en piezas de código extraordinariamente confiables.
* **Recursos**: Las interfaces gráficas suelen consumir muchos recursos mientras que los programas que funcionan en línea de comandos suelen ser extraordinariamente livianos y rápidos. Este poco uso de recursos facilita, por ejemplo, que se utilice a través de la red.

El problema de la terminal es que para poder utilizarla debemos saber previamente qué queremos hacer y cómo. Es habitual descubrir cómo funciona un programa con una interfaz gráfica sin tener que leer un manual, esto no sucede en la terminal.

Para usar la línea de comandos hay que abrir una terminal. Se abrirá una ventana con un mensaje similar a:

```
usuario $
```

Este pequeño mensaje se denomina prompt y el cursor parpadeante que aparece junto al él indica que el servidor está esperando una orden. El mensaje exacto que aparece en el prompt puede variar ligeramente, pero en sistemas GNU/Linux suele ser similar a:

```
usuario@servidor:~$
```

En el prompt se nos muestra el nombre del usuario, el nombre del servidor y el directorio en el que nos encontramos actualmente, es decir, el directorio de trabajo actual. La virgulilla “~” significa que estás en tu home o en tu carpeta de usuario. Finalmente aparece el símbolo del dólar “$” como final del prompt y luego el cursor que parpadea, desde donde podremos escribir comandos. Cuando el prompt se muestra podemos ejecutar cualquier comando.

## Comandos básicos

### LS

Mediante el comando ls (LiSt) podemos pedir que liste el contenido del directorio en el que nos encontramos:

```bash
~$ ls folder_name
archivo_1.txt archivo_2.txt directorio_1 directorio_2  directorio_3
```

El comando ls, como cualquier otro comando, es en realidad un programa que el servidor ejecuta. Cuando escribimos la orden (y pulsamos intro) el programa se ejecuta. Mientras el programa está ejecutándose el prompt desaparece y no podemos ejecutar ningún otro comando. Pasado el tiempo el programa termina su ejecución y el prompt vuelve a aparecer. En el caso del comando ls el tiempo de ejecución es tan pequeño que suele ser imperceptible. Los programas suelen tener unas entradas y unas salidas. Dependiendo del caso estas pueden ser ficheros o caracteres introducidos o impresos en la pantalla. Por ejemplo, el resultado de ls es simplemente una lista impresa de ficheros y directorios en la interfaz de comandos.

### MKDIR

El comando mkdir (make directory) creará un directorio nuevo con el nombre que le pasemos como argumento, por ejemplo, para crear un directorio nuevo llamado `directorio_1` podremos ejecutarlo de la siguiente manera:

```bash
~$ mkdir directorio_1
```

También podremos ejecutar el comando `mkdir` con la opción o *flag* `-p` para crear directorios intermedios, por ejemplo:

```bash
~$ mkdir -p directorio_1/directorio_2/directorio_3
```

Y si queremos crear directorios con un permiso previamente especificado podremos hacerlo utilizando la opción `-m` seguido de los permisos y el nombre del directorio, por ejemplo:

```bash
~$ mkdir -m 600 directorio_1
```

```bash
~$ ls -lrt
drw------- 2 alice alice    4096 May 14 10:35 directorio_1
```

Los permisos de los archivos y directorios es un tema que se abordará en detalle más adelante en la sección *Gestión de usuarios y permisos*.

### TOUCH

Con el comando touch podremos crear archivos vacíos para su posterior edición, por ejemplo:

```bash
~$ touch archivo_1.txt
```

El ejemplo anterior generará un archivo llamado `archivo_1.txt` sin contenido. Una manera alternativa de hacerlo sin usar el comando `touch` podría ser la siguiente:

```bash
~$ > touch archivo_1.txt
```

### CD

El comando cd (change directory) nos servirá para cambiar de directorio. Por ejemplo, si queremos entrar en el directorio `directorio_1` podremos hacerlo de la siguiente manera.

```bash
~$ cd directorio_1
~/directorio_1$
```

Si quisiéramos cambiarnos al directorio `subdirectorio_1` volveríamos a ejecutar el comando cd pero esta vez indicando el nombre `subdirectorio_1`:

```
~/directorio_1$ cd subdirectorio_1/
~/directorio_1/subdirectorio_1$
```

Si quisiéramos volver un directorio para atrás debemos escribir dos puntos (`..`) de la siguiente manera:

```
~/directorio_1/subdirectorio_1$ cd ..
~/directorio_1$
```

También podemos utilizar rutas absolutas, por ejemplo:

```bash
~$ cd directorio_1/subdirectorio_1/
~/directorio_1/subdirectorio_1$
```

### RM y RMDIR

El comando `rm` nos permite eliminar tanto ficheros como directorios del sistema. Hay que tener cuidado, ya que si borramos un fichero o directorio con este comando los elementos borrados no irán a una papelera de reciclaje como cuando se elimina de forma gráfica pulsando la tecla de borrar o mediante un menú interactivo con el ratón. Eliminar un archivo o directorio con `rm` supone eliminarlo del sistema y no se podrá recuperar (fácilmente). Para eliminar un fichero basta con ejecutar el siguiente comando:

```bash
~$ rm archivo_1.txt
```

Si queremos eliminar un directorio podremos hacerlo con el comando `rm` deberemos especificar el flag `-r` del siguiente modo:

```bash
~$ rm -r directorio_1
```

Otro modo de eliminar un directorio es utilizando el comando `rmdir` de la siguiente manera:

```bash
~$ rmdir directorio_1
```

En ambos casos se eliminará el directorio y todo su contenido de forma recursiva. Son comandos que hay que **utilizar con extremo cuidado**, podríamos perder información valiosa y podría ser muy complicado recuperarla.

### MV

El comando `mv` (MoVe) nos servirá para mover un fichero o directorio de ubicación y también para renombrarlos. Por ejemplo, si queremos el fichero `archivo_1.txt` desde la carpeta `/home/alice` a la carpeta `/tmp` podremos hacerlo de la siguiente manera:

```bash
~$ mv /home/alice/archivo_1.txt /tmp
```

Lo mismo con el directorio `directorio_1`:

```bash
~$ mv /home/alice/directorio_1 /tmp
```

Si queremos renombrar un fichero o directorio basta con especificar primero la ubicación del elemento a renombrar y después especificar la misma ubicación pero con el nuevo nombre, por ejemplo:

```bash
~$ mv /home/alice/archivo_1.txt /home/alice/nuevo_nombre_del_archivo.txt
```

Si especificamos diferentes rutas en origen y destino se moverá el archivo o directorio y también se le cambiará el nombre, por ejemplo:

```bash
~$ mv /home/alice/archivo_1.txt /tmp/nuevo_nombre_del_archivo.txt
```

### PWD

El comando `pwd` (Print Working Directory) nos imprimirá por pantalla la ruta absoluta del directorio en el que nos encontramos actualmente, por ejemplo:

```bash
~$ pwd
/home/alice
```

### FIND

El comando `find` nos servirá para buscar archivos y directorios que cumplan una serie de requisitos como puede ser un nombre exacto, un patrón en el nombre, unos permisos concretos, un tamaño, fecha de creación, antigüedad, ubicación, etc específicos. A continuación algunos ejemplos de uso:

Buscar archivos con extensión `.txt` en una profundidad máxima de dos directorios:

```bash
~$ find /home/alice/* -maxdepth 2 -name "*.txt" -type d
```

Buscar archivos con que su nombre comience por `archivo` en una profundidad de un único directorio:

```bash
~$ find /home/alice/* -prune -name "archivo*" -type d
```

Buscar un fichero llamado `mi_archivo` en el directorio `/tmp/example`:

```bash
~$ find /tmp/example -name "mi_archivo"
```

Buscar en todos los directorios a partir del directorio actual:

```bash
~$ find . -name "mi_archivo"
```

Omitir todos los mensajes en los que dice que no tenemos permiso para buscar en esos directorios hay que ejecutar:

```bash
~$ find . -name 2>/dev/null "nombre_del_archivo"
```

Buscar archivos con dos patrones de búsqueda en el nombre, por ejemplo todos los archivos que tengan extensión `.txt` y `.pdf`:

```bash
~$ find . \( -name "*.txt" -o -name "*.pdf" \) -type f -ls
```

Buscar archivos con más de tres días de antigüedad desde su creación:

```bash
~$ find . -name "*.txt" -mtime +3 -type f
```

Buscar archivos con más de dos años de antigüedad y listarlos con un comando `ls -lrt`:

```bash
~$ find / -name "*" -mtime +730 -type f -exec ls -lrt {} \;
```

Buscar, y ejecutar un comando:

```bash
~$ find /tmp/example/mp3 -name "*.mp3" -type f -exec chmod 644 {} \;
```

Buscar archivos por tamaño, por ejemplo, todos los que su tamaño está entre 100k y 500k:

```bash
~$ find . -size +100k -a -size -500k
```

Buscar archivos y mostrar nombre del archivo, nombre con la ruta absoluta, usuario, grupo y permisos:

```bash
~$ find . -printf "%f %p %u %g %m\n"
```

Excluye todos los archivos que se encuentren en estos directorios en concreto:

```bash
~$ find /home/alice -type f -not -path "*example1/archivo*" -not -path "example3/archivo*"
```

### DATE

El comando `date`, si no se le pasa ninguna opción devuelve un formato de fecha y hora estándar como el siguiente:

```bash
~$ date
Fri May  1 10:50:39 UTC 2020
```

Sin embargo, podemos añadir una serie de parámetros para modificar el formato de salida haciéndolo así muy versátil. Por ejemplo, si quisiéramos obtener únicamente el año, mes y día separado por el carácter guión medio (-) podríamos ejecutarlo de la siguiente manera:

```bash
~$ date '+%Y-%m-%d'
2020-05-01
```

Si además quisiéramos añadir la hora, minutos y segundos bastaría con hacerlo de la siguiente forma:

```bash
~$ date '+%Y-%m-%d %H:%M:%S'
2020-05-01 10:55:45
```

Otro formato interesante es obtener el formato de fecha *epoch* o *Unix timestamp*:

```bash
~$ date '+%s'
1588330643
```

### PING

El comando `ping` es una herramienta de diagnóstico que permite hacer una verificación del estado de una determinada conexión de un host local con al menos un equipo remoto contemplado en una red de tipo TCP/IP. Sirve para determinar si una dirección IP específica o host es accesible desde la red o no. Por ejemplo, si queremos saber si el host `gnu.org` está "vivo" o accesible podríamos ejecutar el siguiente comando:

```bash
~$ ping gnu.org
PING gnu.org (209.51.188.148) 56(84) bytes of data.
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=1 ttl=50 time=113 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=2 ttl=50 time=113 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=3 ttl=50 time=113 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=4 ttl=50 time=112 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=5 ttl=50 time=113 ms
...
--- gnu.org ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4001ms
rtt min/avg/max/mdev = 112.955/113.390/113.739/0.419 ms
```

Con el flag `-c` podremos especificar el número de paquetes a enviar en la conexión, por ejemplo:

```bash
~$ ping -c 3 gnu.org
PING gnu.org (209.51.188.148) 56(84) bytes of data.
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=1 ttl=50 time=114 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=2 ttl=50 time=113 ms
64 bytes from wildebeest.gnu.org (209.51.188.148): icmp_seq=3 ttl=50 time=113 ms

--- gnu.org ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 113.178/113.779/114.600/0.601 ms
```

En realidad tiene muchas más opciones que nos permitirán hacer un uso más avanzado de este comando. También podemos hacer `ping` a la dirección IP a la que apunta el dominio `gnu.org`, en este caso la IP `209.51.188.148`:

```bash
~$ ping -c 3 209.51.188.148
PING 209.51.188.148 (209.51.188.148) 56(84) bytes of data.
64 bytes from 209.51.188.148: icmp_seq=1 ttl=50 time=113 ms
64 bytes from 209.51.188.148: icmp_seq=2 ttl=50 time=112 ms
64 bytes from 209.51.188.148: icmp_seq=3 ttl=50 time=113 ms

--- 209.51.188.148 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 112.941/113.111/113.339/0.167 ms
```

### MAN

El comando `man` (MANual) nos servirá para acceder a al manual o documentación de muchos de los programas que tengamos instalados en la computadora, solo hay que pasarle el nombre del programa como argumento para acceder a su documentación, por ejemplo, para acceder a la documentación del comando `ping` tendremos que escribir el siguiente comando:

```bash
~$ man ping
```

Al pulsar intro accederemos a una página en la shell como esta:

```bash
PING(8)                      System Managers Manual: iputils

NAME
       ping - send ICMP ECHO_REQUEST to network hosts

SYNOPSIS
       ping  [-aAbBdDfhLnOqrRUvV46]  [-c count] [-F flowlabel] [-i interval] [-I interface] [-l preload] [-m mark] [-M pmtudisc_option] [-N nodeinfo_option] [-w deadline] [-W timeout] [-p pattern] [-Q tos]
       [-s packetsize] [-S sndbuf] [-t ttl] [-T timestamp option] [hop ...] destination

DESCRIPTION
       ping uses the ICMP protocols mandatory ECHO_REQUEST datagram to elicit an ICMP ECHO_RESPONSE from a host or gateway.  ECHO_REQUEST datagrams (pings) have an IP and ICMP header,  followed  by  a
       struct timeval and then an arbitrary number of pad bytes used to fill out the packet.

       ping works with both IPv4 and IPv6. Using only one of them explicitly can be enforced by specifying -4 or -6.

       ping can also send IPv6 Node Information Queries (RFC4620).  Intermediate hops may not be allowed, because IPv6 source routing was deprecated (RFC5095).

OPTIONS
       -4     Use IPv4 only.

       -6     Use IPv6 only.

       -a     Audible ping.
...
```

Si la página de documentación es muy larga podremos bajar con el scroll, con las flechas del teclado o pulsando la barra espaciadora. Para salir de la página de un manual solo hay que pulsar la letra `q`.

**Nota**: No todos los programas tienen su documentación accesible mediante `man`, algunos tienen su propia documentación que se puede consultar con flags u opciones como `-h` o `--help`.

### EXIT

El comando `exit` lo utilizaremos para salir de un programa, en ocasiones el programa es la propia ejecución de la shell en la que estamos trabajando, lo que quiere decir que si escribimos el comando exit y pulsamos intro se cerrará la consola o terminal.

```bash
~$ exit
```

## Estructura de directorios y archivos

Los sistemas de archivos tienen directorios en los que se organizan los archivos y estos directorios suelen estar organizados jerárquicamente. La jerarquía implica que un directorio puede contener subdirectorios. El directorio más alto en la jerarquía del que cuelgan todos los demás se denomina raíz (root). Este directorio raíz se representa con una barra “/” y sólo existe una jerarquía, es decir, sólo existe un directorio raíz, incluso aunque haya distintos discos duros en el servidor. Dentro del directorio raíz podemos encontrar diversos subdirectorios, por ejemplo en GNU/Linux existe el directorio “home”. Así pues, el directorio “home” es un subdirectorio del directorio raíz “/”. Esta relación se representa como:

```
/home
```

home es el directorio dónde se encuentran los directorios de los usuarios en un sistema GNU/Linux. Imaginemos que tiene los subdirectorios “bob” y “alice”. Se representaría como:

```
/home/bob
/home/alice
```

Existe un estándar, denominado Filesystem Hierarchy Standard que define la estructura de directorios de los sistemas Unix. Los sistemas Unix y GNU/Linux suelen seguir este estándar, aunque a veces lo violan en algunos aspectos. Por ejemplo en el sistema MacOS X de Apple el directorio donde se encuentran los directorios de los usuarios se denomina “Users” y no “home”. En el directorio raíz hay diversos directorios que, en la mayoría de los casos, sólo deberían interesarnos si estamos administrando el servidor. Los usuarios normalmente sólo escriben dentro de un directorio de su propiedad localizado dentro de “/home” y denominado como su nombre de usuario. Los usuarios también pueden escribir en “/tmp” aunque normalmente son los procesos lanzados por estos lo que hacen esta escritura. Es importante revisar el espacio libre en la partición en la que se encuentra “/tmp” para que no se colapse el sistema. Recuerda que el directorio “/tmp” es borrado habitualmente por el sistema.

Los archivos pueden tener prácticamente cualquier nombre. Existe la convención de acabar los nombres con un punto y una pequeña extensión que indica el tipo de archivo. Pero esto es sólo una convención, en realidad podríamos no utilizar este tipo de nomenclatura. Si deseamos utilizar nombres de archivos que no vayan a causar extraños comportamientos en el futuro lo mejor sería seguir unas cuantas reglas al nombrar un archivo:

Añadir una extensión para recordarnos el tipo de archivo, por ejemplo .txt para los archivos de texto.

No utilizar en los nombres:

    * Espacios.
    * Caracteres no alfanuméricos.
    * Ni caracteres no ingleses como letras acentuadas o eñes.

Por supuesto, podríamos crear un archivo denominado “$ñ 1.txt” para referirnos a un archivo de sonido, pero esto conlleva una serie de problemas que aunque son solventables nos dificultará el trabajo.
Además es importante recordar que en Unix las mayúsculas y las minúsculas no son lo mismo. Los ficheros “documento.txt”, “Documento.txt” y “DOCUMENTO.TXT” son tres ficheros distintos.
Otra convención utilizada en los sistema Unix es la de ocultar los archivos cuyos nombres comienzan por punto “.”. Por ejemplo el archivo “.oculto” no aparecerá normalmente cuando pedimos el listado de un directorio. Esto se utiliza normalmente para guardar archivos de configuración que no suelen ser utilizados directamente por los usuarios. Para listar todos los archivos (All), ya sean éstos ocultos o no se puede ejecutar:

```bash
~$ ls -a
.   archivo_1.txt  .bash_logout  directorio_1  directorio_3
..  archivo_2.txt  .bashrc       directorio_2  .profile
```

## Empaquetado y compresión de ficheros y subdirectorios

### GZIP y GUNZIP

El comando `gzip` lo utilizaremos para comprimir un archivo reduciendo así su tamaño y haciéndolo más ligero. Por ejemplo:

```bash
~$ gzip archivo_1.txt
```

Si listamos los archivos que hay en el directorio actual veremos que el archivo que hemos comprimido tiene el mismo nombre y además extensión `.gz`, y ya no veremos el archivo original.

```bash
alice@localhost:~$ ls -lrt
-rw-rw-r-- 1 alice alice    34 May 17 10:25 archivo_1.txt.gz
```

Si queremos que además de comprimir un archivo este permanezca en el mismo directorio sin comprimir podremos utilizar el flag `-k` (keep) del siguiente modo:

```bash
~$ gzip -k archivo_1.txt
```

Y si queremos la mayor compresión posible podremos utilizar el flag `-9`, por ejemplo:

```bash
~$ gzip -9 archivo_1.txt
```

Para descomprimir un archivo con extensión `.gz` podremos hacerlo con el comando `gunzip` del siguiente modo:

```bash
~$ gunzip archivo_1.txt.gz
```

### TAR

El comando `tar` nos permite empaquetar (no comprimir) un conjunto de archivos y directorios. Por ejemplo, si queremos empaquetar un único archivo con extensión `.tar` todos los archivos que tengan extensión `.txt` que existan en el directorio actual podríamos hacerlo con el siguiente comando:

```bash
~$ tar -cf archivos_empaquetados.tar *.txt
```

En este caso usamos los flags `-cf` (*create*, *file*) para crear un nuevo archivo `.tar` que empaquete los archivos que le hemos indicado. Si listamos el contenido del directorio actual veremos que los archivos originales siguen donde están y además se ha creado un archivo nuevo llamado `archivos_empaquetados.tar`.

```bash
~$ ls -lrt
-rw-rw-r-- 1 alice alice     0 May 17 10:25 archivo_3.txt
-rw-rw-r-- 1 alice alice     0 May 17 10:25 archivo_2.txt
-rw-rw-r-- 1 alice alice     0 May 17 10:25 archivo_1.txt
-rw-rw-r-- 1 alice alice 10240 May 17 10:26 archivos_empaquetados.tar
```

Para desempaquetar los archivos podemos hacerlo de la siguiente forma:

```bash
~$ tar -xf archivos_empaquetados.tar
```

Esto volverá a colocar los archivos previamente empaquetados en el directorio en el que nos encontremos. Si los archivos ya existen, como es el caso de este ejemplo, los archivos se sobreescribirán. Si queremos desempaquetar los archivos en otro directorio podremos especificar dicho directorio de salida mediante el flag `-C` de la siguiente manera:

```bash
~$ tar -xf archivos_empaquetados.tar -C /tmp
```

Puedes listar el contenido del directorio que hayas elegido, en este ejemplo el directorio `/tmp`, para ver que los archivos que empaquetamos antes están ahí:

```bash
~$ ls -lrt /tmp
-rw-rw-r-- 1 alice alice    0 May 17 10:25 archivo_3.txt
-rw-rw-r-- 1 alice alice    0 May 17 10:25 archivo_2.txt
-rw-rw-r-- 1 alice alice    0 May 17 10:25 archivo_1.txt
```

El comando `tar` nos permite combinar la función de empaquetado con la de compresión mediante el los flags `-zcf`, por ejemplo:

```bash
~$ tar -zcf archivos_empaquetados.tar.gz *.txt
```

Y con los flags `-zxf` podremos desempaquetar y descomprimir, ambas cosas a la vez:

```bash
~$ tar -zxf archivos_empaquetados.tar.gz
```

Existen muchas otras opciones que te permiten utilizar el comando `tar` con más precisión. Para ver la documentación basta con escribir la opción `--help` de la siguiente manera:

```bash
~$ tar --help
```

## Transferencia de archivos

Una tarea muy habitual cuando estamos trabajando en un servidor remoto es enviar o traer ficheros desde el mismo. Esto también lo podemos hacer utilizando el protocolo SSH por lo que podremos hacerlo de un modo seguro en cualquier servidor que no dé acceso SSH.

### SCP

El programa más sencillo para hacerlo desde Unix o GNU/Linux es `scp` (*Secure Copy Protocol*). El comando `scp` tiene una interfaz muy similar a `cp` pero acepta que los ficheros de origen y destino estén en distintos servidores:

```bash
~$ scp alice@servidor:/remote/directory/file.txt /local/directory
```

```bash
~$ scp /local/directory/file.txt alice@servidor:/remote/directory/
```

Hay varios clientes SCP de entorno gráfico para Windows (y otros sistemas como Mac y GNU/Linux), uno de los más populares es [WinSCP](https://winscp.net/eng/download.php).

### SFTP

Otro programa muy utilizado para transferir archivos entre varias computadoras Unix o GNU/Linux es `sftp` (*Secure File Tranfer Protocol*). El programa `sftp` realiza todas las operaciones en una sesión `ssh` cifrada. Utiliza muchas de las características de `ssh`, como la autenticación de clave pública y la compresión de datos.

Hay varias formas básicas de usar `sftp`. Una de ellas es una sesión interactiva. En este modo, `sftp` se conecta y registra en la máquina remota especificada donde podremos escribir varios comandos en un shell específica. Por ejemplo:

```bash
~$ sftp alice@servidor
```

Al pulsar intro nos pedirá introducir la contraseña. Una vez dentro accederemos a un *prompt* del programa `sftp` que tendrá el siguiente aspecto:

```bash
~$ sftp alice@123.123.123.123
Connected to 123.123.123.123.
sftp>
```

En esta shell de `sftp` podremos escribir algunos comandos para listar archivos, enviarlos desde nuestra máquina al servidor remoto o traerlos desde el servidor remoto a nuestra máquina, entre muchas otras acciones. Por ejemplo, para listar los archivos que tenemos en local, es decir en nuestra máquina, podríamos hacerlo con el comando `lls` (*local list*) en la shell de `sftp` de la siguiente forma:

```bash
sftp> lls
fichero_local1.txt  fichero_local2.txt  fichero_local3.txt
```

Para listar los archivos remotos podríamos hacerlo con el comando `ls` tal y como lo haríamos en una shell de Unix o GNU/Linux, por ejemplo:

```bash
sftp> ls
fichero_remoto1.txt  fichero_remoto2.txt  fichero_remoto3.txt
```

Para traernos un fichero remoto a nuestra máquina local podemos usar el comando `get` especificando el fichero remoto que queremos traer, por ejemplo:

```bash
sftp> get fichero_remoto1.txt
Fetching /home/alice/fichero_remoto1.txt to fichero_remoto1.txt
/home/alice/fichero_remoto1.txt                 100%   22     1.2KB/s   00:00
```

Y para enviar un fichero desde nuestra máquina local hasta el servidor remoto podemos hacerlo con el comando `put` especificando el fichero local que queremos enviar, por ejemplo:

```bash
sftp> put fichero_local1.txt
Uploading fichero_local1.txt to /home/alice/fichero_local1.txt
fichero_local1.txt                              100%   12     1.3KB/s   00:00
```

Para salir de la shell `sftp` basta con escribir el comando `bye` o `exit`.

```bash
sftp> bye   
```

Otra manera de usar `sftp` es de un modo en el que no tengamos que entrar en una shell interactiva, sino indicarle al programa que simplemente queremos descargar un fichero desde el servidor remoto, es la siguiente:

```bash
~$ sftp alice@123.123.123.123:/home/alice/fichero_remoto1.txt
Password:
Connected to 123.123.123.123.
Fetching /home/alice/fichero_remoto1.txt to fichero_remoto1.txt
/home/alice/fichero_remoto1.txt                 100%   22     1.3KB/s   00:00
```

Enviar ficheros desde nuestra máquina local al servidor remoto de esta manera es un poco más complicado, ya que tendríamos que hacerlo mediante el flag `-b` y pasarle un *batchfile* que contenga todos los comandos de la shell propia de `sftp`.

Existen varios clientes FTP y SFTP de entorno gráfico para sistemas Windows (y otros sistemas como Mac y GNU/Linux), uno de los más populares es  [Filezilla Client](https://filezilla-project.org/download.php?type=client).

## Gestión de usuarios y permisos

Los sistemas Unix y GNU/Linux son multiusuario, es decir soportan que varios usuarios los utilicen simultáneamente. Todos los usuarios, excepto uno, tienen unos privilegios restringidos de base y no pueden modificar el sistema fácilmente. De este modo unos usuarios se ven protegidos de las acciones de los otros.

Existe un usuario especial llamado root con privilegios de administración absolutos sobre el sistema. Para realizar las tareas cotidianas no es recomendable acceder al sistema como root. En algunos sistemas GNU/Linux este usuario está deshabilitado por defecto y sólo se pueden adquirir los privilegios de administrador temporalmente.

Para conseguir que cada usuario pueda trabajar en sus archivos, pero que no pueda interferir accidental o deliberadamente con los archivos de otros usuarios se establece un sistema de permisos. Por defecto un usuario tiene permiso para leer y modificar sus propios archivos y directorios, pero no los de los demás. En los sistemas GNU/Linux los ficheros pertenecen a un usuario concreto y existen unos permisos diferenciados para este usuario y para el resto. Además el usuario pertenece a un grupo de trabajo. Por ejemplo, imaginemos que el usuario alice puede pertenecer al grupo “claseuno”. Si alice crea un fichero este tendrá unos permisos diferentes para alice, para el resto de miembros de su grupo y para el resto de usuarios del servidor. Podemos ver los permisos asociados a los ficheros utilizando el comando ls con la opción -l (Long):

```bash
~$ ls -l
total 12
-rw-rw-r-- 1 alice claseuno    0 Apr 25 09:44 archivo_1.txt
-rw-rw-r-- 1 alice claseuno    0 Apr 25 09:44 archivo_2.txt
drwxrwxr-x 2 alice claseuno 4096 Apr 25 09:44 directorio_1
drwxrwxr-x 2 alice claseuno 4096 Apr 25 09:44 directorio_2
drwxrwxr-x 2 alice claseuno 4096 Apr 25 09:45 directorio_3
```

En este caso, los ficheros listados pertenecen al usuario alice y al grupo claseuno. Los permisos asignados al usuario, a los miembros del grupo y al resto de usuarios están resumidos en la primeras letras de cada línea:

```
-rw-rw-r--
drwxrwxr-x
```

La primera letra indica el tipo de fichero listado: (d) directorio, (-) fichero u otro tipo especial. Las siguientes nueve letras muestran, en grupos de tres, los permisos para el usuario, para el grupo y para el resto de usuarios del servidor. Cada grupo de tres letras indica los permisos de lectura (Read), escritura (Write) y ejecución (eXecute). En el caso anterior el usuario alice tiene permiso de lectura y escritura (rw-), el grupo tiene permiso de lectura y escritura también (rw-), es decir puede modificar el fichero o el directorio, y el resto de usuarios solo tienen permisos de lectura (r--). En los ficheros normales el permiso de lectura indica si el fichero puede ser leído, el de escritura si puede ser modificado y el de ejecución si puede ser ejecutado. En el caso de los directorios el de escritura indica si podemos añadir o borrar ficheros del directorio y el de ejecución si podemos listar los contenidos del directorio.

Estos permisos pueden ser modificados con la instrucción chmod. En chmod cada grupo de usuarios se representa por una letra:

* `u`: usuario dueño del fichero
* `g`: grupo de usuarios del dueño del fichero
* `o`: todos los otros usuarios
* `a`: todos los tipos de usuario (dueño, grupo y otros)

Los tipos de permisos también están abreviados por letras:

* `r`: lectura
* `w`: escritura
* `x`: ejecución

Con estas abreviaturas podemos modificar los permisos existentes, por ejemplo, para hacer un fichero ejecutable:

```bash
~$ chmod u+x
```

O:

```bash
~$ chmod a+x
```

También podemos mediante chmod indicar los permisos para un tipo de usuario determinado.

```bash
~$ chmod a=rwx
```

Un modo algo menos intuitivo, pero más útil de utilizar chmod es mediante los números octales que representan los permisos.

* Lectura: `4`
* Escritura: `2`
* Ejecución: `1`

Para modificar los permisos de este modo debemos indicar el número octal que queremos que represente los permisos del fichero. La primera cifra representará al dueño, la segunda al grupo y la tercera al resto de usuarios. Por ejemplo si queremos que único permiso para el dueño y su grupo sea la lectura y que no haya ningún permiso para el resto de usuarios:

```bash
~$ chmod 110 fichero.txt
```

También podemos combinar permisos sumando los números anteriores. Por ejemplo, permiso para leer y escribir para el dueño y ningún permiso para el resto.

```bash
~$ chmod 300 fichero.txt
```

Permisos de lectura, escritura y ejecución para el dueño y su grupo y ninguno para el resto.

```bash
~$ chmod 770 fichero.txt
```

Las restricciones para los permisos no afectan al usuario root, el superusuario del sistema. root también puede modificar el dueño y el grupo al que pertenecen los ficheros mediante los comandos `chown` y `chgrp`.

```bash
~$ chown alice fichero.txt
```

```bash
~$ chgrp claseuno fichero.txt
```

## Creación y edición de ficheros con “vi”

En GNU/Linux existen varios editores de texto que funcionan con entorno gráfico, como [gedit](https://es.wikipedia.org/wiki/Gedit), [KWrite](https://es.wikipedia.org/wiki/KWrite) o [Leafpad](https://es.wikipedia.org/wiki/Leafpad), y editores que funcionan en la terminal, como el [emacs](https://es.wikipedia.org/wiki/Emacs), [nano](https://es.wikipedia.org/wiki/GNU_Nano), [vim](https://es.wikipedia.org/wiki/Vim) o [vi](https://es.wikipedia.org/wiki/Vi). El programa `vi` es un editor de texto potente y versátil que ya está instalado por defecto en la mayoría de las distribuciones de GNU/Linux. Vamos a ver cómo usar `vi` para crear o modificar archivos de texto, así como alguna de sus funcionalidades. Para abrir un archivo de texto con `vi` basta con ejecutar el siguiente comando:

```bash
~$ vi fichero.txt
```

Si el archivo llamado `fichero.txt` no existe se creará como nuevo archivo, y si existe se abrirá para ver su contenido y podremos editarlo. En este ejemplo el archivo no existe y se creará, así que veremos en la consola lo siguiente:

```bash
░
~
~
~
~
~
"fichero.txt" [New File]
```

Para entrar en modo escritura o inserción debemos pulsar la letra `i` (insert), a continuación podremos escribir cualquier texto, por ejemplo `Hola mundo!`:

```bash
Hola mundo!░
~
~
~
~
~
"fichero.txt" [New File]
```

Para salir del modo edición o inserción debemos pulsar la tecla `esc`, y finalmente para guardar y deberemos escribir `:wq` (write & quit) y pulsar intro.

Si listamos los archivos que tenemos en la carpeta actual aparecerá el archivo `fichero.txt` que acabamos de crear:

```bash
~$ ls -lrt
-rw-rw-r-- 1 alice alice 12 May 17 17:42 fichero.txt
```

Si lo vemos su contenido con el comando `cat` veremos que se ha guardado con éxito el texto que hemos escrito desde el editor `vi`:

```bash
~$ cat fichero.txt
Hola mundo!
```

Ahora vamos a editarlo, volvemos a abrirlo con `vi`:

```bash
~$ vi fichero.txt
```

Veremos el archivo de la siguiente manera:

```bash
Hola mundo!
~
~
~
~
~
"fichero.txt" 1 line, 12 characters
```

Ahora puedes desplazar el cursor por la única línea de texto que hay, la línea 1, con las flechas izquierda y derecha del teclado. Estás en modo lectura, pero si quisieras entrar en modo edición tendrías que pulsar la tecla `i` (insert). Si quisiéramos añadir texto a continuación del último carácter, antes de pulsar la `i` (si ya la has pulsado debes pulsar `esc`) puedes desplazarte con la flecha derecha hasta el final, o pulsar `$` para ir automáticamente al final de una línea. Con el cursor situado en el último carácter pulsa `a` y se desplazará un carácter a la derecha en modo inserción. A continuación pulsamos un intro para escribir una segunda línea, y en esta escribiremos el mensaje `Estoy editando con vi.`:

```bash
Hola mundo!
Estoy editando con vi.
~
~
~
~
"fichero.txt" 1 line, 12 characters
```

Una vez más, para salir del modo inserción pulsamos `esc` y escribimos `:wq` (write & quit) pulsando finalmente intro. Ahora ya sabemos cómo abrir un archivo con `vi` y añadir texto en él nos estaremos preguntando cómo podemos borrar una letra, una palabra o una línea entera. Para practicarlo vamos a abrir de nuevo el archivo con `vi` y nos posicionamos al inicio de la segunda línea con el cursor:

```bash
Hola mundo!
Estoy editando con vi.
~
~
~
~
"fichero.txt" 2 lines, 35 characters
```

Ahora nos vamos con el cursor a la derecha hasta posicionarnos en la `e` de la palabra `editando`. Vamos a borrar esta palabra y escribir en su lugar la palabra `escribiendo`. Para borrar un carácter en `vi` pulsaremos la tecla `x`, no hace falta entrar en modo inserción, pues vamos a borrar no insertar. Pulsaremos la `x` ocho veces hasta haber borrado todos los carácteres de la palabra `editando`.

```bash
Hola mundo!
Estoy ░con vi.
~
~
~
~
"fichero.txt" 2 lines, 35 characters
```

Ahora es cuando tenemos que pulsar `i` para entrar en modo inserción y escribir la palabra `escribiendo`. Finalmente pulsamos `esc` para salir del modo inserción y escribimos `:wq` e intro para guardar y salir.

```bash
Hola mundo!
Estoy escribiendo░con vi.
~
~
~
~
"fichero.txt" 2 lines, 35 characters
```

Podemos volver a abrir el archivo con `cat` para comprobar que lo hemos editado correctamente:

```bash
~$ cat fichero.txt
Hola mundo!
Estoy escribiendo con vi.
```

En `vi` existen varias combinaciones de teclas y comandos que nos permiten ejecutar acciones más avanzadas como por ejemplo eliminar una línea entera de principio a fin, para ello solo hay que posicionarse en la línea a borrar y pulsar `dd`. Otra función interesante es la de eliminar todas las líneas del archivo, por lo que puedes posicionarte al inicio de la primera línea y pulsar `dG`. También se pueden escribir comandos un poco más complejos como el reemplazo de una cadena por otra en todos los casos en los que aparezca la expresión a sustituir, por ejemplo escribiendo el siguiente comando dentro de `vi`:

```bash
:1,$ s/cadena1/cadena2/g
```

Esto cambiará automáticamente la cadena de texto `cadena1` en todos lo lugares del archivo donde lo encuentre por la cadena de texto `cadena2`.

Realmente crear y editar archivos de texto con `vi` puede resultar un poco complicado al principio, pues todo son comandos y atajos de teclado, pero hay que tener en cuenta que es un programa de edición de texto en una terminal, no en modo gráfico al que quizás mucha gente pueda estar acostumbrada.

## Gestión de procesos

Un proceso es una instancia de un programa en ejecución. Programas y procesos son entidades distintas. En un sistema operativo multitarea, múltiples instancias de un programa pueden ejecutarse simultáneamente. Cada instancia es un proceso separado.
Prácticamente todo lo que se está ejecutando en el sistema en cualquier momento es un proceso, incluyendo la shell, el ambiente gráfico, que puede tener múltiples procesos, etc.

GNU/Linux, como ya hemos visto, es un sistema operativo multitarea y multiusuario. Esto quiere decir que múltiples procesos pueden operar simultáneamente sin interferir unos con otros. Por ejemplo, si cinco usuarios desde equipos diferentes, ejecutan el mismo programa al mismo tiempo, habría cinco instancias del mismo programa, es decir, cinco procesos distintos. Cada proceso que se inicia es identificado con un número de identificación único conocido como PID (Process ID), que es siempre un número natural.

Haciendo análisis muchas veces ejecutaremos programas, crearemos procesos, que duren un tiempo considerable. Es interesante que durante el tiempo que dure el proceso podamos consultar su estado. Los entornos UNIX y GNU/Linux tienen una serie de herramientas para poder conocer el estado de los procesos y del sistema en general.

### PS

El comando ps (process status) nos informa sobre el estado de los procesos. Dependiendo de los parámetros que le demos nos mostrará un tipo de información u otra y unos procesos u otros.

```bash
~$ ps
```

Si queremos obtener la lista completa de procesos podemos usar las opciones -ef:

```bash
~$ ps -ef
```

En este caso la segunda columna nos indicará el PID o identificador único del proceso.

### KILL

El comando kill, a pesar de su nombre, no sólo sirve para matar o terminar procesos sino también para enviar señales a los procesos. La señal por defecto (cuando no se indica ninguna es terminar o matar el proceso), y la sintaxis es kill PID, siendo PID el número de ID del proceso. Pero hay otras señales que podemos enviar. Así, por ejemplo, es posible enviar una señal de STOP al proceso y se detendrá su ejecución, después cuando se quiera reanudar su ejecución podemos enviar la señal CONTinuar y el proceso continuará desde donde se quedó detenido. Con kill -l podemos acceder a una lista de todas las señales que podemos mandar a un proceso:

```bash
~$ kill -l
```

El modo más convencional de matar un proceso es intentar primero que muera ordenadamente con un -15 (Termination signal) y sino lo conseguimos matarlo con un -9 (Kill signal):

```bash
~$ kill -15 4719
```

```bash
~$ kill -9 4719
```

## Gestión de recursos

### DF

El sistema de archivos puede abarcar una o más particiones. Una partición es una región de un disco o de cualquier otro medio de almacenamiento. Las instalaciones de Windows tienen normalmente una partición por disco, pero en Unix y GNU/Linux esto no es tan habitual. Cada partición tiene un sistema de archivos propio, en Unix estos sistemas deben estar montados en algún lugar dentro de la jerarquía que cuelga de la raíz “/”. En Windows cada partición tiene por defecto una jerarquía independiente.

Podemos pedir información sobre el espacio ocupado por las distintas particiones que tenemos actualmente montadas usando el comando df (Disk Free).

```bash
~$ df -h
Filesystem             Size  Used Avail Use% Mounted on
udev                   2.0G     0  2.0G   0% /dev
tmpfs                  394M   11M  383M   3% /run
/dev/mapper/vg00-lv01   77G  1.6G   72G   3% /
tmpfs                  2.0G     0  2.0G   0% /dev/shm
tmpfs                  5.0M     0  5.0M   0% /run/lock
tmpfs                  2.0G     0  2.0G   0% /sys/fs/cgroup
/dev/sda1              464M  104M  332M  24% /boot
```

### FREE

El comando free nos muestra información sobre el uso y disponibilidad de la memoria. Es aconsejable usar la opción -h ya que así generará la información en una forma más fácil de leer.

```bash
~$ free -h
```

### TOP

Una herramienta muy usada y muy útil para el monitoreo en tiempo real del estado de los procesos y de otras variantes del sistema es el programa llamado top, se ejecuta desde la línea de comandos, es interactivo y por defecto se actualiza cada 3 segundos.
Estando dentro de la aplicación, presionando “h“ se accede a una ayuda de los posibles comandos que permiten configurar top, por ejemplo, al presionar “s” pregunta por el tiempo en segundos de actualización, etc.

## Ejecución de programas ShellScript

Un programa ShellScript es un programa pensado para ser ejecutado en una terminal o shell de Unix o GNU/Linux, es básicamente un intérprete de línea de comandos. Por ejemplo, podríamos hacer un script muy básico llamado `mi_primer_script.sh` que ejecute una única línea en la que se liste todos los archivos del directorio en el que nos encontramos, por ejemplo:

```bash
#!/bin/bash

ls -lrt
```

En este ejemplo encontramos como primera línea `#!/bin/bash`, es la declaración del intérprete de comandos que se desea utilizar en su ejecución. Los dos primeros carácteres se denominan [Shebang](https://es.wikipedia.org/wiki/Shebang) y a continuación viene la ruta absoluta del ejecutable o intérprete, en este caso `/bin/bash`, pero podría ser otro como `/bin/ksh`, `/bin/sh`, `/usr/bin/awk`, etc. Nosotros usaremos para los ejemplos siempre `/bin/bash`.

La siguiente línea es una simple línea de comandos shell, solo que la hemos escrito en un archivo. Se ejecutará igual que si lo tecleamos en la terminal. Para poder ejecutar nuestro primer script antes deberemos darle permisos de ejecución y para ello será necesario escribir el siguiente comando:

```bash
~$ chmod 700 mi_primer_script.sh
```

Ahora ya podremos ejecutarlo como un programa, la sintaxis para ejecutarlo es la siguiente:

```bash
~$ ./mi_primer_script.sh
```

Tendremos como resultado un listado de archivos que se encuentran en el directorio en el que estamos, en este caso únicamente el propio script:

```bash
-rwx------ 1 alice alice  8 May 18 11:15 mi_primer_script.sh
```

Ahora probemos a modificar nuestro script y añadamos las siguientes líneas haciendo así el programa un poquito más complejo:

```bash
#!/bin/bash

for i in {1..5}
do
    echo "Ahora i vale: $i"
done
```

Este programa simplemente imprime un mensaje un número de veces, concretamente 5 veces, y en cada iteración la variable `$i` se va teniendo un valor distinto. Imprime el siguiente mensaje por pantalla:

```bash
Ahora i vale: 1
Ahora i vale: 2
Ahora i vale: 3
Ahora i vale: 4
Ahora i vale: 5
```

## Entrada, salida y error standard de datos (stdin/stdout/stderr)

Tanto en Unix como en GNU/Linux todo es un archivo. Cada vez que ejecutamos una orden, el sistema operativo le abre automáticamente tres interfaces, y estas son:

* Entrada estándar (stdout)
* Salida estándar (stdin)
* Error estándar (stderr)

La entrada estándar se refiere al archivo por el que una orden recibe su entrada, por defecto es el teclado. La salida estándar se refiere al archivo por el que una orden presenta sus resultados, por defecto es la pantalla o más concretamente la ventana en la que se está ejecutando el intérprete de órdenes. El error estándar se refiere al archivo por el que una orden presenta los mensajes que va generando cuando ocurre un error, y en este caso por defecto también es la pantalla. Estos archivos se crean en el directorio `/dev`.

Antes de que se ejecute una orden, es posible redirigir cualquiera de sus archivos de salida, esto es la redirección. Para llevarla a cabo es necesario utilizar los operadores de redirección que se procesan en el orden en el que aparecen.

### Redirección de la entrada estándar

Cuando se quiere redirigir la entrada estándar de una orden a un archivo, es necesario utilizar el operador de redirección “<” seguido del nombre del archivo. En este caso, una orden lee los datos de entrada que necesita desde el archivo señalado, en vez de desde el teclado.

### Redirección de la salida estándar

La salida por defecto de cualquier orden dada en el bash es el monitor. Por ejemplo, si utilizas la orden cal, te mostrará un calendario en la pantalla. Sin embargo puedes hacer que te envíe esos datos a un documento de texto por ejemplo escribiendo lo siguiente:

```bash
~$ echo "Hola mundo!" > salida.txt
```

Vemos cómo se utiliza el operador “>” para establecer que la dirección de la salida será un archivo de texto llamado mitexto.txt, si el archivo no existe se crea, y si existe se vacía antes de hacer la redirección. Para que al redireccionar no borre el contenido de un archivo que ya existe, es decir, que se agregue el calendario al final del archivo, se debe utilizar el operador “>>”.

### Redirección del error estándar

Cuando se quiere redirigir el error estándar de una orden a un archivo, es necesario utilizar el operador de redirección “2>” seguido del nombre del archivo. Como en la salida estándar, en el caso de no existir el archivo se crea, y en el caso de existir, el archivo se vacía antes de hacer la redirección. En el caso de que se quiera añadir el error estándar de una orden sin borrar el contenido de un archivo que ya existe, el operador de redirección a utilizar debe ser “2>>”.

Por ejemplo, si introduces en el terminal:

```bash
~$ abcdefghi1234567
-su: abcdefghi1234567: command not found
```

Te mostrará un error por pantalla ya que la orden no existe (command not found). Si redireccionamos el error estándar a un archivo:

```bash
~$ abcdefghi1234567 2> error.txt
```

El error ahora no se muestra por pantalla, como si nada pasara; se guarda en el archivo error.txt.

## Instalación de paquetes

La instalación de paquetes, librerías o programas desde la terminal o consola la haremos con la utilidad `apt-get`, pero antes deberemos actualizar la lista de paquetes disponibles en los repositorios que tenemos configurados en nuestro sistema. Para ello ejecutaremos el siguiente comando antes de instalar un paquete:

```bash
~$ apt-get update
root@localhost:~# apt-get update
Get:1 http://archive.ubuntu.com/ubuntu xenial InRelease [247 kB]
Get:2 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]       
Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]           
Get:4 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]                   
Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 Packages [1,201 kB]                 
Get:6 http://archive.ubuntu.com/ubuntu xenial/main i386 Packages [1,196 kB]                   
Get:7 http://archive.ubuntu.com/ubuntu xenial/main Translation-en [568 kB]                                          
Get:8 http://archive.ubuntu.com/ubuntu xenial/restricted amd64 Packages [8,344 B]       
Get:9 http://archive.ubuntu.com/ubuntu xenial/restricted i386 Packages [8,684 B]
Get:10 http://archive.ubuntu.com/ubuntu xenial/restricted Translation-en [2,908 B]
Get:11 http://archive.ubuntu.com/ubuntu xenial/universe amd64 Packages [7,532 kB]     
Get:12 http://archive.ubuntu.com/ubuntu xenial/universe i386 Packages [7,512 kB]       
Get:13 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [864 kB]
Get:14 http://archive.ubuntu.com/ubuntu xenial/universe Translation-en [4,354 kB]             
Get:15 http://archive.ubuntu.com/ubuntu xenial/multiverse amd64 Packages [144 kB]                          
Get:16 http://archive.ubuntu.com/ubuntu xenial/multiverse i386 Packages [140 kB]            
Get:17 http://archive.ubuntu.com/ubuntu xenial/multiverse Translation-en [106 kB]           
Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,149 kB]             
Get:19 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [923 kB]            
Get:20 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [434 kB]         
Get:21 http://archive.ubuntu.com/ubuntu xenial-updates/restricted amd64 Packages [7,616 B]        
Get:22 http://archive.ubuntu.com/ubuntu xenial-updates/restricted i386 Packages [7,580 B]    
Get:23 http://archive.ubuntu.com/ubuntu xenial-updates/restricted Translation-en [2,272 B]   
Get:24 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [798 kB]            
Get:25 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [721 kB]      
Get:26 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [334 kB]          
Get:27 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse amd64 Packages [17.1 kB]   
Get:28 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse i386 Packages [16.3 kB]
Get:29 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse Translation-en [8,632 B]
Get:30 http://archive.ubuntu.com/ubuntu xenial-backports/main amd64 Packages [7,280 B]        
Get:31 http://archive.ubuntu.com/ubuntu xenial-backports/main i386 Packages [7,288 B]   
Get:32 http://archive.ubuntu.com/ubuntu xenial-backports/main Translation-en [4,456 B]  
Get:33 http://archive.ubuntu.com/ubuntu xenial-backports/universe amd64 Packages [8,064 B]    
Get:34 http://archive.ubuntu.com/ubuntu xenial-backports/universe i386 Packages [7,744 B]
Get:35 http://archive.ubuntu.com/ubuntu xenial-backports/universe Translation-en [4,328 B]
Get:36 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [660 kB]          
Get:37 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [325 kB]
Get:38 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [491 kB]
Get:39 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [422 kB]
Get:40 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [201 kB]
Get:41 http://security.ubuntu.com/ubuntu xenial-security/multiverse amd64 Packages [6,092 B]
Get:42 http://security.ubuntu.com/ubuntu xenial-security/multiverse i386 Packages [6,248 B]
Get:43 http://security.ubuntu.com/ubuntu xenial-security/multiverse Translation-en [2,888 B]
Fetched 30.8 MB in 6s (4,825 kB/s)
Reading package lists... Done
```

Ahora que ya tenemos todos los repositorios actualizados podremos proceder a instalar los paquetes que tengamos disponibles. En los siguientes puntos veremos la instalación de algunos paquetes de ejemplo que nos pueden resultar bastante útiles para trabajar en un sistema GNU/Linux.

### TREE

El paquete tree es una utilidad que nos ayudará a listar de forma recursiva todos los archivos y directorios que hay en una ubicación específica, y lo hará en forma de árbol desglosado, bastante más visual para poder ver la jerarquía de archivos en toda su profundidad. Para instalarlo basta con escribir el siguiente comando:

```bash
~# apt-get install tree
```

Al pulsar intro se comenzará a instalar y veremos por pantalla unos mensajes como los siguientes:

```bash
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following NEW packages will be installed:
  tree
0 upgraded, 1 newly installed, 0 to remove and 18 not upgraded.
Need to get 40.6 kB of archives.
After this operation, 138 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu xenial/universe amd64 tree amd64 1.7.0-3 [40.6 kB]
Fetched 40.6 kB in 0s (323 kB/s)
Selecting previously unselected package tree.
(Reading database ... 89411 files and directories currently installed.)
Preparing to unpack .../tree_1.7.0-3_amd64.deb ...
Unpacking tree (1.7.0-3) ...
Setting up tree (1.7.0-3) ...
```

Una vez finalizada la instalación podremos probar a ejecutarlo de la siguiente manera:

```bash
~# tree /home/alice/
/home/alice/
├── archivo_01.txt
├── archivo_02.txt
├── archivo_03.txt
└── directorio_1
    ├── archivo_11.txt
    ├── archivo_12.txt
    ├── archivo_13.txt
    └── directorio_2
        ├── archivo21.txt
        ├── archivo22.txt
        ├── archivo23.txt
        └── directorio_3
            ├── archivo31.txt
            ├── archivo32.txt
            └── archivo33.txt

3 directories, 12 files
```

Como se puede ver, la manera de listar los archivos y directorios es mucho mejor y se ve más claro. La utilidad o herramienta `tree` no viene instalada por defecto en el sistema GNU/Linux, por lo que si queremos utilizarla primero deberemos instalarla como ya hemos visto antes.

Para eliminar `tree` del sistema basta con ejecutar el siguiente comando:

```bash
~# apt-get remove tree
```

Aparecerán unos mensajes como los siguientes y pedirá confirmación, deberemos pulsar `Y` e intro:

```bash
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages will be REMOVED:
  tree
0 upgraded, 0 newly installed, 1 to remove and 18 not upgraded.
After this operation, 138 kB disk space will be freed.
Do you want to continue? [Y/n] Y
(Reading database ... 89418 files and directories currently installed.)
Removing tree (1.7.0-3) ...
```
