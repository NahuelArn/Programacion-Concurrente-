// 2.  Existen N personas que deben ser chequeadas por un detector de metales antes de poder 
// ingresar al avión.  
	// a. Analice el problema y defina qué procesos, recursos y semáforos serán 
	// necesarios/convenientes,  además  de  las  posibles  sincronizaciones  requeridas  para 
	// resolver el problema. 
	// b. Implemente una solución que modele el acceso de las personas a un detector (es decir, 
	// si el detector está libre la persona lo puede utilizar; en caso contrario, debe esperar).  
	// c. Modifique su solución para el caso que haya tres detectores. 



	//==========Parte-A===============
	sem detectorFree = 1;
	// int cantPasaron = 0;
procees persona[id: 0..N-1] {
	// while (cantPasaron < N) { //aca el lifeTime, esta definido por la cantidad de personas/procesos, por eso no va un while
		P(detectorFree); //pido el detector
		//paso por el detector
		pasandoPorDetector(id); //uso el detector
		// cantPasaron++; //aumento la cantidad de personas que pasaron
		V(detectorFree); //libero el detector para que otros procesos puedan usarlo
	// }

}


	//==========Parte-B===============
		sem detectorFree = 1;
procees persona[id: 0..N-1] {
		P(detectorFree); //pido el detector
		//paso por el detector
		pasandoPorDetector(id); //uso el detector
		V(detectorFree); //libero el detector para que otros procesos puedan usarlo
}


	//==========Parte-C===============
sem detectorFree = 3; //3 detectores  //<await(detectorFree > 0 ) detectorFree-= 1;> Equivalente

procees persona[id: 0..N-1] {
		P(detectorFree); //pido el detector
		//paso por el detector
		pasandoPorDetector(id); //uso el detector
		V(detectorFree); //libero el detector para que otros procesos puedan usarlo
}