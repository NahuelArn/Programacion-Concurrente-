5.  En  una  clínica  existe  un  médico  de  guardia  que  recibe  continuamente  peticiones  de 
atención de las E  enfermeras que trabajan en su piso y de las  P  personas que llegan a la 
clínica ser atendidos

Cuando una persona necesita que la atiendan espera a lo sumo 5 minutos a que el médico lo 
haga, si pasado ese tiempo no lo hace, espera 10 minutos y vuelve a requerir la atención del 
médico. Si no es atendida tres veces, se enoja y se retira de la clínica. 

Cuando una enfermera requiere la atención del médico, si este no lo atiende inmediatamente 
le  hace  una  nota  y  se  la  deja  en  el  consultorio  para  que  esta  resuelva  su  pedido  en  el 
momento  que  pueda  (el  pedido  puede  ser  que  el  médico  le  firme  algún  papel).  Cuando  la 
petición  ha  sido  recibida  por  el  médico  o  la  nota  ha  sido  dejada  en  el  escritorio,  continúa 
trabajando y haciendo más peticiones. 
El médico atiende los pedidos dándole prioridad a los enfermos que llegan para ser atendidos. 
Cuando atiende un pedido, recibe la solicitud y la procesa durante un cierto tiempo. Cuando 
está libre aprovecha a procesar las notas dejadas por las enfermeras. 



program eje5 is

Task Consultorio is
  Entry dejarNota(n: in Pedido);
  Entry retirarNota(n: out Pedido);
end Consultorio;

Task body Consultorio is
  Queue notas;
begin
  loop
    select 
      Accept dejarNota(n: in Pedido) do
        notas.push(n);
      end dejarNota;
    OR
      Accept retirarNota(o: out Pedido)do
        if(notas.esEmpty()){
          o = null;
        else
          o = notas.pop();
        end if;
      end retirarNota;
  end loop;
end Consultorio;

Task Medico is
  Entry atenderPeticionEnfermera(pedidoEnf: in Pedido);
  Entry atenderPeticionPersona; //el when'count=0... solo se puede hacer si el Entry pertenece al proceso.. Ejemplo aca "atenderPeticionPersona" es propiedad/Entry de Medico... si fuera de otro proceso no puedo consultar  
end Medico;

Task body Medico is

begin
  loop
    select
      Accept atenderPeticionPersona do
        //atiende persona 
      atenderPeticionPersona end;
    OR
      When (atenderPeticionPersona'count= 0) => // si atenderPeticionPersona no tiene peticiones entra aca "que serian atender a las enfermeras que recientemente solicitaron un pedo(inmediatamente)" //es muy fino este caso.. justo se tiene que dar que matcheen/esten listos al mismo tiempo enfermera/doctor.. sino ambos no entran aca.. enfermera deja una nota.. y doctor se pone a ver pacientes o resolver notas
        //analogia.. un doctor entra a su despacho y ni bien cierra la puerta, entra una enfermera solicitando un pedido.. justo ahi entraria en este bloque...
        //si la enfermera entra y no esta el doctor, deja una nota y no entraria aca
        Accept atenderPeticionEnfermera(ped) do
          procesarPeticionEnfermera(ped); //fn
        end atenderPeticionEnfermera;
    else //aca entro si, no hay enfermeras recientes esperando en el atenderPeticionEnfermera... "no hay enfermeras para atender de forma inmediata" "Aca estan las notas que dejaron las enfermeras que no atendio inmediatamente"
      Select 
        Consultorio.retirarNota(ped); //si lo dejo asi... y no tengo nota.. se queda esperando aca...
        if(ped != null)then
          resolverNota(ped);
        end if;
      ELSE  //si la otra tarea no está lista para realizar el accept a su pedido inmediatamente, entonces lo cancela (lo quita de la cola implícita) y realiza otra cosa..... Si no tuviera el select anidado... con else null... entonce el Medico se quedaria demorado hasta que alguien realice su accept retirarNota.. En el caso de que Consultorio, el proceso no gana el procesador por mas tiempo de lo esperado... quedo demorado aca, se provoca deadlock
        null;
      end select;
    end Select;
    
  end loop;
end;

Task type Enfermera;
vecEnf: array(1..E) of Enfermera; //aviso que tengo E procesos enfermera 

Task body Enfermera is
begin
  loop
    select 
      Pedido pedido = generarPedido();
      Medico.atenderPeticionEnfermera(pedido);
    else //el else se activa si no se acepta inmediatamente lo de arriba... "si este no lo atiende inmediatamente, te da el pie para meter un else"
      Pedido pedido = generarNota();
      Consultorio.dejarNota(pedido);
    end select;
  end loop;
end Enfermera;


"Cuando una persona necesita que la atiendan espera a lo sumo 5 minutos a que el médico lo 
haga, si pasado ese tiempo no lo hace, espera 10 minutos y vuelve a requerir la atención del 
médico. Si no es atendida tres veces, se enoja y se retira de la clínica. "

Task Type Persona;

vecPer: array(1..P) of Persona;

Task body Persona is
  continuar: boolean;
  intentos: int;
begin
  loop
    continuar = false; intentos = 0;
    loop (!continuar AND intentos < 3)
      select
        Medico.atenderPeticionPersona; //requiero atencion
        continuar = false; //ya me atendio corto
      or delay 60*5 //espero 5  minutos
        intentos++;
        if(intentos < 3)then
          delay(60*10); //espero 10 minutos
        end if;
      end select; //solo se permite 1 select no se puede anidar selects
    end loop;
  loop end;
end Persona;

begin
  null;
end eje5;




//================
program eje5 is

Task Medico is
  Entry atenderPeticionEnfermera(pedidoEnf: in Pedido);
  Entry atenderPeticionPersona; //el when'count=0... solo se puede hacer si el Entry pertenece al proceso.. Ejemplo aca "atenderPeticionPersona" es propiedad/Entry de Medico... si fuera de otro proceso no puedo consultar  
end Medico;

Task body Medico is

begin
  loop
    select
      Accept atenderPeticionPersona do
        //atiende persona 
      atenderPeticionPersona end;
    OR
      When (atenderPeticionPersona'count= 0) => // si atenderPeticionPersona no tiene peticiones entra aca "que serian atender a las enfermeras que recientemente solicitaron un pedo(inmediatamente)" //es muy fino este caso.. justo se tiene que dar que matcheen/esten listos al mismo tiempo enfermera/doctor.. sino ambos no entran aca.. enfermera deja una nota.. y doctor se pone a ver pacientes o resolver notas
        //analogia.. un doctor entra a su despacho y ni bien cierra la puerta, entra una enfermera solicitando un pedido.. justo ahi entraria en este bloque...
        //si la enfermera entra y no esta el doctor, deja una nota y no entraria aca
        Accept atenderPeticionEnfermera(ped) do
          procesarPeticionEnfermera(ped); //fn
        end atenderPeticionEnfermera;
    else //aca entro si, no hay enfermeras recientes esperando en el atenderPeticionEnfermera... "no hay enfermeras para atender de forma inmediata" "Aca estan las notas que dejaron las enfermeras que no atendio inmediatamente"
      Consultorio.nota(ped); //si lo dejo asi... y no tengo nota.. se queda esperando aca...
      if(ped != null)then
        resolverNota(ped);
      end if;
    end Select;
    
  end loop;
end;

Task type Enfermera;
vecEnf: array(1..E) of Enfermera; //aviso que tengo E procesos enfermera 

Task body Enfermera is
begin
  loop
    select 
      Pedido pedido = generarPedido();
      Medico.atenderPeticionEnfermera(pedido);
    else //el else se activa si no se acepta inmediatamente lo de arriba... "si este no lo atiende inmediatamente, te da el pie para meter un else"
      Pedido pedido = generarNota();
      Medico.atenderPeticionEnfermera(pedido);
    end select;
  end loop;
end Enfermera;


"Cuando una persona necesita que la atiendan espera a lo sumo 5 minutos a que el médico lo 
haga, si pasado ese tiempo no lo hace, espera 10 minutos y vuelve a requerir la atención del 
médico. Si no es atendida tres veces, se enoja y se retira de la clínica. "

Task Type Persona;

vecPer: array(1..P) of Persona;

Task body Persona is
  continuar: boolean;
  intentos: int;
begin
  loop
    continuar = false; intentos = 0;
    loop (!continuar AND intentos < 3)
      select
        Medico.atenderPeticionPersona; //requiero atencion
        continuar = false; //ya me atendio corto
      or delay 60*5 //espero 5  minutos
        intentos++;
        if(intentos < 3)then
          delay(60*10); //espero 10 minutos
        end if;
      end select; //solo se permite 1 select no se puede anidar selects
    end loop;
  loop end;
end Persona;

begin
  null;
end eje5;