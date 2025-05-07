// 7.  Existen N personas que deben imprimir un trabajo cada una. Resolver cada ítem usando 
// semáforos: 


// a) Implemente una solución suponiendo que existe una única impresora compartida por 
// todas las personas, y las mismas la deben usar de a una persona a la vez, sin importar 
// el orden. Existe una función Imprimir(documento) llamada por la persona que simula el 
// uso de la impresora. Sólo se deben usar los procesos que representan a las Personas.


// b) Modifique la solución de (a) para el caso en que se deba respetar el orden de llegada. 


// c) Modifique  la  solución  de  (a)  para  el  caso  en  que  se  deba  respetar  estrictamente  el 
// orden dado por el identificador del proceso (la persona X no puede usar la impresora 
// hasta que no haya terminado de usarla la persona X-1). 


// d) Modifique la solución de (b) para el caso en que además hay un proceso Coordinador 
// que le indica a cada persona que es su turno de usar la impresora. 


// e) Modificar la solución (d) para el caso en que sean 5 impresoras. El coordinador le 
// indica a la persona cuando puede usar una impresora, y cual debe usar.


//A 
sem printer = 1 // semáforo para la impresora
process personas[id: 0..N-1]{
	Documento documento;
	P(printer)	// solicitar la impresora
	imprimir(documento); //usando la impresora
	V(printer)	// liberar la impresora
}

//B
queueFree = true;
Queue q;
sem semaforoPersonas [0..N-1] = ([N-1] 0);
sem printerUsed = 1;

process personas[id: 0..N-1]{
	Documento documento;
	int i;
	P(printerUsed) // solicitar la impresora
	if(queueFree){
		queueFree = False;
		V(printerUsed); // liberar la impresora
	}else{
		q.enqueue(id); // agregar a la cola
		V(printerUsed); // liberar la impresora
		V(semaforoPersonas[id]); // duermo a la persona
	}
	Imprimir(documento); // usando la impresora
	P(printerUsed); // solicitar la impresora //esto protege a la variable queueFree
	if(q.isEmpty()){
		queueFree = true; // la cola está vacía
	}else{
		int id = q.dequeue(); // sacar de la cola
		V(semaforoPersonas[id]); // despierto a la persona
	}
	V(printerUsed); // liberar la impresora
}


// C
int actual = 0;
// sem printerUsed = 1; AL saber que no se va usar x mas de una persona a la vez, no es necesario
sem semaforoPersonas [0..N-1] = ([N-1] 0);
process persona[id: 0..N-1]{
	Documento documento;
	if(id != actual){ //en la primera iteración, solo va dejar pasar a la persona 0
		P(semaforoPersonas[id]); //al no ser x > 0, se queda durmiendo
	}
	// P(printerUsed); // Redundante
	imprimiendo(documento); // usando la impresora
	actual++;
	V(semaforoPersonas[actual]); // despierto a la persona
	// V(printerUsed); // Redundante
}



// d) Modifique la solución de (b) para el caso en que además hay un proceso Coordinador 
// que le indica a cada persona que es su turno de usar la impresora. 

// queueFree = true;
Queue q;
sem queuUsed = 1;

sem semaforoPersonas [0..N-1] = ([N] 0);
sem printerUsed = 1;

process personas[id: 0..N-1]{
	Documento documento;
	P(queuUsed)
	q.enqueue(id); // agregar a la cola
	V(queuUsed)
	P(semaforoPersonas[id]); // duermo a la persona  //ACA ESPERO HASTA QUE ME DESPIERTEN
	P(printerUsed) // solicitar la impresora
	usandoImpresora(documento); // usando la impresora
	V(printerUsed) // liberar la impresora
}

int cantUsos =0;
process coordinador{
	while(cantUsos < N){
		P(queuUsed)
		if(!queue.isEmpty){
			// P(queuUsed)
			int id = q.dequeue(); // sacar de la cola
			V(queuUsed)
			V(semaforoPersonas[id]); // despierto a la persona INDICADA
			cantUsos++;
		}else{
			V(queuUsed)
		}
	}
}


// e) Modificar la solución (d) para el caso en que sean 5 impresoras. El coordinador le 
// indica a la persona cuando puede usar una impresora, y cual debe usar. 

Queue q;
sem queuUsed = 1;

sem semaforoPersonas [0..N-1] = ([N] 0);
sem printerUsed = 1;

impresoras[0..N-1] = ([N] null); //donde voy poniendo mis impresoras 
sem impresorasUsed = 5;
Queue<Impresora> stockImpresoras; // cola de impresoras disponibles
sem semaforoImpresoras = 1; // semáforo para la cola de impresoras

process personas[id: 0..N-1]{
	Documento documento;
	P(queuUsed)
	q.enqueue(id); // agregar a la cola
	V(queuUsed)
	P(semaforoPersonas[id]); // duermo a la persona  //ACA ESPERO HASTA QUE ME DESPIERTEN

	Impresora impresora = impresoras[id];

	P(semaforoImpresoras) // solicitar la impresora
	usandoImpresora(documento); // usando la impresora

	stockImpresoras.enqueue(impresora); // devolver la impresora a la cola de impresoras

	V(semaforoImpresoras) // liberar la impresora
	V(impresorasUsed);

}

int cantUsos =0;
process coordinador{
	while(cantUsos < N){
		P(queuUsed)
		if(!queue.isEmpty){
			// P(queuUsed)
			int id = q.dequeue(); // sacar de la cola
			V(queuUsed)
			P(impresorasUsed); 
			P(semaforoImpresoras); // solicitar la impresora
			impresoras[id] = stockImpresoras.dequeue(); // sacar de la cola de impresoras
			V(semaforoImpresoras); // liberar la cola de impresoras

			V(semaforoPersonas[id]); // despierto a la persona INDICADA
			cantUsos++;
		}else{
			V(queuUsed)
		}
	}
}

//otra solucion

ColaOrdenDeLlegada cola[N];
ColaImpresoras impresorasLibres[5] = (0, 1, 2, 3, 4);
int impresoraPorPersona[N] = ([N] -1);
sem accesoColaPedidos = 1; sem accesoColaImpresoras = 1; sem pedidosImpresion = 0; sem catidadImpresoras = 5; sem espera[N] = ([N] 0);

Process Persona[id: 0 ... N-1] {
    Documento documento;
    P(accesoColaPedidos);
    cola.push(id);
    V(accesoColaPedidos);
    V(pedidosImpresion); // aviso que hay un pedido de impresión
    P(espera[id]); // espero aviso de que se puede usar la impresora
    Imprimir(documento, impresoraPorPersona[id]); // imprime en la impresora asignada
    P(accesoColaImpresoras);
    impresorasLibres.push(impresoraPorPersona[id]);
    V(accesoColaImpresoras); // repongo la impresora usada
    V(catidadImpresoras); // aviso que hay una impresora libre
}

Process Coordinador {
    int sig; int impresora;
    for (int i = 0; i < N; i++) {
        P(pedidosImpresion); // espero pedido de impresión
        P(accesoColaPedidos);
        sig = cola.pop();
        V(accesoColaPedidos);
        P(cantidadImpresoras); // espero que haya impresoras libres
        P(accesoColaImpresoras);
        impresora = impresoras.pop();
        V(accesoColaImpresoras);
        impresoraPorPersona[sig] = impresora; // asigno impresora al siguiente
        V(espera[sig]); // habilito el uso de la impresora
    }
}