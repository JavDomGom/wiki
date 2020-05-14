# Sistema GNU/Linux

## Breve introducción al sistema operativo

GNU, también conocido como GNU/Linux o Linux sin más, es un sistema operativo de tipo Unix, así como una gran colección de programas informáticos que componen al sistema, desarrollado por Richard Stallman para el Proyecto GNU y auspiciado por la Free Software Foundation. Está formado por software libre, mayoritariamente bajo términos de copyleft. GNU es el acrónimo recursivo de "GNU's Not Unix" (en español: GNU no es Unix), nombre elegido debido a que GNU sigue un diseño tipo Unix y se mantiene compatible con éste, pero se distingue de Unix por ser software libre y por no contener código de Unix.

El sobrenombre Linux es realmente el nombre del kernel o núcleo del sistema, es el principal responsable de facilitar a los distintos programas acceso seguro al hardware de la computadora o en forma básica, es el encargado de gestionar recursos, a través de servicios de llamada al sistema. Muchas personas se refieren al sistema operativo simplemente con el nombre “Linux”, y aunque es aceptado realmente es inexacto e injusto, el nombre del sistema es GNU, pero se acepta igualmente llamarlo GNU/Linux, pues ambos forman un sistema operativo completo y funcional.
## Acceso al sistema

### Protocolo SSH

Una de las grandes ventajas de utilizar la terminal es que podemos acceder a terminales en otros ordenadores muy fácilmente. El protocolo más utilizado para acceder a terminales de forma remota es SSH (Secure SHell). El protocolo SSH tiene un gran número de posibilidades, pero el uso más habitual es utilizarlo para abrir terminales en ordenadores remotos que tienen un servicio ssh. Este protocolo SSH es seguro porque cifra las comunicaciones entre el cliente y el servidor. Se diseñó como una alternativa segura a telnet. No debemos usar el protocolo telnet porque las comunicaciones en telnet, incluidas las claves de acceso, no están cifradas y cualquiera puede tener acceso a ellas.
Para acceder a una computadora que implemente el protocolo SSH podemos usar el programa ssh, pero previamente tenemos que tener una cuenta en esa computadora. Imaginemos que Alice tiene una cuenta en un ordenador que tiene un servicio ssh. Para conectarse puede hacer:

```bash
~$ ssh alice@ordenador
```

En este ejemplo podemos cambiar ordenador por la dirección IP o el nombre de la computadora a la que nos queremos conectar. Si el nombre de la cuenta de usuario (alice) en el ordenador cliente y en el servidor es el mismo podemos obviar el nombre de usuario, de modo que se podría hacer la conexión SSH de la siguiente manera:

```bash
~$ ssh ordenador
```
A continuación el servidor nos pedirá la clave correspondiente al usuario con el que nos queremos conectar. Existen clientes SSH para Windows con los que nos podemos conectar. Uno muy común es putty y otro mejor aún es MobaXterm.

Una tarea muy habitual cuando estamos trabajando en un ordenador remoto es enviar o traer ficheros desde el mismo. Esto también lo podemos hacer utilizando el protocolo SSH por lo que podremos hacerlo de un modo seguro en cualquier ordenador que no dé acceso SSH. El programa más sencillo para hacerlo desde Unix o GNU/Linux es scp (Secure Copy Protocol). El comando scp tiene una interfaz muy similar a cp pero acepta que los ficheros de origen y destino estén en distintos ordenadores:

```bash
~$ scp alice@ordenador:/remote/directory/file.txt /local/directory
```

```bash
~$ scp /local/directory/file.txt alice@ordenador:/remote/directory/
```

En Windows también hay distintos clientes SCP, uno de ellos es WinSCP.

### Protocolo SFTP

xxx

## Manejo del sistema a través de la Shell o terminal

La Shell (o terminal) es un intérprete de comandos. Es simplemente un modo alternativo de controlar un ordenador basado en una interfaz de texto. La terminal nos permite ejecutar software escribiendo el nombre del programa que queremos ejecutar en la terminal. Podemos pedirle al ordenador que ejecute un programa mediante el ratón haciendo click en distintos lugares del escritorio o podemos escribir una orden para conseguir el mismo objetivo. Ninguna de las dos formas de comunicarse con el ordenador es mejor que la otra aunque en ciertas ocasiones puede resultar más conveniente utilizar una u otra. Las ventajas de la línea de comandos son:

Necesidad. Existe mucho software que está sólo disponible en la terminal. Esto es especialmente cierto en el área de la administración de sistemas y ciberseguridad.
Flexibilidad. Los programas gráficos suelen ser muy adecuados para realizar la tarea para la que han sido creados, pero son difíciles de adaptar para otras tareas. Los programas diseñados para ser usados en la línea de comandos suelen ser muy versátiles.
Reproducibilidad. Documentar y repetir el proceso seguido para realizar un análisis con un programa gráfico es muy costoso puesto que es difícil describir la secuencia de clicks y doble clicks que hemos realizado. Por el contrario, los procesos realizados mediante la línea de comandos son muy fáciles de documentar puesto que tan sólo debemos guardar el texto que hemos introducido en la pantalla.
Fiabilidad. Los programas básicos de Unix fueron creados en los años 70 y han sido probados por innumerables usuarios por lo que se han convertido en piezas de código extraordinariamente confiables.
Recursos. Las interfaces gráficas suelen consumir muchos recursos mientras que los programas que funcionan en línea de comandos suelen ser extraordinariamente livianos y rápidos. Este poco uso de recursos facilita, por ejemplo, que se utilice a través de la red.

El problema de la terminal es que para poder utilizarla debemos saber previamente qué queremos hacer y cómo. Es habitual descubrir cómo funciona un programa con una interfaz gráfica sin tener que leer un manual, esto no sucede en la terminal.

Para usar la línea de comandos hay que abrir una terminal. Se abrirá una ventana con un mensaje similar a:

```
usuario $
```

Este pequeño mensaje se denomina prompt y el cursor parpadeante que aparece junto al él indica que el ordenador está esperando una orden. El mensaje exacto que aparece en el prompt puede variar ligeramente, pero en sistemas GNU/Linux suele ser similar a:

```
usuario@ordenador:~$
```

En el prompt se nos muestra el nombre del usuario, el nombre del ordenador y el directorio en el que nos encontramos actualmente, es decir, el directorio de trabajo actual. La virgulilla “~” significa que estás en tu home o en tu carpeta de usuario. Finalmente aparece el símbolo del dólar “$” como final del prompt y luego el cursor que parpadea, desde donde podremos escribir comandos. Cuando el prompt se muestra podemos ejecutar cualquier comando.

## Comandos básicos

### LS

Mediante el comando ls (LiSt) podemos pedir que liste el contenido del directorio en el que nos encontramos:

```bash
~$ ls folder_name
archivo_1.txt archivo_2.txt directorio_1 directorio_2  directorio_3
```

El comando ls, como cualquier otro comando, es en realidad un programa que el ordenador ejecuta. Cuando escribimos la orden (y pulsamos enter) el programa se ejecuta. Mientras el programa está ejecutándose el prompt desaparece y no podemos ejecutar ningún otro comando. Pasado el tiempo el programa termina su ejecución y el prompt vuelve a aparecer. En el caso del comando ls el tiempo de ejecución es tan pequeño que suele ser imperceptible. Los programas suelen tener unas entradas y unas salidas. Dependiendo del caso estas pueden ser ficheros o caracteres introducidos o impresos en la pantalla. Por ejemplo, el resultado de ls es simplemente una lista impresa de ficheros y directorios en la interfaz de comandos.

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

El comando cd (change directory) nos servirá para cambiar de directorio. Por ejemplo, si queremos entrar en el directorio directorio_1 podremos hacerlo de la siguiente manera.

```bash
~$ cd directorio_1
~/directorio_1$
```

Si quisiéramos cambiarnos al directorio subdirectorio_1 volveríamos a ejecutar el comando cd pero esta vez indicando el nombre subdirectorio_1:

```
~/directorio_1$ cd subdirectorio_1/
~/directorio_1/subdirectorio_1$
```

Si quisiéramos volver un directorio para atrás debemos escribir dos puntos (..) de la siguiente manera:

```
~/directorio_1/subdirectorio_1$ cd ..
~/directorio_1$
```

También podemos utilizar rutas absolutas, por ejemplo:

```bash
~$ cd directorio_1/subdirectorio_1/
~/directorio_1/subdirectorio_1$
```

### DATE

El comando date, si no se le pasa ninguna opción devuelve un formato de fecha y hora estándar como el siguiente:

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

Otro formato interesante es obtener el formato de fecha epoch o Unix timestamp:

```bash
~$ date '+%s'
1588330643
```

## Estructura de directorios y archivos

Los sistemas de archivos tienen directorios en los que se organizan los archivos y estos directorios suelen estar organizados jerárquicamente. La jerarquía implica que un directorio puede contener subdirectorios. El directorio más alto en la jerarquía del que cuelgan todos los demás se denomina raíz (root). Este directorio raíz se representa con una barra “/” y sólo existe una jerarquía, es decir, sólo existe un directorio raíz, incluso aunque haya distintos discos duros en el ordenador. Dentro del directorio raíz podemos encontrar diversos subdirectorios, por ejemplo en GNU/Linux existe el directorio “home”. Así pues, el directorio “home” es un subdirectorio del directorio raíz “/”. Esta relación se representa como:

```
/home
```

home es el directorio dónde se encuentran los directorios de los usuarios en un sistema GNU/Linux. Imaginemos que tiene los subdirectorios “bob” y “alice”. Se representaría como:

```
/home/bob
/home/alice
```

Existe un estándar, denominado Filesystem Hierarchy Standard que define la estructura de directorios de los sistemas Unix. Los sistemas Unix y GNU/Linux suelen seguir este estándar, aunque a veces lo violan en algunos aspectos. Por ejemplo en el sistema MacOS X de Apple el directorio donde se encuentran los directorios de los usuarios se denomina “Users” y no “home”. En el directorio raíz hay diversos directorios que, en la mayoría de los casos, sólo deberían interesarnos si estamos administrando el ordenador. Los usuarios normalmente sólo escriben dentro de un directorio de su propiedad localizado dentro de “/home” y denominado como su nombre de usuario. Los usuarios también pueden escribir en “/tmp” aunque normalmente son los procesos lanzados por estos lo que hacen esta escritura. Es importante revisar el espacio libre en la partición en la que se encuentra “/tmp” para que no se colapse el sistema. Recuerda que el directorio “/tmp” es borrado habitualmente por el sistema.

Los archivos pueden tener prácticamente cualquier nombre. Existe la convención de acabar los nombres con un punto y una pequeña extensión que indica el tipo de archivo. Pero esto es sólo una convención, en realidad podríamos no utilizar este tipo de nomenclatura. Si deseamos utilizar nombres de archivos que no vayan a causar extraños comportamientos en el futuro lo mejor sería seguir unas cuantas reglas al nombrar un archivo:

Añadir una extensión para recordarnos el tipo de archivo, por ejemplo .txt para los archivos de texto.
No utilizar en los nombres:
    * espacios,
    * caracteres no alfanuméricos,
    * ni caracteres no ingleses como letras acentuadas o eñes.

Por supuesto, podríamos crear un archivo denominado “$ñ 1.txt” para referirnos a un archivo de sonido, pero esto conlleva una serie de problemas que aunque son solventables nos dificultará el trabajo.
Además es importante recordar que en Unix las mayúsculas y las minúsculas no son lo mismo. Los ficheros “documento.txt”, “Documento.txt” y “DOCUMENTO.TXT” son tres ficheros distintos.
Otra convención utilizada en los sistema Unix es la de ocultar los archivos cuyos nombres comienzan por punto “.”. Por ejemplo el archivo “.oculto” no aparecerá normalmente cuando pedimos el listado de un directorio. Esto se utiliza normalmente para guardar archivos de configuración que no suelen ser utilizados directamente por los usuarios. Para listar todos los archivos (All), ya sean éstos ocultos o no se puede ejecutar:

```bash
~$ ls -a
.   archivo_1.txt  .bash_logout  directorio_1  directorio_3
..  archivo_2.txt  .bashrc       directorio_2  .profile
```

## Gestión de usuarios y permisos

Los sistemas Unix y GNU/Linux son multiusuario, es decir soportan que varios usuarios los utilicen simultáneamente. Todos los usuarios, excepto uno, tienen unos privilegios restringidos de base y no pueden modificar el sistema fácilmente. De este modo unos usuarios se ven protegidos de las acciones de los otros.

Existe un usuario especial llamado root con privilegios de administración absolutos sobre el sistema. Para realizar las tareas cotidianas no es recomendable acceder al sistema como root. En algunos sistemas GNU/Linux este usuario está deshabilitado por defecto y sólo se pueden adquirir los privilegios de administrador temporalmente.

Para conseguir que cada usuario pueda trabajar en sus archivos, pero que no pueda interferir accidental o deliberadamente con los archivos de otros usuarios se establece un sistema de permisos. Por defecto un usuario tiene permiso para leer y modificar sus propios archivos y directorios, pero no los de los demás. En los sistemas GNU/Linux los ficheros pertenecen a un usuario concreto y existen unos permisos diferenciados para este usuario y para el resto. Además el usuario pertenece a un grupo de trabajo. Por ejemplo, imaginemos que el usuario alice puede pertenecer al grupo “claseuno”. Si alice crea un fichero este tendrá unos permisos diferentes para alice, para el resto de miembros de su grupo y para el resto de usuarios del ordenador. Podemos ver los permisos asociados a los ficheros utilizando el comando ls con la opción -l (Long):

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

La primera letra indica el tipo de fichero listado: (d) directorio, (-) fichero u otro tipo especial. Las siguientes nueve letras muestran, en grupos de tres, los permisos para el usuario, para el grupo y para el resto de usuarios del ordenador. Cada grupo de tres letras indica los permisos de lectura (Read), escritura (Write) y ejecución (eXecute). En el caso anterior el usuario alice tiene permiso de lectura y escritura (rw-), el grupo tiene permiso de lectura y escritura también (rw-), es decir puede modificar el fichero o el directorio, y el resto de usuarios solo tienen permisos de lectura (r--). En los ficheros normales el permiso de lectura indica si el fichero puede ser leído, el de escritura si puede ser modificado y el de ejecución si puede ser ejecutado. En el caso de los directorios el de escritura indica si podemos añadir o borrar ficheros del directorio y el de ejecución si podemos listar los contenidos del directorio.

Estos permisos pueden ser modificados con la instrucción chmod. En chmod cada grupo de usuarios se representa por una letra:

* u: usuario dueño del fichero
* g: grupo de usuarios del dueño del fichero
* o: todos los otros usuarios
* a: todos los tipos de usuario (dueño, grupo y otros)

Los tipos de permisos también están abreviados por letras:

* r: lectura
* w: escritura
* x: ejecución

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
* Lectura: 4
* Escritura: 2
* Ejecución: 1

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

Las restricciones para los permisos no afectan al usuario root, el superusuario del sistema. root también puede modificar el dueño y el grupo al que pertenecen los ficheros mediante los comando chown y chgrp.

```bash
~$ chown alice fichero.txt
```

```bash
~$ chgrp claseuno fichero.txt
```

## Creación y edición de ficheros con “vi”

En GNU/Linux existen editores de texto que funcionan con entorno gráfico, como el gedit, KWrite o Leafpad, y editores que funcionan en la terminal, como el emacs, nano, vim o vi. El programa vi es un editor de texto potente y versátil que ya está instalado por defecto en la mayoría de las distribuciones de GNU/Linux. Vamos a ver cómo usar vi para crear o modificar archivos de texto, así como alguna de sus funcionalidades.
xxx

## Gestión de procesos

Un proceso es una instancia de un programa en ejecución. Programas y procesos son entidades distintas. En un sistema operativo multitarea, múltiples instancias de un programa pueden ejecutarse simultáneamente. Cada instancia es un proceso separado.
Prácticamente todo lo que se está ejecutando en el sistema en cualquier momento es un proceso, incluyendo la shell, el ambiente gráfico, que puede tener múltiples procesos, etc.

GNU/Linux, como ya hemos visto, es un sistema operativo multitarea y multiusuario. Esto quiere decir que múltiples procesos pueden operar simultáneamente sin interferir unos con otros. Por ejemplo, si cinco usuarios desde equipos diferentes, ejecutan el mismo programa al mismo tiempo, habría cinco instancias del mismo programa, es decir, cinco procesos distintos.
Cada proceso que se inicia es identificado con un número de identificación único conocido como PID (Process ID), que es siempre un número natural.
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

xxx

## Entrada, salida y error standard de datos (stdin/stdout/stderr)

Tanto en Unix como en GNU/Linux todo es un archivo. Cada vez que ejecutamos una orden, el sistema operativo le abre automáticamente tres interfaces, y estas son:

* Entrada estándar (stdout)
* Salida estándar (stdin)
* Error estándar (stderr)

La entrada estándar se refiere al archivo por el que una orden recibe su entrada, por defecto es el teclado. La salida estándar se refiere al archivo por el que una orden presenta sus resultados, por defecto es la pantalla o más concretamente la ventana en la que se está ejecutando el intérprete de órdenes. El error estándar se refiere al archivo por el que una orden presenta los mensajes que va generando cuando ocurre un error, y en este caso por defecto también es la pantalla. Estos archivos se crean en el directorio /dev.

Antes de que se ejecute una orden, es posible redirigir cualquiera de sus archivos de salida, esto es la redirección. Para llevarla a cabo es necesario utilizar los operadores de redirección que se procesan en el orden en el que aparecen, por ejemplo:

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

xxx
