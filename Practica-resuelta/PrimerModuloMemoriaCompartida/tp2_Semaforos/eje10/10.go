// 10. A una cerealera van T camiones a descargarse trigo y M camiones a descargar maíz. Sólo 
// hay lugar para que 7 camiones a la vez descarguen, pero no pueden ser más de 5 del mismo 
// tipo de cereal. Nota: no usar un proceso extra que actué como coordinador, resolverlo 
// entre los camiones




//
Esta bien esta solucion, no hay deadlock pero puede llegar a aparecer starvation pero es algo que no puedo controlar, 
lo controla la ta tecnica de scheduling del sistema operativo, por lo que no es un problema de la solucion.

//causa deadlock
sem lugaresDescarga = 7;
sem trigoSem = 5;
sem maizSem = 5;

process trigo[id: 0..T-1]{
	Trigo trigo;
	P(trigoSem);
	P(lugaresDescarga);
	//Descargando trigo
	DescargandoTrigo(trigo);
	V(lugaresDescarga);
	V(trigoSem);
}

process maiz[id: 0..M-1]{
	Maiz maiz;
	P(maizSem);
	P(lugaresDescarga);
	//Descargando maiz
	DescargandoTrigo(maiz);
	V(lugaresDescarga);
	V(maizSem);
}





Aca hay una solucion que contempla starvation




sem mutex = 1;
sem puedeEntrarTrigo = 0;
sem puedeEntrarMaiz = 0;

int trigoEsperando = 0;
int maizEsperando = 0;
int cantTrigo = 0;
int cantMaiz = 0;
int lugaresOcupados = 0;

process trigo[id: 0..T-1] {
    P(mutex);
    if (lugaresOcupados < 7 && cantTrigo < 5) {
        lugaresOcupados++;
        cantTrigo++;
        V(mutex);
    } else {
        trigoEsperando++;
        V(mutex);
        P(puedeEntrarTrigo);
    }

    // --- Descarga ---
    DescargandoTrigo();

    P(mutex);
    lugaresOcupados--;
    cantTrigo--;
    
    // Despierto a alguien si corresponde
    if (trigoEsperando > 0 && lugaresOcupados < 7 && cantTrigo < 5) {
        trigoEsperando--;
        lugaresOcupados++;
        cantTrigo++;
        V(puedeEntrarTrigo);
    } else if (maizEsperando > 0 && lugaresOcupados < 7 && cantMaiz < 5) {
        maizEsperando--;
        lugaresOcupados++;
        cantMaiz++;
        V(puedeEntrarMaiz);
    }
    V(mutex);
}



//repetir para maiz