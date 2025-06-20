1.  Se requiere modelar un puente de un único sentido que soporta hasta 5 unidades de peso. 
El peso de los vehículos depende del tipo: cada auto pesa 1 unidad, cada camioneta pesa 2 
unidades  y  cada  camión  3  unidades.  Suponga  que  hay  una  cantidad  innumerable  de 
vehículos  (A  autos,  B  camionetas  y  C  camiones).  Analice  el  problema  y  defina  qué  tareas, 
recursos y sincronizaciones serán necesarios/convenientes para resolver el problema. 
a. Realice  la  solución  suponiendo  que  no  se  tiene  ningún  orden  ni  prioridad  entre  los 
diferentes tipos de vehículos. 
b. Modifique la solución de (a) para que tengan mayor prioridad los camiones que el resto 
de los vehículos. 

tipo  | weight
auto-> 1
caminoneta -> 2
camion -> 3


A)
Procedure 1 is

Task Admin is
  Entry vaSalir(Peso: in Integer);
  Entry quierePasarAuto;
  Entry quierePasarCamioneta;
  Entr quierePasarCamion;
End Admin;

Task type Auto;
Task type Camioneta;
Task type Camion;

arrClientes: array(1..N) of Auto;
arrClientes: array(1..N) of Camioneta;
arrClientes: array(1..N) of Camion;

Task body Admin is
  pesoActual: integer;
Begin
  pesoActual = 0;
  LOOP
    SELECT
      WHEN (pesoActual+1 < 5) => 
        ACCEPT quierePasarAuto do //aca es necesario los do, ya que al tener varios procesos que pueden usar pesoActual.. sin el do no vaser atomico 
          pesoActual+= 1;
        end quierePasarAuto;
      OR WHEN (pesoActual+2 < 5) =>
        ACCEPT quierePasarCamioneta do
          pesoActual+= 2;
        end quierePasarCamioneta;
      OR WHEN (pesoActual+3 < 5) =>
        ACCEPT quierePasarCamion do
          pesoActual+= 3;
        end quierePasarCamion;
      OR  ACCEPT vaSalir(peso: in integer) do
          pesoActual -= peso;
        end vaSalir;
    end SELECT;
  end loop;
end Admin;

Task body Auto is
Begin
  Admin.quierePasarAuto;
  Admin.vaSalir(1);
end Auto;

Task body Camioneta is
Begin
  Admin.quierePasarCamioneta;
  Admin.vaSalir(2);
end Camioneta;

Task body Camion is
Begin
  Admin.quierePasarCamion;
  Admin.vaSalir(3);
end Camion;

Begin 
  null
End aaa1;


//========================================================B===================================================================

A)
Procedure 1 is

Task Admin is
  Entry vaSalir(Peso: in Integer);
  Entry quierePasarAuto;
  Entry quierePasarCamioneta;
  Entr quierePasarCamion;
End Admin;

Task type Auto;
Task type Camioneta;
Task type Camion;

arrClientes: array(1..N) of Auto;
arrClientes: array(1..N) of Camioneta;
arrClientes: array(1..N) of Camion;

Task body Admin is
  pesoActual: integer;
Begin
  pesoActual = 0;
  LOOP
    SELECT
      WHEN (pesoActual+1 < 5) => 
        ACCEPT quierePasarAuto do //aca es necesario los do, ya que al tener varios procesos que pueden usar pesoActual.. sin el do no vaser atomico 
          pesoActual+= 1;
        end quierePasarAuto;
      OR WHEN (pesoActual+2 < 5) =>
        ACCEPT quierePasarCamioneta do
          pesoActual+= 2;
        end quierePasarCamioneta;
      OR WHEN (quierePasarCamion'COUNT==0 && pesoActual+3 < 5) => //Le doy mas prioridad
        ACCEPT quierePasarCamion do
          pesoActual+= 3;
        end quierePasarCamion;
      OR  ACCEPT vaSalir(peso: in integer) do
          pesoActual -= peso;
        end vaSalir;
    end SELECT;
  end loop;
end Admin;

Task body Auto is
Begin
  Admin.quierePasarAuto;
  Admin.vaSalir(1);
end Auto;

Task body Camioneta is
Begin
  Admin.quierePasarCamioneta;
  Admin.vaSalir(2);
end Camioneta;

Task body Camion is
Begin
  Admin.quierePasarCamion;
  Admin.vaSalir(3);
end Camion;

Begin 
  null
End 1;