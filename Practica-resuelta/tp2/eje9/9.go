// 9.  Resolver el funcionamiento en una fábrica de ventanas con 7 empleados (4 carpinteros, 1 
// vidriero y 2 armadores) que trabajan de la siguiente manera: 

// • Los carpinteros continuamente hacen marcos (cada marco es armando por un único 
// carpintero) y los deja en un depósito con capacidad de almacenar 30 marcos. 

// • El vidriero continuamente hace vidrios y los deja en otro depósito con capacidad para 
// 50 vidrios. 

// • Los  armadores  continuamente  toman  un  marco  y  un  vidrio  (en  ese  orden)  de  los 
// depósitos correspondientes y arman la ventana (cada ventana es armada por un único 
// armador).

//=================================
//estos 3 tal vez no son necesarios?
sem carpinteroSem = 4;
sem vidrieroSem = 1;
sem armadorSem = 2;
//=================================

Queue <Marco> depositoMarcos;
sem protexDepositoMarcos = 1;
sem maxMarcos = 30
sem cantMarcosHechos = 0;


Queue <Vidrio> depositoVidrios;
sem protexDepositoVidrios = 1;
sem maxVidrios = 50
sem cantVidriosHechos = 0;

// sem armandoVentanaSem = 1; 

process carpintero[id: 0..3]{
	while(true){
		ArmandoMarco marco = armarMarco();
		P(maxMarcos);
		P(protexDepositoMarcos); 
		depositoMarcos.enqueue(marco);
		V(cantMarcosHechos);
		V(protexDepositoMarcos); 
}

//estos 2 procesos son equivalentes?
process vidriero{
	while(true){
		ArmandoVidrio vidrio = armarVidrio();
		P(maxVidrios);
		P(protexDepositoVidrios);
		depositoVidrios.enqueue(vidrio);
		V(cantVidriosHechos);
		V(protexDepositoVidrios);
	}
}

process armador[id: 0..1]{
	while(true){
		P(cantMarcosHechos); //si hay marcos en el deposito, no espero a que haya uno
		P(protexDepositoMarcos);
		Marco marco = depositoMarcos.dequeue(); // saco un marco de la cola
		V(protexDepositoMarcos); // libero el semaforo para usar la Queue de marcos
		V(maxMarcos); // Ya use un marco, que vuelva a tener stock el deposito
		
		P(cantVidriosHechos); 
		P(protexDepositoVidrios);
		Vidrio vidrio = depositoVidrios.dequeue(); // saco un vidrio de la cola
		V(protexDepositoVidrios); // libero el semaforo para usar la Queue de vidrios
		V(maxVidrios); // Ya use un vidrio, que vuelva a tener stock el deposito
		
		//Esta parte nose bien como hacerla 
		// "cada ventana es armada por un único armador"
		// P(armandoVentanaSem);
		armandoVentana(marco, vidrio); // armo la ventana

		}
	}
}