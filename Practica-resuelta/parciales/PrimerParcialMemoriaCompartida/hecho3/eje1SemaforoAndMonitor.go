//Semaforo

Semaforo
1) Resolver los problemas siguientes: 
a) En una estación de trenes, asisten P personas que deben realizar una carga de su tarjeta SUBE 
en la terminal disponible. La terminal es utilizada en forma exclusiva por cada persona de acuerdo 
con el orden de llegada. Implemente una solución utilizando únicamente procesos Persona. Nota: 
la función UsarTerminal() le permite cargar la SUBE en la terminal disponible.  
b) Resuelva el mismo problema anterior pero ahora considerando que hay T terminales disponibles. 
Las personas realizan  una única fila y la  carga la realizan  en  la primera terminal que se libera. 
Recuerde que sólo debe emplear procesos Persona. Nota: la función UsarTerminal(t) le permite 
cargar la SUBE en la terminal t


A)

boolean libre = true;
Queue espera;
sem persona[P] = ([P] 0);
sem mutex = 1;
subes[P];
Process Persona[id : 0..P-1]{

	P(mutex);
	if(libre){
		libre = false;
		V(mutex);
	}else{
		espera.push(id);
		V(mutex);
		P(persona[id]);
	}
	//aca solo llegan de a 1
	subes[id] = UsarTerminal();
	P(mutex);
	if(espera.size()>0){
		int sigId = espera.pop();
		V([persona[sigId]);
	}else{
		libre = true;
	}
	V(mutex);
}


b) Resuelva el mismo problema anterior pero ahora considerando que hay T terminales disponibles. 
Las personas realizan  una única fila y la  carga la realizan  en  la primera terminal que se libera. 
Recuerde que sólo debe emplear procesos Persona. Nota: la función UsarTerminal(t) le permite 
cargar la SUBE en la terminal t

Queue espera;
sem persona[P] = ([P] 0);
sem mutex = 1;
subes[P];

Queue terminal;
terminalPaLaPersona[P] = ([P] null);

Process Persona[id : 0..P-1]{

	P(mutex);
	if(terminal.size()>0){
		terminalPaLaPersona[id] = terminal.pop();
		V(mutex);
	}else{
		espera.push(id);
		V(mutex);
		P(persona[id]);
	}
	//aca solo llegan de a 1
	subes[id] = terminalPaLaPersona[id].UsarTerminal();

	P(mutex);
	if(espera.size()>0){
		int sigId = espera.pop();
		terminalPaLaPersona[sigId] = terminalPaLaPersona[id]; //le doy la terminal
		V([persona[sigId]);
	}else{
		terminal.push(terminalPaLaPersona[id]); //devuelvo la terminal... 
	}
	V(mutex);
}

//==============================================================================================================================

//Monitores

Monitores 
1) Resolver  el  siguiente  problema.  En  una  elección  estudiantil,  se  utiliza  una  máquina  para  voto 
electrónico. Existen N Personas que votan y una Autoridad de Mesa que les da acceso a la máquina 
de acuerdo con el orden de llegada, aunque ancianos y embarazadas tienen prioridad sobre el resto. 
La máquina de voto sólo puede ser usada por una persona a la vez. Nota: la función Votar() permite 
usar la máquina.


Monitor Votacion{
	Quue fila;
	cond persona, coordinador;
	cond aDormir[N];
	Maquina maquina;
	boolean libre = true;

	procedure votacionLlegada(id : in int, machine : out Maquina){
		fila.pushConPrioridadEmbarazadasAndAncianas(id); //si es una anciana o embaraza, las pone adelante de la fila
		wait(aDormir[id]);
		machine = maquina;
		// libre = false;
	}

	procedure liberarMaquina(maquinaI in: Maquina){
		signal(coordinador);
		maquina = maquinaI;
		libre = true;
	}

	procedure darAcceso(){
		if(fila.size() == 0 || !libre){
			wait (coordinador);
		}
		int sigId = fila.pop();
		signal(aDormir[sigId]);
		libre = false;
	}
	
}

Process Persona[id: 0..N]{
	Maquina machine;
	Votacion.votacionLlegada(id,machine);
	maquina.Votar();
	Votacion.liberarMaquina(machine);
}

Process coordinador{
	for (int u = 0; u < N; u++){
		Votacion.darAcceso();
	}
}