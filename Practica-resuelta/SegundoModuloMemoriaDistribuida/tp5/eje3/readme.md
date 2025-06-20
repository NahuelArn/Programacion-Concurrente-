

### ============= Ejemplo de especificacion de las tareas

```ADA

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

  end Periferico1;

  Task body Periferico2 is

  end Periferico2;

  Task body Central is

  end Central;


begin


end 3;

```