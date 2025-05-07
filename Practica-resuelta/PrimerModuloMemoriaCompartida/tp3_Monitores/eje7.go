Se  debe  simular  una  maratón  con  C  corredores  donde  en  la  llegada  hay  UNA  máquina 
expendedoras de agua con capacidad para 20 botellas. Además, existe un repositor encargado 
de reponer las botellas de la máquina. Cuando los C corredores han llegado al inicio comienza 
la carrera. Cuando un corredor termina la carrera se dirigen a la máquina expendedora, espera 
su turno (respetando el orden de llegada), saca una botella y se retira. Si encuentra la máquina 
sin  botellas,  le  avisa  al  repositor  para  que  cargue  nuevamente  la  máquina  con  20  botellas; 
espera a que se haga la recarga; saca una botella y se retira.  Nota: mientras se reponen las 
botellas se debe permitir que otros corredores se encolen. 


//llegan lo C corredores llegan al inicio
//comienza la carrera
//termina la carrera
//se dirigen a la maquina expendedora
//espera su turno
//sacar una botella
//se retira



Monitor Carrera{
	int cantCorredores = 0;

	cond corredoresEsperando;

	procedure coordinarSalida(){
		cantCorredores++;
		if(cantCorredores == C){
			signal_all(corredoresEsperando);
		}else{
			wait(corredoresEsperando);
		}
	}
}

Monitor PosCarrera{
	int cantBotellas = 20;
	int cantEnEspera = 0;
	boolean libre = true;
	cond corredoresEsperando;
	cond esperandoReStockeo;
	cond repositor; 
	boolean hayStock = true;

	procedure sacarBotella(id : in int){
		if(!libre){
			cantEnEspera++;
			wait(corredoresEsperando);
		}else{
			libre = false;
		}
		if(cantBotellas  == 0){
			hayStock = false;
			signal(repositor);
			wait(esperandoReStockeo);
		}
		if(cantEnEspera > 0){
			cantEnEspera--;
			signal(corredoresEsperando);
		}else{
			libre = true;
		}
		cantBotellas--;
	}


	procedure reponerBotellas(){
		if(hayStock){
			wait(repositor);
		}
		cantBotellas = 20; //repongo stock
		hayStock = true;
		signal(esperandoReStockeo);
	}
}


process Corredor(int id: 0..C-1){
	Carrera.coordinarSalida();
	PosCarrera.sacarBotella(id);
}

process Repositor(){
	while(true){
		PosCarrera.reponerBotellas();
	}
}




































Monitor Carrera{
	int cantCorredores = 0;

	cond corredoresEsperando;

	procedure coordinarSalida(){
		cantCorredores++;
		if(cantCorredores == C){
			signal_all(corredoresEsperando);
		}else{
			wait(corredoresEsperando);
		}
	}
}

Monitor PosCarrera{
	int cantBotellas = 20;
	boolean hayStock = true;
	cond botellas;
	int cantEnEsperaDeStock = 0;
	cond repositor;
	// Queue corredoresEsperandoStock;

	cond corredoresEsperandoStock;
	boolean hayFila = false;

	cond haciendoFila;

	procedure sacarBotella(id : in int){
		if(hayStock && !hayFila){
			cantBotellas--;
		}else{
			if(hayStock && hayFila){
				wait(haciendoFila);
			}else{
				if(!hayStock){ //si hay o no hay fila tengo que avisar al repositor y esperar
					signal(repositor);
					conStock = false;
					wait(haciendoFila);
				}
			}
			cantBotellas--;
		}
	}


	procedure reponerBotellas(){
		if(hayStock){
			wait(repositor);
		}
		cantBotellas = 20; //repongo stock
		hayStock = true;
		signal(haciendoFila);
	}
}


process Corredor(int id: 0..C-1){
	Carrera.coordinarSalida();
	Carrera.sacarBotella(id);
}

process Repositor(){
	while(true){
		PosCarrera.reponerBotellas();
	}
}
























































Se  debe  simular  una  maratón  con  C  corredores  donde  en  la  llegada  hay  UNA  máquina 
expendedoras de agua con capacidad para 20 botellas. Además, existe un repositor encargado 
de reponer las botellas de la máquina. Cuando los C corredores han llegado al inicio comienza 
la carrera. Cuando un corredor termina la carrera se dirigen a la máquina expendedora, espera 
su turno (respetando el orden de llegada), saca una botella y se retira. Si encuentra la máquina 
sin  botellas,  le  avisa  al  repositor  para  que  cargue  nuevamente  la  máquina  con  20  botellas; 
espera a que se haga la recarga; saca una botella y se retira.  Nota: mientras se reponen las 
botellas se debe permitir que otros corredores se encolen. 


//llegan lo C corredores llegan al inicio
//comienza la carrera
//termina la carrera
//se dirigen a la maquina expendedora
//espera su turno
//sacar una botella
//se retira



Monitor Carrera{
	int cantCorredores = 0;

	cond corredoresEsperando;

	procedure coordinarSalida(){
		cantCorredores++;
		if(cantCorredores == C){
			signal_all(corredoresEsperando);
		}else{
			wait(corredoresEsperando);
		}
	}
}

Monitor PosCarrera{
	int cantBotellas = 20;
	boolean conStock = true;
	cond botellas;
	int cantEnEsperaDeStock = 0;
	cond repositor;
	// Queue corredoresEsperandoStock;

	cond corredoresEsperandoStock;
	boolean hayFila = false;

	procedure sacarBotella(id : in int){
		if(hayFila){
			wait(corredoresEsperando);
		}else{
			if(cantBotellas == 0){	//no hay stock y no habia fila
				conStock = false;
				signal(botellas);
				cantEnEsperaDeStock++;
				hayFila = true;
				wait(corredoresEsperandoStock);
			}else{ //hay stock y no hay fila
				cantBotellas--;
			}
		}
		
		if(cantBotellas == 0){
			conStock = false;
			signal(botellas);
			cantEnEsperaDeStock++;
			wait(corredoresEsperandoStock);
		}
		if(cantEnEsperaDeStock > 0){
			cantEnEsperaDeStock--;
			signal(corredoresEsperandoStock);
			wait(corredoresEsperando);
		}
		cantBotellas--;
	}
	// procedure sacarBotella(id : in int){


	// 	if(cantBotellas == 0){
	// 		conStock = false;
	// 		signal(botellas);
	// 		cantEnEsperaDeStock++;
	// 		wait(corredoresEsperandoStock);
	// 	}
	// 	if(cantEnEsperaDeStock > 0){
	// 		cantEnEsperaDeStock--;
	// 		signal(corredoresEsperandoStock);
	// 		wait(corredoresEsperando);
	// 	}
	// 	cantBotellas--;
	// }

	procedure reponerBotellas(){
		if(conStock){
			wait(repositor);
		}
		cantBotellas = 20; //repongo stock
		conStock = true;
		signal(corredoresEsperandoStock);
	}
}


process Corredor(int id: 0..C-1){
	Carrera.coordinarSalida();
	Carrera.sacarBotella(id);
}

process Repositor(){
	while(true){
		PosCarrera.reponerBotellas();
	}
}