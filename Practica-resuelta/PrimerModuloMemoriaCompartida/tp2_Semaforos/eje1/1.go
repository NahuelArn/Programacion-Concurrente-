// 1.  Un  sistema  operativo  mantiene  5  instancias  de  un  recurso  almacenadas  en  una  cola. 
// Además, existen P procesos que necesitan usar una instancia del recurso. Para eso, deben 
// sacar la instancia de la cola antes de usarla. Una vez usada, la instancia debe ser encolada 
// nuevamente. 
//sem -> sem = "semaforos"

 
Queue q; //instancias del recurso
sem recursosDispnibles = 5; //cantidad de recursos disponibles
sem mutexCola = 1; //semáforo para proteger la cola// me indica si esta en uso o no
process consumidor [id: 0..P-1] {
	Recurso r; //recurso que se va a usar
	while(true){ // el lifeTime no esta definido, siempre termianan pidiendo y liberando
		P(recursosDispnibles); //pido un recurso
		P(mutexCola); //me aseguro de que la cola no se use mientras yo la uso
		r = .q.dequeue(); //saco un recurso de la cola
		V(mutexCola); //libero la cola para que otros procesos puedan usarla
		//uso el recurso
		usandoRecurso(r); //uso el recurso
		P(mutexCola); //me aseguro de que la cola no se use mientras yo la uso
		q.enqueue(r); //vuelvo a poner el recurso en la cola
		V(mutexCola); //libero la cola para que otros procesos puedan usarla
		V(recursosDispnibles); //libero el recurso para que otros procesos puedan usarlo
	}
}



p vendria a ser un await, tomo y -1 

V libero el recurso tomado +1

// mirar notion o exp practica