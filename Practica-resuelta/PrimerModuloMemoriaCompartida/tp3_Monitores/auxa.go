Monitor corralon{
    cond enEspera; cond empleado; cond hayLista; cond hayComprobante; cond retiroComprobante;
    int clienteEsperando=0;
    listaMaterialesActual; comprobanteActual;
    boolean empleadoLibre = true;
    
    procedure clienteSePresenta(listaMateriales: in lista){
        clienteEsperando ++;                                            //si el empleado esta trabajando, lo espero hasta mi turno . 
        signal(empleado);
        wait(enEspera);
        listaMaterialesActual = listaMateriales;                            //la lista de materiales actual= lista de clietne. 
        signal(hayLista);                                                   //informo que hay lista
    }

    procedure atenderCliente(){                                             //vendedor si en algun momento no tiene cola de espera, se duerme  e informa que esta libre. 
        if(clienteEsperando==0){
            wait(empleado);
        }
        clienteEsperando--;                                                 //al ser despertado es por que llego un cliente nuevo.
        signal(enEspera);
        wait(hayLista);                                                     //me quedo esperando a que el cliente me de su lista de materiales.
        comprobanteActual = generarComprobante(listaMaterialesActual);
        signal(hayComprobante);                                             //informo al cliente que ya esta su tickey (para que se vaya)
        wait(retiroComprobante);                                            // me quedo esperando que el cliente retire su ticket
    }
    procedure retiroComprobante(comprobante:OUT Comprobante){
        wait(hayComprobante);                                               //espera al informe que esta el comprobante.
        comprobante = comprobanteActual;                                    //lo extraigo
        signal(retiroComprobante);                                          //lo saco y me retiro
    }
}

process Cliente [1..n]{
    Lista listaMateriales; comprobante;
    corralon.clienteSePresenta(listaMateriales);            //cliente llega con la lista de listaMateriales
    corralon.clienteRetirarComprobante(comprobante);        //cliente se retira con el ticket de la compra
}

process empleado{
    for (int i=1; i<N; i++){
        corralon.atenderCliente();                          //empleado atiende a los N clientes.
    }
}