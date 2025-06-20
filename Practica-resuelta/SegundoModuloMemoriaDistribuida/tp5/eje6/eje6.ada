6.  En  un  sistema  para  acreditar  carreras  universitarias,  hay  UN  Servidor  que  atiende  pedidos 
de U Usuarios de a uno a la vez y de acuerdo con el orden en que se hacen los pedidos.  
Cada  usuario  trabaja  en  el  documento  a  presentar,  y  luego  lo  envía  al  servidor;  espera  la 
respuesta de este que le indica si está todo bien o hay algún error. Mientras haya algún error, 
vuelve a trabajar con el documento y a enviarlo al servidor. Cuando el servidor le responde 
que está todo bien, el usuario se retira. Cuando un usuario envía un pedido espera a lo sumo 
2 minutos a que sea recibido por el servidor, pasado ese tiempo espera un minuto y vuelve a 
intentarlo (usando el mismo documento).


Procedure Ejercicio4.2 is

  U: int := U;
TASK Servidor is
    ENTRY peticionUsuario(doc: IN TEXT , OK : OUT Boolean);
end servidor;
TASK BODY Servidor is
ok: boolean
cant : int;
Begin
    LOOP (cant < U);
        SELECT 
            ACCEPT peticionUsuario(doc: IN TEXT , OK : OUT Boolean) do
                OK = ProcesarDcoumento (doc);
                if(!ok)
                  cant++1;
                end inf;
            end peticionUsuario;
        end Select;
    END LOOP;
END BODY servidor;

TASK TYPE Usuario;
TASK BODY Usuario is
Documento : TEXT; OK:  boolean;
Begin
    OK=false;
    Documento = new Documento();
    while (not OK) LOOP
        SELECT
            Servidor.peticionUsuario(Documento, OK);
            If (not OK) then
                Documento = ArreglarDocumento(Documento);
            end IF;
        OR  DELAY 120.0;
            DELAY 60.0;
        end Select;
    END LOOP;
END BODY Usuario;

Begin
    NULL;
end Ejercicio4.2;


//esta solucion estaria bien, si tuviera que hacer algo hasta que me de una respuesta, o tendria un intermediario
//===== "y  luego  lo  envía  al  servidor;  espera  la 
respuesta de este que le indica si está todo bien o hay algún error"
Procedure eje6 is

Task Servidor is
  Entry presentarDocumento(id: in int, doc: in Text);
  
end Servidor;

Task body Servidor is
  cantFinalizados: int;
  hayError : boolean;
begin
  loop cantFinalizados < U 
    Accept presentarDocumento(id in: int, doc: in Text) do
      hayError = revisandoDocumento(doc);
      vecUsu(id).entregarDocumento(doc,hayError);
      if(!hayError)then
        cantFinalizados++;
      end if;
    end presentarDocumento;
  end loop;
end Servidor;


//----------------------------

Task type Usuario is
  Entry entregarDocumento(doc: out Text, hayError: out boolean);
end Usuario;

vecUsu: array(1..U)of Usuario;

Task body Usuario is
  doc: Tex;
  id: int;
  aceptado, hayError : boolean;
begin
  Accept Ident(Pos: in int) do
    id = Pos;
  end Ident;
  elaborandoDocumento(doc);
  loop (hayError)
    aceptado = false;
    loop (!aceptado)
      select 
        Servidor.presentarDocumento(id,doc);
        aceptado = true;
      OR DELAY 2*60
        delay 60
    end loop;
    Accept entregarDocumento(doc, hayError);
    if(hayError)then
      doc = arreglandoErrores(doc);
    end if;
  end loop;


end Usuario;

begin
  for i in 1..U loop
    vecUsu(i).Ident(i);
  end loop;
end eje6;