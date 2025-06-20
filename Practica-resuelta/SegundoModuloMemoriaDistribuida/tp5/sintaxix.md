//sintaxis pascal/Semaforos/monitores/sql

Procedure nombreDelEjeArchivo is

//se debe especificar los Entry de cada proceso
//los entry son una mezcla entre una función/método y un semáforo con cola de espera.
Task Admin is
  Entry hola(casa: in Integer)
  Entry adios(casa: in Integer)
End Admin;


//zona de declaracion de tipos de tareas (el proceso como tal)
Task Type Auto;
Task type camion;

arrClientes: array(1..10) of Auto;
arrCamiones: array(1..10) of camion;


Task body Admin is
  //se pueden declarar variables locales

  casa: Integer;
  //se pueden declarar variables de tipo tarea
  auto: Auto;
  camion: camion;
  //se pueden declarar variables de tipo array
  arr: array(1..10) of Integer;
Begin
  casa = 0;
  Loop //es un while true
    SELECT
      WHEN(1>2) => Aceppt
      WHEN(2>1) => Reject
      WHEN(3>4) => Accept
      ELSE => Reject;
    END SELECT;
  end Loop;

End body Admin;

//lo mismo para el resto de procesos

//programa principal 
begin 
  null
end nombreDelEjeArchivo;