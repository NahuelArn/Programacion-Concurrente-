5.  En un estadio de fútbol hay una máquina expendedora de gaseosas que debe ser usada por 
E Espectadores de acuerdo al orden de llegada. Cuando el espectador accede a la máquina 
en su turno usa la máquina y luego se retira para dejar al siguiente.  Nota: cada Espectador 
una sólo una vez la máquina. 

process Administrador {
	cola pedidos
	bool ocupado = false
	int actual = -1
	do 
		[] Espectador[*]?pedir(idP) -> 
			if (ocupado){
				pedidos.push(idP) ;
			} else {
				Persona[idP]!recibir();
				actual = idP;
				ocupado = true;
			}
		[] Espectador[actual]?salir(); -> 
			if (pedidos.empty()) {
				ocupado = false;
				actual = -1;
			} else {
				actual = pedidos.pop();
				Persona[actual]!recibir();
			}
	od
}
process Espectador[id:0..E-1]{
	Administrador!pedir(id);
	Administrador?recibir();
	// usando maquina();
	Administrador!salir();
}