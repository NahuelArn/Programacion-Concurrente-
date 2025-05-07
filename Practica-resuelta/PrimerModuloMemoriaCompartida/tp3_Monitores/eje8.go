En un examen de la secundaria hay  un preceptor y una profesora que deben tomar un examen 
escrito a 45 alumnos. El preceptor se encarga de darle el enunciado del examen a los alumnos 
cundo los 45 han llegado (es el mismo enunciado para todos). La profesora se encarga de ir 
corrigiendo los exámenes de acuerdo con el orden en que los alumnos van entregando. Cada 
alumno al llegar espera a que le den el enunciado, resuelve el examen, y al terminar lo deja 
para que la profesora lo corrija y le envíe la nota. Nota: maximizar la concurrencia; todos los 
procesos deben terminar su ejecución; suponga que la profesora tiene una función 
corregirExamen que recibe un examen y devuelve un entero con la nota.

Monitor BeforeExamen{
  Enunciado enunciado;
  cond alumno;
  cond preceptor;
  int cantAlumnos;
  //boolean pasaronTodos = false;
  
  procedure esperandoAlosDemas(enunciadO : out Enunciado){
    cantAlumnos++;
    if(cantAlumnos == 45){
      signal(preceptor);
      //pasaronTodos = true;
    }
    wait(alumno);
    enunciadO = enunciado;
  {  
  
  procedure darExamenes(enunciadoI: in Enunciado){
    if(cantAlumnos != 45){
      wait(preceptor);
    }
    enunciado = enunciadoI;
    signal_all(alumno);
  }
}

Monitor InExamen{
  cond profesora;
  Queue examenesEntregados;
  
  procedure resolviendoExamen(id : in int, enunciado : in Enunciado){
    //resolviendo Examen
    examenesEntregados.push(id);
    signal(profesora);
  }
  
  procedure corregirExamen(){
    if(examenesEntregados.isEmpty(){
      wait(profesora);
    }
    int idAlumno = examenesEntregados.pop()
    int nota = corregirExamen(idAlumno);
    enviarNota(idAlumno ); //nunca dice que el alumno espera la nota
  }
  
}

Process Alumno[id: 1..45]{
  Enunciado enunciado;
  BeforeExamen.esperandoAlosDemas(enunciado);
  InExamen.resolviendoExamen(id,enunciado);
}

Process Preceptor{
  Enunciado enunciado;
  BeforeExamen.darExamenes(enunciado);
}

Process Profesora{
  //int cantExamenes = 45;
  for(int i = 0; i < 45; i++{
    InExamen.corregirExamen();
  }
}
