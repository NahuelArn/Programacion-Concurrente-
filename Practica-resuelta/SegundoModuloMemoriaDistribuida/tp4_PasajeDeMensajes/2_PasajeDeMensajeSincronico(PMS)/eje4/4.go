4.  En  una  exposición  aeronáutica  hay  un  simulador  de  vuelo  (que  debe  ser  usado  con 
	exclusión  mutua)  y  un  empleado  encargado  de  administrar  su  uso.  Hay  P  personas  que 
	esperan a que el empleado lo deje acceder al simulador, lo usa por un rato y se retira. 
	a. Implemente  una  solución  donde  el  empleado  sólo  se  ocupa  de  garantizar  que  el 
	simulador es usado por una persona a la vez. 
	b. Modifique  la  solución  anterior  para  que  el  empleado  además  considere  el  orden  de 
	llegada para dar acceso al simulador. 
	Nota: cada persona usa sólo una vez el simulador.


//A

process Empleado{
	for(i=1; i<=P;i++){
			Persona[*]?SolicitarUso(idPersona);
			Persona[idPersona]!darAcceso();
			Persona[idPersona]?fin();
	}
}

procces Persona[id: 0..P-1]{
	Empleado!solicitarUso(id);
	Empleado?darAcceso();
	usandoSimuladorConExclusionMutua();
	Empleado!fin();
}

//B

Process Empleado(){
	int idP;
	for(int i=0; i++; i<P){
		Administrador!estoyLibre();
		Administrador?aQuienAtiendo(idP);
		Persona[idP]!darAcceso();
		Persona[idP]?termine();
	}
}

Process Persona(id: 0..P-1){
	Administrador!solicitarUso(id);
	Empleado?darAcceso();
	//usando simulador();
	Empleado!termine();
}

Process Administrador(){
	Queue personas; int idP;
	int cantP=0;
	do
		[] (cantP<P); Persona[*]?solicitarUso(idP) ->
			cantP++;
			personas.push(idP);
		[] (!personas.isEmpty()); Empleado?estoyLibre() ->
			Empleado!aQuienAtiendo(personas.pop());
	od
}