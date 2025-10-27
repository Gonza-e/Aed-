Accion insertar es 
 Ambiente 
	nodo = registro 
		num: entero 
		prox,ant: puntero a nodo 
	FReg 
	prim,ult,p,q,e: puntero a nodo 

	Procedimiento insertar() 
		Nuevo(q); *q.num:= numero
		si prim = nil entonces 
			prim:= q 
			ult:= q
			*q.ant:= nil 
			*q.prox:= nil 
		sino 
			p:= prim 
			Mientras p <> nil y *q.num < *p.num hacer 
				p:= *p.prox 
			FM 
			si p = prim entonces 
				*q.ant:= nil 
				*q.prox:= prim 
				*p.ant:= q 
				prim:= q
			sino 
				si p = nil entonces 
					*q.prox:= nil 
					*q.ant:= ult 
					*ult.prox:= q 
					ult:= q 
				sino 
					*(*p.ant).prox:= q 
					*q.ant:= *p.ant 
					*p.ant:= q
					*q.prox:= p 
				Fsi 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento eliminar() 
		si p = prim entonces 
			prim:= *p.prox
			*(p.prox).ant:= nil
		sino 
			si p = ult entonces 
				*(*ult.ant).prox:= nil
				ult:= *p.ant
			sino 
				si prim = ult entonces 
					prim:= nil 
					ult:= nil 
				sino 
					*(*p.ant).prox:= *p.prox 
					*(*p.prox).ant:= *p.ant
				Fsi 
			Fsi 
		Fsi 
		e:= p 
		p:= *p.prox
		Disponer(e)
	FProcedimiento
fin


Accion ejercicio es 
 Ambiente 
	Funcion esMultiplo(num,cont: entero): booleano
		si num = 0 entonces 
			si cont mod 3 = 0 entonces 
				esMultiplo:= verdadero 
			sino 
				esMultiplo:= falso
			Fsi 
		sino 
			si num mod 10 = 0 entonces
				esMultiplo:= esMultiplo(num div 10,cont+1)
			sino
				esMultiplo:= esMultiplo(num div 10,cont)
			Fsi 
		Fsi 
	FFuncion 
FAccion 


Accion ejercico es 
 Ambiente 
	Funcion cumple(a: arreglo[1...4] de reales; ind,cont:enteros): booleano
		si ind = 5 entonces 
			si cont <> 0 entonces 
				cumple:= falso 
			sino 
				cumple:= verdadero
			Fsi 
		sino 
			si a[ind] >= 36,5 entonces 
				cumple:= cumple(a,ind+1,cont+1)
			sino 
				cumple:= cumple(a,ind+1,cont)
			Fsi 
		Fsi 
finaccion


accion suma es
	Ambiente
		funcion multiplicacion(n1,n2,suma:entero):entero
			si n1=0 entonces	
				multiplicacion:=suma
			sino
				multiplicacion(n1-1,n2,suma+n2)
			Fsi
		FFuncion
fin

accion suma es
	Ambiente
		funcion suma(n,sum:entero):entero
			si n=0 entonces
				suma:=sum
			sino 
				suma:=suma(ndiv10, sum+nmod10)
			finsi
		FFuncion
fin

Accion verificar (prim: puntero a nodo) es 
	Ambiente 
		nodo = registro 
			num: entero 
			prox: puntero a nodo 
		FReg 
		ant,p: puntero a nodo 

		nodo2 = registro 
			num: entero 
			prox,ant: puntero a nodo2 
		FReg 
		prim2,p2,q2: puntero a nodo2 

		Procedimiento carga()
			si prim = nil entonces 
				*q.ant:= nil 
				*q.prox:= nil 
				prim:= q 
				ult:= q 
			sino 
				*q.prox:= prim
				*prim.ant:= q 
				*q.ant:= nil
				prim:= q
			Fsi 
		FProcedimiento

		Funcion esPrimo(n1,n2: entero): booleano
			si n1 < 2 entonces 
				esPrimo:= falso 
			sino 
				si n2 = 0 entonces 
					esPrimo:= verdadero 
				sino 
					si ((n1 mod n2) = 0) y (n2 <> 1) entonces 
						esPrimo:= falso
					sino 
						esPrimo:= esPrimo(n1,n2-1)
					Fsi
				Fsi 
			Fsi 
		FFuncion
finaccion


Accion ejercicio es 
	Ambiente
		Funcion esCapicua(n1,n2,n1_original:entero):booleano
			si n1 = 0 entonces 
				si n2 = n1_original entonces 
					esCapicua:= verdadero
				sino 
					esCapicua:= falso
				Fsi
			sino 
				esCapicua:= esCapicua(n1 div 10,(n2 * 10)+(n1 mod 10),n1_original)
			Fsi 
		FFuncion
FAccion 



