Accion capicua es 
	Ambiente 

	sec,sal: secuencia de caracteres 
	s: caracter 

	nodoprin = registro	
		prin: puntero a nodo
		ult1: puntero a nodo  
		prox: puntero a nodoprin 
	Freg 
	prim1,p1,q1: puntero a nodoprin 

	nodo = registro
		let: AN(1)
		prox,ant: puntero a nodo 
	Freg
	prim,p,q,ult: puntero a nodo 

	Procedimiento cargarSec()
		si prim = nil entonces 
			*q.prox:= nil 
			*q.ant:= nil 
			ult:= q 
			prim:= q 
		sino 
			*q.prox:= *ult.prox 
			*q.ant:= ult 
			*ult.prox:= q 
			ult:= q 
		Fsi 
	FProcedimiento 

	Procedimiento cargarPrin()
		Nuevo(q1); *q.prin:= prim; *q.ult1:= ult; prim:= nil; ult:= nil 
		*q1.prox:= prim1
		prim1:= q1 
	FProcedimiento

	Funcion capicua(p,u: puntero a nodo): booleano 
		si p = u o (*p.ant = u) entonces 
			capicua:= verdadero 
		sino 
			si *p.let = *u.let entonces 
				capicua:= capicua(*p.prox,*u.ant)
			sino 
				capicua:= falso 
			Fsi 
		Fsi 
	FFuncion 

	Proceso 
		Arr(sec); Avz(sec,s); Crear(sal)
		prim:= nil; prim1:= nil 

		Mientras NFDS(sec) hacer
			Mientras s <> "+" hacer 
				Nuevo(q); *q.let:= s 
				cargarSec()
				Avz(sec,s) 
			FM 
			cargarPrin()
			Avz(sec,s)
		FM 
		Cerrar(sec)

		p1:= prim1 
		Mientras p1 <> nil hacer 
			prim:= *p1.prin; ult:= *p1.ult1 
			si capicua(prim,ult) entonces 
				Mientras prim <> nil hacer 
					Esc(sal,*prim.let)
				FM 
			Fsi 
			p1:= *p1.prox 
		FM 
		Cerrar(sal)
	FProceso 
FAccion 