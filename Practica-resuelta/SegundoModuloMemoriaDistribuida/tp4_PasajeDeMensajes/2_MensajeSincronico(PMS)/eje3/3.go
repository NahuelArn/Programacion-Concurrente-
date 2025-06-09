3.  En  un  examen  final  hay  N  alumnos  y  P  profesores.  Cada  alumno  resuelve  su  examen,  lo 
entrega  y  espera  a  que  alguno  de  los  profesores  lo  corrija  y  le  indique  la  nota.  Los 
profesores corrigen los exámenes respetando el orden en que los alumnos van entregando.  
a. Implemente una solución con PMS considerando que P=1. 
b. Implemente una solución con PMS considerando que P>1. 
c. Modifique (b) considerando que los alumnos no comienzan a realizar su examen hasta 
que todos hayan llegado al aula. 
Nota: maximizar la concurrencia y no generar demora innecesaria.

//A
Process Alumno(id: 0..N-1){
	Text examen, examenCorregido;
	Profesor!termineExamen(examen, id);
	Profesor?esperandoCorreccion(examenCorregido);
}

Process Profesor(){
	Text examenEnEsperaDeNota, examenCorregido; int idAlumnoActual;
	for(int i=0; i < N; i++){
		Administrador!solicitarExamen();
		Administrador?examenActual(examenEnEsperaDeNota, idAlumnoActual);
		examenCorregido = corrigiendoExamen(examenEnEsperaDeNota);
		Alumno[idAlumnoActual]!esperandoCorreccion(examenCorregido);
	}
}

Process Administrador(){
	Queue examenesEsperandoCorreccion;
	int cantCorregidos = 0; Text examen;
	int idAlumnoActual;
	do
		[] (cantCorregidos < N); Alumno[*]?termineExamen(examen, idAlumnoActual)-> //quiero escuchar, a cualquier alumno *
			examenesEsperandoCorreccion.push(examen, idAlumnoActual);
		[] (!examenesEsperandoCorreccion.isEmpty()); Profesor?solicitarExamen() ->
			//al hacer pop es como "si tuviera un registro/enum" examen, idAlumnoActual... y el receptor, capta directamente esos 2 parametros
			Profesor!examenActual(examenesEsperandoCorreccion.pop()); 
	od
	esperandoCorreccion();
}

// Queue<Examen> examenesEsperandoCorreccion;
// type Examen{
// 	id int
// 	examen Text
// }

//================================================================================================================


//B

Process Alumno(id: 0..N-1){
	//haciendoExamen();
	Examen examen = haciendoExamen();
	Administrador!corregime(examen, id);
	Text examenCorregido;
	Profesor[*]?meCorrigio(examenCorregido); //de que profesor? de cualquiera 
	//mirando Correciones
}

//en el mejor caso N=P por cada alumno hay 1 profesor.. pero como P es desconocido
Process Profesor(id: 0..P-1){
	Text examen, examenCorregido;
		Administrador!estoyLibre(id);
		Administrador?reciboElExamen(examen, idA);
		while(idA != -1){ //si recibo -1.. significa que todos los examenes ya fueron corregidos/entregados
			examenCorregido = corrigiendoExamen(examen);
			Alumno[idA]!meCorrigio(examenCorregido);
			Administrador!estoyLibre(id);
			Administrador?reciboElExamen(examen, idA);
		}
	
}

Process Administrador(){
	Queue eec; //examenes esperando correccion
	Text examen; int idA, idP;
	int cantExamenes = 0;
	//corta cuando tenga todos los examenes de los alumnos y derive todos los examenes a los profesores correspondientes
	do
		[] (cantExamenes < N); Alumno[*]?corregime(examen, idA) -> 
			eec.push(examen,idA);
		[] (!eec.isEmpty()); Profesor[*]?estoyLibre(idP) ->
			Profesor[idP]!reciboElExamen(eec.pop());
	od

	//si estoy aca, sali del "do od" exp arriba cuando voy a salir
	//Todos los profesores P van a estar activos... entonces tengo que avisar a los P profesores que termino
	for(int i=0; i<P; i++){
		Profesor[*]?estoyLibre(idP);
		//examen va tener "basura"... va ser el mismo examen del Alumno que atendi anteriormente
		Profesor[idP]!reciboElExamen(examen,-1);
	}
}



//==================================================================================================================
//C

//la famosa barrera
//misma logica de semaforos, adaptada a canales
//aca no tengo variables compartidas.. creo un proceso(o pongo en el Administrador los fors) que capture N mensajes de alumnos y despues libere a los N alumnos...
Process Barrera(){
	for (int i = 0; i< N; i++){ // "duerme a todos" cualquier alumno que me haya mandado un mensaje lo duermo
		Alumno[*]?formandoBarrera();
	}

	for(int i = 0; i< N; i++){
		Alumno[i]!liberaronAtodos(); //voy liberando a los alumnos 
		// Alumno[*]!liberaronAtodos(); el * es solo para recepcion
	}
}

Process Alumno(id: 0..N-1){
	//llego al examen
	//espero que todos lleguen
	//------------------------------
	Barrera!formandoBarrera();
	Barrera?liberaronAtodos();
	//-----------------
	//haciendoExamen();
	Examen examen = haciendoExamen();
	Administrador!corregime(examen, id);
	Text examenCorregido;
	Profesor[*]?meCorrigio(examenCorregido); //de que profesor? de cualquiera 
	//mirando Correciones
}

//en el mejor caso N=P por cada alumno hay 1 profesor.. pero como P es desconocido
Process Profesor(id: 0..P-1){
	Text examen, examenCorregido;
		Administrador!estoyLibre(id);
		Administrador?reciboElExamen(examen, idA);
		while(idA != -1){ //si recibo -1.. significa que todos los examenes ya fueron corregidos/entregados
			examenCorregido = corrigiendoExamen(examen);
			Alumno[idA]!meCorrigio(examenCorregido);
			Administrador!estoyLibre(id);
			Administrador?reciboElExamen(examen, idA);
		}
	
}

Process Administrador(){
	Queue eec; //examenes esperando correccion
	Text examen; int idA, idP;
	int cantExamenes = 0;
	//corta cuando tenga todos los examenes de los alumnos y derive todos los examenes a los profesores correspondientes
	do
		[] (cantExamenes < N); Alumno[*]?corregime(examen, idA) -> 
			cantExamenes++;
			eec.push(examen,idA);
		[] (!eec.isEmpty()); Profesor[*]?estoyLibre(idP) ->
			Profesor[idP]!reciboElExamen(eec.pop());
	od

	//si estoy aca, sali del "do od" exp arriba cuando voy a salir
	//Todos los profesores P van a estar activos... entonces tengo que avisar a los P profesores que termino
	for(int i=0; i<P; i++){
		Profesor[*]?estoyLibre(idP);
		//examen va tener "basura"... va ser el mismo examen del Alumno que atendi anteriormente
		Profesor[idP]!reciboElExamen(examen,-1);
	}
}



















////=========
//no me sirve, tengo q avisarle dinamicamente cuando cortar
//B

Process Alumno(id: 0..N-1){
	//haciendoExamen();
	Examen examen = haciendoExamen();
	Administrador!corregime(examen, id);
	Text examenCorregido;
	Profesor[*]?meCorrigio(examenCorregido); //de que profesor? de cualquiera 

}

//en el mejor caso N=P por cada alumno hay 1 profesor.. pero como P es desconocido
Process Profesor(id: 0..P-1){
	Text examen, examenCorregido;
	for (int i = 0; i< N; i++){
		Administrador!estoyLibre(id);
		Administrador?reciboElExamen(examen, idA);
		
		examenCorregido = corrigiendoExamen(examen);

		Alumno[idA]!meCorrigio(examenCorregido);
	}
}

Process Administrador(){
	Queue eec; //examenes esperando correccion
	Text examen; int idA, idP;
	int cantExamenes = 0;
	//corta cuando tenga todos los examenes de los alumnos y derive todos los examenes a los profesores correspondientes
	do
		[] (cantExamenes < N); Alumno[*]?corregime(examen, idA) -> 
			eec.push(examen,idA);
		[] (!eec.isEmpty()); Profesor[*]?estoyLibre(idP) ->
			Profesor[idP]!reciboElExamen(eec.pop());
	od
}




















//problema popeo
//B

Process Alumno(id: 0..N-1){
	//haciendoExamen();
	Examen examen = haciendoExamen();
	Administrador!corregime(examen, id);
	Text examenCorregido;
	Administrador[id]?meCorrigio();

}

//en el mejor caso N=P por cada alumno hay 1 profesor.. pero como P es desconocido
Process Profesor(id: 0..P-1){
	Text examen, examenCorregido;
	Administrador!estoyLibre(id);
	Administrador![id]mandameElExamen();
	Administrador?[id]reciboElExamen(examen, idA);
}

Process Administrador(){
	Queue eec; //examenes esperando correccion
	Queue pl; //profesores libres;
	Text examen; int idC, idP;

	do
		[] Alumno[*]?corregime(examen, idC) -> 
			eec.push(examen,idC);
		[] Profesor[*]?estoyLibre(idP) ->
			pl.push(idP);
		[] (!pl.isEmpty() && !eec.isEmpty()); Profesor[pl.pop()]?.mandameElExamen() ->
			Alumno

	od
}










//aaaaaaaaaaaa
Process Alumno(id: 0..N-1){
	//haciendoExamen();
	Examen examen = haciendoExamen();
	Administrador!examenEsperandoCorreccion(examen, id);


}

//en el mejor caso N=P por cada alumno hay 1 profesor.. pero como P es desconocido
Process Profesor(id: 0..P-1){

}

Process Administrador(){
	Queue eec; //examenes esperando correccion
	Text examen; int idC;
	do
		[] Alumno[*]?examenEsperandoCorreccion(examen, idC) -> 
			eec.push(examen,idC);
	od
}