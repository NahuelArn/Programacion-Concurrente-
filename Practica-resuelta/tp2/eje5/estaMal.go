// 5.  Suponga que se tiene un curso con 50 alumnos. Cada alumno debe realizar una tarea y 
// existen  10  enunciados  posibles.  
// 
// Una  vez  que  todos  los  alumnos  eligieron  su  tarea, 
// comienzan a realizarla. 
// 
// Cada vez que un alumno termina su tarea, le avisa al profesor y se 
// queda esperando el puntaje del grupo, el cual está dado por todos aquellos que comparten 
// el mismo enunciado. 
// 
// Cuando todos los alumnos que tenían la misma tarea terminaron, el 
// profesor les otorga un puntaje que representa el orden en que se terminó esa.  


// Nota: para elegir la tarea suponga que existe una función elegir que le asigna una tarea a 
// un alumno (esta función asignará 10 tareas diferentes entre 50 alumnos, es decir, que 5 
// alumnos tendrán la tarea 1, otros 5 la tarea 2 y así sucesivamente para las 10 tareas). 
int barrera = 50; // cantidad de alumnos

notas[0..9] = ([9] -1); // puntajes disponibles, seteo los 9 en -1
sem semaforoNotas[0..9] = ([9] 1); // puntajes disponibles, seteo los 9 en 1


cantTarea[0..9] = ([9] 0); // tareas entregadas para corregir, una vez que alcance 5 alumnos(cada []), se corrige
sem semaforoVTareaCant = ([9] 1); // semaforo para proteger el vector de tareas

sem protexVPuntaje = 1; // semaforo para proteger el vector de puntajes // si solo hago OP de lectura, tengon q protegerlo?

Queue <Tarea> tareas; // cola de tareas
sem protexQueuTareas = 1; // semaforo para proteger la cola de tareas
tarea = {
	int numero: 0..9; 
}

process alumno[id: 0..49] {
	Tarea tarea = elegir(); 
	P(barrera); 
	realizarTarea(tarea); 
	//entregando
	// P(protexVTareaCant); 
	// cantTarea[tarea.getNumero()]++; 
	// V(protexVTareaCant); 
	P(protexQueuTareas); // pido el semaforo para usar la Queue de tareas
	tarea.enqueue(tarea); // encolando la tarea
	V(protexQueuTareas); // libero el semaforo para usar la Queue de tareas

	P(semaforoNotas[tarea.getNumero()]); //aca va ver directo el puntaje del grupo, no el de la tarea
	mirandoNota(notas[tarea.getNumero()]);
	V(semaforoNotas[tarea.getNumero()]); // libero el semaforo para usar el vector de puntajes
}

process profesor {
	int contCorrecciones = 0; 
	while (contCorrecciones < 10) { 
		P(protexQueuTareas);
		if(Queue.size() != 0){
			Tarea tarea = tareas.dequeue(); // saco la tarea de la cola (tarea es local al proceso)
			V(protexQueuTareas); // libero el semaforo para usar la Queue de tareas
			tarea = corrigiendoTarea(tarea); // corrijo la tarea
			P(semaforoVTareaCant[tarea.getNumero()]); // pido el semaforo para usar el vector de puntajes
			cantTarea[tarea.getNumero()]++; 
			if(cantTarea[tarea.getNumero()] == 5){
				double nota =dandoNotaGrupal();

				P(semaforoNotas[tarea.getNumero()]); // pido el semaforo para usar el vector de puntajes
				notas[tarea.getNumero()] = nota; // guardo la nota
				V(semaforoNotas[tarea.getNumero()]); // libero el semaforo para usar el vector de puntajes
				
			}
			V(protexVTareaCant); 
		}
		V(protexQueuTareas);
	}
}




process profesor {
	int contCorrecciones = 0; 
	while (contCorrecciones < 10) { 
		P(protexVTareaCant); 
		for (i = 0; i < 9; i++) {
			if (cantTarea[i] == 5) { 
				contCorrecciones++; 
				cantTarea[i] = 0; 
				V(protexVTareaCant); 
				double nota = corrigiendoTarea(tarea); 
				P(protexVPuntaje); 
				notas[i] = nota; 
				V(semaforoPuntaje[i]); 
				V(protexVPuntaje); 
				P(protexVTareaCant);
			}
		}
		V(protexVTareaCant);
	}
}

