2) Resolver el siguiente problema con PMS. En la estación de trenes hay una terminal de SUBE que 
debe ser usada por P personas de acuerdo con el orden de llegada. Cuando la persona accede a la 
terminal,  la  usa  y  luego  se  retira  para  dejar  al  siguiente.  Nota:  cada  Persona  usa  sólo  una  vez  la 
terminal. 

Process Persona(id: 1..P){
	Coordinador!quieroUsarLaTerminal(id);
	Coordinador?usarTerminal(estacion);
	Coordinador!termineDeUsarTerminal(estacion);
}

Process Coordinador{
	boolean terminalLibre = true;
	Queue fila;
	integer cantPersonas =0 ;
	integer idP;
	Terminal terminal; Integer cantTerminaron = 0;
	do
		[] (cantPersonas < P); Persona[*]?quieroUsarLaTerminal(idP) ->
			if(terminalLibre == true){
				terminalLibre = false;
				Persona[idP]!usarTerminal(terminal);
			}else{
				fila.push(idP);
			}
			cantPersonas++;
		[] (cantTerminaron < P); Persona[*]? termineDeUsarTerminal(estacion) ->
			if(fila.isEmpty){
				terminalLibre = true;
			}else{
				Persona[fila.pop()]!usarTerminal(estacion);
			}
			cantTerminaron++
	od

}