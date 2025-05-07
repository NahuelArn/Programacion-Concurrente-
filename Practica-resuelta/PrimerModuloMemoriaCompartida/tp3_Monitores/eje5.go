 Existe  una  comisión  de  50  alumnos  que  deben  realizar  tareas  de  a  pares,  las  cuales  son 
corregidas por un JTP. Cuando los alumnos llegan, forman una fila. Una vez que están todos 
en fila, el JTP les asigna un número de grupo a cada uno. Para ello, suponga que existe una 
función AsignarNroGrupo() que retorna un número “aleatorio” del 1 al 25. Cuando un alumno 
ha recibido su número de grupo, comienza a realizar su tarea. Al terminarla, el alumno le avisa 
al JTP y espera por su nota. Cuando los dos alumnos del grupo completaron la tarea, el JTP les asigna un puntaje (el primer grupo en terminar tendrá como nota 25, el segundo 24, y así 
sucesivamente hasta el último que tendrá nota 1). Nota: el JTP no guarda el número de grupo 
que le asigna a cada alumno.

Monitor Esperando{
	int cantAlumnosEsperando = 0;
	cond alumno;
	cond jtp;

	cond esperandoNroGrupo;
	Queue alumnos;

	cond avisoDeInicio[50];

	procedure haciendoFila(id: out int){
		if(cantAlumnosEsperando < 51){
			cantAlumnosEsperando ++;
			
		}else{
			signal(jtp);
		}

	}
	procedure esperandoNroGrupo(idI: in int){
		wait(avisoInicio[idI]);
	}

	procedure esperandoAlumnos{
		if(cantAlumnos != 50){
			wait(jtp);
		}
	}
}

Process alumno[id: 1..50]{
	Esperando.haciendoFila(id);
	Esperando.esperandoNroGrupo(id);
}

Process jtp{
	int puntaje = 25;
	for (int i = 0; i < 50; i++){
		Esperando.esperandoAlumnos();
		Esperando.asignarNroGrupo();
	}
}