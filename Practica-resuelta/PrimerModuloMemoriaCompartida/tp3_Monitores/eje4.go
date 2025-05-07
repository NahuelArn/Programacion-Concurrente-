En un corralón de materiales se deben atender a N clientes de acuerdo con el orden de llegada. 
Cuando  un  cliente  es  llamado  para  ser  atendido,  entrega  una  lista  con  los  productos  que 
comprará, y espera a que alguno de los empleados le entregue el comprobante de la compra 
realizada. 
a. Implemente una solución considerando que el corralón tiene un único empleado y cada 
cliente hace un único pedido. El empleado debe terminar su ejecución



Llega el cliente: 
	->pasa directo
	->espera en la fila

Llaman a cliente
	->atiendo al que recien esta llegando
	-> atiendo si ya habia alguien en la fila

Monitor Compra{
	cond cliente;
	cond empleado;

	cond recibioComprobante;
	cond hayComprobante;
	cond hayLista;

	int esperando;
	Comprobante comprobante;
	Lista listaActual;

	procedure entrando(listaProductos: in Lista){
		esperando++; //por a o por b vas a esperar... si no se ejecuto anteriormente atenderCliente vas a tener q esperar
		signal(empleado);
		wait(cliente); //espero que alguien me atienda
		listaActual = listaProductos;
		signal(hayLista); //doy aviso que mi lista esta disponible
	}

	procedure esperarComprobante(){
		wait(hayComprobante); // espero que me den el comprobante
		signal(recibioComprobante);	//aviso que ya retire el comprobante
	}

	procedure atenderCliente{
		if(esperando == 0){
			wait(empleado); //si no hay nadie esperando me duermo
		}
		signal(cliente); //atiendo a un cliente
		wait(hayLista);
		comprobante = generandoComprobante(listaActual);
		esperando --;
		signal(hayComprobante);
		wait(recibioComprobante); //espero confirmacion que ya recibio el comprobante
	}

}

Process Cliente [id: 0..N-1]{
	Lista listaProductos;
	Compra.entrando(listaProductos);
	Compra.esperarComprobante();
}

Process Empleado{
	for(int i = 0; i < N; i++){
		Compra.atenderCliente();
	}
}

//=========================================================================================

En un corralón de materiales se deben atender a N clientes de acuerdo con el orden de llegada. 
Cuando  un  cliente  es  llamado  para  ser  atendido,  entrega  una  lista  con  los  productos  que 
comprará, y espera a que alguno de los empleados le entregue el comprobante de la compra 
realizada.

b. Implemente una solución considerando que el corralón tiene E empleados (E > 1). Los 
empleados no deben terminar su ejecución.


Monitor Conversacion{
	int cantEmpleadosDisponibles=0;
	int esperandoAserAtendido =0;
	cond esperaAtencion;
	Queue stockEmpleados;

	procedure entrando(idE : out int){
		if(cantEmpleadosDisponibles == 0){ //si no hay empleados disponibles, espero
			esperandoAserAtendido++;
			wait(esperaAtencion);
		}else{	//si hay empleados disponibles
			cantEmpleadosDisponibles --;
			int idE = stockDeEmpleados.pop();
		}
	}

	procedure atendiendo(idE : in int){ //paso el id del empleado que ya se libero
		stockDeEmpleados.push(idE);
		if(esperandoAserAtendido > 0){
			esperandoAserAtendido--;
			signal(esperaAtencion); //despierto al cliente
		}else{
			cantEmpleadosDisponibles++;
		}
	}
}

Monitor Accion{
	boolean llegoCliente = false;
	cond cliente;
	cond empleado;
	Lista lista;
	Comprobante comprobante;
	procedure realizandoAtencion(listaP: in Lista, comprobanteR : out Comprobante ){
		llegoCliente = true;
		lista = listaP;
		signal(empleado);
		wait(cliente);
		comprobanteR = comprobante;
		signal(empleado);
	}

	procedure esperarLista(listaR : out Lista){
		if(!llegoCliente){
			wait(empleado);
		}
		listaR = lista;
	}
	procedure entregarComprobante(comprobanteI : in Comprobante){
		comprobante = comprobanteI;
		signal(cliente);
		wait(empleado);
		lista = false; //el proximo cliente no llego todavia
	}
}

Process Cliente [id: 0..N-1]{
	Lista listaProductos;
	Comprobante comprobante;
	int idE;
	Compra.entrando(idE);
	Accion[idE].realizandoAtencion(listaProductos,comprobante); 
	//ya tengo mi comprobante
}

Process Empleado[id: 0..N-1]{
	Lista lista;
	Comprobante comprobante;
	while(true){
		Compra.atenderCliente();
		Accion[id].esperarLista(lista);
		comprobante = generarComprobante(lista);
		Accion[id].entregarComprobante(comprobante);
	}
}



//C

Process Empleado[id:1..E]{
	text lista
	termino= false
	while(!termino)
		corralon.llamarCliente(id, termino)
		if (!termino)
			mostrador[id].esperarLista(lista)
			comprobante = generarComprobante(lista) 
			mostrador[id].entregarComprobante(comprobante)
}
Monitor Corralon{
	cola elibres;
	cond esperaC;
	int esperando = 0, cantlibres = 0, atendidos = 0 //llevo la cant de antendidos

	Procedure llegar(idE: out int){
		if(cantLibres == 0) 
			esperando++
			wait(esperaC)
		else cantlibres-- 
		elibres.pop() 
	}
	Procedure llamarCliente(idE: in int, termino: out bool){
		if (atendidos == N)
			termino = true //si atendidos = cant de clientes termino
		else
			atendidos++ //sino sumo atendidos
			elibres.push(idE) 
			if(esperando > 0) 
				esperando--
				signal(esperaC)
			else cantlibres++ 
	}
}














































	
// Monitor Compra{
// 	int cantEnEspera = 0;
	
// 	cond cliente;
// 	cond empleado;
// 	cond comprobante

// 	Queue lista;
// 	Queue comprobantes;

// 	procedure esperarLLamado(listarda : in ListaCompra){
// 		cantEnEspera ++;
// 		if(cantEnEspera != 0){ //si ya hay fila
// 			// cantEnEspera ++;
// 			wait(cliente);
// 		}
// 		lista.push(listarda);
// 		signal(empleado);
// 		wait(comprobante);

// 		mirandoComprobante(comprobantes.pop());
// 	}

// 	procedure atenderCliente(){
// 		if(cantEnEspera != 0){
// 			signal(cliente);
// 		}
// 		wait(empleado);
// 		ListaCompra listaDeCompras = lista.pop();
// 		Comprobante comprobante = armandoPedido(listaDeCompras); //toma el pedido //sarasa
// 		comprobantes.push(comprobante);
// 		signal(comprobante);
// 	}

// }

// Process Cliente [id: 0..N-1]{
// 	ListaCompra listarda = generandoLoqueNecesito();
// 	Compra.esperarLlamado(listarda);
// 	// esperarComprobante();
// }

// Process Empleado{
// 	for(int i = 0; i < N; i++){
// 		Compra.atenderCliente();
// 	}
// }







//comportamiento encapsulado

// Llega el cliente: 
// 	->pasa directo
// 	->espera en la fila

// Llaman a cliente
// 	->atiendo al que recien esta llegando
// 	-> atiendo si ya habia alguien en la fila


// Monitor Compra{
// 	int cantEnEspera = 0;
	
// 	cond cliente;
// 	cond empleado;
// 	cond comprobante

// 	Queue lista;
// 	Queue comprobantes;

// 	procedure esperarLLamado(listarda : in ListaCompra){
// 		if(cantEnEspera != 0){ //si ya hay fila
// 			cantEnEspera ++;
// 			wait(cliente);
// 		}
// 		lista.push(listarda);
// 		signal(empleado);
// 		wait(comprobante);

// 		mirandoComprobante(comprobantes.pop());
// 	}

// 	procedure atenderCliente(){
// 		if(cantEnEspera == 0){
// 			wait(empleado);
// 			ListaCompra listaDeCompras = lista.pop();
// 			Comprobante comprobante = armandoPedido(listaDeCompras); //toma el pedido //sarasa
// 			comprobantes.push(comprobante);
// 			signal(comprobante);
// 		}else{
// 			signal(cliente);
// 			wait(empleado);
// 			ListaCompra listaDeCompras = lista.pop();
// 			Comprobante comprobante = armandoPedido(listaDeCompras); //toma el pedido //sarasa
// 			comprobantes.push(comprobante);
// 			signal(comprobante);
// 		}
// 		//en este punto 
// 	}

// }

// Process Cliente [id: 0..N-1]{
// 	ListaCompra listarda = generandoLoqueNecesito();
// 	Compra.esperarLlamado(listarda);
// 	// esperarComprobante();
// }

// Process Empleado{
// 	for(int i = 0; i < N-1; i++){
// 		Compra.atenderCliente();
// 	}
// }











// Monitor Compra{
// 	int cantEnEspera = 0;
	
// 	cond Cliente;
// 	cond Empleado;
	
// 	Queue lista;

// 	procedure esperarLLamado(listarda : in ListaCompra){
// 		if(cantEnEspera != 0){ //si ya hay fila
// 			cantEnEspera ++;
// 			wait(Cliente);
// 		}else{
// 			signal(empleado);
// 		}
// 		lista.push(listarda);
// 	}

// 	procedure atenderCliente(){
// 		if(cantEnEspera == 0){
// 			wait(Empleado);
// 		}else{
// 			signal(Cliente);
// 		}
		
// 	}

// }

// Process Cliente [id: 0..N-1]{
// 	ListaCompra listarda = generandoLoqueNecesito();
// 	Compra.esperarLlamado(listarda);
// 	// esperarComprobante();
// }

// Process Empleado{
// 	for(int i = 0; i < N-1; i++){
// 		Compra.atenderCliente();
// 	}
// }
















// Orden de llegada, Unico Empleado
//

Monitor Compra{
	boolean libre = true;
	Queue esperaComprobante;
	
	cond cliente;
	cond empleado;

	cond entregaComprobante;
	cond reciboComprobante
	procedure esperarLlamado(){
		if(!libre){
			signal(empleado);
			wait(cliente);
		}else{
			libre = false;
		}
		//entrego lista y espero comprobante
		
		//entrego lista..
		esperaComprobante.push(); //pusheoComprobante
		signal(entregaComprobante);

		wait(reciboComprobante);
	}

	procedure atenderCliente{
		if(esperandoClientes.isEmpty()){
			wait(empleado);
		}
		signal(cliente); //aviso al proximo cliente que pase
		wait(entregaComprobante);
		Lista lista = esperaComprobante.pop();
		//arma el pedido
		signal(reciboComprobante);
	}

}

Process Cliente [id: 0..N-1]{
	Compra.esperarLlamado();
	// esperarComprobante();
}

Process Empleado{
	for(int i = 0; i < N-1; i++){
		Compra.atenderCliente();
	}
}











































