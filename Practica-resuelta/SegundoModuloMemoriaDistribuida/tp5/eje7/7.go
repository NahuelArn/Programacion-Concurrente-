7.  En una playa hay 5 equipos de 4 personas cada uno (en total son 20 personas donde cada 
	una  conoce  previamente  a  que  equipo  pertenece).  Cuando  las  personas  van  llegando 
	esperan  con  los  de  su  equipo  hasta  que  el  mismo  esté  completo  (hayan  llegado  los  4 
	integrantes), a partir de ese momento el equipo comienza a jugar. El juego consiste en que 
	cada integrante del grupo junta 15 monedas de a una en una playa (las monedas pueden ser 
	de  1,  2  o  5  pesos)  y  se  suman  los  montos  de  las  60  monedas  conseguidas  en  el  grupo.  Al 
	finalizar  cada  persona  debe  conocer  el  grupo  que  más  dinero  junto.  Nota:  maximizar  la 
	concurrencia.  Suponga  que  para  simular  la  búsqueda  de  una  moneda  por  parte  de  una 
	persona existe una función Moneda() que retorna el valor de la moneda encontrada.

Procedure eje7 is

//==========================================Equipo
Task type Equipo is
		Entry llegoAbarrera;
		Entry salirDebarrera;
		Entry recibirSumaIndividual(monto: in integer);
		Entry ident(pos: in integer);
end Equipo;

vecEquipo: array(1..5) of Equipo;

Task body Equipo is
		sumaGrupal: integer;
		integer idEquipo;
begin
		accept ident(pos: in integer) do
			idEquipo = pos;
		end ident;

		for i in (1..4) loop //levanto la barrera
			Accept llegoAbarrera;
		end loop;
		for i in 1..4 loop //tiro la barrera ... todos los del equipo son liberados para que empiecen a juntar monedas
			Accept salirDebarrera;
		end loop;
		sumaGrupal = 0;
		for i in 1..4 loop
			Accept recibirSumaIndividual(monto: in integer)do
				sumaGrupal = sumaGrupal + monto;
			end recibirSumaIndividual;
		end;
		Administrador.recibirSumaDeEquipo(idEquipo,sumaGrupal);
end Equipo;
//==========================================Persona

Task type Persona is
		Entry queEquipoGano(idEquipoGanador: in integer);
end Persona;

vecPersona: array(1..20) of Persona;

Task body Persona is
		idEquipo: integer;
		sumaIndividual: integer;
		moneda Moneda;
		equipoGanador: integer;
begin
	idEquipo := yoSeCualEsMiIDdeEquipo();
	sumaIndividual = 0;
	vecEquipo(idEquipo).llegoAbarrera;
	vecEquipo(idEquipo).salirDebarrera;
	for 1 in 1..15 loop
		moneda = buscandoMonerada(Moneda(moneda));
		sumaIndividual = sumaIndividual + moneda;
	end loop;
	vecEquipo(idEquipo).recibirSumaIndividual(sumaIndividual);
	Accept queEquipoGano(equipoGanador: in integer) do//aca me entero el que equipoGano
		if(idEquipoGanador == idEquipo)then
			puts "Easy";
		end if;
	end queEquipoGano;
end Persona;
//==========================================Administrador

Task Administrador is
		Entry recibirSumaDeEquipo(id: in integer, sumaEquipo: in integer);
end Administrador;

Task body Administrador is
		vecResultadosEquipos: array(1..5) of Integer;
		max: Integer; integer idEquipoGanador;
begin
	// ponerEn0Vector(vecEquipos); ya recibo la totalidad
	for i in 1..5 loop
		Accept recibirSumaDeEquipo(idEquipo: in integer, sumaEquipo: in integer)do
			vecResultadosEquipos(idEquipo) = sumaEquipo;
		end recibirSumaDeEquipo;
	end loop;
	idEquipoGanador = vecResultadosEquipos(1);
	max = vecResultadosEquipos(1);
	for i in (1..5) loop
		if vecResultadosEquipos(i) > max then
			max = vecResultadosEquipos(i);
			idEquipoGanador = i;
		end if;
	end;
	for i in 1..20 loop	//no me interesa mandar a ids especificos me dice que todos se tienen enterar que equipo gano.. aviso a todos
		vecPersona(i).queEquipoGano(idEquipoGanador);
	end loop;
end Administrador;

//cada Jugador sabe a que equipo pertenece...
//Pero cada Equipo no sabe que nro de equipo es... entonces tengo que avisarles
begin
		for i in 1..5 loop
			vecEquipo(i).ident(i);
		end loop;
end eje7;























//asdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
with Ada.Text_IO; use Ada.Text_IO;
with Ada.Numerics.Discrete_Random;

procedure Eje7 is

   function Moneda return Integer is
      subtype Moneda_Valor is Integer range 1 .. 3;
      package Aleatorio is new Ada.Numerics.Discrete_Random(Moneda_Valor);
      use Aleatorio;
      Gen : Generator;
      V : Moneda_Valor;
   begin
      Reset(Gen);
      V := Random(Gen);
      case V is
         when 1 => return 1;
         when 2 => return 2;
         when 3 => return 5;
         when others => return 1;
      end case;
   end Moneda;

   function YoSeCualEsMiIDdeEquipo(i : Integer) return Integer is
   begin
      return ((i - 1) / 4) + 1; -- Persona 1-4 → Equipo 1, 5-8 → Equipo 2, ...
   end;

   -- ===================== Administrador =====================
   task Administrador is
      entry RecibirSumaDeEquipo(Id : in Integer; Suma : in Integer);
   end Administrador;

   task body Administrador is
      VecResultados : array(1 .. 5) of Integer := (others => 0);
      Max, IdGanador : Integer := 0;
   begin
      for I in 1 .. 5 loop
         accept RecibirSumaDeEquipo(Id : in Integer; Suma : in Integer) do
            VecResultados(Id) := Suma;
         end;
      end loop;

      IdGanador := 1;
      Max := VecResultados(1);
      for I in 2 .. 5 loop
         if VecResultados(I) > Max then
            Max := VecResultados(I);
            IdGanador := I;
         end if;
      end loop;

      for I in 1 .. 20 loop
         VecPersona(I).QueEquipoGano(IdGanador);
      end loop;
   end Administrador;

   -- ===================== Equipo =====================
   task type Equipo is
      entry LlegoABarrera;
      entry SalirDeBarrera;
      entry RecibirSumaIndividual(Monto : in Integer);
      entry Ident(Pos : in Integer);
   end Equipo;

   VecEquipo : array(1 .. 5) of Equipo;

   task body Equipo is
      SumaGrupal, IdEquipo : Integer;
   begin
      accept Ident(Pos : in Integer) do
         IdEquipo := Pos;
      end;

      for I in 1 .. 4 loop
         accept LlegoABarrera;
      end loop;

      for I in 1 .. 4 loop
         accept SalirDeBarrera;
      end loop;

      SumaGrupal := 0;
      for I in 1 .. 4 loop
         accept RecibirSumaIndividual(Monto : in Integer) do
            SumaGrupal := SumaGrupal + Monto;
         end;
      end loop;

      Administrador.RecibirSumaDeEquipo(IdEquipo, SumaGrupal);
   end Equipo;

   -- ===================== Persona =====================
   task type Persona is
      entry QueEquipoGano(IdEquipoGanador : in Integer);
   end Persona;

   VecPersona : array(1 .. 20) of Persona;

   task body Persona is
      IdEquipo, Suma : Integer;
      Mon : Integer;
   begin
      -- Identificar equipo según índice
      IdEquipo := YoSeCualEsMiIDdeEquipo(Integer'Value(Task_Identity));

      VecEquipo(IdEquipo).LlegoABarrera;
      VecEquipo(IdEquipo).SalirDeBarrera;

      Suma := 0;
      for I in 1 .. 15 loop
         Mon := Moneda;
         Suma := Suma + Mon;
      end loop;

      VecEquipo(IdEquipo).RecibirSumaIndividual(Suma);

      accept QueEquipoGano(IdEquipoGanador : in Integer) do
         if IdEquipoGanador = IdEquipo then
            Put_Line("Easy");
         end if;
      end;
   end Persona;

begin
   -- Identificar cada equipo
   for I in 1 .. 5 loop
      VecEquipo(I).Ident(I);
   end loop;
end Eje7;
