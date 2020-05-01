# Exercise #7: CLI Task Manager

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/task) <!--[![demo](https://img.shields.io/badge/demo-%E2%86%92-yellow.svg?style=for-the-badge)](https://gophercises.com/demos/cyoa/)-->

## Exercise details

En este ejercicio vamos a construir una herramienta CLI que se puede utilizar para administrar sus tareas pendientes en el terminal. El uso básico de la herramienta se verá más o menos así:
```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
```

*Nota: Las líneas con el prefijo $ son líneas donde escribimos en la terminal, y otras líneas salen de nuestro programa.*

Su CLI final no tendrá que verse exactamente así, pero así es como espero que sea la mía. En la sección de bonificación también discutiremos algunas características adicionales que podríamos agregar, pero por ahora nos quedaremos con los tres programas anteriores:

- `add` - agrega una nueva tarea a nuestra lista
- `list` - enumera todas nuestras tareas incompletas
- `do` - marca una tarea como completa

Para construir esta herramienta, necesitaremos explorar algunos temas diferentes. En particular, necesitaremos:

1. Obtenga información sobre cómo crear interfaces de línea de comandos (CLI)
2.  Interactuar con una base de datos. Usaremos BoltDB en este ejercicio para que podamos aprender al respecto.
3. Descubra cómo almacenar nuestro archivo de base de datos en diferentes sistemas operativos. Esto básicamente se reducirá a aprender sobre directorios de inicio.
4. Códigos de salida (brevemente)
5.Y probablemente más. Actualizaré esta lista una vez que termine el ejercicio.

Le invitamos a abordar el problema como mejor le parezca, pero a continuación se muestra el orden en que recomendaría comenzar.

### 1. Construir el shell CLI

Para construir la CLI, le recomiendo usar un paquete de terceros (biblioteca, marco o como quiera llamarlo). Puede hacer este ejercicio sin uno, pero hay muchos casos extremos que deberá manejar por su cuenta y, en este caso, creo que es mejor elegir una biblioteca existente para usar.

Hay muchas bibliotecas de CLI, y puede encontrar la mayoría de ellas aquí: <https://github.com/avelino/awesome-go#command-line>

Cuando codifico este ejercicio, pretendo usar [spf13/cobra](https://github.com/spf13/cobra).No es necesariamente mejor que otros, pero es uno que he usado en el pasado y sé que satisfará mis necesidades.

Una vez que elija una biblioteca, úsela para crear el comando original `task` que muestre todos sus subcomandos, y luego cree subcomandos stubbed para cada una de las acciones que discutimos anteriormente. Las acciones aún no tienen que hacer nada con una base de datos, pero queremos asegurarnos de que el usuario que escriba cada comando individual resulte en una ejecución diferente de código.

Por ejemplo, supongamos que definimos el comando `task list` para ejecutar el siguiente código Go:

```go
fmt.Println("This is a fake \"list\" command")
```

Luego, cuando usamos ese comando con nuestra CLI, deberíamos ver lo siguiente:

```
$ task list
This is a fake "list" command
```

Después de eliminar los 3 comandos, intente también ver cómo analizar argumentos para los comandos `task do` y` task add`.

### 2. Escribe las interacciones de BoltDB

Después de eliminar los comandos de la CLI, intente escribir código que lea, agregue y elimine datos en una base de datos BoltDB. Puede encontrar más información sobre el uso de Bolt aquí <https://github.com/boltdb/bolt>

* Nota: Sé que muchas personas afirman que `bolt` está abandonado, pero eso es inexacto en mi opinión. En cambio, lo consideraría un proyecto estable y completo que ya no necesita ningún desarrollo activo. Dicho esto, hay una bifurcación de la biblioteca creada por el equipo de CoreOS que se puede encontrar aquí: <https://github.com/coreos/bbolt>*

Por ahora, no se preocupe por dónde almacena la base de datos a la que se conecta Bolt. En esta etapa, tengo la intención de usar el directorio desde el que se ejecutó el comando `task`, por lo que utilizaré el código más o menos así:

```go
db, err := bolt.Open("tasks.db", 0600, nil)
```

Más adelante, puede investigar cómo instalar la aplicación para que pueda ejecutarse desde cualquier lugar y continuará con nuestras tareas independientemente de dónde ejecutemos la CLI.

### 3. Poniendolo todo junto

Finalmente, junte las dos piezas que escribió para que cuando alguien escriba `task add some task` agregue esa tarea al boltdb.

Después de eso, explore cómo configurar e instalar la aplicación para que pueda ejecutarse desde cualquier directorio en su terminal. Esto puede requerir que busque cómo encontrar el directorio de inicio de un usuario en cualquier sistema operativo (Windows, Mac OS, Linux, etc.).

Si lo desea, puede investigar cómo determinar esto por su cuenta, pero le recomiendo que tome este paquete: <https://github.com/mitchellh/go-homedir>. Puede leer el código para ver cómo funciona, son solo 137 líneas de código, pero debe ocuparse de todas las rarezas entre los diferentes sistemas operativos para nosotros.

Después de eso, necesitará investigar cómo instalar un binario en su computadora. El primer lugar que sugiero comenzar es el comando `go install`. (* Sugerencia: intente `go install --help` para ver qué hace este comando. *). Es probable que esta sea la ruta más simple, pero hay otras opciones (como copiar manualmente un binario a un directorio en su `$ PATH`).

* Nota: sospecho que muchos usuarios tendrán problemas por aquí que son específicos del sistema operativo. Si lo hace, primero verifique los [problemas de Github] (https://github.com/gophercises/task/issues?utf8=%E2%9C%93&q=is%3Aissue) para ver si hay problemas abiertos o cerrados que son similares a tu problema Espero usar eso como una buena sección de preguntas y respuestas para este ejercicio. *

Si todo va bien, debe tener una CLI completa para administrar sus tareas instaladas una vez que haya terminado con esta sección.


## Bono

Como ejercicio adicional, recomiendo trabajar en los siguientes dos comandos nuevos:

```
$ task rm 1
You have deleted the "review talk proposal" task.

$ task completed
You have finished the following tasks today:
- wash the dishes
- clean the car
```

El comando `rm` eliminará una tarea en lugar de completarla.

El comando `completado` enumerará todas las tareas completadas en el mismo día. Puede definir esto como lo desee (últimas 12 horas, últimas 24 horas o la misma fecha del calendario).

La primera versión de nuestra CLI podría eliminar las tareas de la base de datos, pero si desea que estas características funcionen, es probable que necesite modificar un poco su diseño de base de datos. Lo dejaré como ejercicio para que lo pruebe por su cuenta, pero si necesita ayuda, no dude en ponerse en contacto - <jon@calhoun.io>