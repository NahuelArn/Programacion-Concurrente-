3) Resolver el siguiente problema  con PMA. En  un  negocio de cobros digitales hay P personas que 
deben  pasar  por  la  única  caja  de  cobros  para  realizar  el  pago  de  sus  boletas.  Las  personas  son 
atendidas de acuerdo con el orden de llegada, teniendo prioridad aquellos que deben pagar menos 
de 5 boletas de los que pagan más. Adicionalmente, las personas embarazadas tienen prioridad sobre 
los dos casos anteriores. Las personas entregan sus boletas al cajero y el dinero de pago; el cajero les 
devuelve el vuelto y los recibos de pago.

chan respuestaCajero[P](real,Text);

chan filaEmbarazadas[P](int);
chan filaMenos5Boletas(int);
chan filaMas5Boletas(int);

chan hayPedido(boolean);
Process Persona(id: 1..P){ //embarazada, >5, <5
	boolean embarazada = yoSeSiEstoy(); //?
	boletas [] bo = tengoEstasBoletas();
	real dinero = cuanToTengoQuePagar();
	real vuelto; Queue recibos;
	if(embarazada){
		send filaEmbarazadas(id,bo,dinero);
	}else{
		if(bo.size() < 5){
			send fila.filaMenos5Boletas(id,bo,dinero);
		}else{
			send fila.filaMas5Boletas(id,bo,dinero);
		}
	}
	send hayPedido(true); //algo tengo que ser no? lo pongo aca para ahorrarme ponerlo en cada if else... Si o si, soy algo...
	receive respuestaCajero[id](vuelto,recibos);
}

Process Cajero(){
	integer cantidadDeAtenciones = 0;
	int id; boletas [] bo; real dinero;
	real vuelto; Queue recibos;
	boolean aux;
	while(cantidadDeAtenciones < P){
		receive hayPedido(aux); //para no generar bussy waiting ni contar cantidades de atencion erroneas
		if(!filaEmbarazadas.isEmpty){
			receive filaEmbarazadas(id,bo,dinero);
		}else{
			if(!filaMenos5Boletas.isEmpty()){
				receive filaMenos5Boletas(id,bo,dinero);
			}else{
				if(!filaMas5Boletas.isEmpty()){
					receive filaMas5Boletas(id,bo,dinero);
				}
			}
		}
		cantidadDeAtenciones++;
		(vuelto, recibos) = cajeroLaburando(id,bo,dinero);
		send respuestaCajero[id](vuelto,recibos);
	}
}
