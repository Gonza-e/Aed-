"7. A partir de un arreglo de entrada de 100 posiciones de enteros, se pide cargar en
una lista los elementos del arreglo que son impares. Para la verificación de paridad
se pide desarrollar una función recursiva."

Accion recursiva (numeros: arreglo[1...10] de enteros) es 
 Ambiente
	nodo = registro 
		num: entero 
		prox: puntero a nodo 
	FReg
	prim,q,p: puntero a nodo 

	Procedimiento cargarNodo(numero)
		Nuevo(q); *q.num:= numero
		*q.prox:= prim 
		prim:= q
	FProcedimiento

	Funcion esPar(n:entero):booleano
		si n mod 2 = 0 entonces 
			esPar:= verdadero
		sino 
			esPar:= falso 
		Fsi
	FFuncion

	Procedimiento verificar(a: arreglo[1...100] de enteros; indice: entero)
		si ind < 101 entonces 
			si esPar(a[indice]) entonces 
				cargarNodo(a[indice])
			Fsi
		Fsi

		verificar(a,indice + 1)
	FProcedimiento

	Proceso 
		prim:= nil 
		verificar(numeros,0)

		p:= prim 
		Mientras p <> nil hacer 
			Esc(*p.num)
		FM 
	FProceso 
FAccion

