// 11. En un vacunatorio hay un empleado de salud para vacunar a 50 personas. El empleado 
// de salud atiende a las personas de acuerdo con el orden de llegada y de a 5 personas a la 
// vez.  Es  decir,  que  cuando  está  libre  debe  esperar  a  que  haya  al  menos  5  personas 
// esperando, luego vacuna a las 5 primeras personas, y al terminar las deja ir para esperar 
// por otras 5. Cuando ha atendido a las 50 personas el empleado de salud se retira. Nota: 
// todos  los  procesos  deben  terminar  su  ejecución;  asegurarse  de  no  realizar  Busy  Waiting; 
// suponga que el empleado tienen una función VacunarPersona() que simula que el empleado 
// está vacunando a UNA persona.

Queue<Persona> colaVacunacion;
sem barrera = 1;

sem mutexCola = 1;

sem hay5 = 0;
int cantFila = 0;
sem mePuedoRetirar[0..49] = ([50] 0); // semaforos para que cada persona se retire cuando la vacunan


int cantTotalAtenciones =0;
process persona[id: 0..49]{
	P(mutexCola);
	cantFila++;
	colaVacunacion.enqueue(id);
	if(cantFila == 5){
		for (i = 0; i< 5; i++){
			V(barrera);
		}
		cantFila = 0;
	}
	V(mutexCola);
	P(barrera);
	//V(hay5); // le avisa al empleado que hay 5 personas //uno solo estaria mal, llegan 5 queriando hacer V(hay5)
	// for(int i = 0; i<5; i++){ //Tambien estaria mal, x cada uno haria esto... osea 5*5 = 25 veces, la primera estaba bien, implicitamente ya lo hago con el V(hay5), esto ya manda 5 mensajes al empleado
	// 	V(hay5); // espero a que el empleado me vacune
	// }
	V(hay5);
	P(mePuedoRetirar[id]);
}

process empleadoSalud{
	while (cantTotalAtenciones < 50) {
		// P(hay5) //tiene que recibir los 5 mensajes de las 5 personas
		for (int i = 0; i < 5; i++){
			P(hay5); // espero a que haya 5 personas
		}
		P(mutexCola);
		for (int i = 0; i < 5; i++){
			Persona p = colaVacunacion.dequeue();
			V(mutexCola);
			vacunarPersona(p);
			V(mePuedoRetirar[p.id]); // le avisa a la persona que ya se puede ir
			P(mutexCola);
		}
		V(mutexCola);
		cantTotalAtenciones += 5; //no necesita el semaforo porque no hay otro proceso que lo modifique
	}
}




































//================================================================

//Si no controlo el tema de las 50, va funcionar, pero si serian 60 personas que llegan y cuando completa las 50 se tiene que ir, 
// ahi si necesito controlarlo, igual lo dejo controlando tambien para >= 50
Queue<Persona> colaVacunacion;
sem protexQueue = 1;
sem limiteDeAtencion = 5;

sem barrera = 0;
int cantAtenciones =0;
int cantAtencionesTotal = 0;
sem protexEmpleadoSC = 1;


process Persona[id: 0..49]{
	while (cantAtencionesTotal < 50) { //podes tener 60 personas, pero solo atendes 50... las otras 10 pasan de largo
		P(protexQueue);
		colaVacunacion.enqueue(id);
		V(protexQueue);

		P(protexEmpleadoSC);
		cantAtenciones++;
		if(cantAtenciones == 5){
			for(i =0; i < 5; i++){
				V(barrera); // El último toca el timbre
			}
			cantAtenciones = 0;
			cantAtencionesTotal+= 5;
		}
		V(protexEmpleadoSC);
		P(barrera);// se espera a que 5 lleguen

		int i = 0;
		P(protexQueue);
		// while (queue.isEmpty() == false && i < 5 && (cantAtencionesTotal < 50)) { //SI ESTOY ACA, SE QUE HAY 5 PERSONAS EN LA COLA
		while (i < 5 ) {
			vacunarPersona(colaVacunacion.dequeue()); // vacuno a la persona
			V(protexQueue);		
			i++;
			P(protexQueue);
		}
		V(protexQueue); // libero el semaforo para usar la Queue de piezas
	}
}













//Con coordinador
Queue<Persona> colaVacunacion;
sem mutexCola = 1;
sem mutexContador = 1;

sem barrera = 0;           // Para liberar a las personas vacunadas
sem listoParaVacunar = 0; // Para avisar al empleado cuando hay 5

int personasEsperando = 0;
int vacunados = 0;

process Persona[id: 0..49] {
    P(mutexCola);
    colaVacunacion.enqueue(id);
    V(mutexCola);

    P(mutexContador);
    personasEsperando++;
    if (personasEsperando == 5) {
        V(listoParaVacunar); // Le avisa al empleado que hay 5
    }
    V(mutexContador);

    P(barrera); // Espera a ser vacunado
}

process Empleado {
    while (vacunados < 50) {
        P(listoParaVacunar); // Espera a que haya 5

        for (int i = 0; i < 5; i++) {
            P(mutexCola);
            Persona p = colaVacunacion.dequeue(); // Obtiene el ID de la persona
            V(mutexCola);

            vacunarPersona(p); // Vacuna a la persona
        }

        P(mutexContador);
        personasEsperando = 0;
        vacunados += 5;
        V(mutexContador);

        // Libera a las 5 personas
        for (int i = 0; i < 5; i++) {
            V(barrera);
        }
    }
}





















//Si no controlo el tema de las 50, va funcionar, pero si serian 60 personas que llegan y cuando completa las 50 se tiene que ir, 
// ahi si necesito controlarlo, igual lo dejo controlando tambien para >= 50
Queue<Persona> colaVacunacion;
sem protexQueue = 1;
sem limiteDeAtencion = 5;

sem barrera = 0;
int cantAtenciones =0;
int cantAtencionesTotal = 0;
sem protexEmpleadoSC = 1;


process empleadoSaludo[id: 0..49]{
	P(protexQueue);
	colaVacunacion.enqueue(id);
	V(protexQueue);

	P(protexEmpleadoSC);
	cantAtenciones++;
	if(cantAtenciones == 5){
		for(i =0; i < 5; i++){
			V(barrera); // El último toca el timbre
		}
		cantAtenciones = 0;
		cantAtencionesTotal+= 5;
	}
	V(protexEmpleadoSC);
	P(barrera);// se espera a que 5 lleguen
}

