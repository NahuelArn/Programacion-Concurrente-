// Resolver con SENTENCIAS AWAIT (<> y/o <await B; S>) el siguiente problema. En un 
// examen final hay P alumnos y 3 profesores. Cuando todos los alumnos han llegado comienza 
// el examen. Cada alumno resuelve su examen, lo entrega y espera a que alguno de los 
// profesores lo corrija y le indique la nota. Los profesores corrigen los exámenes respectando 
// el orden en que los alumnos van entregando.  

const Prof = 3;
const Alu = x; // x es un valor N
bool comenzoExamen = false;
int cantActual = 0;
int cantCorregidos = 0;

process esperandoAlumnos[id = 0..P-1]{
	while P != cantActual{
		<cantActual++;>
		<await (comenzoExamen);>
	}
}

process corrigiendo[id = 0..2]{ //van a corregir tantos examenes como alumnos haya
	while cantCorregidos < P{
		<cantCorregidos +=1;>
		
	}
}
