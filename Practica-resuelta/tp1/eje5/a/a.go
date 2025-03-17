// En  cada  ítem  debe  realizar  una  solución  concurrente  de  grano  grueso  (utilizando  <>  y/o 
// <await B; S>) para el siguiente problema, teniendo en cuenta las condiciones indicadas en el 
// item. Existen N personas que deben imprimir un trabajo cada una.  


// a) Implemente  una  solución  suponiendo  que  existe  una  única  impresora  compartida  por 
// todas las personas, y las mismas la deben usar de a una persona a la vez, sin importar el 
// orden. Existe una función Imprimir(documento) llamada por la persona que simula el uso 
// de la impresora. Sólo se deben usar los procesos que representan a las Personas.

int N = ...;
boolean ocupada = false;


process imprimiendo[id = 0..N-1]{ // id es el identificador de la persona
	while(true){	 //representa que los procesos están continuamente solicitando y devolviendo recursos.
		<await (!ocupada); 
			ocupada = true;
		>

		Imprimir(documento);

		<ocupada = false;>
	}	
}

//Alternativa? no es necesario simular la peticion de recursos y devolucion de los mismos, ya que la impresora es un recurso compartido ?
process persona[id: 0..N-1] {
    Documento documento;
    <imprimir(documento); >
}