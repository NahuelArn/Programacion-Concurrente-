En un entrenamiento de fútbol hay 20 jugadores que forman 4 equipos (cada jugador conoce 
el equipo al cual pertenece llamando a la función DarEquipo()). Cuando un equipo está listo 
(han llegado los 5 jugadores que lo componen), debe enfrentarse a otro equipo que también 
esté listo (los dos primeros equipos en juntarse juegan en la cancha 1, y los otros dos equipos 
juegan en la cancha 2). Una vez que el equipo conoce la cancha en la que juega, sus jugadores 
se dirigen a ella. Cuando los 10 jugadores del partido llegaron a la cancha comienza el partido, 
juegan  durante  50  minutos,  y  al  terminar  todos  los  jugadores  del  partido  se  retiran  (no  es 
necesario que se esperen para salir).


Monitor PreparandoEquipo{
	int cantJugadoresPorEquipo[4] = ([4],0);
	cond jugadoresListos[4];
	// int cantEquiposListos = 0 ;
	Queue equiposListos;

	cond hayRival;

	Queue canchas;

	int canchaQueJuega[4] = ([4],0);

	procedure avisarQueLlegueJugador(nroEquipo : out int){
		nroEquipo = DarEquipo();
		cantJugadoresPorEquipo[nroEquipo]++; //cada jugador sabe a que equipo pertenece
		if(cantJugadoresPorEquipo[nroEquipo] < 5){ //el 5to avisa que ya estan completos
			wait(jugadoresListos[nroEquipo])
		}else{	//el 5to jugador en llegar avisa que ya estan completos
			equiposListos.enqueue(nroEquipo);
			if(equiposListos.size()  > 1 ){
				signal(hayRival);
			}
			wait(jugadoresListos[nroEquipo]); //para el ultimo jugador
		}
	}

	procedure hayRival(nroCancha : in int){
		if(equiposListos.size()  < 2 ){
			wait(hayRival);
		}
		for (int i = 0; i < 2; i++){
			Partido.AsignarCancha(nroCancha);
		}
	}
	

	procedure enQueCanchaJuego(nroEquipo : in int, nroCancha : out int){
		nroCancha = canchaQueJuega[nroEquipo];
	}

	procedure asignarCancha(nroCancha : in int){
		int nroEquipo = equiposListos.dequeue();
		canchaQueJuega[nroEquipo] = nroCancha;
		signal_all(jugadoresListos[nroEquipo]);
	}
}

Monitor Partido{

	cond jugadoresEnCancha[10];
	int cantJugadoresEnCancha[2] = ([2],0);

	cond comenzarPartido;
	cond finalizarPartido;


	procedure AsignarCancha(nroCancha : int){
		PreparandoEquipo.asignarCancha(nroCancha);
	}

	procedure irHaciaCancha(nroCancha : int){
		cantJugadoresEnCancha[nroCancha]++;
		if(cantJugadoresEnCancha[nroCancha] < 10){
			wait(jugadoresEnCancha[nroCancha]);
		}else{
			signal(comenzarPartido); //el ultimo jugador indica que se puede empezar el partido
			wait(jugadoresEnCancha[nroCancha]); //duermo al ultimo jugador
		}
	}

	procedure iniciarPartido(nroCancha in int){
		if(cantJugadoresEnCancha[nroCancha] < 10){
			wait(comenzarPartido);
		}
		signal_all(jugadoresEnCancha[nroCancha]); //despierta a todos los jugadores que esperan en esa cancha
		wait(finalizarPartido);
	}

	procedure terminarPartido(nroCancha in int){
		cantJugadoresEnCancha[nroCancha] = 0;
		signal_all(finalizarPartido);
	}
}




Process jugador[id: 0..19]{
	int nroEquipo; 
	PreparandoEquipo.avisarQueLlegueJugador(nroEquipo);
	PreparandoEquipo.enQueCanchaJuego(nroEquipo,nroCancha);
	Partido.irHaciaCancha(nroCancha);
}

Process coordinador[nroCancha: 1..2]{
	PreparandoEquipo.hayRival(nroCancha);
	// for (int i = 0; i < 2; i++){
	// 	Partido.AsignarCancha(nroCancha);
	// }
	Partido.iniciarPartido(nroCancha);
	delay(90); //jugando el partido
	Partido.terminarPartido(nroCancha);
}

//con el for limito la concurrencia de los partidos... tengo 2 recursos ->" cancha1 y cancha2"
// Process coordinador{
// 	int nroCancha = 1;
// 	for(int i = 0; i < 2; i++){
// 		PreparandoEquipo.hayRival();
// 		Partido.AsignarCancha(nroCancha);
// 		Partido.iniciarPartido(nroCancha);
// 		delay(90); //jugando el partido
// 		Partido.terminarPartido(nroCancha);
// 	}
// }




















En un entrenamiento de fútbol hay 20 jugadores que forman 4 equipos (cada jugador conoce 
el equipo al cual pertenece llamando a la función DarEquipo()). Cuando un equipo está listo 
(han llegado los 5 jugadores que lo componen), debe enfrentarse a otro equipo que también 
esté listo (los dos primeros equipos en juntarse juegan en la cancha 1, y los otros dos equipos 
juegan en la cancha 2). Una vez que el equipo conoce la cancha en la que juega, sus jugadores 
se dirigen a ella. Cuando los 10 jugadores del partido llegaron a la cancha comienza el partido, 
juegan  durante  50  minutos,  y  al  terminar  todos  los  jugadores  del  partido  se  retiran  (no  es 
necesario que se esperen para salir).


Monitor PreparandoEquipo{
	int cantJugadoresPorEquipo[4] = ([4],0);
	cond jugadoresListos[4];
	// int cantEquiposListos = 0 ;
	Queue equiposListos;

	cond hayRival;

	Queue canchas;

	int canchaQueJuega[4] = ([4],0);

	procedure avisarQueLlegueJugador(nroEquipo : out int){
		int nroEquipo = DarEquipo();
		cantJugadoresPorEquipo[nroEquipo]++; //cada jugador sabe a que equipo pertenece
		if(cantJugadoresPorEquipo[nroEquipo] < 5){ //el 5to avisa que ya estan completos
			wait(jugadoresListos[nroEquipo])
		}
  	equiposListos.enqueue(nroEquipo);
		wait(jugadoresListos[nroEquipo]); //para el ultimo jugador
		if(equiposListos.size()  > 1 ){
			signal(hayRival);
		}
	}

	procedure hayRival(){
		if(equiposListos.size()  < 2 ){
			wait(hayRival);
		}
	}

	procedure enQueCanchaJuego(nroEquipo : in int, nroCancha : out int){
		nroCancha = canchaQueJuega[nroEquipo];
	}

	procedure asignarCancha(nroCancha : in int){
		int nroEquipo = equiposListos.dequeue();
		canchaQueJuega[nroEquipo] = nroCancha;
		signal_all(jugadoresListos[nroEquipo]);
	}
}

Monitor Partido{

	cond jugadoresEnCancha[10];
	int cantJugadoresEnCancha[2] = ([2],0);

	cond comenzarPartido;
	cond finalizarPartido;


	procedure AsignarCancha(nroCancha : int){
		PreparandoEquipo.asignarCancha(nroCancha);
	}

	procedure irHaciaCancha(nroCancha : int){
		cantJugadoresEnCancha[nroCancha]++;
		if(cantJugadoresEnCancha[nroCancha] < 10){
			wait(jugadoresEnCancha[nroCancha]);
		}else{
			signal(comenzarPartido); //el ultimo jugador indica que se puede empezar el partido
		}
		wait(jugadoresEnCancha[nroCancha]); //duermo al ultimo jugador
		// signal_all(jugadoresEnCancha[nroCancha]); //el encargo de iniciar el partido es el coordinador
	}

	procedure iniciarPartido(nroCancha in int){
		if(cantJugadoresEnCancha[nroCancha] < 10){
			wait(comenzarPartido);
		}
		signal_all(jugadoresEnCancha[nroCancha]); //despierta a todos los jugadores que esperan en esa cancha
		wait(finalizarPartido);
	}

	procedure terminarPartido(nroCancha in int){
		cantJugadoresEnCancha[nroCancha] = 0;
		signal_all(finalizarPartido);
	}
}




Process jugador[id: 0..19]{
	int nroEquipo; 
	PreparandoEquipo.avisarQueLlegueJugador(nroEquipo);
	PreparandoEquipo.enQueCanchaJuego(nroEquipo,nroCancha);
	Partido.irHaciaCancha(nroCancha);
}

Process coordinador[nroCancha: 1..2]{
	PreparandoEquipo.hayRival();
	Partido.AsignarCancha(nroCancha);
	Partido.iniciarPartido(nroCancha);
	delay(90); //jugando el partido
	Partido.terminarPartido(nroCancha);
}































































// En un entrenamiento de fútbol hay 20 jugadores que forman 4 equipos (cada jugador conoce 
// el equipo al cual pertenece llamando a la función DarEquipo()). Cuando un equipo está listo 
// (han llegado los 5 jugadores que lo componen), debe enfrentarse a otro equipo que también 
// esté listo (los dos primeros equipos en juntarse juegan en la cancha 1, y los otros dos equipos 
// juegan en la cancha 2). Una vez que el equipo conoce la cancha en la que juega, sus jugadores 
// se dirigen a ella. Cuando los 10 jugadores del partido llegaron a la cancha comienza el partido, 
// juegan  durante  50  minutos,  y  al  terminar  todos  los  jugadores  del  partido  se  retiran  (no  es 
// necesario que se esperen para salir).


// Monitor PreparandoEquipo{
// 	int cantJugadoresPorEquipo[4] = ([4],0);
// 	cond jugadoresListos[4];
// 	Queue equiposListos;

// 	cond hayRival;
// 	procedure avisarQueLlegueJugador(){
// 		int nroEquipo = DarEquipo();
// 		cantJugadoresPorEquipo[nroEquipo]++; //cada jugador sabe a que equipo pertenece
// 		if(cantJugadoresPorEquipo[nroEquipo] < 5){ //el 5to avisa que ya estan completos
// 			wait(jugadoresListos[nroEquipo])
// 		}
// 		equiposListos.enqueue(nroEquipo);
// 		wait(jugadoresListos[nroEquipo]); //para el ultimo jugador
// 		if(equiposListos.size()  > 1 ){
// 			signal(hayRival);
// 		}
// 	}

// 	procedure hayRival(){
// 		if(equiposListos.size()  < 2 ){
// 			wait(hayRival);
// 		}
// 	}
// }

// Monitor Partido{

// 	procedure buscarRival(){

// 	}
// }




// Process jugador[id: 0..19]{
// 	PreparandoEquipo.avisarQueLlegueJugador();
// }

// Process coordinador{
// 	int nroCancha = 1;
// 	for(int i = 0; i < 2; i++){
// 		PreparandoEquipo.hayRival();
// 	}
// }