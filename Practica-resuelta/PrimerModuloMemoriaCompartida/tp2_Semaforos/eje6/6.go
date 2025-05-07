// 6.  Suponga  que  existe  una  BD  que  puede  ser  accedida  por  6  usuarios  como  máximo  al 
// mismo  tiempo.  Además,  los  usuarios  se  clasifican  como  usuarios  de  prioridad  alta  y 
// usuarios de prioridad baja. Por último, la BD tiene la siguiente restricción: 
// • no puede haber más de 4 usuarios con prioridad alta al mismo tiempo usando la BD. 
// • no puede haber más de 5 usuarios con prioridad baja al mismo tiempo usando la BD. 
// Indique si la solución presentada es la más adecuada. Justifique la respuesta.



6.  Suponga  que  existe  una  BD  que  puede  ser  accedida  por  6  usuarios  como  máximo  al 
mismo  tiempo.  Además,  los  usuarios  se  clasifican  como  usuarios  de  prioridad  alta  y 
usuarios de prioridad baja. Por último, la BD tiene la siguiente restricción: 
• no puede haber más de 4 usuarios con prioridad alta al mismo tiempo usando la BD. 
• no puede haber más de 5 usuarios con prioridad baja al mismo tiempo usando la BD. 
Indique si la solución presentada es la más adecuada. Justifique la respuesta


Var 
 sem: semaphoro := 6; 
 alta: semaphoro := 4; 
 baja: semaphoro := 5; 
 
Process Usuario-Alta [I:1..L]::  {    
	P (sem); 
  P (alta); 
  //usa la BD 
  V(sem); 
  V(alta); 
} 

Process Usuario-Baja [I:1..K]::  {    
	P (sem); 
  P (baja); 
  //usa la BD 
  V(sem); 
  V(baja); 
} 


El problema de la solución es que no se debería hacer P(total) antes que hacer P(alta) o P(baja) ya que esto puede 
bloquear la ejecución de otros procesos que necesiten usar la BD y que puedan hacerlo según las restricciones. 
Por ejemplo, supongamos que existen 5 procesos de prioridad alta y 5 procesos de prioridad baja, si se llegaran a ejecutar los 5 procesos 
de prioridad alta antes que los 5 de prioridad baja, se bloquearía la ejecución de los procesos de prioridad baja, 
ya que el 5to de prioridad alta disminuye con P(total) el valor de total cosa que estaría mal ya que según las 
restricciones no puede haber más de 4 usuarios de prioridad alta al mismo tiempo. Por lo tanto, la solución no es la más adecuada.