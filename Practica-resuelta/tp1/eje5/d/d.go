// Modifique la solución de (b) para el caso en que además hay un proceso Coordinador que 
// le indica a cada persona que es su turno de usar la impresora.

int N = ...;
boolean ocupada = false;
Persona pActual = null;

Persona queue[N];

process coordinador[]{
	while true {
			<await (!ocupada);
				ocupada = true;
				pActual = queue.pop();
			>
	}
}

process imprimiendo[id = 0..N-1]{ // id es el identificador de la persona
	while(true){	 //representa que los procesos están continuamente solicitando y devolviendo recursos.
		<queue.push(id);> //pushea a la primer persona que llega
		<await (pActual != null);>

		Imprimir(pActual.getDocumento(documento));

		<ocupada = false;>
	}	
}

