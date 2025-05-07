Se dispone de un puente por el cual puede pasar un solo auto a la vez. Un auto pide permiso 
para pasar por el puente, cruza por el mismo y luego sigue su camino


Monitor  Puente 
    cond cola;  
    int cant= 0; 
 
    Procedure entrarPuente () 
         while ( cant > 0) wait (cola); 
         cant = cant + 1;    
    end; 
 
    Procedure salirPuente () 
        cant = cant – 1; 
        signal(cola); 
    end; 
End Monitor;  
 
Process Auto [a:1..M] 
   Puente. entrarPuente (a); 
   “el auto cruza el puente” 
   Puente. salirPuente(a); 
End Process; 


a. ¿El código funciona correctamente? 
Justifique su respuesta. 
Si, funciona correctamente.
"
	Puente.entrarPuente(a); //se ejecuta
	si cant > 0 // se pushea el wait(cola) y se queda esperando //no vuelve a chequear en el while si hay autos en la cola, porque el semaforo no lo deja pasar
	cando no no hay autos en la cola, el auto entra al puente y ejecuta
	cant = cant + 1; //simboliza que el auto entra al puente
	//el auto cruza el puente (simulado por el print)
	Puente.salirPuente(a); //se ejecuta
	cant = cant - 1; //simboliza que el auto sale del puente
	signal(cola); //despierta al siguiente auto que estaba esperando en la cola
"

 
	 b. ¿Se  podría  simplificar  el  programa?  ¿Sin 
monitor? ¿Menos procedimientos? ¿Sin 
variable condition? En caso afirmativo, 
rescriba el código. 

Tiene que ser con Monitor, se puede simplificar, como aca no importa el orden podemos
manejarlo solo con exclusion mutua perse del monitor. Sin variables condition, Menos Procedimientos

"
Monitor puente
 procedure cruzandoPuente(){ 
	//recoriedo el puente
}

Process Auto [a:1..M]{ //por definicion, va ir mandando de a un proceso a la vez
	puente.cruzandoPuente(a); //se ejecuta 
}
"

c. ¿La  solución  original  respeta  el  orden  de 
llegada de los vehículos? Si rescribió el código 
en el punto b), ¿esa solución respeta el orden 
de llegada?

No, la solucion original no respeta el orden de llegada, el orden en el que van a cruzar los autos no es posible de determinarlo con ese codigo.

B: No, tampoco respeta el orden de llegada. No es posible de determinarlo con ese codigo, "Por definicion de Monitores".










//sarasa
a. ¿El código funciona correctamente? 
Justifique su respuesta. 
No, casos para los que genere inconsistencia:
1.  hay 1 Auto
	cant+1 // en entrarPuente
	ejecuta el puente.salirPuete(a);
	cant-1; // hay 0 autos en el puente
	aplica un signal(cola); // despierta a un auto que no existe
2. hay 2 autos
	cant+1 // en entrarPuente
	el 2do hace bussy waiting en el while (cant > 0) wait (cola); // se pushea N veces el wait(cola)
3. hay 2 autos
	 cant+1 // en entrarPuente
	 //se pushea N veces el 2do a N autos
	 ejecuta el puente.salirPuete(a);
	 cant-1; // resta al auto actual
	 signal(cola); // despierta a un auto, pero nunca verifica si hay autos en la Queue 
	 //si no hay autos en la queue deberia pasar directame el proximo que llegue, no lleva ningun condicional para hacer el Passing p2p

4. etc etc etc... 
	 b. ¿Se  podría  simplificar  el  programa?  ¿Sin 
monitor? ¿Menos procedimientos? ¿Sin 
variable condition? En caso afirmativo, 
rescriba el código. 
c. ¿La  solución  original  respeta  el  orden  de 
llegada de los vehículos? Si rescribió el código 
en el punto b), ¿esa solución respeta el orden 
de llegada?

