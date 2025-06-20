1) Resolver el siguiente problema. La página web del Banco Central exhibe las diferentes cotizaciones 
del dólar oficial de 20 bancos del país, tanto para la compra como para la venta. Existe una tarea 
programada que se ocupa de actualizar la página en forma periódica y para ello consulta la cotización 
de cada uno de los 20 bancos. Cada banco dispone de una API, cuya única función es procesar las 
solicitudes de aplicaciones externas. La tarea programada consulta de a una API por vez, esperando 
a lo sumo 5 segundos por su respuesta. Si pasado ese tiempo no respondió, entonces se mostrará 
vacía la información de ese banco.


Procedure s is

Task Tarea is
  -- Entry miCotizacion(cotizacionActual: in real);
end Tarea;

Task body Tarea is

begin
  loop 
    for i in (1..20)loop
      select 
          vecBanco(i).miCotizacion(cotizacionCompra: out real, cotizacionVenta: out real);
          actualizarValorParaUnBancoPagina(cotizacionCompra,cotizacionVenta);
      or delay 5
        actualizarValorPagina(null);
      end select;
    end loop;
    delay (5*60) --cada cuanto tiempo le voy a pedir actualizacion de cotizaciones de los 20 bancos
  end loop;
end;

//====

vecBanco: array(1..20) of Banco;
Task type Banco is
  Entry miCotizacion(cotizacionCompra: out real, cotizacionVenta: out real);
  Entry recibirId(id: in integer);
end banco;

task body Banco is
  id: integer;
  v,c : real;
begin
  Accept recibirId(pos: in integer)do
    id = pos;
  end recibirId;
  loop
    v = consultarApiCotizacionV(); --cuestionable... de esta forma no calcularia la cotizacion en el momento que tarea hizo la solicitud... podria devolverle valores calculados 3 horas atras
    c = consultarApiCotizacionV(); --consultar... en este caso creo que seria mejor calcular dentro del do... no me importa cuanto tarde la api.. necesito el valor en el momento que se me pide el accept
    Accept miCotizacion(cotizacionCompra: out real, cotizacionVenta: out real)do
      -- cotizacionCompra = consultarApiCotizacionC(); //esto estaria mal, el llamado a la funcion puede tardar N y estaria demorando la conexion... lo calculo anteriormente
      -- cotizacionVenta = consultarApiCotizacionV();
      cotizacionCompra =c;
      cotizacionVenta = v;
    end do;
  end loop;
end Banco;



begin
  for i in (1..20)loop
    vecBanco(i).recibirId(i);
  end loop;

end s;