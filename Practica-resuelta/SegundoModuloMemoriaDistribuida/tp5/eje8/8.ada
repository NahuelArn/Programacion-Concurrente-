8.  Hay un sistema de reconocimiento de huellas dactilares de la policía que tiene 8 Servidores 
para realizar el reconocimiento, cada uno de ellos trabajando con una Base de Datos propia; 
a su vez hay un Especialista que utiliza indefinidamente. El sistema funciona de la siguiente 
manera: el Especialista toma una imagen de una huella (TEST) y se la envía a los servidores 
para que cada uno de ellos le devuelva el código y el valor de similitud de la huella que más 
se  asemeja  a  TEST  en  su  BD;  al  final  del  procesamiento,  el  especialista  debe  conocer  el 
código  de  la  huella  con  mayor  valor  de  similitud  entre  las  devueltas  por  los  8  servidores. 
Cuando  ha  terminado  de  procesar  una  huella  comienza  nuevamente  todo  el  ciclo.  Nota: 
suponga  que  existe  una  función  Buscar(test,  código,  valor)  que  utiliza  cada  Servidor  donde 
recibe  como  parámetro  de  entrada  la  huella  test,  y  devuelve  como  parámetros  de  salida  el 
código  y  el  valor  de  similitud  de  la  huella  más  parecida  a  test  en  la  BD  correspondiente. 
Maximizar la concurrencia y no generar demora innecesaria. 


... no se implementa el proceso BaseDeDatos ya que 
"Nota: 
suponga  que  existe  una  función  Buscar(test,  código,  valor)  que  utiliza  cada  Servidor"
//es una funcion propia de cada servidor.. no es un proceso... no gano mas concurrencia poniendo un proceso BaseDeDatos...

{Servidores, Especialista,}
Procedure 8 is

vecServers: array(1..8) of Servidor;

Task type Servidor is
  Entry recibirId(id: in integer);
end Servidor;

Task body Servidor is
  codigo: integer; valor: real;
  testActual : Text;
begin
  loop
    Especialista.recibirHuella(testActual); //estoy listo, dame un test
    Buscar(testActual,codigo,valor);
    Especialista.recibirResultado(codigo, valor);
    Especialista.termine; //tengo que pararlos.. ya que me puede contar dentro de los 16 q espero
  end loop;
end Servidor;

//====


Task Especialista is
  Entry recibirResultado(cod: in, vSimilitud: in real);
  Entry recibirHuella(test: out Text);
  Entry termine;
end Especialista;

Task body Especialista is
  huellaTest: Text;
  co: integer; vS: real;
  Queue resultados;
  codMaxSimilitud: integer; valorMaxSimilitud: real;
begin
  loop 
    huellaTest = tomandoHuella();
    valorMaxSimilitud:= -1;
    for i in 1..16 loop
      Select
        Accept recibirHuella(test: out Text)do //bidireccional
          test = huellaTest;
        end recibirHuella;
      OR 
        Accept recibirResultado(cod: in integer, vSimilitud: in real) do
          if(vSimilitud > valorMaxSimilitud)then
            valorMaxSimilitud = vSimilitud;
            codMaxSimilitud = cod;
          end if;
        end recibirResultado;
    end loop;
    put_line("valor maxDeSimilitud "+ valorMaxSimilitud + "con codigo: "+codMaxSimilitud);
    //este es el codigo con maximo valor de similitud valorMaxSimilitud;

    for i in 1..8 loop //aca le indico a los servidores que pueden pedir una nueva huella... sin esto puedo llegar "se da el caso que 1 servidor solo responde y resuelve... entonces va completar las 16 iteraciones del for ese mismo servidor con informacion repetida.. ya que cada server tiene su propia db"
      accept termine;
    end loop;
  end loop;
end Especialista;

begin
  null;
end 8;



///aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

//v2
Procedure 8 is

vecServers: array(1..8) of Servidor;

Task type Servidor is
  Entry recibirId(id: in integer);
end Servidor;

Task body Servidor is
  codigo: integer; valor: real;
  testActual : Text;
begin
  loop
    Especialista.recibirHuella(testActual); //estoy listo, dame un test
    Buscar(testActual,codigo,valor);
    Especialista.recibirResultado(codigo, valor);

  end loop;
end Servidor;

//====


Task type Especialista is
  Entry recibirResultado(cod: in, vSimilitud: in real);
  Entry recibirHuella(test: out Text);
end Especialista;

Task body Especialista is
  huellaTest: Text;
  co: integer; vS: real;
  Queue resultados;
  codMaxSimilitud: integer; valorMaxSimilitud: real;
begin
  loop 
    huellaTest = tomandoHuella();
    for i in 1..16 loop
      Select
        Accept recibirHuella(itest: out Text)do //bidireccional
          test =huellaTest;
        end recibirHuella;
      OR 
        Accept recibirResultado(cod: in integer, vSimilitud: in real) do
          co:= cod;
          vS:= vSimilitud;
          resultados.push(co,Vs);
        end recibirResultado;
    end loop;
    codMaxSimilitud, valorMaxSimilitud = resultados.pop();
    for i in 1..7 loop
      int auxC;
      int auxS;
      auxC, auxS = resultado.pop();
      if(auxS > valorMaxSimilitud)then
        codMaxSimilitud = auxC;
        valorMaxSimilitud = auxS;
      end if;
    end loop;

    //este es el codigo con maximo valor de similitud valorMaxSimilitud;
  end loop;
end Especialista;

begin
  null;
end 8;




//v1
Procedure 8 is

vecServers: array(1..8) of Servidor;

Task type Servidor is
  Entry recibirId(id: in integer);
end Servidor;

Task body Servidor is
  codigo: integer; valor: real;
  testActual : Text;
begin
  loop
    Especialista.recibirHuella(testActual); //estoy listo, dame un test
    Buscar(testActual,codigo,valor);
    Especialista.recibirResultado(codigo, valor);

  end loop;
end Servidor;

//====


Task type Especialista is
  Entry recibirResultado(cod: in, vSimilitud: in real);
  Entry recibirHuella(test: out Text);
end Especialista;

Task body Especialista is
  huellaTest: Text;
  co: integer; vS: real;
  Queue resultados;
  codMaxSimilitud: integer; valorMaxSimilitud: real;
begin
  loop 
    huellaTest = tomandoHuella();
    for i in 1..8 loop
      Accept recibirHuella(itest: out Text)do //bidireccional
        test =huellaTest;
      end recibirHuella;
    end loop;
    for i in 1..8 loop
      Accept recibirResultado(cod: in integer, vSimilitud: in real) do
        co:= cod;
        vS:= vSimilitud;
        resultados.push(co,Vs);
      end recibirResultado;
    end loop;
    codMaxSimilitud, valorMaxSimilitud = resultados.pop();
    for i in 1..7 loop
      int auxC;
      int auxS;
      auxC, auxS = resultado.pop();
      if(auxS > valorMaxSimilitud)then
        codMaxSimilitud = auxC;
        valorMaxSimilitud = auxS;
      end if;
    end loop;

    //este es el codigo con maximo valor de similitud valorMaxSimilitud;
  end loop;
end Especialista;

begin
  null;
end 8;








//========================================================
Procedure 8 is

vecServers: array(1..8) of Servidor;

Task type Servidor is
  Entry recibirId(id: in integer);
end Servidor;

Task body Servidor is
  id: integer;
  codigo: integer; valor: real;
  testActual : Text;
begin
  Accept recibirId(pos: in integer) do --el servidor no necesita saber si id... logica inncesario... lo que me interesa es saber el codigo de la huella que tiene mas % de compatibilidad
    id :=  pos;
  end recibirId;
  loop
    -- Accept recibirHuella(test: in Text) do .... Si hago esto, atraso cada proceso.. cada llamdda esperia que cada Servidor busque en su db, y hasta que no termine uno, no busque el otro...
    --   Buscar(test,codigo,valor); ... malll, retraso todo... puede tardar 2 hs, en buscar en su db y hasta que no finalicen esas 2 horas los demes servers no empezaron a buscar
    -- end recibirHuella;
    Especialista.recibirHuella(id, testActual); //estoy listo, dame un test
    Buscar(testActual,codigo,valor);
    Especialista.recibirResultado(codigo, valor);
end Servidor;

//====


Task type Especialista is
  Entry recibirResultado(cod: in, vSimilitud: in real);
  Entry recibirHuella(idS: in integer, test: out Text);
end Especialista;

Task body Especialista is
  huellaTest: Text;
  co: integer; vS: real;
begin
  huellaTest = tomandoHuella();
  for i in 1..8 loop
    Accept recibirHuella(idS: in integer, test: out Text)do //bidireccional
      test =huellaTest;
    end recibirHuella;
  end loop;
  for i in 1..8 loop
    Accept recibirResultado(cod: in integer, vSimilitud: in real) do
      co:= cod;
      vS:= vSimilitud;
    end recibirResultado;
  end loop;

end Especialista;

begin
  for i in 1..8 loop
    vecServers(i).recibirId(i);
  end loop;
end 8;



//=====================================================
Procedure 8 is

vecServers: array(1..8) of Servidor;

Task type Servidor is
  -- Entry recibirHuella(test: in Text);
  Entry recibirId(id: in integer);
end Servidor;

Task body Servidor is
  id: integer;
  codigo: integer; valor: real;
  testActual : Text;
begin
  Accept recibirId(pos: in integer) do --el servidor no necesita saber si id... logica inncesario... lo que me interesa es saber el codigo de la huella que tiene mas % de compatibilidad
    id :=  pos;
  end recibirId;
  loop
    -- Accept recibirHuella(test: in Text) do .... Si hago esto, atraso cada proceso.. cada llamdda esperia que cada Servidor busque en su db, y hasta que no termine uno, no busque el otro...
    --   Buscar(test,codigo,valor); ... malll, retraso todo... puede tardar 2 hs, en buscar en su db y hasta que no finalicen esas 2 horas los demes servers no empezaron a buscar
    -- end recibirHuella;
    Especialista.recibirHuella(test: in Text)do --solo recibo
      testActual = test;
    end recibirHuella
    Buscar(testActual,codigo,valor);
    Especialista.recibirResultado(codigo, valor);
end Servidor;

//====


Task type Especialista is
  Entry recibirResultado(cod: out, vSimilitud: out real);
  Entry recibirHuella(test: in Text);
end Especialista;

Task body Especialista is
  huellaTest: Text;
begin
  huellaTest = tomandoHuella();
  for i in 1.. do

  end Especialista;

end Especialista;

begin
  for i in 1..8 loop
    vecServers(i).recibirId(i);
  end loop;
end 8;