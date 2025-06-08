1.  Suponga  que  existe  un  antivirus  distribuido que  se  compone  de  R  procesos  robots 
Examinadores y 1 proceso Analizador. Los procesos Examinadores están buscando 
continuamente  posibles  sitios  web  infectados;  cada  vez  que  encuentran  uno  avisan  la 
dirección y luego continúan buscando. El proceso Analizador se encarga de hacer todas las 
pruebas necesarias con cada uno de los sitios encontrados por los robots para determinar si 
están o no infectados.  
a. Analice el problema y defina qué procesos, recursos y comunicaciones serán 
necesarios/convenientes para resolver el problema. 
b. Implemente una solución con PMS sin tener en cuenta el orden de los pedidos. 
c. Modifique el inciso (b) para que el Analizador resuelva los pedidos en el orden en que 
se hicieron.

//A
?

//B
Process Examinador(id: 0..R-1){
	//buscando web potencialmete infectada
	String url = buscandoWebInfectada();

	Administrador!aviso(url);
}

Process Analizador(){
	while(true){
		Administrador!estoyListoParaAnalizar();
		String url;
		Administrador?confirmacion(url);

		analizandoWeb(url);
	}
}

Process Administrador(){
	String url;
	Queue sitiosARevisar;
	//el do od, "ya tiene incluido un while con una logica particular de corte "LogicaGuarda3CuandoTermina.png"
	do
		[] Examinador?aviso(url) -> sitiosARevisar.push(url);
		[] (!sitiosARevisar.isEmpty()); Analizador?estoyListoParaAnalizar() -> Analizador!confirmacion(sitiosARevisar.popRandom());
	od
}
//necesito un administrador para resolver el problema de perdida de tiempo
//Una vez que 1 o N Examinadores encuentran el potencial sitio infectado, envian la url y tienen que seguir buscando... no tienen que quedar "frenados"


//========================================================================================================================

//C

Process Examinador(id: 0..R-1){
	//buscando web potencialmete infectada
	String url = buscandoWebInfectada();

	Administrador!aviso(url);
}

Process Analizador(){
	Administrador!estoyListoParaAnalizar();
	Administrador?confirmacion();
}

Process Administrador(){
	String url;
	Queue sitiosARevisar;
	
	do
		[] Examinador?aviso(url) -> sitiosARevisar.push(url);
		[] (!sitiosARevisar.isEmpty()); Analizador?estoyListoParaAnalizar() -> Analizador!confirmacion(sitiosARevisar.pop());
	od
}
//necesito un administrador para resolver el problema de perdida de tiempo
//Una vez que 1 o N Examinadores encuentran el potencial sitio infectado, envian la url y tienen que seguir buscando... no tienen que quedar "frenados"
