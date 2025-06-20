5.  Resolver la administración de las impresoras de una oficina. Hay 3 impresoras, N usuarios y 
1  director.  Los  usuarios  y  el  director  están  continuamente  trabajando  y  cada  tanto  envían 
documentos  a  imprimir.  Cada  impresora,  cuando  está  libre,  toma  un  documento  y  lo 
imprime,  de  acuerdo  con  el  orden  de  llegada,  pero  siempre  dando  prioridad  a  los  pedidos 
del  director.  Nota:  los  usuarios  y  el  director  no  deben  esperar  a  que  se  imprima  el 
documento.

//clean.. abajo explicado
chan impresoraLibre(int);

chan colaDeImpresionDirector(Text);
chan colaDeImpresionUsuario(Text);

Process Impresora(id: 0..3-1){
	Text hoja;
	while(true){
		send impresoraLibre(id); //mando que estoy libre
		receive hayLaburo[id](hoja); //los usuarios y director.. no esperan.. entonces Impresora no les avisa
		imprimiendo(hoja);
	}
}

Process Usuario(id: 0..N-1){
	while (true){
		Text doc = generandoQueQuieroImprimir();
		send colaDeImpresionUsuario(doc);
		//nada
	}
}

Process Director(){
	while(true){
		Text doc = generandoQueQuieroImprimir();
		send colaDeImpresionDirector(doc);
		//nada
	}
}

//opcion 1
Process Coordinador(){
	while(true){
		int idImpresoraLibre
		receive impresoraLibre(idImpresoraLibre); //espero a que haya alguna impresora libre, si hay la quito como libre
		Text documento;
		boolean hayPeticion = false;
		if(!colaDeImpresionDirector.isEmpty() || !colaDeImpresionUsuario.isEmpty()){hayPeticion = true}
		if(!colaDeImpresionDirector.isEmpty()){ 
			receive colaDeImpresionDirector(documento);
		}else if(!colaDeImpresionUsuario.isEmpty()) {
			receive colaDeImpresionUsuario(documento);
		}
		if(hayPeticion){
			send hayLaburo[idImpresoraLibre](documento);
		}	
	} 
}	

//opcion2
Process Coordinador(){
	while(true){
		int idImpresoraLibre
		receive impresoraLibre(idImpresoraLibre); //espero a que haya alguna impresora libre, si hay la quito como libre
		Text documento;
		if(!colaDeImpresionDirector.isEmpty()) { 
			receive colaDeImpresionDirector(documento);
			send hayLaburo[idImpresoraLibre](documento);
		} else if(!colaDeImpresionUsuario.isEmpty()) {
			receive colaDeImpresionUsuario(documento);
			send hayLaburo[idImpresoraLibre](documento);
		}
	}
}



//========================================

//explicaciones



//en las colasDeImpresionX no nos importa el idDelUsuario/director nos importa el documento...
// puede quedar colaImpresionDirector(idDirector, Documento) = igual para usuario
// o puede quedar solo = colaImpresionDirector(Documento);
//tambien tenia un error de logica, los ids y documentos/parametros, tipos
chan impresoraLibre(int);


chan colaDeImpresionDirector(id);
chan colaDeImpresionUsuario(id);

Process Impresora(id: 0..3-1){
	Text hoja;
	while(true){
		send impresoraLibre(id); //mando que estoy libre
		receive hayLaburo[id](hoja); //los usuarios y director.. no esperan.. entonces Impresora no les avisa
		// if(!hoja.equals("noHayLaburo")){ //si hay laburo, lo imprimo.. en este caso como no se requiere que el coordinador le responda si no hay laburo no modelo esto
		// 	imprimiendo(hoja);
		// }
		imprimiendo(hoja);
	}
	//como no tengo que avisar que termine de de imprimir, no hago nada despues de imprimir
}

Process Usuario(id: 0..N-1){
	while (true){
		send colaDeImpresionUsuario(id);
		//nada
	}
}

Process Director(){
	while(true){
		send colaDeImpresionDirector(id);
		//nada
	}
}


//sarasav4
//v3 REDUNDANCIA, con esta version no tengo bussy waiting en el coordinador "cuando no tenga impresora libre"
//pasaba que al no tener impresoras libres, siempre preguntaba !if(!impresora.isEmpty()){logica}..

//necesito un coordinador, no puede resolver por si solo impresora IsEmpty
//aparece el problema de empty, pero diferente... aca no necesito la respuesta del coordinador, "cuando no hay laburo"

Process Coordinador(){
	while(true){
		int idImpresoraLibre
		receive impresoraLibre(idImpresoraLibre); //espero a alguna impresora libre, si hay la quito como libre
		Text documento;
		boolean hayPeticion = false;

		if(!colaDeImpresionDirector.isEmpty() || !colaDeImpresionUsuario.isEmpty()){hayPeticion = true}

		if(!colaDeImpresionDirector.isEmpty()){ //si hay un pedido de un director y una impresora libre
			receive colaDeImpresionDirector(documento);
		
		}else if(!colaDeImpresionUsuario.isEmpty()) {
			receive colaDeImpresionUsuario(documento);
			
		}
		if(hayPeticion){
			send hayLaburo[idImpresoraLibre](documento);
		}
		
	} 
}



//sarasav3
//v3 REDUNDANCIA, con esta version no tengo bussy waiting en el coordinador "cuando no tenga impresora libre"
//pasaba que al no tener impresoras libres, siempre preguntaba !if(!impresora.isEmpty()){logica}..

//necesito un coordinador, no puede resolver por si solo impresora IsEmpty
//aparece el problema de empty, pero diferente... aca no necesito la respuesta del coordinador, "cuando no hay laburo"

Process Coordinador(){
	while(true){
		int idImpresoraLibre
		receive impresoraLibre(idImpresoraLibre); //espero a alguna impresora libre, si hay la quito como libre
		Text documento;
		if(!colaDeImpresionDirector.isEmpty()){ //si hay un pedido de un director y una impresora libre
			receive colaDeImpresionDirector(documento);
			send hayLaburo[idImpresoraLibre](documento); //x1
		}else if(!colaDeImpresionUsuario.isEmpty()) {
			receive colaDeImpresionUsuario(documento);
			send hayLaburo[idImpresoraLibre](documento); //x1
		}
		//x1 si bien, esto es repetido, no puedo ponerlo en esta linea... me puede pasar que entro en el if o else... y no tengo forma de saber si 
		//hice un receive... esto no va funcionar if(!colaDeImpresionDirector.isEmpty || !colaDeImpresionUsuario.isEmpty) ...perooo se puede agregar un boolean
	} 
}




//=================================================================

//Esta bien, pero puede mejorarse (arriba mejor version)
chan impresoraLibre(int);


chan colaDeImpresionDirector(id);
chan colaDeImpresionUsuario(id);

Process Impresora(id: 0..3-1){
	Text hoja;
	while(true){
		send impresoraLibre(id); //mando que estoy libre
		receive hayLaburo[id](hoja); //los usuarios y director.. no esperan.. entonces Impresora no les avisa
		// if(!hoja.equals("noHayLaburo")){ //si hay laburo, lo imprimo.. en este caso como no se requiere que el coordinador le responda si no hay laburo no modelo esto
		// 	imprimiendo(hoja);
		// }
		imprimiendo(hoja);
	}
	//como no tengo que avisar que termine de de imprimir, no hago nada despues de imprimir
}

Process Usuario(id: 0..N-1){
	while (true){
		send colaDeImpresionUsuario(id);
		//nada
	}
}

Process Director(){
	while(true){
		send colaDeImpresionDirector(id);
		//nada
	}
}


//v3 REDUNDANCIA, con esta version no tengo bussy waiting en el coordinador "cuando no tenga impresora libre"
//pasaba que al no tener impresoras libres, siempre preguntaba !if(!impresora.isEmpty()){logica}..

//necesito un coordinador, no puede resolver por si solo impresora IsEmpty
//aparece el problema de empty, pero diferente... aca no necesito la respuesta del coordinador, "cuando no hay laburo"

Process Coordinador(){
	while(true){
		int idImpresoraLibre
		receive impresoraLibre(idImpresoraLibre); //espero a alguna impresora libre, si hay la quito como libre
		if(!colaDeImpresionDirector.isEmpty()){ //si hay un pedido de un director y una impresora libre
			int idDirector; 
			receive colaDeImpresionDirector(idDirector);
			send hayLaburo[idDirector](idImpresoraLibre);
		}else if(!colaDeImpresionUsuario.isEmpty()) {
			int idUsuario; 
			receive colaDeImpresionUsuario(idUsuario);
			send hayLaburo[idUsuario](idImpresoraLibre); 
		}
	} 
}


//======================================================================================================




//v2
//necesito un coordinador, no puede resolver por si solo impresora IsEmpty
//aparece el problema de empty, pero diferente... aca no necesito la respuesta del coordinador, "cuando no hay laburo"

Process Coordinador(){
	while(true){
		int idImpresoraLibre
		if(!impresora.isEmpty()){
			receive impresoraLibre(idImpresoraLibre); //espero a alguna impresora libre, si hay la quito como libre
			if(!colaDeImpresionDirector.isEmpty()){ //si hay un pedido de un director y una impresora libre
				int idDirector; 
				receive colaDeImpresionDirector(idDirector);
				send hayLaburo[idDirector](idImpresoraLibre);
			}else if(!colaDeImpresionUsuario.isEmpty()) {
				int idUsuario; 
				receive colaDeImpresionUsuario(idUsuario);
				send hayLaburo[idUsuario](idImpresoraLibre); 
			}
		}
	} 
}




//v1

//en este caso no es necesario "noHayLaburo" (creo) ya que me sirve que se queden dormidos en el receive las impresoras

//pregunta esto
por si solo esto ya resuelve con lo pedido, necesario agregar la logica en este caso "noHayLaburo"
Process Coordinador(){
	while(true){
		int idImpresoraLibre
		if(!colaDeImpresionDirector.isEmpty() && !impresora.isEmpty()){ //si hay un pedido de un director y una impresora libre
			int idDirector; 
			receive colaDeImpresionDirector(idDirector);
			receive impresoraLibre(idImpresoraLibre);
			send hayLaburo[idDirector](idImpresoraLibre);
		}
		if(!colaDeImpresionUsuario.isEmpty() && !impresora.isEmpty()){
			int idUsuario; 
			receive colaDeImpresionUsuario(idUsuario);
			receive impresoraLibre(idImpresoraLibre);
			send hayLaburo[idUsuario](idImpresoraLibre); 
		}
	} 
}