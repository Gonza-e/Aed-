Procedimiento cargaOrdenada()
	si prim = nil entonces 
		prim:= q
		*q.prox:= nil 
	sino 
		ant:= nil; p:= prim 
		Mientras (p <> nil) y (*q.dato < *p.dato) hacer 
			ant:= p 
			p:= *p.prox 
		FM 

		si ant = nil entonces // o prim = nil 
			*q.prox:= p 
			prim:= q 
		sino // insercion en el medio o al final
			*ant.prox:= q 
			*q.prox:= p 
		Fsi 
	Fsi 
FProcedimiento

Procedimiento eliminarNodo()
	si prim = nil entonces 
		Esc("Error lista vacia")
	sino 
		p:= prim; ant:= nil 
		Mientras (p <> nil) y (*p.dato <> dato) hacer 
			ant:= p 
			p:= *p.prox 
		FM 

		si p = prim entonces 
			prim:= *prim.prox 
		sino
			*ant.prox:= *p.prox 
		fsi 
		
		e:= p 
		p:= *p.prox 
		Disponer(e)
	Fsi 
FProcedimiento
		
Procedimiento cargaApilada() 
	*q.prox:= prim 
	prim:= q 
FProcedimiento

Procedimiento cargaEncolada()
	si prim = nil entonces 
		prim:= q 
		*q.prox:= nil 
	sino 
		ant:= nil; p:= prim 
		Mientras (p <> nil) hacer 
			ant:= p 
			p:= *p.prox 
		FM 

		*ant.prox:= q 
		*q.prox:= p 
	Fsi 
FProcedimiento

Procedimiento enOrden(p: puntero a nodo)
	si p <> nil entonces
		enOrden(*p.izquierda)
		tratarRaiz()
		enOrden(*p.derecha)
	Fsi 
FProcedimiento

Procedimiento preOrden(p: puntero a nodo)
	si p <> nil entonces
		tratarRaiz()
		preOrden(*p.izquierda)
		preOrden(*p.derecha)
	Fsi 
FProcedimiento

Procedimiento postOrden(p: puntero a nodo)
	si p <> nil entonces
		postOrden(*p.izquierda)
		postOrden(*p.derecha)
		tratarRaiz()
	Fsi 
FProcedimiento

// Algoritmo de corte de control 
Abrir E/(arch), Leer(arch,reg)
res_N:= reg.claveN
res_1:= reg.clave1
res_2:= reg.clave2 
Mientras NFDA(arch) hacer 
	tratarCorte()
	tratarRegistro()
	Leer(arch,reg)
FM 
corte_N()

Procedimiento tratarCorte()
	si res_N <> reg.claveN entonces 
		corte_N()
	sino 
		si res_2 <> reg.clave2 entonces 
			corte_2()
		sino 
			si res_1 <> reg.clave1 entonces 
				corte_1()
			Fsi 
		Fsi 
	Fsi 
FProcedimiento

Procedimiento corte_N()
	corte_1() // Llamada del corte de inmediato inferior
	mostrarSubtotales()
	acumular() // Acumular para el nivel superior
	iniciarlizarContadoresDeNivel() // Poner a cero los contadores mostrados/usados
	resguardarClave() // Actualizar el resguardo 
FProcedimiento

// Ciclo incluyente

Mientras (reg1.clave <> HV) o (reg2.clave <> HV) hacer 
	proceso()
FM

// Ciclo excluyente 

Mientras NFDA(arch1) y NFDA(arch2) hacer 
	procesoRegistrosComunes()
FM 

// Uno de estos ciclos por cada fichero interviniente
Mientras NFDA(arch1) hacer 
	procesoRegdelArchivo1()
FM 
Mientras NFDA(arch2) hacer 
	procesoRegdelArchivo2()
FM 

// Busqueda lineal

Para i:= 1 hasta N hacer 
	si dato = A[i] entonces 
		tratarExito()
		j:= j+1
	Fsi 
FPara 

si j = 0 entonces 
	tratarFracaso()
Fsi 

// Busqueda lineal con centinela

Mientras (i < N) y (dato <> A[i]) hacer 
	i:= i + 1 
FM 

si dato = A[i] entonces 
	tratarExito()
sino 
	tratarFracaso()
Fsi 

// Busqueda binaria o dicotomica 

izq:= 1
der:= N 
medio:= (izq+der) div 2

Mientras (izq < der) y (A[medio] <> dato) hacer 
	si dato < A[medio] hacer 
		der:= medio-1
	sino	
		izq:= medio-1 
	Fsi 
	
	medio:= (izq+der) div 2 
FM 
	
// Algoritmo de seleccion directa

Para i:= 1 hasta N-1 hacer 
	x:= A[i]
	Para j:= i+1 hasta N hacer 
		si x < A[j] entonces 
			min:= j 
			x:= A[j]
		Fsi 

		A[min]:= A[i]
		A[i]:= x
	FPara 
FPara

Procedimiento invertirArreglo(A: arreglo[1...10] de enteros; ind,x: entero)
	si ind < 11 entonces 
		x:= A[11-ind]
		invertirArreglo(A,ind+1,x)
		A[ind]:= x 
	Fsi 
FProcedimiento 

Procedimiento cargaEncolada()
	si prim = nil entonces 
		prim:= q
		*q.prox:= nil  
	sino 
		*q.prox:= nil 
		*ult.prox:= q 
	Fsi 
	ult:= q 
FProcedimiento 

Procedimiento insercionDirecta()
	Para i:= 2 hasta N hacer 
		actual:= A[i]
		j:= i-1
		Mientras (j>0) y A[j]<actual hacer 
			A[j+1]:= A[j]
			j:= j-1
		FM 
		
		A[j+1]:= actual 
	FPara 
FProcedimiento

Procedimiento intercambioDirecto()
	bandera:= falso 
	Mientras NO bandera hacer 
		bandera:= verdadero 
		Para i:= 1 hasta N hacer 
			si A[i] > A[i+1] entonces 
				x:= A[i]
				A[i]:= A[i+1]
				A[i+1]:= x 
				bandera:= falso 
			Fsi 
		FPara 
	FM 
FProcedimiento

nodo1 = registro 
	dato: AN(30)
	prox: puntero a nodo1 
Freg 
p: puntero a nodo1

nodo2 = registro 
	prim: puntero a nodo1 
	prox: puntero a nodo2 
Freg 
q: puntero a nodo2

p:= *q.*prim.dato  