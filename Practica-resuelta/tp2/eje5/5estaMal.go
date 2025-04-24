// 5.  Suponga que se tiene un curso con 50 alumnos. Cada alumno debe realizar una tarea y 
// existen  10  enunciados  posibles.  Una  vez  que  todos  los  alumnos  eligieron  su  tarea, 
// comienzan a realizarla. Cada vez que un alumno termina su tarea, le avisa al profesor y se 
// queda esperando el puntaje del grupo, el cual está dado por todos aquellos que comparten 
// el mismo enunciado. Cuando todos los alumnos que tenían la misma tarea terminaron, el 
// profesor les otorga un puntaje que representa el orden en que se terminó esa.  


// Nota: para elegir la tarea suponga que existe una función elegir que le asigna una tarea a 
// un alumno (esta función asignará 10 tareas diferentes entre 50 alumnos, es decir, que 5 
// alumnos tendrán la tarea 1, otros 5 la tarea 2 y así sucesivamente para las 10 tareas). 



// MIRARLO BIEN, CREO QUE LO ENCARE MAL, NO ESTOY CONTROLANDO EL VER LA TAREA, SOLO MANEJO LASC CANTIDAS, PERO NO REFLEJO EL CORREGIR CADA TAREA DE CADA ALUMNO

int barrera = 50; // cantidad de alumnos
cantTarea[0..9] = ([9] 0); // tareas entregadas para corregir, una vez que alcance 5 alumnos(cada []), se corrige
sem puntaje[0..9] = ([9] 0); // puntajes disponibles, seteo los 9 en -1
notas[0..9] = ([9] -1); // puntajes disponibles, seteo los 9 en -1
sem protexVTarea = 1; // semaforo para proteger el vector de tareas
sem protexVPuntaje = 1; // semaforo para proteger el vector de puntajes // si solo hago OP de lectura, tengon q protegerlo?

tarea = {
	int numero: 0..9; // numero de la tarea
}
process alumno[id: 0..49] {
	//asumo que elegir(), resuelve bien la entrega de tareas/stocks de tareas, y no se repiten mas de 5 veces
	Tarea tarea = elegir(); // SNC nose como se maneja la fn, entonces asumo q no es una SC
	P(barrera); // espero a que todos los alumnos elijan su tarea         BARRERAAAAAA
	realizarTarea(tarea); // realizo la tarea
	P(protexVTarea); // pido el semaforo para usar el vector de tareas
	cantTarea[tarea.getNumero()]++; // aumento la cantidad de tareas entregadas
	V(protexVTarea); // libero el semaforo para usar el vector de tareas
// espero el puntaje de la tarea
	P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes (creo q no es necesario, PREGUNTAR)
	P(puntaje[tarea.getNumero()]);// espero a que el profesor me de el puntaje
	mirandoPuntaje(notas[tarea.getNumero()]); // miro el puntaje
}


process profesor {
	int contCorrecciones = 0; // contador de correcciones
	while (contCorrecciones < 10) { // mientras no haya corregido todas las tareas
		P(protexVTarea); // pido el semaforo para usar el vector de tareas
		for (i = 0; i < 9; i++) { // recorro el vector de tareas
			if (cantTarea[i] == 5) { // si la tarea no esta completa, sigo
				contCorrecciones++; // aumento el contador de correcciones
				// Tarea tarea = tarea[i]; // guardo la tarea
				cantTarea[i] = 0; // reseteo la tarea para que no se vuelva a corregir
				V(protexVTarea); // libero el semaforo para usar el vector de tareas
				double nota = corrigiendoTarea(tarea); // corrijo la tarea
				P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes
				notas[i] = nota; // guardo la nota
				V(puntaje[i]); // libero el semaforo para usar el vector de puntajes
				V(protexVPuntaje); // libero el semaforo para usar el vector de puntajes
				P(protexVTarea);
			}
		}
		V(protexVTarea); // libero el semaforo para usar el vector de tareas
	}
}





























int barrera = 50; // cantidad de alumnos
tarea[0..9] = ([9] 0); // tareas entregadas para corregir, una vez que alcance 5 alumnos(cada []), se corrige
sem puntaje[0..9] = ([9] 0); // puntajes disponibles, seteo los 9 en -1
sem protexVTarea = 1; // semaforo para proteger el vector de tareas
sem protexVPuntaje = 1; // semaforo para proteger el vector de puntajes // si solo hago OP de lectura, tengon q protegerlo?

tarea = {
	int numero: 0..9; // numero de la tarea
}
process alumno[id: 0..49] {
	//asumo que elegir(), resuelve bien la entrega de tareas/stocks de tareas, y no se repiten mas de 5 veces
	Tarea tarea = elegir(); // SNC nose como se maneja la fn, entonces asumo q no es una SC
	P(barrera); // espero a que todos los alumnos elijan su tarea         BARRERAAAAAA
	realizarTarea(tarea); // realizo la tarea
	P(protexVTarea); // pido el semaforo para usar el vector de tareas
	tarea[tarea.getNumero()]++; // aumento la cantidad de tareas entregadas
	V(protexVTarea); // libero el semaforo para usar el vector de tareas
	// P(puntaje[tarea.getNumero()]); 
// espero el puntaje de la tarea
	P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes (creo q no es necesario, PREGUNTAR)
	P(puntaje[tarea.getNumero()]);// espero a que el profesor me de el puntaje
	mirandoPuntaje(puntaje[tarea.getNumero()]); // miro el puntaje
}

process profesor {
	int contCorrecciones = 0; // contador de correcciones
	while (contCorrecciones < 10) { // mientras no haya corregido todas las tareas
		P(protexVTarea); // pido el semaforo para usar el vector de tareas
		for (i = 0; to 9; st tarea[i] == 5) { // recorro el vector de tareas
			contCorrecciones++; // aumento el contador de correcciones
			Tarea tarea = tarea[i]; // guardo la tarea
			tarea[i] = 0; // reseteo la tarea para que no se vuelva a corregir
			V(protexVTarea); // libero el semaforo para usar el vector de tareas
			double puntaje = corrigiendoTarea(tarea); // corrijo la tarea
			P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes
			puntaje[i] = puntaje; // guardo el puntaje
			V(protexVPuntaje); // libero el semaforo para usar el vector de puntajes
			P(protexVTarea);
		}
		V(protexVTarea); // libero el semaforo para usar el vector de tareas
	}
}

















int barrera = 50; // cantidad de alumnos
tarea[0..9] = ([9] 0); // tareas entregadas para corregir, una vez que alcance 5 alumnos(cada []), se corrige
sem puntaje[0..9] = ([9] -1); // puntajes disponibles, seteo los 9 en -1
sem protexVTarea = 1; // semaforo para proteger el vector de tareas
sem protexVPuntaje = 1; // semaforo para proteger el vector de puntajes // si solo hago OP de lectura, tengon q protegerlo?

process alumno[id: 0..49] {
	//asumo que elegir(), resuelve bien la entrega de tareas/stocks de tareas, y no se repiten mas de 5 veces
	Tarea tarea = elegir(); // SNC nose como se maneja la fn, entonces asumo q no es una SC
	P(barrera); // espero a que todos los alumnos elijan su tarea         BARRERAAAAAA
	realizarTarea(tarea); // realizo la tarea
	P(protexVTarea); // pido el semaforo para usar el vector de tareas
	tarea[tarea.getNumero()]++; // aumento la cantidad de tareas entregadas
	V(protexVTarea); // libero el semaforo para usar el vector de tareas
	// P(puntaje[tarea.getNumero()]); 
// espero el puntaje de la tarea
	P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes (creo q no es necesario, PREGUNTAR)
	while (puntaje[tarea.getNumero()] == -1) { skip}// espero a que el profesor me de el puntaje
	mirandoPuntaje(puntaje[tarea.getNumero()]); // miro el puntaje
}

process profesor {
	int contCorrecciones = 0; // contador de correcciones
	while (contCorrecciones < 10) { // mientras no haya corregido todas las tareas
		P(protexVTarea); // pido el semaforo para usar el vector de tareas
		for (i = 0; to 9; st tarea[i] == 5) { // recorro el vector de tareas
			contCorrecciones++; // aumento el contador de correcciones
			Tarea tarea = tarea[i]; // guardo la tarea
			tarea[i] = 0; // reseteo la tarea para que no se vuelva a corregir
			V(protexVTarea); // libero el semaforo para usar el vector de tareas
			double puntaje = corrigiendoTarea(tarea); // corrijo la tarea
			P(protexVPuntaje); // pido el semaforo para usar el vector de puntajes
			puntaje[i] = puntaje; // guardo el puntaje
			V(protexVPuntaje); // libero el semaforo para usar el vector de puntajes
			P(protexVTarea);
		}
		V(protexVTarea); // libero el semaforo para usar el vector de tareas
	}
}