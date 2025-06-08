4.  Simular  la  atención  en  un  locutorio  con  10  cabinas  telefónicas,  el  cual  tiene  un  empleado 
que se encarga de atender a N clientes. Al llegar, cada cliente espera hasta que el empleado 
le  indique  a  qué  cabina  ir,  la  usa  y  luego  se  dirige  al  empleado  para  pagarle.  El  empleado 
atiende a los clientes en el orden en que hacen los pedidos, pero siempre dando prioridad a 
los  que  terminaron  de  usar  la  cabina.  A  cada  cliente  se  le  entrega  un  ticket  factura.  Nota: 
maximizar la concurrencia; suponga que hay una función  Cobrar() llamada por el empleado 
que simula que el empleado le cobra al cliente. 


//Prioriza antes de atender a un nuevo cliente, si ya termino un cliente anterior... Prioriza que le pague el/los cliente que ya terminaron

//tiene sentido atender a todos los libres que encuentre.. hago un barrido de los q terminaron y despues atiendo a un nuevo cliente
chan fila(int);
chan cabinaAsignada[N](int);

chan termine(N, int);

chan factura[N](text);

//cabinas deberia ser una queue, modela mejor el problema, quito abstraccion

//abstraccion 
Process Empleado(){
	int cabinas[10]; 
	while(true){
		int idCliente, idClienteTermino;
		int cabinaLibre;
		receive fila(idCliente);
		//Primero me fijo si tengo ya termino alguien
		while(!termine.isEmpty()){ //hago mi barrido
			receive termine(idClienteTermino, cabinaLibre);
			ponerEnLibreCabina(cabinaLibre, cabinas); //funcion que dado "cabinaLibre" que es el id, de la cabina que esta libre 
			//insertar la cabina libre en el vector de cabinas para reponer stock, abstracion

			Text ticked = Cobrar(idClienteTermino, cabinaLibre);

			send factura[idClienteTermino](ticked);
		}
		//en este punto, si tenia clientes esperando, su ticket.. ya los atendi y quedaron las cabinas libres
		if(hayCabinaLibre(cabinas)){ //abstraccion
			cabinaAsignada[idCliente](cabinas[idCabinaLibre(cabinas)]); //idCabinaLibre(cabinas) me devuelve la pos de la cabina libre
		}
	}
}

Process Cliente(id: 0..N-1){
	while(true){
		int numCabina; Text text;
		send fila(id);
		receive cabinaAsignada[id](numCabina);

		usandoCabina(cabinaAsignada[id](numCabina));

		send termine(id, numCabina);
		receive factura[id](text);
	}
}
	









// Solucion, con bussy waiting
//que pasa si no tengo Cabinas libres y tengo solicitudes
//esta constantemente entrando en el else if //la solucion de arriba tiene el mismo problema
chan solicitarCabina(int);
chan obtenerCabina[N](int);
chan pagarEmpleado(int,int);
chan recibirComprobante[N](int);

process cliente[id: 1..N]{
    int cabina, ticket;
    send solicitarCabina(id);
    receive obtenerCabina[id](cabina);
    //USO CABINA
    send pagarEmpleado(id, cabina);
    receive recibirComprobante[id](ticket);
}

process empleado{
    int idCliente, cantCabinas=10, ticketAct, cabinaUsada;
    cola colaCabina=(10 cabinas); 
    while (true){
        if(!pagarEmpleado.empty()){
            receive pagarEmpleado(idCliente, cabinaUsada);
            ticketAct = Cobrar(cabinaUsada);
            cantCabinas++;
            colaCabina.push(cabinaUsada);
            send recibirComprobante[idCliente](ticketAct);
        }else{
            if((!solicitarCabina.empty()) and (cantCabinas>0)){
                receive solicitarCabina(idCliente);
                cantCabinas--;
                send obtenerCabina[idCliente](colaCabina.pop());
            }
        }
    }
}
