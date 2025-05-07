// 4.  En  una  empresa  de  logística  de  paquetes  existe  una  sala  de  contenedores  donde  se 
// preparan las entregas. Cada contenedor puede almacenar un paquete y la sala cuenta con 
// capacidad para N contenedores. Resuelva considerando las siguientes situaciones: 
// a. La  empresa  cuenta  con  2  empleados:  un  empleado  Preparador  que  se  ocupa  de 
// preparar los paquetes y dejarlos en los contenedores; un empleado Entregador que 
// se ocupa de tomar los paquetes de los contenedores y realizar la entregas. Tanto el 
// Preparador como el Entregador trabajan de a un paquete por vez.

// b. Modifique la solución a) para el caso en que haya P Preparadores y E Entregadores. 

//esta parte no salio en el pdf, pero tiene estos dos (C,D) puntos mas
c) Modifique la solución a) para el caso en que haya E empleados Entregadores.
d) Modifique la solución a) para el caso en que haya P empleados Preparadores y E
empleadores Entregadores.
//


Queue contenedor;
sem usandoContenedor = 0; //semafora para que me diga si hay algun contedor en uso
sem contenedoresLibres = N; // semaforo para controlar la cantidad de contenedores libres
sem protecQueue = 1; // semaforo para proteger la Queue, protege del uso al mismo tiempo que la Queue
process preparador{
	while(true){
		Paquete paquete = preparandoPaquete();
		p(contenedoresLibres); // disminuyo la cantidad de contenedores libres
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		contenedor.enqueue(paquete);
		V(protecQueue); //libero el semaforo para usar la Queue
		V(usandoContenedor);
	}
}

process entregador{
	Paquete paquete;
	while(true){
		P(usandoContenedor);
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		paquete = contenedor.dequeue();
		V(protecQueue); //libero el semaforo para usar la Queue
		V(contenedoresLibres); // aumento la cantidad de contenedores libres
		entregarPaquete(paquete);	
	}
}





//A--- mejor explicado 
Queue contenedor;
sem usandoContenedor = 0;       // paquetes disponibles
sem contenedoresLibres = N;     // espacios disponibles
sem mutex = 1;                  // protección para la cola (sección crítica)

process preparador {
	while(true){
		Paquete paquete = preparandoPaquete();
		P(contenedoresLibres);      // espera si no hay espacio
		P(mutex);                   // entra a la sección crítica
		contenedor.enqueue(paquete);
		V(mutex);                   // sale de la sección crítica
		V(usandoContenedor);       // indica que hay un paquete más
	}
}

process entregador {
	Paquete paquete;
	while(true){
		P(usandoContenedor);       // espera si no hay paquetes
		P(mutex);                  // entra a la sección crítica
		paquete = contenedor.dequeue();
		V(mutex);                  // sale de la sección crítica
		V(contenedoresLibres);     // indica que hay un espacio más
		entregarPaquete(paquete);	
	}
}










//A
//Solucion NO correcta, no contrala Los N contenedores, puedo pushear a la Queue N+N elementos y esta mal
Queue contenedor;
sem usandoContenedor = 1;

process preparador{
	while(true){
		Paquete paquete = preparandoPaquete();
		P(usandoContenedor);
		// Seccion critica
		contenedor.enqueue(paquete);
		V(usandoContenedor);
	}
}

process entregador{
	Paquete paquete;
	while(true){
		P(usandoContenedor);
		// Seccion critica
		if(!contenedor.isEmpty()){
			paquete = contenedor.dequeue();
		}
		V(usandoContenedor);
		if(paquete != null){
			entregarPaquete(paquete);	
		}
	}
}




//===================================== -B- ====================================
// b. Modifique la solución a) para el caso en que haya P Preparadores y E Entregadores. 


Queue contenedor;
sem usandoContenedor = 0; //semafora para que me diga si hay algun contedor en uso
sem contenedoresLibres = N; // semaforo para controlar la cantidad de contenedores libres
sem protecQueue = 1; // semaforo para proteger la Queue, protege del uso al mismo tiempo que la Queue

process preparador[id: 0..P-1]{
	while(true){
		Paquete paquete = preparandoPaquete();
		p(contenedoresLibres); // disminuyo la cantidad de contenedores libres
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		contenedor.enqueue(paquete);
		V(protecQueue); //libero el semaforo para usar la Queue
		V(usandoContenedor);
	}
}


process entregador[id: 0..E-1]{
	while(true){
		Paquete paquete;
		P(usandoContenedor);
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		paquete = contenedor.dequeue();
		V(protecQueue); //libero el semaforo para usar la Queue
		V(contenedoresLibres); // aumento la cantidad de contenedores libres
		entregarPaquete(paquete);	
	}
}

//===================================== -C- ====================================

c) Modifique la solución a) para el caso en que haya P empleados Preparadores

Queue contenedor;
sem usandoContenedor = 0; //semafora para que me diga si hay algun contedor en uso
sem contenedoresLibres = N; // semaforo para controlar la cantidad de contenedores libres
sem protecQueue = 1; // semaforo para proteger la Queue, protege del uso al mismo tiempo que la Queue

process preparador[id: 0..E-1]{
	while(true){
		Paquete paquete = preparandoPaquete();
		p(contenedoresLibres); // disminuyo la cantidad de contenedores libres
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		contenedor.enqueue(paquete);
		V(protecQueue); //libero el semaforo para usar la Queue
		V(usandoContenedor);
	}
}


process entregador{
	while(true){
		Paquete paquete;
		P(usandoContenedor);
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		paquete = contenedor.dequeue();
		V(protecQueue); //libero el semaforo para usar la Queue
		V(contenedoresLibres); // aumento la cantidad de contenedores libres
		entregarPaquete(paquete);	
	}
}


//===================================== -D- ====================================

d) Modifique la solución a) para el caso en que haya E empleados Entregadores.


Queue contenedor;
sem usandoContenedor = 0; //semafora para que me diga si hay algun contedor en uso
sem contenedoresLibres = N; // semaforo para controlar la cantidad de contenedores libres
sem protecQueue = 1; // semaforo para proteger la Queue, protege del uso al mismo tiempo que la Queue

process preparador{
	while(true){
		Paquete paquete = preparandoPaquete();
		p(contenedoresLibres); // disminuyo la cantidad de contenedores libres
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		contenedor.enqueue(paquete);
		V(protecQueue); //libero el semaforo para usar la Queue
		V(usandoContenedor);
	}
}


process entregador[id: 0..E-1]{
	while(true){
		Paquete paquete;
		P(usandoContenedor);
		// Seccion critica
		P(protecQueue); //pido el semaforo para usar la Queue
		paquete = contenedor.dequeue();
		V(protecQueue); //libero el semaforo para usar la Queue
		V(contenedoresLibres); // aumento la cantidad de contenedores libres
		entregarPaquete(paquete);	
	}
}