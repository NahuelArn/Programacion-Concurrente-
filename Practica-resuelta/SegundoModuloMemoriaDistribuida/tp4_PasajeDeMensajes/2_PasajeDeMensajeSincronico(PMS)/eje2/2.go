2.  En un laboratorio de genética veterinaria hay 3 empleados. El primero de ellos 
continuamente prepara las muestras de ADN; cada vez que termina, se la envía al segundo 
empleado  y  vuelve  a  su  trabajo.  El  segundo  empleado  toma  cada  muestra  de  ADN 
preparada,  arma  el  set  de  análisis  que  se  deben  realizar  con  ella  y  espera  el  resultado  para 
archivarlo.  Por  último,  el  tercer  empleado  se  encarga  de  realizar  el  análisis  y  devolverle  el 
resultado al segundo empleado. 



Process Empleado1(){
	while(true){
		Text muestra = preparandoMuestra();
		Empleado2!muestraPreparada(muestra);
	}
}

Process Empleado2(){
	Text muestra;
	Text resultadoSetAnalisis;
	Text resultadoAnalisis;
	Queue archivo;
	while(true){
		Administrador!solicitarMuestra();
		Administrador?recibirMuestra(muestra);
		resultadoSetAnalisis = realizarSetAnalisis(muestra);
		Empleado3!teMandoAnalisis(resultadoSetAnalisis);
		Empleado3?tenesMiAnalisis(resultadoAnalisis)
		archivo.push(resultadoAnalisis);
	}
}

Process Empleado3(){	
	Text muestra, analisisResultado;
	while(true){
		Empleado2?teMandoAnalisis(muestra);
		analisisResultado = realizandoAnalisis(muestra);
		Empleado2!tenesMiAnalisis(analisisResultado);
	}
}

Process Administrador(){
	Queue muestras;
	Muestra muestraActual;
	do
		[] Empleado1?muestraPreparada() -> muestras.push(muestraActual);
		[] (!muestras.isEmpty());Empleado2?solicitarMuestra() -> Empleado2!recibirMuestra(muestras.pop());
	od
}











//==============================================================================

//creo 3 procesos empleado1,2,3
//no tengo forma de que dinamicamente entienda un solo proceso todos los mensajes (creo)


//no va funcionar, replantearlo con 4 procesos.. 1 admin

Process Empleado1(){
	while(true){
		Text muestra = preparandoMuestra();
		Empleado2!muestraPreparada(muestra);
	}
}

Process Empleado2(){
	Queue muestras;
	Text muestra;
	Text resultadoAnalisis;
	do
		[] Empleado1?muestraPreparada(muestra) -> 
			muestras.push();
		[] (!muestras.isEmpty()); Empleado3?sePuedeHacerAnalisis() ->
			Empleado3!teMandoAnalisis(muestras.pop());
		[] Empleado3?tenesMiAnalisis(resultadoAnalisis) -> //Aca ya se rompe algo
			mirandoElResultado(resultadoAnalisis);
	od
}

Process Empleado3(){	
	Text muestra, analisisResultado;
	while(true){
		Empleado2!sePuedeHacerAnalisis();
		Empleado2?teMandoAnalisis(muestra);
		analisisResultado = realizandoAnalisis(muestra);
		Empleado2!tenesMiAnalisis(analisisResultado);
	}
}