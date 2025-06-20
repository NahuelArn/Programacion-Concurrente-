3.  Se  debe  modelar  el  funcionamiento  de  una  casa  de  comida  rápida,  en  la  cual  trabajan  2 
cocineros  y  3  vendedores,  y  que  debe  atender  a  C  clientes.  El  modelado  debe  considerar 
que: 
- Cada cliente realiza un pedido y luego espera a que se lo entreguen. 
- Los pedidos que hacen los clientes son tomados por cualquiera de los vendedores y se 
lo pasan a los cocineros para que realicen el plato. Cuando no hay pedidos para atender, 
los vendedores aprovechan para reponer un pack de bebidas de la heladera (tardan entre 
1 y 3 minutos para hacer esto). 
- Repetidamente cada cocinero toma un pedido pendiente dejado por los vendedores, lo 
cocina y se lo entrega directamente al cliente correspondiente. 
Nota: maximizar la concurrencia. 

//se puede consultar si es vacio, pero no la cantidad 
chan pedidos(int, Pedido); //simula el linkeo de cada id con su PedidoContenido

chan pedidoListo[C](Pedido); //pedidos listos

chan vendedorLibre(int); //medio de aviso vendedor-coordinador

chan meTocaTarea(int, Pedido); //comunicacion coordinador-vendedor

chan avisoCocinero(int, Pedido);

Process Vendedor(id: 0..3-1){
	int idCliente; Pedido pedidoSolicitado;
	while(true){
		send vendedorLibre(id);
		receive meTocaTarea(idCliente, pedidoSolicitado);
		if(idCliente == -1){ //si es igual a -1 entonces no me fue asignada una tarea
			delay(1..3);
		}else{ //me fue asignada una tarea.. idCliente <> -1
			send avisoCocinero(idCliente, pedidoSolicitado);
		}
	}
}

Process Cocinero(id: 0..2-1){
	int idCliente; 
	Pedido pedidoSolicitado;
	while(true){
		receive avisoCocinero(idCliente, pedidoSolicitado);
		realizandoPlato(pedidoSolicitado); //sarasa
		send pedidoListo[idCliente](pedidoSolicitado);
	}
}

Process Cliente(id: 0..C-1){
	Pedido pedidoRecibido;
	Pedido pedidoSolicitado = generandoQueVoyApedir();
	send pedidos(id, pedidoSolicitado);
	receive pedidoListo[id](pedidoRecibido);
}

Process Coordinador(){
	int idVendedorLibre; 
	int idCliente; 
	while(true){
		Pedido pedidoSolicitado; //si lo pongo arriba, potencialmete pueden pasar comportamientos inesperados en las iteracicione 0+1..N
		receive vendedorLibre(idVendedorLibre); //bloqueante
		//verifico que haya algun pedido, (si estoy aca, ya tengo un vendedor disponible para responder un potencial pedido)
		if(pedidos.isEmpty()){
			idCliente = -1;
		}else{ //tenemos un pedido
			receive pedidos(idCliente, pedidoSolicitado);
		}
		send meTocaTarea(idCliente, pedidoSolicitado);
	}
	
}

expl
process coordinador{                            //EXPLICACION DEL USO DEL COORD: el proceso coordinador hace de pasamanos y regula el comportamiento de los vendedores, esto pasa por que 
	while(true){                                //los vendedores no pueden evaluar de forma correcta que un canal sea empty, ya que puede pasar, que 2 procesos pregunten por
			receive vendedorLibre(idVend);          //un canal que no sea vacio en un if(lo cual es verdad) y al entrar haber un mensaje solo pudiendo suceder que uno de los dos procesos lo   
			if (pedido.empty()){                    //toma y el otro se queda esperando en el receive. Este 2do proceso no hace el delay al quedarse trabado, pero si ademas nunca llegase
					id=-1;                              //otro mensaje el proceso nunca finalizaria. 
			}else{
					receive pedido(id,pedido);
			}
			send vendedor(id,pedido);
	}   
}