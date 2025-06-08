//semaforos
1. Se debe simular el uso de un sistema virtual de venta de entradas para un evento musical. El sistema cuenta con C 
cajeros  virtuales  que  atienden  indefinidamente.  Sin  embargo,  como  la  venta  de  entradas  comienza  a  una  hora 
determinada, sólo atienden a partir del aviso de un Timer. Una vez que reciben dicho aviso, los cajeros atienden de 
acuerdo  con  el  orden  de  llegada  de  los  compradores.  La  atención  consiste  en  recibir  la  solicitud  del  comprador 
(datos para el pago) y responderle si pudo comprar (o no) junto al comprobante de la operación. Para este evento 
se cuenta con E entradas y N compradores, donde cada comprador puede solicitar a lo suma una entrada. Resuelva 
usando SEMÁFOROS.


//semaforo
sem sem_cant=1; 
sem sem_cola=1; 
sem sem_atencion=0; 
sem sem_vendedor[C] = [C]{0}; 
sem sem_comprador[N] = [C]{0}; 
 
int cant_disponible = E; 
Queue solicitudes; 
Tuple respuestas[N]; 
 
Process Timer { 
 // espera la hora indicada 
 wait(); 
 // les avisa a los cajeros que pueden atender 
 for i in 1..C  
  V(sem_vendedor[i]); 
} 
 
Process Comprador [i=1..N] { 
 
 Pago pago = generarPago(); 
Comprobante comp;  
bool pudo_comprar; 
 // ingresar solicitud 
 P(sem_cola); 
 push (id,pago); 
 V(sem_cola); 
 // solicitar atencion 
 V(sem_atencion); 
 // esperar respuesta 
 P(sem_comprador[id]); 
 // verificar si pudo comprar 
 (pudo_comprar,comp) = respuestas[i]; 
  
} 

Process Vendedor [i=1..C] { 
	Pago pago; 
 Comprobante comp;  
	int id; 
	bool pudo_comprar; 
	// esperar aviso del Timer 
	P(sem_vendedor[i]); 
	// atender 
	while (true) { 
	 // esperar solicitud 
	 P(sem_atencion); 
	 // desencolar solicitud 
	 P(sem_cola); 
	 (id, pago) = pop (solicitudes); 
	 V(sem_cola); 
	 pudo_comprar = false; 
	 // analizar disponibilidad  
	 P (sem_cant); 
	 if (cant_disponible > 0) { 
		cant_disponible --; 
		pudo_comprar = true; 
	 } 
	 V(sem_cant); 
	 // cobrar y responder 
	 comp = (pudo_comprar ? cobrar(pago) : null); 
	 respuestas[id] = (pudo_comprar, comp); 
	 // avisarle al comprador 
	 V(sem_comprador[id]); 
	} 
 } 
 //asdasdaskdqwdqpdqwkdllallllllllllllllllllllllllllllllllllllllllllllllllllllllllll

//con monitor
Monitor StockDeEntradas{
	int cant = E;
	Queue entrada;

	procedure estado(cantPedida : in int,hayStock: out boolean, entradaO : out Entrada){
		if(cantPedida > cant){
			cant -= cantPedida;
			entradaO = entrda.pop(); //un ticked de entrada puede valer x N ... este ticked es valido para N personas..
			hayStock = true;
		}else{
			hayStock = false;
		}
	}
}




Monitor LlegandoXCajero[id : 0..C-1]{
	Queue espera;
	cond compradores[N];
	solicitud[N] = ([N] 0);
	cond empleado;

	procedure esperandoAtencion(id : in int, cantEntradas : in int){
		espera.push(id);
		solicitud[id] = cantEntradas;
		signal(empleado);
		wait(compradores[id]); //atencion
		
	}

	procedure atenderClientes(id : in int, cant : out int){
		if(espera.size() == 0){
			wait(empleado);
		}
		id = espera.pop();
		cant = solicitud[id];
	}

}
Monitor Cajero[id : 0..C-1]{
	comprabantes[N] = ([N] null);
	Queue esperandoComprante;
	cond compradoresEsperando;

	procedure esperarComprante(id :in int, comprobante : out Comprobante){
		if(comprobantes[id] == null){
			wait(esperandoComprante[id])
		}
		comprobante = comprobantes[id];
	}

	procedure entregarComprobante(comprobante : in Comprobante, idAtencion : in int){
		comprobantes[idAtencion] = comprobante;
		signal(compradoresEsperando[idAtencion]);
	}

}



Process Comprador(id: 0..N-1){
	Cajero cajero; int idCajero; //cada comprador sabe a que cajero llego... (Asumo)
	int cantEntradas//ya se cuantas entradas voy a querer
	llegandoXCajero[idCajero].EsperarAtencion(id,cantEntradas);
	Comprobante comprobante;
	Cajero[idCajerp].esperarComprante(id,comprobante);
}

Process Empleado[id: 0..C-1]{
	timer(); //pasa X tiempo
	int idAtencion, cant,cantPedida;
	Entrada entradaO; 
	hayStock = true;
	while(hayStock){
		LlegandoXCajero[id].atenderClientes(idAtencion, cant);
		//haciendoComprobante
		StockDeEntradas.estado(cantPedida,hayStock,entradaO);
		if(hayStock()){
			Comprobante comprante = generarComprante(cantPedida,entrada0);
		}else{
			Comprobante comprante = generarCompranteN(cantPedida);
		}
		Cajero[id].entregarComprabante(comprante,idAtencion);
	}
a
}

//a


// Monitor StockDeEntradas{
// 	int cant = E;
// 	Queue entrada;

// 	procedure estado(hayStock: out boolean, entradaO : out Entrada){
// 		if(cant > 0){
// 			cant--;
// 			entradaO;
// 			hayStock = true;
// 		}else{
// 			hayStock = false;
// 		}
// 	}
// }

// Monitor llegando[id : 0..N-1]{
// 	Queue espera;
// 	cond compradores[N];
	
// 	aQueCajeroMeDesignaron[N] = ([N] -1);
// 	procedure esperandoAtencion(id : in int, idCajero: out int){
// 		espera.push(id);
// 		wait(compradores[id]);
// 		idCajero = aQueCajeroMeDesignaron[id];
// 	}
// }
// Monitor Cajero[id : 0..C-1]{

// }



// Process Comprador(id: 0..N-1){
// 	Cajero cajero; int idCajero;
// 	Cajero.EsperarAtencion(id, idCajero);
// }

// Process Empleado[id: 0..C-1]{
// 	timer();
// }






//monitores
2. Existen N personas que desean acceder a un mirador al borde del lago Nahuel Huapi en Bariloche. Como el mirador 
es angosto, sólo puede ser usado por una persona a la vez. Resuelva con MONITORES los dos casos siguientes: 
a. El acceso al mirador es por orden de llegada.

b. El acceso al mirador es por orden de llegada, pero dando prioridad a los mayores de 60 años.