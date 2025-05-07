 Existe  una  comisión  de  50  alumnos  que  deben  realizar  tareas  de  a  pares,  las  cuales  son 
corregidas por un JTP. Cuando los alumnos llegan, forman una fila. Una vez que están todos 
en fila, el JTP les asigna un número de grupo a cada uno. Para ello, suponga que existe una 
función AsignarNroGrupo() que retorna un número “aleatorio” del 1 al 25. Cuando un alumno 
ha recibido su número de grupo, comienza a realizar su tarea. Al terminarla, el alumno le avisa 
al JTP y espera por su nota. Cuando los dos alumnos del grupo completaron la tarea, el JTP les asigna un puntaje (el primer grupo en terminar tendrá como nota 25, el segundo 24, y así 
sucesivamente hasta el último que tendrá nota 1). Nota: el JTP no guarda el número de grupo 
que le asigna a cada alumno.


//en este ejercicio no se maximisa la concurrencia haciendolo con 2 monitores
// ya que se espera primero a los 50 alumnos, y esos 50 no vuelven a entrar.. 

// "la comunicacion entre Monitores se realiza con parametros in/out"

//se hace con 2 para mejor legibilidad 
Monitor Esperando{
	// int cantAlumnosEsperando = 0; //lo puedo saber con dimF qyeye
	Queue alumnosEsperando;

	cond alumno;
	cond jtp;

	cond esperandoNroGrupo;
	Queue alumnos;

	cond avisoDeInicio[50];
	int nroDeGrupoAlumno[50];

	procedure haciendoFila(id: in int){
		alumnosEsperando.push(id);
		if(alumnosEsperando.size() == 50){
			signal(jtp);
		}
	}
	procedure esperandoNroGrupo(idI: in int, tarea : out int){
		wait(avisoDeInicio[idI]);
		tarea = nroDeGrupoAlumno[idI];
	}

	procedure esperandoAlumnos{
		if(alumnosEsperando.size() != 50){
			wait(jtp);
		}
	}

	procedure asignarNroGrupo(){
		int id = alumnosEsperando.pop();
		int nroDeGrupo = AsignarNroGrupo();
		nroDeGrupoAlumno[id] = nroDeGrupo;
		signal(avisoDeInicio[id]); //loko ya tenes tu nro de grupo,
		
		//tengo que esperar que algun alumno termine, no va esto pertenece a las primeras 50 iteraciones de asignar y esperando
	}
}

Monitor ManejoNotas{
	int cantTareasXgrupo = ([25], 0)
	Queue tareasEntregadas;
	int puntajePorGrupo[25];

	cond esperandoNotaGrupo[25];	//cada pos va poder tenr 2 procesos dormidos
	cond avisoProfesor;

	procedure entregarTarea(nroGrupo : in int){
		tareasEntregadas.push(nroGrupo); //aviso como integrante del grupo que termine la tarea
		signal(avisoProfesor);
		wait(esperandoNotaGrupo[nroGrupo]);
		int puntaje = puntajePorGrupo[nroGrupo]; //cada proceso lo va chequear 1 vz "viendo nota"
	}

	procedure esperandoEntregas{
		if(tareasEntregadas.isEmpty()){
			wait avisoProfesor;
		}

	}
	procedure corrigiendo(puntaje : in int){
		int nroGrupo = tareasEntregadas.pop();
		cantTareasXgrupo[nroGrupo] ++;
		if(cantTareasXgrupo[nroGrupo] == 2){
			puntajePorGrupo[nroGrupo] = puntaje;
			singal_All(esperandoNotaGrupo[nroGrupo]);
		}
	}
}


Process alumno[id: 1..50]{
	int tarea;
	Esperando.haciendoFila(id);
	Esperando.esperandoNroGrupo(id, tarea);
	//entregar nota, tengo que mandar mi id y mi tarea""
	//# haciendo el tp
	ManejoNotas.entregarTarea(tarea);
}

Process jtp{
	int puntaje = 25;

	for (int i = 0; i < 50; i++){ //esto corresponde al comportamiento del monitor Esperando
		Esperando.esperandoAlumnos();
		Esperando.asignarNroGrupo();
	}

	for(int i = 0; i < 50; i++){
		ManejoNotas.esperandoEntregas();
		ManejoNotas.corrigiendo(puntaje);
		puntaje--; //si el decremento se realiza dentro del proceso no impacta al main, se pasa puntaje como IN
	}

}