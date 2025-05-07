// Realice una solución concurrente de grano grueso (utilizando <> y/o <await B; S>) para el 
// siguiente problema. Un sistema operativo mantiene 5 instancias de un recurso almacenadas 
// en una cola, cuando un proceso necesita usar una instancia del recurso la saca de la cola, la 
// usa y cuando termina de usarla la vuelve a depositar.







int N = 5;

queue q [N]; int cant_recursos_en_uso = 0; 

process instancia[id: 0..N-1]{ //cantidad de procesos concurrentes || Cada uno de ellos ejecuta el código dentro del while (true), pero en momentos diferentes según la disponibilidad del recurso.
    while (true) { //representa que los procesos están continuamente solicitando y devolviendo recursos.
        <await (cant_recursos_en_uso < N); 
        cant_recursos_en_uso = cant_recursos_en_uso + 1; //decremento en 1 mis recursos disponibles>
        int recurso = q.pop();>
        
        usando_recurso(recurso); //fn que utiliza el recurso "imprimiendo"

        <q.push(recurso);
        cant_recursos_en_uso = cant_recursos_en_uso - 1;> //vuelvo a recuperar mi recurso ocupado
    }
}




//explicacion de ejercicio 

// 1. Interpretación del enunciado
// El enunciado dice que:
// ✅ Hay 5 instancias de un recurso almacenadas en una cola.
// ✅ Cuando un proceso necesita usar una instancia, la saca de la cola.
// ✅ Cuando termina de usarla, la vuelve a depositar en la cola.

// 📌 Punto clave:
// El enunciado no menciona explícitamente una variable cant, pero implícitamente sugiere que nunca puede haber más de 5 procesos usando un recurso a la vez, ya que solo hay 5 recursos disponibles.

// 2. ¿Qué representa cant en la solución?
// La variable cant en el código lleva la cuenta de cuántos recursos están en uso.

// cant < 5 → Significa que hay al menos un recurso disponible en la cola.
// cant++ → Significa que un proceso ha tomado un recurso.
// cant-- → Significa que un proceso ha devuelto un recurso.
// 📌 Ejemplo visual:

// plaintext
// Copy
// Edit
// Tiempo 0:  5 recursos en la cola
// Tiempo 1:  Proceso A saca 1 recurso (cant = 1)
// Tiempo 2:  Proceso B saca 1 recurso (cant = 2)
// Tiempo 3:  Proceso C saca 1 recurso (cant = 3)
// Tiempo 4:  Proceso D saca 1 recurso (cant = 4)
// Tiempo 5:  Proceso E saca 1 recurso (cant = 5)  <-- Ya no quedan más
// Tiempo 6:  Proceso F intenta sacar un recurso pero espera (cant = 5)
// Tiempo 7:  Proceso A devuelve el recurso (cant = 4), ahora F puede continuar
// 🛑 Sin await (cant < 5)
// Si no controlamos la cantidad de recursos en uso, más de 5 procesos podrían intentar sacar recursos al mismo tiempo, lo que es incorrecto.

// 3. ¿Es cant realmente necesaria?
// Podemos eliminar cant y controlar el acceso simplemente verificando si la cola está vacía o no, porque la cola ya impone una restricción de 5 recursos.

// Código sin cant
// c
// Copy
// Edit
// colaRecurso c[5]; // La cola contiene 5 recursos inicialmente.

// process pSo[id: 0..N-1]{
//     Recurso recurso;
//     while (true){
//         <await (!c.empty()); // Espera hasta que haya un recurso disponible
//         recurso = c.pop(); >  // Saca un recurso de la cola

//         // Usa recurso

//         <c.push(recurso); >  // Devuelve el recurso
//     }
// }
// Ventajas de esta versión
// ✅ No necesitamos una variable extra (cant).
// ✅ Controlamos el acceso directamente con la estructura de la cola (c.empty()).
// ✅ El código es más claro y más cercano al enunciado.

// Conclusión
// 📌 La variable cant no es estrictamente necesaria porque el control de acceso ya lo hace la cola.
// 📌 La mejor solución es la última, donde usamos await (!c.empty()) en lugar de await (cant < 5).
// 📌 La cola ya limita la cantidad de procesos concurrentes al número de recursos disponibles.

// 🔥 Resumen final:
// 💡 La cola c por sí sola es suficiente para resolver el problema sin necesidad de cant. 🚀




//glosario

//tipo T 
int N = 5
int queue[N]
process instancia[id: 0..N-1]{
		<int recurso = queue.pop()>
		// uso el recurso
		usando_recurso(recurso);
		// fin de uso del recurso
		<queue.push(recurso)>
}




// 1. ¿Qué significa while (true)?
// while (true) es un bucle infinito.
// Significa que el proceso se ejecutará indefinidamente, hasta que el sistema lo termine o lo suspenda.
// En este caso, representa que los procesos están continuamente solicitando y devolviendo recursos.
//======================================================================
// 2. ¿Qué significa id: 0..N-1?
// Define varios procesos con identificadores (id).
// id varía de 0 a N-1, lo que significa que hay N procesos en total.
// 📌 Ejemplo:
// Si N = 3, entonces hay tres procesos:

// pSo[0], pSo[1], pSo[2].
// Cada uno de ellos ejecuta el código dentro del while (true), pero en momentos diferentes según

//======================================================================

// RESUMEN
// while (true): representa que los procesos están continuamente solicitando y devolviendo recursos..
// id: 0..N-1: Hay N procesos concurrentes, cada uno con su propio identificador.