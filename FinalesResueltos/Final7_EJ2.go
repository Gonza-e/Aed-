Accion ej1 es 
	Ambiente 

	sec: secuencia de caracteres
	s: caracter 

	vocal: {"A","E","I","O","U","a","e","i","o","u"}
	numeros: {"0","1","2","3","4","5","6","7","8","9"}
	especiales: {"@","/","#","$"}

	nodoprin = registro 
		prin: puntero a nodo 
		voc,con,esp,num: entero 
		prox: puntero a nodoprin 
	Freg 
	prim,p,q,ant: puntero a nodo prin 

	nodo = registro 
		let: AN(1)
		prox: puntero a nodo 
	Freg 
	prim1,p1,q1,ult: puntero a nodo 

	Funcion long(x:puntero a nodo): entero 
		si x = nil entonces 
			long:= 0
		sino 
			long:= 1+long(*x.prox)
		Fsi
	FFuncion 

	Funcion consec(x:puntero a nodo; cont:entero; d1,d2:AN): booleano 
		si *x.prox = nil entonces 
			consec:= falso 
		sino 
			si cont > 3 entonces 
				consec:= verdadero 
			sino 
				d1:= *x.let; d2:= *(*x.prox).let
				si d1 = d2 entonces 
					consec:= consec(*x.prox,cont+1,d1,d2)
				sino 
					consec:= consec(*x.prox,1,d1,d2)
				Fsi 
			Fsi 
		Fsi 
	FFuncion 

	Funcion esMayor(n,a: puntero a nodo): booleano 
		si n = nil entonces
			esMayor:= falso 
		sino 
			si a = nil entonces 
				esMayor:= verdadero 
			sino 
				si *n.let = *a.let entonces 
					esMayor:= esMayor(*n.prox,*a.prox)
				sino 
					si *n.let > *a.let entonces 
						esMayor:= verdadero
					sino 
						esMayor:= falso
					Fsi 
				Fsi 
			Fsi 
		Fsi 
	FFuncion 

	A: arreglo[1...4] de entero 
	may,Nmay,num,esp,con,voc,tot: entero
	porc: real  

	Procedimiento cargarPrin()
		Nuevo(q); *q.prin:= prim1; *q.voc:= voc; *q.con:= con; *q.num:= num; *q.esp:= esp 

		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil
			Mientras (p<>nil) y esMayor(*q.*prin.let,*p.*prin.let) hacer 
				ant:= p 
				p:= *p.prox 
			FM 
			
			si prim = p entonces 
				*q.prox:= prim 
				prim:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= p 
			Fsi 
		Fsi 

		voc:= 0; con:= 0; num:= 0; esp:= 0
		prim1:= nil 
	FProcedimiento

	Procedimiento cargarSec()
		si s en numeros entonces 
			num:= num+1
		sino 
			si s en esp entonces 	
				esp:= esp+1 
			sino 
				si s en vocal entonces 
					voc:= voc+1
				sino 
					con:= con+1 
				Fsi 
			Fsi 
		Fsi 
		Nuevo(q1); *q1.let:= s

		si prim1 = nil entonces 
			*q1.prox:= prim1 
			prim1:= q1 
			ult:= q1 
		sino 
			*q1.prox:= *ult.prox 
			*ult.prox:= q1 
			ult:= q1 
		Fsi 
	FProcedimiento 

	Procedimiento verificar()
		p1:= *p.prin 
		si (long(p1)>10) y (*p.voc<>0 y *p.con<>0 y *p.esp<>0 y *p.num<>0) entonces 
			A[4]:= A[4]+1
		sino 
			si (long(p1)>8 y long(p1)<10) y (*p.voc<>0 *p.con<>0 y *p.num<>0) entonces 
				A[3]:= A[3]+1
			sino 
				si (long(p1)>6 y long(p1)<8) y (*p.voc+*p.con > *p.num) y (*p.num>0)
					A[2]:= A[2]+1
				sino 
					A[1]:= A[1]+1
				Fsi 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento buscarMay()
		Para i:= 1 hasta 4 hacer 
			si A[i] > may entonces 
				Nmay:= i 
				may:= A[i]
			Fsi 
		FPara 
	FProcedimiento

	Procedimiento borrar()
		p1:= prim1 
		Mientras p1<>nil hacer 
			prim1:= *prim1.prox 
			Disponer(p1)
			p1:= prim1 
		FM 
		voc:= 0; con:= 0; esp:= 0; num:= 0
	FProcedimiento

	Proceso 
		Arr(sec); Avz(sec,s)
		may:= 0; Nmay:= 0; num:= 0; con:= 0; esp:= 0; voc:= 0; tot:= 0
		prim:= nil; prim1:= nil 

		Mientras NFDS(sec) hacer 
			Mientras s <> "+" hacer 
				cargarSec()
				Avz(sec,s)
			FM 

			si prim1 = ult entonces 
				cargarPrin()
			sino 
				borrar()
			Fsi 
			Avz(sec,s)
		FM 
		Cerrar(sec)

		p:= prim 
		Mientras p <> nil hacer 
			p1:= *p.prin 
			si consec(p1,0,nil,nil) entonces 
				tot:= tot+1 
			Fsi 
			p:= *p.prox 
		FM 

		porc:= (tot/long(prim))*100
		Esc("Porcentaje: ",porc,"%")

		buscarMay()
		Esc("Nivel con mayor cantidad de contraseñas: ",i)
	FProceso 
FAccion 

