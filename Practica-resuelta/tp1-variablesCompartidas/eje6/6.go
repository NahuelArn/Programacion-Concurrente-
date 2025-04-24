// Resolver con SENTENCIAS AWAIT (<> y/o <await B; S>) el siguiente problema. En un 
// examen final hay P alumnos y 3 profesores. Cuando todos los alumnos han llegado comienza 
// el examen. Cada alumno resuelve su examen, lo entrega y espera a que alguno de los 
// profesores lo corrija y le indique la nota. Los profesores corrigen los exámenes respectando 
// el orden en que los alumnos van entregando.  


{
	P alumnos
	3 profesores
	"Cuando todos los alumnos han llegado comienza el examen" -> ""Comienza el examen""-> lo determina un profesor
	"Cada alumno resuelve su examen, 
	lo entrega y espera a que alguno de los profesores lo corrija y le indique la nota" -> "El profesor corrige el examen y le indica la nota al alumno"
	necesito linkear al alumno-profesor- a travez de su examen
}

int cantActual = 0;
boolean comenzoExamen = false;
Queue<id> q = new Queue<>(); // cola de examenes
int notas[P] = -1; // inicializadas todas en -1
int corregidos = 0; // cantidad de alumnos corregidos
process alumno(id:= 0..P-1){ //limito mis procesos a P alumnos
	<cantActual++;>
	<await (comenzoExamen)>
	//resolviendo examen
	q.enqueue(id); //entrego el examen
	//espero la nota
	if(await (notas[id] != -1)){
		//veo nota
		verNota(notas[id]); //veo la nota
	}
}

process profesor(id: = 0..2){ //limito mis procesos a 3 profesores
	<await (cantActual == P)>; //espero a que lleguen todos
	<if comenzoExamen == false; comenzoExamen = true;> //comienza el examen //para q no reasigne lo mismo 3 veces
	// int corregidos = 0;
	while (corregidos < P){ //espero a que lleguen todos
		int actualAlumno = q.dequeue(); //saco el examen
		//correccion
		corrige(int nota,actualAlumno); //correccion del examen
		//entrego la nota
		<notas[actualAlumno] = nota;> //entrego la nota
		<corregidos++;> //incremento la cantidad de corregidos
	}
}









//poca concurrencia
const Prof = 3;
const Alu = x; // x es un valor N
bool comenzoExamen = false;
int cantActual = 0;
int cantCorregidos = 0;
Queue<Examen> q = new Queue<Examen>(P); // cola de examenes
int nota = -1; // nota del examen
Examen {
	int id; // id del alumno
}
process esperandoAlumnos[id = 0..P-1]{
	while P != cantActual{
		<cantActual++;>
		<await (comenzoExamen);> //aca se van quedando los alumnos "ponele q esperan sentados"
		resolviendoExamen(); //esto es atomico, cada alumno resulve su examen
		//entrego el examen
		<entregarExamen(q,id);> //entrego el examen
		//espero la nota
		<await (nota != -1);> //espero la nota
		//recibo la nota
		verNota(nota); //veo la nota
		<nota = -1;> //la reinicio para el proximo examen
	}

}

process profesor[id = 0..2]{ //van a corregir tantos examenes como alumnos haya
	//espero a que lleguen todos
	while (true){
		//no hago nada
		if(p == cantActual)){
			<comenzoExamen = true;>
			//comienza el examen
			//salgo del while
			<break;>
		}
	}
	while (cantCorregidos < P){
		if(q != q.isEmpty()){
			<cantCorregidos +=1;>
			<Examen examen = q.dequeue();> //saco el examen
			//correccion
			corrige(examen); //correccion del examen
			//entrego la nota
			<examen.nota = nota;> //entrego la nota
		}
	}

// process corrigiendo[id = 0..2]{ //van a corregir tantos examenes como alumnos haya
// 	while cantCorregidos < P{
// 		<cantCorregidos +=1;>
		
// 	}
}



//Esta solucion es media rara, es como reservar N sillas para N estudiantes para la muestra de examenes y al final solo
// hay capacidad para que como maximo vean 3 alumnos la nota al mismo tiempo
//pero creo que esta solucion incrementa la concurrencia, ya que el alumno puede estar viendo su nota y el profesor ya tener disponible 
//para corregir otro examen / informar la nota de otro alumno
//como que el profesor no queda esperando una respuesta del alumno a la correccion del examen
- Ejercicio 06

int llegados = 0;
bool comenzar = false;
cola cola_entregados[P];
int notas[P] = -1;
int entregados;
int corregidos = 0;

Process Alumno[ id = 1 .. P ] {
    <llegados ++>;
    <await comenzar>;
    // Hace el examen
    <cola_entrega.push(id)>;
    <await notas[id] <> -1>;
    // ver nota
}

Process Profesor:: [ id = 1 .. Q ]{
  int id_alumno = -1;
  <await (llegados = P)>;
  <if (comenzar = false) comenzar = true>;
  while (corregidos < P) {
    <if (not cola_entrega.isEmpty()); id_alumno = cola_entrega.pop()>;
    // corregir examen
    if (id_alumno <> -1) {
      notas[id_alumno] = generarNota();
      <corregidos++>;
    }
  }
}




































int llegaron = 0;
Cola c;
Double[P-1] notas; // inicializadas todas en -1


Process rendidor[id 0..P-1]{
	while(true){
		//espero a que lleguen todos
	< llegaron := llegaron +1;>
	<await (llegaron == P)>
//realizo examen
//entrego
<termineExamen(c,id)>
	//espero mi nota
	<await (notas[id] != -1)>
}
	
}

Process profesor[id 0..2]{
	while(true){
	int siguiente = -1;
	//espero a que entregue algun alumno
	<await (colaNoVacia); siguiente = sacarAlumno(c)>
//corijo
//entrego nota
	notas[siguiente] = nota;
}
}