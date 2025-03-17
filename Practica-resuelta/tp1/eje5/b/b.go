int N = ...;
boolean ocupada = false;

Persona queue[N];

process imprimiendo[id = 0..N-1]{ // id es el identificador de la persona
	while(true){	 //representa que los procesos están continuamente solicitando y devolviendo recursos.
		<await (!ocupada && !queue.empty()); 
			ocupada = true;
			Persona pActual = queue.pop();
		>

		Imprimir(pActual.getDocumento(documento));

		<ocupada = false;>
	}	
}