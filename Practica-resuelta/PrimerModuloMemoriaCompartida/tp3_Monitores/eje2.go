Existen N procesos que deben leer información de una base de datos, la cual es administrada 
por un motor que admite una cantidad limitada de consultas simultáneas. 
a. Analice el problema y defina qué procesos, recursos y monitores serán 
necesarios/convenientes, además de las posibles sincronizaciones requeridas para 
resolver el problema. 
b. Implemente el acceso a la base por parte de los procesos, sabiendo que el motor de base 
de datos puede atender a lo sumo 5 consultas de lectura simultáneas.

A/B 

//===================================BIEN====================================================
//logica
//me despertaron, me dijeron que tenia lugar, pero me gano otro proceso el lugar, me tengo que volver a dormir (While)
Monitor baseDeDatos{
	int numLectores = 5; // maximo de lectores simultaneos
  cond v; // variable de condicion para los lectores
	procedure leer(){
		while(numLectores < 1){
			wait(v); // espero a que haya lugar
		}
		numLectores--;
	}

	procedure salir(){
		if(numLectores < 6){
			numLectores++;
			signal(v); // despierto a un lector
		}
}


process lector [id: 0..N-1]{
	baseDeDatos.leer(); // el lector entra a la base de datos
	//leyendo
	baseDeDatos.salir(); // el lector sale de la base de datos
}


//==========================================================================================



//MALLLLLLLLLLLLLLLLLLLLLLLLL
//esta mal, en el if, si estoy con capacidad maxima y se libera 1 slot y quieren entrar 2, 1 va poder entrar pero el otro lo tengo q volver a dormir
Monitor baseDeDatos{
	int numLectores = 5; // maximo de lectores simultaneos
  cond v; // variable de condicion para los lectores
	procedure leer(){
		if(numLectores > 0){
			numLectores--;
		}else{
			wait(v);
		}
	}

	procedure salir(){
		if(numLectores < 5){
			numLectores++;
			signal(v); // despierto a un lector
		}
}


process lector [id: 0..N-1]{
	baseDeDatos.leer(); // el lector entra a la base de datos
	//leyendo
	baseDeDatos.salir(); // el lector sale de la base de datos
}