1. Suponga  que  N  clientes  llegan  a  la  cola  de  un  banco  y  que  serán  atendidos  por  sus 
empleados.  Analice  el  problema  y  defina  qué  procesos,  recursos  y  comunicaciones  serán 
necesarios/convenientes  para  resolver  el  problema.  Luego,  resuelva  considerando  las 
siguientes situaciones: 
a. Existe un único empleado, el cual atiende por orden de llegada. 
b. Ídem a) pero considerando que hay 2 empleados para atender, ¿qué debe 
modificarse en la solución anterior? 
c. Ídem  b)  pero  considerando  que,  si  no  hay  clientes  para  atender,  los  empleados 
realizan  tareas  administrativas  durante  15  minutos.  ¿Se  puede  resolver  sin  usar 
procesos adicionales? ¿Qué consecuencias implicaría?


/div----------------------------------------------------

//=========================[Problema-A]=========================
a. Existe un único empleado, el cual atiende por orden de llegada.


chan Fila(int)

Process Empleado{
	int id;
	for(int i = 0; i< N; i++){
		receive Fila(id) //si llega primero, se queda esperando automaticamente
		atendiendoEmpleado(id);
		//siguiente Iteraciin
	}
}

Process Cliente[id: 0..N-1]{
	send Fila(id) //se encola en la fila, el canal tiene estructura FIFO
}

//=========================[Problema-B]=========================
b. Ídem a) pero considerando que hay 2 empleados para atender, ¿qué debe 
modificarse en la solución anterior?

chan Fila(int)

Process Empleado[idE: 0..2-1]{
	int id;
	while(true){
		receive Fila(id) //si llega primero, se queda esperando automaticamente
		atendiendoEmpleado(id);
		//siguiente Iteraciin
	}
}

Process Cliente[id: 0..N-1]{
		send Fila(id) //se encola en la fila, el canal tiene estructura FIFO
}

//=========================[Problema-C]=========================
c. Ídem  b)  pero  considerando  que,  si  no  hay  clientes  para  atender,  los  empleados 
realizan  tareas  administrativas  durante  15  minutos.  ¿Se  puede  resolver  sin  usar 
procesos adicionales? ¿Qué consecuencias implicaría?

chan Fila(int) //simula la fila que hacen los clientes
chan Siguiente[2](int); //lo utilizo para saber si tengo un cliente asignado
chan Pedido(int); //canal para notificar al coordinador que estoy disponible (Empleado -> Coordinador)

Process Coordinador(){ //es el encargado de fijarse si hay alguien en la fila
	int idE, idC;
	// int respuesta;

	while(true){
		receive Pedido(idE); //agarro a un empleado
		if(Fila.isEmpty()){
			// respuesta = "vacio"; //como utilizo un int, para representar que no hay nadie haciendo la fila devuelvo un -1
			idC = -1;		
		}else{
			receive Fila(idC) //hay un cliente esperando, retorno su id valido
		}
		send Siguiente[idE](respuesta);
	}
}

Process Empleado[idE: 0..2-1]{
	int idC;
	while(true){
		send Pedido(idE);
		receive Siguiente[idE](idC); //si tengo un cliente asignado
		
		if(idC != -1){
			atiendoCliente(idC);
		}else{
			delay(15); //15 minutos
		}
		//siguiente Iteraciin
	}
}

Process Cliente[id: 0..N-1]{
		send Fila(id) //se encola en la fila, el canal tiene estructura FIFO
}

//------------------c---------------------------

//error
chan Fila(int)

Process Empleado[idE: 0..2-1]{
	int id;
	while(true){
		if(Fila.isEmpty()){ //pero con esta solucion tengo el problema de 1 solo Cliente en fila y llegan los 2 empleados a querer atenderlo (BusyWaiting)
			delay(15); //15 minutos
		}else{
			receive Fila(id) //si llega primero, se queda esperando automaticamente
			atendiendoEmpleado(id);
		}
		//siguiente Iteraciin
	}
}

Process Cliente[id: 0..N-1]{
		send Fila(id) //se encola en la fila, el canal tiene estructura FIFO
}