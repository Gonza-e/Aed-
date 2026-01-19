Accion ej1 es 
 Ambiente 
	NOTAS = registro 
		carrera: (1...5)
		legajo: N(5)
		materia: (1...50)
		nota: N(2)
		f_aprobacion: Fecha 
		plan: N(4)
	Freg 
	arch_n: archivo de NOTAS ordenado por carrera, legajo y materia 
	re_n: NOTAS 

	ALUMNOS = registro 
		legajo: N(5)
		nombre: AN(30)
		f_ingreso: Fecha 
		mat_aprobadas: N(3)
		email: AN(30)
	Freg 
	arch_a: archivo de ALUMNOS indexado por legajo 
	reg_a: ALUMNOS

	Funcion cond(x,n:entero): AN 
		si (x > 6) o (n = 1) entonces 
			cond:= "Promocionado"
		sino 
			si (x = 6) o (n = 2) entonces 
				cond:= "Regular"
			sino 
				si ((x < 6) y (x > 0)) o (n = 3)
					cond:= "Recupera"
				Fsi 
			Fsi 
		Fsi 
	FFuncion 

	Procedimiento cargarCond()
		si x > 6 entonces 
			i:= 1 
		sino 
			si x = 6 entonces 
				i:= 2 
			sino 
				i:= 3 
			Fsi 
		Fsi 
		A[i]:= A[i] + 1
	FProcedimiento

	nodo = registro 
		nombre: AN(20)
		prom: N(1,2)
		prox: puntero a nodo 
	Freg 
	prim, p, q, ant: puntero a nodo 

	Procedimiento cargarLista()
		Nuevo(q); *q.prom:= prom_leg
		reg_a.legajo:= res_legajo 
		Leer(arch_a,reg_a)
		si EXISTE entonces 
			*q.nombre:= reg_a.nombre 
		Fsi 
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras p <> nil y *q.prom < *p.prom hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			si prim = p entonces 
				*q.prox:= p 
				prim:= q 
			sino 
				*q.prox:= *ant.prox 
				*ant.prox:= q 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento corte_mat()
		prom:= prom/cont
		reg_a.legajo:= res_legajo 
		Leer(arch_a,reg_a)
		si EXISTE entonces  
			Esc("Nombre: ",reg_a.nombre,"Condicion: ",cond(prom,0))
		Fsi 
		Si prom >= 6 entonces 
			cont_mat:= cont_mat + 1 
		Fsi 
		prom_leg:= prom_leg + prom 
		prom:= 0 
		cont:= 0
		res_mat:= reg_n.mat 
	FProcedimiento

	Procedimiento corte_legajo()
		corte_mat()
		Para i:= 1 hasta 3 hacer 
			A[i]:= (A[i]/50) * 100
			Esc(cond(-1,i),":",A[i],"%")
			A[i]:= 0
		FPara
		si res_carrera = 1 entonces 
			cargarLista()
		Fsi 
		reg_a.legajo:= res_legajo 
		Leer(arch_a,reg_a)
		si EXISTE entonces 
			reg_a.mat_aprobadas:= cont_mat
			cont_mat:= 0
		Fsi 
		Esc("Promedio general: ",prom_leg/50)
		prom_carrera:= prom_carrera + prom_leg  
		cont_alumnos:= cont_alumnos + 1
		prom_leg:= 0
		res_legajo: reg_n.legajo 
	FProcedimiento 

	Procedimiento corte_carrera()
		Esc("Promedio general carrera",res_carrera,":",prom_carrera/cont_alumnos)
		prom_carrera:= 0 
		cont_alumnos:= 0 
		res_carrera:= reg_n.carrera 
	FProcedimiento

	Procedimiento inicializar()
		prim:= nil 
		Para i:= 1 hasta 3 hacer 
			A[i]:= 0  
		FPara 
		res_carrera:= reg_n.carrera; res_legajo:= reg_n.legajo; res_mat:= reg_n.mat
		prom:= 0; prom_leg:= 0; prom_carrera:= 0
		cont:= 0; cont_alumnos:= 0
	FProcedimiento

	res_carrera, res_legajo, res_mat, cont, cont_alumnos, cont_mat: entero 
	prom, prom_leg, prom_carrera: real 
	A: arreglo[1...3] de real

	Proceso 
		Abrir E/ (arch_n); Leer(arch_n,reg_n)
		Abrir E/S (arch_a)
		inicializar()

		Mientras NFDA(arch_n) hacer 
			si res_carrera <> reg_n.carrera entonces 
				corte_carrera()
			sino 
				si res_legajo <> reg_n.legajo entonces 
					corte_legajo()
				sino 
					si res_mat <> reg_n.materia entonces 
						corte_mat()
					Fsi 
				Fsi 
			Fsi 

			prom:= prom + reg_n.nota
			cont:= cont + 1

			Leer(arch_n,reg_n)
		FM 
		corte_carrera()
		Cerrar(arch_a); Cerrar(arch_n)

		p:= prim 
		Para i:= 1 hasta 10 hacer 
			Esc("Nombre:",*p.nombre;"Promedio:",*p.prom)
			p:= *p.prox 
		FPara 
	FProceso 
FAccion

