3.  Se  dispone  de  un  sistema  compuesto  por  1  central  y  2  procesos  periféricos,  que  se 
comunican continuamente. Se requiere modelar su funcionamiento considerando las 
siguientes condiciones: 
- La  central  siempre  comienza  su  ejecución  tomando  una  señal  del  proceso  1;  luego 
toma  aleatoriamente  señales  de  cualquiera  de  los  dos  indefinidamente.  Al  recibir  una 
señal de proceso 2, recibe señales del mismo proceso durante 3 minutos. 
- Los  procesos  periféricos  envían  señales  continuamente  a  la  central.  La  señal  del 
proceso  1  será  considerada  vieja  (se  deshecha)  si  en  2  minutos  no  fue  recibida.  Si  la 
señal del proceso 2 no puede ser recibida inmediatamente, entonces espera 1 minuto y 
vuelve a mandarla (no se deshecha).

program 3 is

  Task Central is   //defino los mensajes que entiende Central
    Entry senhalPeriferico1(dato: in Dato);
    Entry senhalPeriferico2(dato: in Dato);
    Entry strobe;
  end Central
  
  Task body Central is
    sigo: boolean;
  begin
    Accept senhalPeriferico1(senhial: in Senhal);
    LOOP 
      SELECT 
        Accept senhalPeriferico1(senhial: in Senhial);
      OR 
        Accept senhalPeriferico2(senhial: in Senhial);
        sigo = true;
        Timer.iniciarTimer;
        LOOP (sigo)
          select 
            when (strobe'count =0) =>
              Accept senhalPeriferico2(senhial: in Senhial);
            OR
              Accept strobe; //no lo hago con exclusion mutua, en este punto es atomico
              sigo= false;
          end select;
        end loop;
      end select;
    END loop;
  end Central;

  //=======

  TASK Periferico1; //defino los mensajes q entiene Periferico1

  Task body Periferico1 is
    senhial: Senhial;
  begin
    loop 
      senhial = captandoUnaSenhal();
      select 
        Central.senhalPeriferico1(senhial);
      or delay (120)
        null;
      end select;
    end loop;
  end Periferico1;

  //========

  TASK Periferico2; //defino los mensajes q entiene Periferico2

  Task body Periferico2 is
  begin
    senhial = captandoUnaSenhal();
    loop
      select 
        Central.senhalPeriferico2(senhial);
      else
        delay (60)
      end select;
    end loop;
  end Periferico2;

  //=========

  TASK Timer is //defino los mensajes que entiende Timer
    Entry iniciarTimer;
  END Timer;

  Task Body Timer is
  begin
    Accept iniciarTimer;
    delay(180);
    Central.strobe;
  end Timer;

begin
  null;
end 3;









//=========================== Ejemplo de especificacion de las tareas


program 3 is

  Task Central is   //defino los mensajes que entiende Central
    Entry senhalPeriferico1(dato: in Dato);
    Entry senhalPeriferico2(dato: in Dato);
    Entry fin;
  end Central

  TASK Periferico1; //Task type "es para declarar un molde de una tarea, cuando voy a usar un array"
  TASK Periferico2; //task solo es cuando solo voy a tener 1 sola tarea concreta

  TASK Timer is //defino los mensajes que entiende Timer
    Entry iniciarTimer;
  END Timer;

  Task Body Timer is
  begin

  end Timer;


  Task body Periferico1 is
  begin
  end Periferico1;

  Task body Periferico2 is
  begin
  end Periferico2;

  Task body Central is
  begin
  end Central;


begin


end 3;