2.  Se  desea  modelar  el  funcionamiento  de  un  banco  en  el  cual  existen  5  cajas  para  realizar 
pagos.  Existen  P  clientes que  desean  hacer  un  pago.  Para  esto,  cada  una  selecciona  la  caja 
donde hay menos personas esperando; una vez seleccionada, espera a ser atendido. En cada 
caja, los clientes son atendidos por orden de llegada por los cajeros. Luego del pago, se les 
entrega un comprobante. Nota: maximizando la concurrencia. 
 

chan cajas[5];
chan esperandoAsignacionDeCaja(int);
chan hayPedido(boolean);
chan cajaAsignada[P](int)

chan cajas[5](int);
chan comprobante[p](text);

send avisoTermine(int); //me va tirando la data de que Queue se estan liberando


Process Cliente[id: 0..P-1]{
	text comprob; int numCajaAsignada;
	//Pido caja
	send esperandoAsignacionDeCaja(id); //no bloqueante
	send hayPedido(true); //no bloqueante
	receive cajaAsignada[id](numCajaAsignada);
	
	usandoLaCaja(); //que buena caja
	
	//Termine de usar la caja, 
	send cajas[id](numCajaAsignada);
	receive comprobante[id](comprob); //espero que asignen mi comprobante

	//aviso que termine //esto lo utilizo para formar un sistema de prioridad "la caja donde hay menos personas esperando"
	send avisoTermine(numCajaAsignada);
	send hayPedido(true); //se envia senhal, se activa por 2 cosas, pedido esperando caja, pedido termine de usar la caja
}

Process Coordinador(){
	int siguiente; boolean pedido; int cantPorCaja[5] =([5], 0) //1 solo coordinador//
	int cajaActual;
	while (true){
		receive hayPedido(pedido);
		receive esperandoAsignacionDeCaja(siguiente);
		//ahora antes de asignarle una caja verifico si hay alguna caja que se libero
		if(!empty(avisoTermine)){ //si hay gente que termino
			receive hayPedido(cajaActual);
			cantPorCaja[cajaActual] --;
			//si hay gente esperando caja
		}else if (!empty(esperandoAsignacionDeCaja)){
			receive hayPedido(pedido);
			cajaActual = minCaja(cantPorCaja) //una funcion q me devuelve la nro de caja con la cant minima
			send cajaAsignada[pedido](cajaActual)
			cantPorCaja[cajaActual] ++;
		}
	}
}

Process Caja(id: 0..5-1){
	while(true){
		// Recibo los siguientes, los atiendo y les pongo el comprobante donde van
		int idSiguiente
		receive cajas[id](idSiguiente)
		comprobante = pago(idSiguiente)
		send comprobante[idSiguiente](comprobante)
	}
}