// c) Modifique la solución de (a) para el caso en que se deba respetar el orden dado por el 
// identificador del proceso (cuando está libre la impresora, de los procesos que han 
// solicitado su uso la debe usar el que tenga menor identificador). 



int N = ...;
boolean ocupada = false;

Persona queue[N];

process imprimiendo[id = 0..N-1]{ // id es el identificador de la persona
	while(true){	 //representa que los procesos están continuamente solicitando y devolviendo recursos.
		<queue.pushOrdenado(id);> //pushea a la primer persona que llega //pusheaOrdenado va insertando de menor a menor segun id
		<await (!ocupada && !queue.empty()); 
			ocupada = true;
			Persona pActual = queue.pop();
		>

		Imprimir(pActual.getDocumento(documento));

		<ocupada = false;>
	}	
}