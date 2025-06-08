Importante: las soluciones que se muestran a continuación no son las únicas que se pueden considerar correctas para los 
ejercicios planteados. 



//SEMAFOROS



1. Existe una sala de cine 3D, a las que asisten N personas a ver una película. Antes de entrar a la salsa, los asistentes 
deben retirar los anteojos 3D en la máquina repartidora que se encuentra en la entrada. Se debe simular el uso de 
la  máquina  repartidora  de  anteojos  3D,  con  capacidad  para  A  anteojos  (A  <  N).  Además,  existen  un  repositor 
encargado de reponer los anteojos en la máquina cuando se agotan. Los usuarios usan la máquina según el orden 
de llegada. Cuando les toca usarla, sacan un par de anteojos y luego se dirigen a la sala. En caso de que la máquina 
se  quede  sin  anteojos,  entonces  le  debe  avisar  al  repositor  para  que  cargue  nuevamente  la  máquina  en  forma 
completa. Luego de la recarga, saca un par de anteojos y se retira. Implemente un programa que permita resolver el 
problema  anterior  usando  SEMÁFOROS.  Nota:  maximizar  la  concurrencia;  la  reposición  de  anteojos  no  debe 
impedir que otros asistentes puedan agregarse a la fila.



Queue esperaFila;
boolean filaLibre = true;
sem asistentes[N] = 0;

Queue maquina[A];
sem mutex = 1;
sem empleado = ;

sem esperandoReStockEmpleado = 0;
sem estockOk = 0;
Process Asistente[id: 0..N-1]{
	P(mutex);
	if(filaLibre){
		filaLibre = false;
		V(mutex);
	}else{
		esperaFila.push(id);
		V(mutex);
		P(asistentes[id]);
	}
	//aca llegan de a uno, por eso no es necesario proteger el size
	if(maquina.size() <= 0){
		V(esperandoReStockEmpleado);
		P(estockOk);
	}
	Lentes parDeLentes = maquina.pop();
	P(mutex);
	if(esperaFila.size() > 0){
		int idSig = esperaFila.pop();
		V(asistentes[idSig]);
	}else{
		filaLibre = true;
	}
	V(mutex);
}

Process Empleado{
	int cant = N;
	while(cant > 0 ){
		cant--;
		P(esperandoReStockEmpleado);
		for(int i = 0; i < A; i++{
			Lentes parDeLentes; //el empleado ya tiene lentes para reestockear
			maquina.push(parDeLentes);
		}
		V(estockOk);
	}
}






//MONITORESSS

En  el  Registro  de  la  Propiedad  se  pueden  realizar  4  trámites  administrativos  diferentes. Para  cada  trámite,  hay  un 
puesto de atención específico. Existen 100 personas que se dirigen a la oficina para resolver un trámite particular. 
La persona deja su trámite en el puesto correspondiente y espera a que le entreguen el resultado. El puesto atiende 
a  las  personas  que  le  corresponden  de  acuerdo  con  el  orden  de  llegada.  Implemente  un  programa  que  permita 
resolver  el  problema  anterior  usando  MONITORES.  Notas:  maximizar  la  concurrencia;  todos  los  procesos  deben 
terminar; la función obtenerPuesto() retorna el número de puesto al que la persona debe dirigirse para su trámite; 
la  función  obtenerTrámite()  retorna el  trámite  a  realizar;  la  función  procesarTrámite(t)  procesa  el  trámite recibido 
como entrada y retorna su resultado.


Monitor Tramite[id: 0..4-1]{
	Queue tramites[100] = ([100] null);
	resultadoTramite[100] = ([100] null);
	Queue espera;
	cond personas[100];
	cond empleado;

	procedure realizarTramite(id : in int, tramite : in Tramite, rT : out Tramite){
		tramites[id] = tramite;
		espera.push(id);
		wait(personas[id]);
		rT[id] = resultadoTramite[id]
	}

	procedure atenderTramite(id : out int, tramite : out Tramite){
		if(espera.isEmpty()){
			wait(empleado);
		}
		id = espera.pop();
		tramite = tramites[id];
	}

	procedure entregarTramiteFinalizado(id : in int, tR : in Tramite){
		resultadoTramite[id] = tR;
		signal(personas[id]);
	}
}	
Monitor contador {
	int cant = 0;

	procedure mas1{
		cant++;
	}

	
}


//Aca hay un problema de enunciado, el enunciado nunca aclara que la funcion "Obtener Puesto" va repartir equitativamente los puestos...
//Podria pasar que la funcion asigne a las 90 personas que vayan al puesto 0, y los otros 3 puestos quedarian esperando una cantidad que no va llegar
//
Process Empleado[id : 0..4-1]{
	int id; Tramite t, tR;
	for (int i = 1 to 25){
		Tramite[id].atenderTramite(id,t);
		tR = procesarTramite(t);
		Tramite[id].entregarTramiteFinalizado(id, tR);
	}
}

Process Persona[id 0..N-1]{
	int puesto = obtenerPuesto();
	Tramite rT;
	Tramite[puesto].realizarTramite(id, obtenerTrámite(),rT);

	//mirando tramite terminado rT

}


















// estructuas disponibles sem ... P ..... V

// sem protexFilaEsperaMaquina = 1;

// sem protexMaquina = 1;
// Queue filaEsperaMaquina;
// Queue maquinaAnteojos;

// sem asistente[N] = ([N] 1);

// Anteojo anteojoAsistente[N] = ([N] null);


// Queue anteojosDevueltos;
// sem protexAnteojosDevueltos = 1;

// sem repositor = 1;

// Process Asistente[id 0..N-1]{

// 	P(protexMaquina);
// 	P(protexFilaEsperaMaquina);
// 	if(maquinaAnteojos.isEmpty() || filaEsperaMaquina.size() > 0){ //no hay stock de anteojos || o ya hay gente esperando
// 		V(protexMaquina);
// 		// P(protexFilaEsperaMaquina);
// 		filaEsperaMaquina.push(id);
// 		V(repositor);
// 		V(protexFilaEsperaMaquina);
// 		P(asistente[id]);	//me duermo hasta que recarguen la maquina
// 		P(protexMaquina); //pido la maquina
// 		P(protexFilaEsperaMaquina);
// 		filaEsperaMaquina.pop();
// 	}
// 	V(protexFilaEsperaMaquina);
// 	//aca no es necesario llvar un vector semaforo, por cada Asistente... no puede pasar que 2 procesos accedan a los[id] anteojos de ese cliente
// 	anteojoAsistente[id] = maquinaAnteojos.pop();
// 	V(protexMaquina);

// 	//mirando pelicula
// 	P(protexAnteojosDevueltos);
// 	anteojosDevueltos.push(anteojoAsistente[id]);
// 	V(protexAnteojosDevueltos);
// }

// Process Repositor{
// 	while(true){

// 		P(protexMaquina);
// 		if(!maquinaAnteojos.isEmpty()){
// 			V(protexMaquina);
// 			P(repositor);
// 		}
// 		// V(protexMaquina);
// 		P(protexAnteojosDevueltos);
// 		for(int i = 0; i < anteojosDevueltos.size(); i++){ //repongo stock, con los anteojos que me devolvieron... si nadie devolio recorro 0 veces... y espero a que devuelvan
// 			maquinaAnteojos.pus(anteojosDevueltos.pop());
// 		}
// 		V(protexMaquina);
// 		V(protexAnteojosDevueltos);
// 		P(protexFilaEsperaMaquina);
// 		if(!filaEsperaMaquina.isEmpty()){
// 			VP(asistente[filaEsperaMaquina.pop()]);
// 		}

// 		V(protexFilaEsperaMaquina);

// 	}
// }
