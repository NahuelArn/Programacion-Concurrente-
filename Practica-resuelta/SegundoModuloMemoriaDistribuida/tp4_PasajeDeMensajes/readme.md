diferencia entre los 2

PMA
ACA LOS CANALES SE DECLARAN
envio //no bloquante
recepcion //bloquante
-No vimos la barrera "para PMA"

PMS
-ACA LOS CANALES NO SE DECLARAN
-Cuando envio un mensaje, necesito tener al receptor listo para recibirlo, o el emisor se bloquea hasta que el receptor lo reciba.
-las "guardas" solo lo podemos usar para recibir mensajes, no para enviar.
- "GUARDA" si no tengo condicion, siempre es true... por lo tanto, nunca sale del "do od".. seria como un while (true) <- importante
Por lo tanto, la unica forma de salir del "do od" es teniendo todas las "guardas" con condiciones.Y que en un momento todas fallen (sean false)..............(si tengo N gurdas, y N-1 tienen condicion y es false, pero una no tiene condicion(siempre true), entonces nunca sale del "do od")
-Tenemos la barrera
"
Hay que tener cuidado al usar la recepción con [*], ya que si más de un proceso hizo un envío en ese canal, se recibirá el pedido de cualquiera de ellos, sin respetar el orden de llegada. (SOLO EN LA RECEPCION esta el [*])
"

Los canales son punto a punto, no hay una Queue implicita por
envio //bloquante
recepcion //bloqueante



PMA
Potencialmente va a necesitar de un Coordinador -> para resolver su problema de isEmpty()

PMS
Potencialmente va a necesitar de un Administrador -> para resolver su problema de esperaEntreComunicaciones 