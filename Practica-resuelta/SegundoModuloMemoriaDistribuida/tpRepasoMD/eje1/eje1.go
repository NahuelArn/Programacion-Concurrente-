1) En  una  oficina  existen  100  empleados  que  envían  documentos  para  imprimir  en  5  impresoras 
compartidas. Los pedidos de impresión son procesados por orden de llegada y se asignan a la primera 
impresora que se encuentre libre: 
a) Implemente un programa que permita resolver el problema anterior usando PMA. 
b) Resuelva el mismo problema anterior pero ahora usando PMS.


A- PMA
//Esto tambien soluciona el problema,Seria la mejor solucion para el problema...
//si tuviera que la impresora 1 es la mas rapida "podria tomar siempre todos los pedidos" pero en la solucion con coordinador tambien pasaria lo mismo... si es tan rapida, siempre mandaria mensajes al coordinador antes q todas
chan empleadoPeticion(int, Documento); // Empleados envían doc + su id
chan recibirCopia[100](Documento); // Cada empleado tiene su canal para recibir la copia

Process Impresora(id: 1..5){
	Documento doc;
	int idE;
	while (true){
		receive empleadoPeticion(idE, doc);
		doc = imprimiendo(doc);
		send recibirCopia[idE](doc);
	}

}

Process Empleado(id: 1..100){
	Documento doc;
	while(true){
		send empleadoPeticion(id,doc);
		receive recibirCopia[id](doc);
	}
}

//======================con coordinador"innecesario"=============================

chan impresoraLibre(int);
chan pedidos(id, Documento);

chan respuesta[100](Documento);
chan impresoraTarea[5](int,Documento);

Process Empleado(id: 1..100){
	Documento doc;
	while(true){
		generarDocumentoImpresion(doc);
		send pedidos(id,doc);
		receive respuesta[id](doc);
	}
	
}

Process Impresora(id: 1..5){
	Documento doc;
	While(true){
		send impresoraLibre(id);
		receive impresoraTarea[id](idE, doc);
		doc = realizandoImpresion(doc);
		send respuesta[idE](doc);
	}
	
}

Process Coordinador(){
	int idI, idE;
	Documento doc;
	while (true) {
		receive impresoraLibre(idI);
		receive pedidos(idE,doc);
		send impresoraTarea[iE](doc);
	}
}








b) Resuelva el mismo problema anterior pero ahora usando PMS.
//==================================================================PMS==========================================================================


1) En  una  oficina  existen  100  empleados  que  envían  documentos  para  imprimir  en  5  impresoras 
compartidas. Los pedidos de impresión son procesados por orden de llegada y se asignan a la primera 
impresora que se encuentre libre: 

//En este caso si es necesario usar un coordiador.. En PMS nos encontraos con el problema de espera incensario de los mensajes...
//Empleado al hacer un send es bloqueante y no tiene una Queue implicita.. entonces retrasaria todo el pipeline
Process Empleado(id: 1..100){
	Documento doc;
	while(true){
		doc = generandoLoQueQUieroImprimir();
		Coordinador!estoyListo(id, doc);
		Impresora[*]?copiaLista(doc);
		mirandoLaDoc(doc);
	}
	
}

Process Impresora(id: 1..10){
	Documento doc; idE: integer;
	While(true){
		Coordiandor.printerLista(id);
		Coordinador?esperandoDocumento(doc, idE);
		Documento impreso = imprimiendo(doc);
		Empleado[idE]!copiaLista(impreso);
	}
}

Process Coordiandor { //como te dice que nunca debe terminar... el do od se va hacer infinitamente, por que la primera guarda[] su condicion siempre es true
	int idE, idI;
	Documento doc;
	Queue documentosEsperandoImpresion;
	do
		[] Empleado[*]?estoyListo(id, doc) ->
			documentosEsperandoImpresion.push(doc,idE);
		[] (!documentosEsperandoImpresion.isEmpty()); Impresora[*]?printerLista(idI) ->
			Impresora[idI]!esperandoDocumento(documentosEsperandoImpresion.pop());
	od
}

























































//Cache


//SOLUCION QUE NO RESPETA EL ORDEN DE LLEGADA...
chan impresoraLibre(int); //id de impresora libre

chan pedidoDeImpresion[5](int,Documento); //okey esto es un un vector de queues

chan recibirCopia[100](Documento); //por condicion de carrera necesito asegurarme que cada uno reciba bien, su documento

Process Impresora(id: 1..5){
	Documento doc;
	idE: int;
	while(true){
		send impresoraLibre(id);	
		receive pedidoDeImpresion[id](idE,doc);
		imprimiendo(doc);
		send recibirCopia[idE](doc);
	}
}

Process Empleado(id: 1..100){
	int idI;
	while(true){
		Documento documento = generarDocumentoImpresion();
		receive impresoraLibre(idI); // si los 100 esperan... cualquiera puede tomar la impresora.. no cumplo con el orden de llegada... Por eso usabamos coordinador
		send pedidoDeImpresion[idI](id,documento);
		receive recibirCopia[id](documento);
	}
}

// Process Coordinador{
// 	int idImpresoraLibre;
// 	Documento docActual
// 	while(true){
// 		receive impresoraLibre(idImpresoraLibre); //dequeue
// 		receive pedidos(docActual);
// 		send pedidoDeImpresion[idImpresoraLibre](docActual);
// 	}
}

//necesito un coordinador... Por que? si hago en el empleado impresoras.isEmpty() varios Empleado a la vez preguntan lo mismo... inconsistencia con la informacion a recibir exp ppt








//
chan impresoraLibre(int); //id de impresora libre

chan pedidos(int, Documento); //ideEmpleado y documento que quiere imprimir

chan pedidoDeImpresion[5](Documento); //okey esto es un un vector de queues

chan recibirCopia[100](Documento);

Process Impresora(id: 1..5){
	Documento doc;
	while(true){
		send impresoraLibre(id);	
		receive pedidoImpresion[id](doc);
		imprimiendo(doc);
	}
}

Process Empleado(id: 1..100){
	int idI;
	while(true){
		Documento documento = generarDocumentoImpresion();
		receive impresoraLibre(idI);
		send pedidoDeImpresion[idI](documento);
		send pedidos(documento);
	}
}

Process Coordinador{
	int idImpresoraLibre;
	Documento docActual
	while(true){
		receive impresoraLibre(idImpresoraLibre); //dequeue
		receive pedidos(docActual);
		send pedidoDeImpresion[idImpresoraLibre](docActual);
	}
}


















chan pedidos(int, Documento); //ideEmpleado y documento que quiere imprimir

Process Impresora(id: 1..5){
	Documento doc;
	while(true){	
		receive pedidos(doc);
		imprimiendo(doc);
	}

}

Process Empleado(id: 1..100){
	while(true){
		Documento documento = generarDocumentoImpresion();
		send pedidos(documento);
	}
}



aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa


chan pedidos(string)
chan miRespuesta(string)

process Cliente[id: 0..N-1] {
  string pedido, respuesta
  while true {
    send pedidos(id, pedido)
    receive miRespuesta(respuesta)
  }
}

process Servidor[id: 0..3] {
  int idC
  string pedido, respuesta
  while true {
    receive pedidos(idC, pedido)
    respuesta = Procesar(pedido)
    send miRespuesta(respuesta)
  }
}