// b) Modifique la solución de (a) para el caso en que se deba respetar el orden de llegada. 


int N = ...;
boolean ocupada = false;

Persona queue[N];

process imprimiendo[id = 0..N-1]{ // id es el identificador de la persona
	while(true){	 //representa que los procesos están continuamente solicitando y devolviendo recursos.
		<queue.push(id);> //pushea a la primer persona que llega
		<await (!ocupada && !queue.empty()); 
			ocupada = true;
			Persona pActual = queue.pop();
		>

		Imprimir(pActual.getDocumento(documento));

		<ocupada = false;>
	}	
}



ColaPersonas cola;
int siguiente = -1;

Process Persona[id: 0 ... N-1] {
    Documento documento;
    <if (siguiente == -1) 
			siguiente = id;
    else cola.push(id);>
    	<await (siguiente == id);>
    Imprimir(documento);
    <if (cola.isEmpty()) 
			siguiente = -1;
    else 
			siguiente = cola.pop();>
}