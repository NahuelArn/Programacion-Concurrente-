// 8.  Una fábrica de piezas metálicas debe producir T piezas por día. Para eso, cuenta con E 
// empleados que se ocupan de producir las piezas de a una por vez. La fábrica empieza a 
// producir una vez que todos los empleados llegaron. Mientras haya piezas por fabricar, los 
// empleados tomarán una y la realizarán. Cada empleado puede tardar distinto tiempo en 
// fabricar una pieza. Al finalizar el día, se debe conocer cuál es el empleado que más piezas 
// fabricó. 
// a. Implemente una solución asumiendo que T > E. 
// b. Implemente una solución que contemple cualquier valor de T y E (incluso T<E).


// T ---> Piezas
// E ---> Empleados
sem barrera = 0;
int cantEmpleadosActual = 0;
sem empleadoSem = 1;
int cantPiezas = T;
cantPiezasPorEmpleado[0..E] = ([E] 0);
Queue <Pieza> piezas;
sem protexQueuPiezas = 1;

sem cantPiezasVariable = 1;


//para calcular el max

int max = -1;
int idMax = -1;
sem protexMax = 1;


int cantPiezasActual = 0;
process empleado[id: 0..E-1]{
	P(empleadoSem);
	cantEmpleadosActual++;
	if(cantEmpleadosActual == E){
		for (i = 0; i < E; i++){
			V(barrera);	// El último toca el timbre
		}
	}
	V(empleadoSem);
	P(barrera); // se espera a que todos lleguen
	P(cantPiezasVariable); // pido el semaforo para usar la variable de cantPiezas
	while(cantPiezasActual < cantPiezas){
		V(cantPiezasVariable);
		P(protexQueuPiezas); // pido el semaforo para usar la Queue de piezas
		if(piezas.size() != 0){
			Pieza piesa = piezas.dequeue(); // saco una pieza de la cola	
			V(protexQueuPiezas); // libero el semaforo para usar la Queue de piezas
			realizandoPieza(piesa);
			cantPiezasPorEmpleado[id]++; // aumento el contador de piezas por empleado  //ASUMO QUE NO VAN A VER +1 PROCESO CONCURRENTE QUE VA A QUERER ACCEDER A [X] EN EL MISMO MOMENTO
			P(cantPiezasVariable); // pido el semaforo para usar la variable de cantPiezas
			cantPiezasActual++;
			V(cantPiezasVariable); // libero el semaforo para usar la variable de cantPiezas
		}else{
			V(protexQueuPiezas);
		}
		P(cantPiezasVariable); 
	}
	V(cantPiezasVariable);
	//Fin de jornada
	P(protexMax); // pido el semaforo para usar la variable de max
	if(cantPiezasPorEmpleado[id] > max){ // si el empleado hizo más piezas que el maximo
		idMax = id; // guardo el id del empleado
		max = cantPiezasPorEmpleado[id]; // guardo el maximo
	}
	V(protexMax); // libero el semaforo para usar la variable de max
	//al ser concurrentemente como saco el maximo?
}


//Preguntar a alguien
Cuando tengo un vector semaforo, y quiero modificar/leer un campo

debo proteger todo el vector con un semaforo? o solo el campo que quiero modificar/leer? -->si es la segunda opcion lo protegeria con un vector semaforo no? q suportee al vector de cantPiezasPorEmpleado