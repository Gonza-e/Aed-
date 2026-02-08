Accion sufri (primC: puntero a nodoC) es 
 Ambiente 
	nodoC = registro 
		beneficio: AN(30)
		prox: puntero a nodoC 
	Freg 
	pC: puntero a nodoC 

	nodo = registro 
		regional: N(2)
		leg: N(5)
		nombre: AN(30)
		matCursadas: N(3)
		añoIng: N(4)
		prox: puntero a nodo 
	Freg 
	prim,ant,p,q: puntero a nodo 

	ALUMNO = registro 
		legajo: N(5)
		nom: AN(30)
		ape: AN(30)
		año: N(4)
	Freg 
	arch: archivo de ALUMNO indexado por legajo 
	reg: ALUMNO

	sec: secuencia de caracteres 
	s: caracter 
	regional,apeynom: AN 
	legajo,cantM: entero 
	A: arreglo[1...10] de enteros 

	Procedimiento cargarArreglo()
		A[*q.regional].cantAlumnos:= A[*q.regional].cantAlumnos + 1 
		A[*q.regional].cantMaterias:= A[*q.regional].cantMaterias + cantM
	FProcedimiento 

	Funcion conv(x:AN): entero 
		segun x hacer 
			"1": conv:= 1 
			"2": conv:= 2 
			"3": conv:= 3
			"4": conv:= 4
			"5": conv:= 5
			"6": conv:= 6
			"7": conv:= 7 
			"8": conv:= 8
			"9": conv:= 9 
			"0": conv:= 0 
		Fsegun 
	FFuncion 

	Procedimiento obtDatos()
		Mientras s <> "-" hacer 
			regional:= concatenar(regional,s)
			Avz(sec,s)
		FM 
		Avz(sec,s)

		Mientras s <> "+" hacer 
			Mientras s <> "-" hacer 
				legajo:= legajo*10 + conv(s)
				Avz(sec,s)
			FM 
			Avz(sec,s)
			
			Mientras s <> "*" hacer 
				apeynom:= concatenar(apeynom,s)
				Avz(sec,s)
			FM 
			Avz(sec,s)

			Para i:= 1 hasta 3 hacer 
				cantM:= cantM*10 + conv(s)
				Avz(sec,s)
			FPara 
			Avz(sec,s)
		FM 
	FProcedimiento 

	Procedimiento verifBeneficio()
		si cantM > 30 entonces
			pC:= primC 
			Para i:= 1 hasta beneficio(legajo) hacer 
				pC:= *pC.prox 
			FPara 
			Esc("El beneficio del alumno es:",*pC.beneficio) 
		Fsi 
	FProcedimiento 

	Procedimiento cargarLista()
		Nuevo(q); *q.regional:= numero(regional); *q.leg:= legajo; *q.apeynom:= nombre; *q.matCursadas:= cantM
		reg.legajo:= legajo; Leer(arch,reg)
		si EXISTE entonces 
			*q.añoIng:= *q.año 
		Fsi 

		cargarArreglo()

		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q
		sino 
			ant:= nil; p:= prim 
			Mientras p <> nil y (*q.regional > *p.regional) hacer 
				ant:= p 
				p:= *p.prox 
			FM 
			
			si p = prim entonces 
				*q.prox:= p 
				prim:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= p 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento init()
		Para i:= 1 hasta 10 hacer 
			A[i].cantAlumnos:= 0; A[i].cantMaterias:= 0 
		FPara
		regional:= ""; nombre:= ""; legajo:= 0; cantM:= 0
		prim:= nil 
	FProcedimiento

	Proceso 
		init()
		Arr(sec); Avz(sec,s)
		 
		Mientras NFDS(sec) hacer 
			obtDatos()
			verifBeneficio()
			cargarLista()
			regional:= ""; nombre:= ""; legajo:= 0; cantM:= 0
		FM 
		Cerrar(sec)

		p:= prim 
		Mientras p <> nil hacer 
			Esc("Alumno:",*p.nom,"Legajo:",*p.leg)
			si *(*p.prox).regional <> *p.regional entonces
				Esc("Regional:",*p.regional,"Cantidad de alumnos:",A[*p.regional].cantAlumnos,"Cantidad de materias cursadas:",A[*p.regional].cantMaterias)
				Esc("Promedio de materias cursadas:",A[*p.regional].cantMaterias/A[*p.regional].cantAlumnos)
			Fsi 
			p:= *p.prox 
		FM 





