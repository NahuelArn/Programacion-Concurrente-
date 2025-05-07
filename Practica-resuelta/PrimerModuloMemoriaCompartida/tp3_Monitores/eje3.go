3.  Existen N personas que deben fotocopiar un documento. La fotocopiadora sólo puede ser 
usada  por  una  persona  a  la  vez.  Analice  el  problema  y  defina  qué  procesos,  recursos  y 
monitores serán necesarios/convenientes, además de las posibles sincronizaciones requeridas 
para resolver el problema. Luego, resuelva considerando las siguientes situaciones: 

a. Implemente  una  solución  suponiendo  no  importa el  orden  de  uso.  Existe  una  función 
Fotocopiar() que simula el uso de la fotocopiadora.  

b. Modifique la solución de (a) para el caso en que se deba respetar el orden de llegada. 

c. Modifique la solución de (a) para el caso en que se deba dar prioridad de acuerdo con la 
edad de cada persona (cuando la fotocopiadora está libre la debe usar la persona de mayor 
edad entre las que estén esperando para usarla). 

d. Modifique la solución de (a) para el caso en que se deba respetar estrictamente el orden 
dado por el identificador del proceso (la persona X no puede usar la fotocopiadora hasta 
que no haya terminado de usarla la persona X-1). 

e. Modifique la solución de (b) para el caso en que además haya un Empleado que le indica 

a cada persona cuando debe usar la fotocopiadora. 

f. Modificar la solución (e) para el caso en que sean 10 fotocopiadoras. El empleado le indica 
a la persona cuál fotocopiadora usar y cuándo hacerlo.

//=========================================================================================================


//A
Monitor Impresora{
	procedure usarImpresora(){
		//simulo el uso de la impresora
	}
}

Process Persona[id: 0..N-1]{
	Impresora.usarImpresora();
}

//=========================================================================================================


//B

Monitor Impresora{
	cond queue;
	int cantEnQueue = 0;
	boolean ocupado = false;

	procedure usarImpresora(){
		if(!ocupado{
			ocupado = true;
		}else{
			cantEnQueue++;
			wait(queue); // espero a que haya lugar
		}
	}

	procedure salir(){
		if(cantEnQueue > 0){
			cantEnQueue--;
			signal(queue); // despierto a un lector
		}else{
			ocupado = false;
		}
	}
}

process Persona[id: 0..N-1]{
	Impresora.usarImpresora(); // el lector
	Impresora.salir();
}

//=========================================================================================================


//C
//Aca se usa un vector de como semaforos//// CondicionVc

Monitor Impresora{
	cond espera[N];
	boolean libre = true;
	Queue queue;

	procedure usarImpresora(int in id, int in edad){
		if(!libre{
			insertarOrdenado(queue, id, edad);
			wait(espera[id]);
		}else{ //si esta libre, lo uso
			libre = false;
		}
	}

	procedure salir(){
		int id;
		if(empty(espera)){
			libre = true;
		}else{
			id = sacar(espera);
			signal(espera[id]);
		}
	}
}

Process Persona[id: 0..N-1]{
	Impresora.usarImpresora(id, edad); // el lector entra a la base de datos
	//simulo el uso de la impresora
	Impresora.salir(); // el lector sale de la base de datos
}

//=========================================================================================================


//D

Monitor Impresora{
	con vec[N];
	posDormidos[N] = ([N],0); //indica que posicion tiene un proceso durmiendo

	int idActual = 0;
	Queue queue;

	procedure usarImpresora(int id){
		while(id != idActual){
			posDormidos[id] = 1; //indico que estoy dormido
			wait(vec[id]); //dormido	
		}
		queue.push(id); //me agrego a la cola //para avisar que estoy usando la impresora
		//usa la impresora
	}
	procedure salir(){
		if(!queue.empty()){
			int id = queue.pop(); //saco el id de
			idActual++;
			if(posDormidos[idActual] == 1){
				signal(vec[idActual]); // despierto a un lector
			}
		}
	}
}

Process Persona[id: 0..N-1]{
	Impresora.usarImpresora(id);
	//usandoImpresora
	Impresora.salir();

}


//=========================================================================================================
//E
//PreCondiciones

//Respetar el orden de llegada
//1 empleado coordinador
Monitor impresora{
	cantPedidos=0;

	cond pedidos; // usado para simular el aviso de Persona hacia el Coordinador de recibir Permiso de acceso
	cond turno;	//usado para simular la Espera de turno de la Persona
	cond termino; //usado Para la comunicacion entre Persona y Coordinador

	procedure usarImpresora(){
		cantPedidos++;
		signal(pedidos); // Aviso al coordinador, que ya me puede dar permiso
		wait(turno); //Duemo la instancia actual
		//usandoImpresora...
	}

	procedure darAcceso(){
		 
		if(cantPedidos == 0){ //si todavia no hay nadie esperando el acceso, me duermo hasta que me avisen que hay 1 pedido 
			wait(pedidos);
		}
		// else{   #sarasa1
		// 	signal(pedidos); //doy permiso de usar la impresora al que este, esperando permiso de acceso
		// }
		signal(turno); //doy permiso de usar la impresora al que este, esperando permiso de acceso
		cantPedidos--;
		wait(termino); //no doy permiso a un nuevo proceso(duermo al coordinador), hasta que me avise el proceso que esta usando la impresora que ya termino de usarla
	}

	procedure salir(){
		signal(termino); //aviso como Persona, al coordinador que ya termine de usar la impresora, asi otra Persona la puede usar
	}
}

Process Persona[id: 0..N-1]{
	Impresora.usarImpresora();
	//usandoImpresora();
	Impresora.salir();
}

process Coordinador{	
	for(int i =0; i < N-1; i++){
		Impresora.darAcceso();
	}
}

//Flujo de ejecucion-> coordinador espera que cantPedidos>0, si es menor se duerme...

//#sarasa1
Que pasa si el primero en entrar es darAcceso? //coordinador por alguna razon no gano el Monitor Primero
darAcceso hace...
cantPedidos+1;
signal(pedidos);  //"despierta al primer proceso dormido" pero no hay ningun proceso dormido // jace un signal a ningun proceso... ahora
//segun esto, no hay problema: "Hacer signal en una variable condición vacía no causa ningun error."
//entonces hacemos un signal al vacio
wait(turno);

despues de eso puede llegar a ganar el monitor otro proceso Persona y repetir este ciclo 1 o X veces
Cuando gana el monitor el Coordinador va hacer
cantPedidos > 0, entonces no entra al if 
Directamente entra a dar permiso no espera un aviso.. entonces esta bien la solucion planteada.
//=========================================================================================================


//cambiar logica de parametros
//ejemplardo  in -> entrada, out -> salida 

// NombreVar  Salida/Entrda  TipoDeDato
// ---------  ------------   -----------
parametros(id : in int, saludo: out string){
	saludo = saludoParaElId(id);
}


//F

//PreCondiciones

//Respetar el orden de llegada
//1 empleado coordinador
//cant impresoras 10
Monitor impresora{
	Queue queue;
	// cantPedidos=0; en este caso no me sirve, necesito entre procesos tener comunicacion de ids

	cond pedidos;  
	cond terminoImpresora; 
	
	Queue impresoras = {1,2,3,4,5,6,7,8,9,10};
	Queue impresoraAsignadaId = ([N], 0);
	Queue personasEsperando;

	cond personasDormidas[N];

	procedure usarImpresora(idC : in int, impresora: out int){
		personasEsperando.push(idC); //me pusheo (id)
		signal(pedidos);  
		wait(personasDormidas[idC]);  

		impresora = impresoraAsignadaId[idC];
	}

	procedure darAcceso(){
		if(personasEsperando.isEmpty()){ //si no hay personas esperando
			wait(pedidos);
		} //si hay personas esperando
		int idC = personasEsperando.pop();
		if(impresoras.isEmpty()){ //si hay impresoras libres
			wait(terminoImpresora);
		}
		impresoraAsignadaId[id] = impresoras.pop(); //le da una impresora 
		signal(personasDormidas[idC]); //aviso a la persona q ya tiene una impresora asiganda
		// cantPedidos--;
	}

	procedure salir(impresora: in Printer){
		impresoras.push(impresora); //libero a la impresora que estaba en uso
		signal(terminoImpresora); 
	}
}

Process Persona[id: 0..N-1]{
	int impresora;
	Impresora.usarImpresora(id,impresora);

	usandoImpresora(impresora);
	//navegando en chrome
	//tomando un cafe
	//impresion finalizada...

	Impresora.salir(impresora);
}

process Coordinador{	
	for(int i =0; i < N-1; i++){
		Impresora.darAcceso();
	}
}







//clean
// Monitor Fotocopiadora{
//     Cola libres = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
//     Cola asignada = [N][0]
//     Cola llegados
//     cond impresora;
//     cond personas [n] 
//     cond empleado;
//     Procedure asignar(){
//         if (llegados.empty()) 
//             wait(empleado)
//         aux= llegados.pop()
//         if (libres.empty())
//             wait(impresora)
//         asignada[aux] = libres.pop()
//         signal(personas[aux])
//     }
//     Procedure usar (idP: in int; fotocop: out int){
//         llegados.push()
//         signal(empleado)
//         wait(personas[idP])
//         fotocop = asignada[idP]
//     }
//     Procedure salir(fotocop: in int){
//         libres.push(fotocop)
//         signal(impresora)
//     }
// }

// Process Persona [id:1..N]{
//     fotocopiadora.usar(id)
//     //fotocopiar
//     fotocopiadora.dejar(fotocop)
// }

// Process Empleado {
//     for (i=1; i<=N; i++)
//         fotocopiadora.asignar()    
// }