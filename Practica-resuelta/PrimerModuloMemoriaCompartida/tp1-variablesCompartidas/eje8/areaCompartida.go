// Desarrolle una solución de grano fino usando sólo variables compartidas (no se puede usar 
// las sentencias await ni funciones especiales como TS o FA). En base a lo visto en la clase 2 
// de teoría, resuelva el problema de acceso a sección crítica usando un proceso coordinador. 

// En este caso, cuando un proceso SC[i] quiere entrar a su sección crítica le avisa al coordinador, 
// y espera a que éste le dé permiso. Al terminar de ejecutar su sección crítica, el proceso SC[i] 
// le  avisa al  coordinador.  Nota:  puede  basarse en  la solución  para  implementar  barreras  con 
// Flags y Coordinador vista en la teoría 3.


//Soluciona, pero puede llegar a ser un poco confuso, ya que un recurso se utiliza para dos cosas distintas
//A nivel de eficiencia, no es la mejor solución, ya que el coordinador puede ser un cuello de botella
//con una Queue se podria mejorar
process esclavo[id: 1..N]{
	while (true) {
		llego[id] = true; // Indica que el proceso i quiere entrar a la sección crítica
		while (llego[i] == true) { skip} //skip salta a la siguiente iteración del while
		SC;
		llego[id] = true; // le avisa al coordinador que terminó su sección crítica
		llego[id] = false; //vuelve a indicar que no quiere entrar a la sección crítica
		SNC	; // Sección no crítica
	}
}

process coordinador::{
	while (true) {
		for (i = 1; i <= N; i++; st (llego[i] == true) ) {
			llego[i] = false; // INDICA al proceso que puede entrar a la sección crítica
			while (llego[i] == false) { skip} //skip salta a la siguiente iteración del while, LO MANTIENE ESPERANDO, hasta que le avisen que terminó

		}
	}
}

//
Dada  la  siguiente  solución  para  el  Problema  de  la  Sección  Crítica  entre  dos  procesos 
(suponiendo que tanto SC como SNC son segmentos de código finitos, es decir que terminan 
en algún momento), indicar si cumple con las 4 condiciones requeridas:  

int turno = 1;      
Process SC1::  
{ while (true) 
      {   while (turno == 2) skip;  
           SC;  //seccion Critica
           turno = 2; 
           SNC;  //Seccion No Critica
       } 
} 


Process SC2::  
{ while (true) 
      {   while (turno == 1) skip;  
           SC;  //seccion Critica
           turno = 1; 
           SNC;  //Seccion No Critica
       } 
}


