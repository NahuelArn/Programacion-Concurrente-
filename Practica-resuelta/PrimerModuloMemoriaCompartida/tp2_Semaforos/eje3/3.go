// 3.  Un sistema de control cuenta con 4 procesos que realizan chequeos en forma 
// colaborativa. Para ello, reciben el historial de fallos del día anterior (por simplicidad, de 
// tamaño  N).  De  cada  fallo,  se  conoce  su  número  de  identificación  (ID)  y  su  nivel  de 
// gravedad (0=bajo, 1=intermedio, 2=alto, 3=crítico). Para cada item realice una solución 
// adecuada a lo pedido: 
	// a. Se  debe  imprimir  en  pantalla  los  ID  de  todos  los  errores  críticos  (no  importa  el 
	// orden). 
	// b. Se  debe  calcular  la  cantidad  de  fallos  por  nivel  de  gravedad,  debiendo  quedar  los 
	// resultados en un vector global. 
	// c. Ídem b) pero cada proceso debe ocuparse de contar los fallos de un nivel de gravedad 
	// determinado. 


//===========================Parte-A===========================

Fallo = {
	int id: 0..N-1; //id del fallo
	int gravedad: 0..3; //gravedad del fallo
}

// sem procesosFree = 4; //cantidad de procesos disponibles
List <Fallo> listaFallos; //lista de fallos
sem usandoListaFallos = 1; //semáforo para proteger la lista de fallos
int dimFHistorial = N; //dimensio del historial de fallos
int cant = 0;


process procesando [id: 0..3]{ // --> procesos
	P(usandoListaFallos); //pido el fallo
	while (cant < dimFHistorial-1) { //mientras haya fallos en el historial
		Fallo fallo = listaFallos.remove(0); //saco el primer fallo de la lista (fallo es local al proceso)
		cant++; //aumento la cantidad de fallos que se imprimieron
		V(usandoListaFallos); //libero la lista de fallos para que otros procesos puedan usarla
		if (fallo.gravedad == 3) { //si el fallo es critico (como es local al proceso, no afecta a los demas)
			imprimir(fallo.id); //imprimo el id del fallo
		}
		P(usandoListaFallos); //pido el fallo
	}
	V(usandoListaFallos); //pido el fallo
}


//===========================Parte-B===========================

Fallo = {
	int id: 0..N-1; //id del fallo
	int gravedad: 0..3; //gravedad del fallo
}

List <Fallo> listaFallos; //lista de fallos
sem usandoListaFallos = 1; //semáforo para proteger la lista de fallos
int dimFHistorial = N; //dimensio del historial de fallos
int cant = 0;
int contadorFallos[4][0]; //vector de contadores de fallos por gravedad (0,1,2,3)
sem usandoContadores = 1; //semáforo para proteger el vector de contadores
process procesando [id: 0..3]{ // --> procesos
	P(usandoListaFallos); //pido el fallo
	while (cant < dimFHistorial-1) { //mientras haya fallos en el historial
		Fallo fallo = listaFallos.remove(0); //saco el primer fallo de la lista (fallo es local al proceso)
		cant++; //aumento la cantidad de fallos que se imprimieron
		V(usandoListaFallos); //libero la lista de fallos para que otros procesos puedan usarla
		P(usandoContadores); //pido el vector de contadores
		contadorFallos[fallo.gravedad]++; //aumento el contador de fallos por gravedad (como es local al proceso, no afecta a los dem
		V(usandoContadores); //libero el vector de contadores para que otros procesos puedan usarlo
		P(usandoListaFallos); //pido el fallo
	}
	V(usandoListaFallos); //pido el fallo
}

//===========================Parte-C===========================

Fallo = {
	int id: 0..N-1; //id del fallo
	int gravedad: 0..3; //gravedad del fallo
}

List <Fallo> listaFallos; //lista de fallos
sem usandoListaFallos = 1; //semáforo para proteger la lista de fallos
int dimFHistorial = N; //dimensio del historial de fallos
int cant = 0;
int contadorFallos[4][0]; //vector de contadores de fallos por gravedad (0,1,2,3)
sem usandoContadores = 1; //semáforo para proteger el vector de contadores
process procesando [id: 0..3]{ // --> procesos
	P(usandoListaFallos); //pido el fallo
	while (cant < dimFHistorial-1) { //mientras haya fallos en el historial
		Fallo fallo = listaFallos.remove(0); //saco el primer fallo de la lista (fallo es local al proceso)
		cant++; //aumento la cantidad de fallos que se imprimieron
		V(usandoListaFallos); //libero la lista de fallos para que otros procesos puedan usarla
		P(usandoContadores); //pido el vector de contadores
		if(fallo.id == fallo.gravedad){
			contadorFallos[fallo.gravedad]++; //aumento el contador de fallos por gravedad (como es local al proceso, no afecta a los dem
		}else{
			P(usandoListaFallos); //pido el fallo
			listaFallos.add(fallo); //agrego el fallo a la lista de fallos
			cant--; //disminuyo la cantidad de fallos que se imprimieron
			V(usandoListaFallos); //libero la lista de fallos para que otros procesos puedan usarla
		}
		V(usandoContadores); //libero el vector de contadores para que otros procesos puedan usarlo
		P(usandoListaFallos); //pido el fallo
	}
	V(usandoListaFallos); //pido el fallo
}


//otra solucion

cola historial[N]
int ocupado = 0
sem mutex = 1
list contador[4][0]
Process chequear[id: 0..3]{
    P(mutex)
    while(ocupado < N){
        fallo = historial.pop()
        ocupado++
        v(mutex)
        if(fallo.getNivel() == id)
            contador[id]++
        else
            p(mutex)
            ocupado--
            historial.push(fallo)
            v(mutex)
        P(mutex)
    }
    V(mutex)
}
















// fallo = {
// 	int id: 0..N-1; //id del fallo
// 	int gravedad: 0..3; //gravedad del fallo
// }
// sem procesosFree = 4; //cantidad de procesos disponibles

// process procesando [id: 0..3]{ // --> proceso

// }