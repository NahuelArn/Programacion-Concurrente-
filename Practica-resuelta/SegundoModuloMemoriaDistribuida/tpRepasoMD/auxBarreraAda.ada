procedure Contador is
   Task type Worker is
      Entry Listo;
      Entry Comenzar;
      Entry AsignarID(id: in Integer);
   end Worker;

   Task type Coordinador is
      Entry AvisarListo;
      Entry Resultado(res: in Integer);
   end Coordinador;

   workers: array(1 .. 10) of Worker;
   coordinador: Coordinador;

   Task body Worker is
      vector: array(1 .. 100000) of Float;
      res: Integer := 0;
   begin
      accept AsignarID(id: in Integer) do
        id_local := id;
      end AsignarID;
      coordinador.AvisarListo;
      accept Comenzar;
      res := SacarPromedio(vector);
      coordinador.Resultado(res);
   end Worker;

   Task body Coordinador is
      resultados: array(1 .. 10) of Float;
      listos: Integer := 0;
   begin
      -- Esperar a que los 10 workers estén listos
      for I in 1 .. 10 loop
         accept AvisarListo;
         listos := listos + 1;
      end loop;
      -- Todos listos, ahora sí los arrancamos
      for I in 1 .. 10 loop
         workers(I).Comenzar;
      end loop;
      -- Recolectar resultados
      for I in 1 .. 10 loop
         accept Resultado(res: in Integer) do
            resultados(I) := res;
         end Resultado;
      end loop;
   end Coordinador;
begin
   for I in 1 .. 10 loop
    workers(I).AsignarID(I);
   end loop;
end Contador;
