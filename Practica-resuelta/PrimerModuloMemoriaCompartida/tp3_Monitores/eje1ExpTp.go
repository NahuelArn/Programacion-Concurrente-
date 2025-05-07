

Existen N personas que desean utilizar un cajero automatico. En este primer caso no se debe tener en cuenta el orden de llegada de las personas 
(cuando esta libre cualquiera lo puede usar). Supogonga que hay una funcion UsarCajero(); que simula el uso del cajero.

// swapeo de tp2 a tp3
//uso de recurso compartido = monitor

//todo dentro del semaforo se ejecuta de forma atomica
Monitor Cajero{
	String saludo = "hola soy el cajero";
	Procedure PasarAlCajero(){
		print(saludo); // saludo al cliente
		UsarCajero(); // simulo el uso del cajero
	}
}

process Persona[id: 0..N-1]{
	//cualquiera que este esperando para usar Cajero, puede entrar al caajero (no hay un criterio de prioridad)
	//el hecho de usar semaforos, ya me da la exclusion mutua, solo un proceso a la vez puede usar el cajero (el semaforo me garantiza esto)
	Cajero.PasarAlCajero(); 
}





//eje2

Existen N personas que desean utilizar un cajero automatico. En este segundo caso se debe tener en cuenta el orden de llegada de las personas 
(cuando esta libre cualquiera lo puede usar). Supogonga que hay una funcion UsarCajero(); que simula el uso del cajero.
//v1 ESTA MALL
// process Cajero{
// 	cond vc;
// 	procedure Adormir({
// 		wait(vc);
// 	}
// 	process PasarAlCajero(){
// 		signal(vc); // despierto a una persona
// 		UsarCajero(); // simulo el uso del cajero
// 	}
// }

// process Persona[id: 0..N-1]{
// 	Cajero.Adormir();
// 	Cajero.PasarAlCajero();
// }

//V2 Este esta bien
Monitor Cajero {
	cond vc;
	boolean ocupado = false;
	int cantEnQueue = 0;
	Procedure entrarAlBanco(){
		if(ocupado == true){
			cantEnQueue += 1;
			wait(vc); // espero a que el cajero se desocupe
		}else{
			ocupado = true;
		}
	}
	
	Procedure salirCajero(){
		if(cantEnQueue > 0){ //verifico si hay gente esperando en la cola
			cantEnQueue -= 1;
			signal(vc);
		}else{
			ocupado = false; // si no hay gente en la cola, dejo listo el cajero para el prox q llegue
		}
	}

}

Process Persona[id: 0..N-1]{
  entrarAlBanco();
	usarCajero();
	salirCajero();
}