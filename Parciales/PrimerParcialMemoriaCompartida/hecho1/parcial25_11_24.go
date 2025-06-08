//PARTE SEMAFOROS

Resolver con SEMÁFOROS el siguiente problema. La Clave Única de Identificación Tributaria (CUIT) es una clave que 
se utiliza en el sistema tributario de la República Argentina para poder identificar correctamente a las personas físicas 
o  jurídicas.  Consta  de  un  total  de  once  (11)  cifras  numéricas,  siendo  la  última  un  dígito  verificador  (del 0 al  9).  Una 
empresa  cuenta  con  una  lista  de  CUITs  que  debe  procesar,  debiendo  informar  la  cantidad  de  CUITs  por  dígito 
verificador. Para ello, dispone de un software que emplea 5  workers, los cuales trabajan colaborativamente 
procesando  de  a  una  CUIT  por  vez  cada  uno.  Al  finalizar  el  procesamiento,  el  último  worker  en  terminar  debe 
informar  los  resultados  del  procesamiento.  Notas:  la  función  obtenerDV(CUIT)  retorna  el  dígito  verificador  para  la 
CUIT  recibida  como  entrada.  La  lista  de  CUITs  se  almacena  como  una  cola  global  y  la  solución  debe  maximizar  la 
concurrencia.


Queue cuits; //se dispone
sem protexQueue = 1;
sem protexCant = 1;

int cantXdig[9] = ([9], 0); //0..9
sem protexVec[9] = ([9], 1);

// sem protexVec = 1;
int cant = 0;
Process Worker[id: 0..4]{
	p(protexQueue);
	while(!cuits.isEmpty()){
		Cuit cuit = cuits.dequeue();
		V(protexQueue);
		int digVerficador = obtenerDV(cuit);

		P(protexVec[digVerficador]); //de esta forma varios procesos acceden al vector a la vez, pero solo al campo que necesitan, sin parar a los demas
		cantXdig[digVerficador]
		V(protexVec[digVerficador]);
		//
		// P(protexVec); //de esta forma no se maximiza la concurrencia
		// cantXdig[digVerficador] += 1;
		// V(protexVec);
		//
		P(protexQueue);
	}
	V(protexQueue);
	P(protexCant);
	cant++;



	if(cant == 5){
		for(int i = 0; i < 9; i++){
			informar("la cantida del dig "+i+ " es: "+cantXdig[i]); //como es el ultimo no puede pasar que varios accedan a esto
		}
	}
	V(protexCant);
}





//Cualquiera, era con semaforos no con monitores
// Queue cuits; //se dispone

// //tengo 4 empleados, cada empleado tiene un lugar de trabajo donde puede laburar
// Process ProcesarCuit[id: 0..4{	

// }


// Process Worker[id: 0..4]{
// 	while(cuits.size() > 0){

// 	}
// }















Resolver con  MONITORES la siguiente situación. En un negocio hay UN empleado que diseña tarjetas  digitales. 
El  empleado  debe  atender  los  pedidos  de  C  clientes,  de  acuerdo  con  el  orden  en  que  se  hacen  los  pedidos.  El 
cliente  envía  las  indicaciones,  y  el  empleado  en  base  a  eso  diseña  la  tarjeta  y  se  la  envía  al  cliente.  Notas: 
maximizar la concurrencia; existe una función HacerTarjeta(indicaciones) que simula el armado de la tarjeta por 
parte del empleado; todos los procesos deben terminar su ejecución.

Monitor Pedido{
	Queue pedido[C];
	Queue filaTomaPedidoCliente;

	cond empleado;

	procedure realizarPedido(id in int, indicacion in Indicacion){
		filaTomaPedidoCliente.push(id);
		pedido[id] = indicacion;
		signal(empleado);
	}

	procedure tomarPedido(siguienteId : out int ,indicacion out Indicacion){
		if(filaTomaPedidoCliente.isEmpty()){
			wait(empleado);
		}
		int siguienteId = filaTomaPedidoCliente.pop();
		indicacion = pedido[siguienteId];
	}

}

Monitor EsperarPedido{
	cond cliente[C];
	//Tarjeta tarjeta;

	Tarjeta tarjetas[C] = ([C, null]);

	procedure esperandoTarjeta(id : in int,tarjetaO out Tarjeta){
		if(tarjetas[id] == null){
			wait(cliente[id]);
		}
		tarjetaO = tarjetas[id];
	}

	procedure entregarTarjeta(idO : in int, tarjeta : in Tarjeta){
		tarjetas[idO] = tarjeta;
		signal(cliente[idO]); //si estaba esperando le aviso
	}
}

Process Cliente[id: 0..C]{
	Indicacion indicacion;
	Tarjeta tarjetaO;
	Pedido.realizarPedido(id, indicacion);
	EspererandoPedido.esperandoTarjeta(id,tarjetaO);
	//mirandoTarjeta
}

Process Empleado{
	int cant = 0;
	int idO;
	while (cant < C){
		cant++;
		Indicacion indicacion;
		Pedido.tomarPedido(idO,indicacion);
		Tarjeta tarjeta = hacerTarjeta(indicacion); //funcion interna (Se dispone)
		EsperarPedido.entregarTarjeta(idO, tarjeta);

	}
}	