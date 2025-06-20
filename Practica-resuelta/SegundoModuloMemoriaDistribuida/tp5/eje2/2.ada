2.  Se quiere modelar el funcionamiento de un banco, al cual llegan clientes que deben realizar 
un pago y retirar un comprobante. Existe un único empleado en el banco, el cual atiende de 
acuerdo con el orden de llegada.  
a. Implemente  una  solución  donde  los  clientes  llegan  y  se  retiran  sólo  después  de  haber 
sido atendidos. 
b. Implemente  una  solución  donde  los  clientes  se  retiran  si  esperan  más  de  10  minutos 
para realizar el pago. 
c. Implemente una solución donde los clientes se retiran si no son atendidos 
inmediatamente. 
d. Implemente  una  solución  donde  los  clientes  esperan  a  lo  sumo  10  minutos  para  ser 
atendidos. Si pasado ese lapso no fueron atendidos, entonces solicitan atención una vez 
más y se retiran si no son atendidos inmediatamente. 


//================================================ A ================================================================

Procedure eje2 is

Task Empleado is
  Entry atenderCliente(id: in int);
  Entr ack;
End Empleado;

Task type Cliente;

arrClientes: array(0..C-1) of Cliente;

Task Body Empleado is
//aca van las variables locales (en este caso no las necesito)
Begin
  loop
    acept atenderCliente(idC: in integer; c: out comprobante) do //el id no lo necesito, la quueue implicita ya lo maneja
      atendiendoCliente(idC, c);
    end atenderCliente;
  end loop;
end Empleado;

Task body Cliente is
  comprobante Text;
  id: Integer;
Begin 
  //Aca me entero que id tengo  //en este eje no necesito saber mi Id
  Accept Ident(Pos: in Integer) do
    id:= Pos;
  end Ident;

  Empleado.atenderCliente(idC, comprobante);
end Cliente;

Begin
  for (int i =0; i<C; i++){ //esto es para el id
    arrClientes(i).Ident(i);
  }
end eje2;



//================================================ B ================================================================


Procedure eje2 is

Task Empleado is
  Entry atenderCliente(id: in int);
  Entr ack;
End Empleado;

Task type Cliente;

arrClientes: array(0..C-1) of Cliente;

Task Body Empleado is
//aca van las variables locales (en este caso no las necesito)
Begin
  loop
    acept atenderCliente(idC: in integer; c: out comprobante) do //el id no lo necesito, la quueue implicita ya lo maneja
      atendiendoCliente(idC, c);
    end atenderCliente;
  end loop;
end Empleado;

Task body Cliente is
  comprobante Text;
  id: Integer;
Begin 
  //Aca me entero que id tengo  //en este eje no necesito saber mi Id
  Accept Ident(Pos: in Integer) do
    id:= Pos;
  end Ident;
  SELECT 
    Empleado.atenderCliente(idC, comprobante);
  OR DELAY 600.0
    NULL
  END SELECT;
end Cliente;

Begin
  for (int i =0; i<C; i++){ //esto es para el id
    arrClientes(i).Ident(i);
  }
end eje2;


//================================================ C ================================================================

Procedure eje2 is

Task Empleado is
  Entry atenderCliente(id: in int);
  Entr ack;
End Empleado;

Task type Cliente;

arrClientes: array(0..C-1) of Cliente;

Task Body Empleado is
//aca van las variables locales (en este caso no las necesito)
Begin
  loop
    acept atenderCliente(idC: in integer; c: out comprobante) do //el id no lo necesito, la quueue implicita ya lo maneja
      atendiendoCliente(idC, c);
    end atenderCliente;
  end loop;
end Empleado;

Task body Cliente is
  comprobante Text;
  id: Integer;
Begin 
  //Aca me entero que id tengo  //en este eje no necesito saber mi Id
  Accept Ident(Pos: in Integer) do
    id:= Pos;
  end Ident;
  SELECT 
    Empleado.atenderCliente(idC, comprobante);
  ELSE //Si Empleado no acepta inmediatamente su pedido, cancela el entry call 
    NULL
  END SELECT;
end Cliente;

Begin
  for (int i =0; i<C; i++){ //esto es para el id
    arrClientes(i).Ident(i);
  }
end eje2;


//================================================ D ================================================================


d. Implemente  una  solución  donde  los  clientes  esperan  a  lo  sumo  10  minutos  para  ser 
atendidos. Si pasado ese lapso no fueron atendidos, entonces solicitan atención una vez 
más y se retiran si no son atendidos inmediatamente. 

//se resuelve concatenando selects

Procedure eje2 is

Task Empleado is
  Entry atenderCliente(id: in int);
  Entr ack;
End Empleado;

Task type Cliente;

arrClientes: array(0..C-1) of Cliente;

Task Body Empleado is
//aca van las variables locales (en este caso no las necesito)
Begin
  loop
    acept atenderCliente(idC: in integer; c: out comprobante) do //el id no lo necesito, la quueue implicita ya lo maneja
      atendiendoCliente(idC, c);
    end atenderCliente;
  end loop;
end Empleado;

Task body Cliente is
  comprobante Text;
  id: Integer;
Begin 
  //Aca me entero que id tengo  //en este eje no necesito saber mi Id
  Accept Ident(Pos: in Integer) do
    id:= Pos;
  end Ident;
  SELECT 
    Empleado.atenderCliente(idC, comprobante);
  ELSE OR DELAY 600.0
    SELECT
      Empleado.atenderCliente(idC, comprobante);
    ELSE
      NULL;
    End select;
  END SELECT;
end Cliente;

Begin
  for (int i =0; i<C; i++){ //esto es para el id
    arrClientes(i).Ident(i);
  }
end eje2;
