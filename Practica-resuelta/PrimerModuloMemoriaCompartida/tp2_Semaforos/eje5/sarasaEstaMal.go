// MIRARLO BIEN, CREO QUE LO ENCARE MAL, NO ESTOY CONTROLANDO EL VER LA TAREA, SOLO MANEJO LASC CANTIDAS, PERO NO REFLEJO EL CORREGIR CADA TAREA DE CADA ALUMNO

int barrera = 50; // cantidad de alumnos
cantTarea[0..9] = ([9] 0); // tareas entregadas para corregir, una vez que alcance 5 alumnos(cada []), se corrige
sem puntaje[0..9] = ([9] 0); // puntajes disponibles, seteo los 9 en -1
notas[0..9] = ([9] -1); // puntajes disponibles, seteo los 9 en -1
sem protexVTarea = 1; // semaforo para proteger el vector de tareas
sem protexVPuntaje = 1; // semaforo para proteger el vector de puntajes // si solo hago OP de lectura, tengon q protegerlo?

tarea = {
	int numero: 0..9; 
}

process alumno[id: 0..49] {
	Tarea tarea = elegir(); 
	P(barrera); 
	realizarTarea(tarea); 
	P(protexVTarea); 
	cantTarea[tarea.getNumero()]++; 
	V(protexVTarea); 

	P(protexVPuntaje); 
	P(puntaje[tarea.getNumero()]);
	mirandoPuntaje(notas[tarea.getNumero()]);
}


process profesor {
	int contCorrecciones = 0; 
	while (contCorrecciones < 10) { 
		P(protexVTarea); 
		for (i = 0; i < 9; i++) {
			if (cantTarea[i] == 5) { 
				contCorrecciones++; 
				cantTarea[i] = 0; 
				V(protexVTarea); 
				double nota = corrigiendoTarea(tarea); 
				P(protexVPuntaje); 
				notas[i] = nota; 
				V(puntaje[i]); 
				V(protexVPuntaje); 
				P(protexVTarea);
			}
		}
		V(protexVTarea);
	}
}

