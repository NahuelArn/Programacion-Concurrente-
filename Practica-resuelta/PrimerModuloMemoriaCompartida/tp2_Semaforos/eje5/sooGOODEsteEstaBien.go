// 5.  Suponga que se tiene un curso con 50 alumnos. Cada alumno debe realizar una tarea y 
// existen  10  enunciados  posibles.  Una  vez  que  todos  los  alumnos  eligieron  su  tarea, 
// comienzan a realizarla. Cada vez que un alumno termina su tarea, le avisa al profesor y se 
// queda esperando el puntaje del grupo, el cual está dado por todos aquellos que comparten 
// el mismo enunciado. Cuando todos los alumnos que tenían la misma tarea terminaron, el 
// profesor les otorga un puntaje que representa el orden en que se terminó esa.  


// Nota: para elegir la tarea suponga que existe una función elegir que le asigna una tarea a 
// un alumno (esta función asignará 10 tareas diferentes entre 50 alumnos, es decir, que 5 
// alumnos tendrán la tarea 1, otros 5 la tarea 2 y así sucesivamente para las 10 tareas). 

int barrera = 0; 
int contAlumnos = 0; // contador de alumnos

sem semaforoAlumno = 1;

Queue<Tarea> tareas;
sem protexQueuTareas = 1; // semaforo para proteger la cola de tareas

sem estaListNotaGrupal [0..9] = ([9] 0)
notasGrupal[0..9] = ([9] -1); // notas del grupo, si esta en -1, tadavia no se corrigio la tarea

cantEntregasXgrupo[0..9] = ([9] 0); // contador de entregas por grupo
sem protexCantEntregasXgrupo [0..9] =([9] 1) ; // semaforo para proteger el contador de entregas por grupo

tarea = {
	int numero: 0..9; 
}

process alumno[id: 0..49] {
	// Tarea tarea = elegir(); 
	P(semaforoAlumno);
	Tarea tarea = elegir();
	contAlumnos++; // aumento el contador de alumnos
	if(contAlumnos == 50){
		//El ultimo toca el timbre	///Preguntar el FOR ACA, ESE Q V(BARRERA) X50, ES NECESARIO? CON 1 V(CARRERA) NO BASTA?
		For (i = 0..49) -> V(barrera); 
	}
	V(semaforoAlumno); 
	P(barrera);
	realizarTarea(tarea);

	//entregando
	P(protexQueuTareas);
	tareas.enqueue(tarea); // encolando la tarea
	V(protexQueuTareas); // libero el semaforo			//---> aca interviene el profesor
	//esperando la correccion del grupo
	P(estaListNotaGrupal[tarea.getNumero()]); // espero el semaforo de la tarea
	viendoNota(notasGrupal[tarea.getNumero()]); // miro la nota
}

int contCorrecciones = 0;
process profesor {
	while(contCorrecciones < 50){
		P(protexQueuTareas); // pido el semaforo para usar la Queue de tareas
		if(tareas.size() != 0){ // si hay tareas en la cola
			Tarea tarea = tareas.dequeue();
			V(protexQueuTareas); // libero el semaforo para usar la Queue de tareas
			corrigiendoTarea(tarea); // corrijo la tarea
			P(protexCantEntregasXgrupo[tarea.getNumero()]); // pido el semaforo para usar el vector de tareas
			cantEntregasXgrupo[tarea.getNumero()]++; // aumento la cantidad de entregas del grupo
			if(cantEntregasXgrupo[tarea.getNumero()] == 5){ // si el grupo ya entrego todas las tareas //tiene que darles la nota || No se va a volver a repetidor este tarea.getNumero() ya que solo es 1 nro de tarea cada 5 alumnos
				double nota = dandoNotaGrupal(); // doy la nota grupal
				notasGrupal[tarea.getNumero()] = nota; // guardo la nota grupal
				For (i = 0..4) -> V(estaListNotaGrupal[tarea.getNumero()]); // libero el semaforo para que los alumnos puedan ver la nota
				contCorrecciones+=5; // aumento el contador de correcciones, Si estoy aca ya se que hice +5, y si o si, voy a hacer 5*10 = 50
			}
			V(protexCantEntregasXgrupo[tarea.getNumero()]); // libero el semaforo para usar el vector de tareas
		}else{
			V(protexQueuTareas); // libero el semaforo para usar la Queue de tareas
		}
	}
}














---

### ✅ **Aspectos que resolviste bien:**

- **Sincronización inicial con la barrera**:
  - Usás `contAlumnos` para saber cuándo todos eligieron.
  - Liberás la barrera con `V(barrera)` *50 veces* (una por alumno), lo cual es **correcto** si cada uno va a hacer `P(barrera)` para esperar a los demás.

- **Cola de tareas protegida**:
  - Usás bien un semáforo `protexQueuTareas` para proteger la `Queue<Tarea>` al hacer `enqueue` y `dequeue`.

- **Contador de entregas por grupo**:
  - Protegido correctamente con `protexCantEntregasXgrupo[]`.

- **Esperar nota grupal con `estaListNotaGrupal[]`**:
  - Cada alumno se queda bloqueado en `P(estaListNotaGrupal[i])` y lo liberás con un `for (i = 0..4) -> V(...)` una vez que el grupo se completa. ✅

- **Corrección de tareas y conteo global `contCorrecciones`**:
  - Sumás de a 5 cuando se termina un grupo completo, lo cual **evita errores de conteo** y te asegura que el `while` en `profesor` se termina correctamente.

---

### 🔍 **Pequeñas observaciones (opcional):**

1. **Nombre del semáforo `barrera`**:
   - El nombre está bien, pero podrías usar algo como `barreraInicio` para mayor claridad semántica.

2. **Posible mejora: no hacer `V(barrera)` 50 veces**:
   - Podrías usar un semáforo de tipo "barrier clásico", donde los 50 procesos hagan `P(barrera)` y el último haga `V(barrera)` una vez, y ese `P(barrera)` esté en un loop:

     ```c
     for(i = 0..49){
       P(barrera);
     }
     ```

     Pero tu solución también es válida, aunque un poco menos elegante o económica con los recursos.

3. **`notasGrupal[] = ([9] -1)`**:
   - Bien para indicar que aún no fue corregida la tarea. Aunque no lo usás directamente como condición, es útil como indicativo.

---



