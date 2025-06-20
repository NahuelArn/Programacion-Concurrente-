4.  Se  debe  calcular  el  valor  promedio  de  un  vector  de  1  millón  de  números  enteros  que  se 
encuentra  distribuido  entre  10  procesos  Worker  (es  decir,  cada Worker  tiene  un  vector  de 
100  mil  números).  Para  ello,  existe  un  Coordinador  que  determina  el  momento  en  que  se 
debe  realizar  el  cálculo  de  este  promedio  y  que,  además,  se  queda  con  el  resultado.  Nota: 
maximizar la concurrencia; este cálculo se hace una sola vez. 

Procedure Promedio is

Task Admin is //task simple es para un solo proceso
  entry empezar;
  entry recibirSuma(suma: in int);
end Admin;

Task body Admin() is
  sumaDeTodos : int;
  prom : double
begin
  sumaDeTodos = 0;
  -- for(int i = 0 i<20; i++){ //usar la 
  for i in (1..20) loop
    select
      Accept empezar;
    OR
      Accept recibirSuma(suma: in int)do //el que recibe especifica si es IN/OUT los parametros  && el que recibe no es necesario declarar el parametro
        sumaDeTodos+= suma;
      end recibirSuma;
    end select;
  end loop;
  prom = sumaDeTodos / 100 000;
end Admin;



Task type Worker;

ArrW: array(1..10) of Worker; //especifico cuantos procesos pueden haber de una misma tarea

Task body Worker is
  vecDisponible: array(1..100000)of integer ; 
  suma: int; 
begin
  suma = 0;
  Admin.empezar;
  -- for(int i = i < vecDisponible.size(); i++){
  for i in 1..100 000 loop
    -- suma+= vecDisponible[i]; //usar sintaxis de ada
    suma+= vecDisponible(i);
  end loop;
  Admin.recibirSuma(suma);
End Worker;

begin
  null;
end Promedio;


-- Procedure Promedio is

-- Task Admin is //task simple es para un solo proceso

-- end Admin;



-- Task type Worker is //type es cuando voy a tener mas de 1 proceso que va utilizar esta tarea/entender/mensajes

-- end Worker;

-- ArrW: array(1..10) of Worker;

-- Task body Worker is
--   vecRecibido: array(1..100000)of integer := InicializarVector; //inicializa en 0

-- begin

-- End Worker;



-- begin

-- end Promedio;