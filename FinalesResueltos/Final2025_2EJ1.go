Accion ej1 es 
	Ambiente 
	
	sec,sal1,sal2: secuencia de caracteres
	s: caracter 
	A: arreglo[1...4] de enteros 

	nodo = registro 
		char: AN(1)
		prox: puntero a nodo 
	Freg 
	prim,p,q,ant,x: puntero a nodo 

	Procedimiento cargarEncolada()
		Nuevo(q); *q.char:= s 
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			ant:= nil; p:= prim 
			Mientras p <> nil hacer
				ant:= p 
				p:= *p.prox 
			FM 

			*q.prox:= *ant.prox
			*ant.prox:= q 
		Fsi 
	FProcedimiento 

	Procedimiento guardarPatente()
		repetir 
			cargarEncolada()
			Avz(sec,s)
		Hasta que (s = "-")
	FProcedimiento

	Procedimiento grabarSec1()
		p:= prim
		Mientras *p.char <> "+" hacer 
			Esc(sal1,*p.char)
		FM 
		Esc(sal1,*p.char)
	FProcedimiento

	Procedimiento grabarSec2()
		p:= prim 
		Mientras *p.char <> "+" hacer 
			Esc(sal2,*p.char)
		FM 
		Esc(sal2,*p.char)
	FProcedimiento

	Procedimiento verMayor()
		Para i:= 1 hasta 4 hacer 
			si A[i] > mayor entonces 
				mayor:= A[i]
				tipo:= i 
			Fsi 
		FPara 
	FProcedimiento

	Funcion obtTipo(x: entero): AN 
		segun x hacer 
			1: obtTipo:= "A"
			2: obtTipo:= "B"
			3: obtTipo:= "C"
			4: obtTipo:= "D"
		Fsegun 
	FFuncion 

	zona1,inf,masde2: booleano 
	contz1,cont2,porc1,porc2,total,mayor,tipo: entero 
	Proceso 
		Arr(sec); Avz(sec,s)
		Crear(sal1); Crear(sal2)
		zona1:= falso; inf:= falso; masde2:= falso 
		mayor:= 0; tipo:= 0; contz1:= 0; cont2:= 0; porc1:= 0; porc2:= 0 
		Para i:= 1 hasta 4 hacer 
			A[i]:= 0 
		FPara 

		Mientras NFDS(sec) hacer 
			guardarPatente()

			cargarEncolada()
			segun s hacer 
				"A": A[1]:= A[1]+1
				"B": A[2]:= A[2]+1 
				"C": A[3]:= A[3]+1 
				"D": A[4]:= A[4]+1 
			Fsegun
			Avz(sec,s)

			si convEntero(s) > 2 entonces
				masde2:= verdadero 
			Fsi 
			cargarEncolada()
			Avz(sec,s)

			si s = "S" entonces 
				inf:= verdadero 
			Fsi 

			Para i:= 1 hasta 2 hacer 
				cargarEncolada()
				Avz(sec,s)
			FPara  

			si s = "1" entonces 
				zona1:= verdadero 
			Fsi 
			cargarEncolada()
			Avz(sec,s)
			cargarEncolada()
			Avz(sec,s)

			si inf y masde2 entonces 
				grabarSec2()
				inf:= falso; masde2:= falso 
				cont2:= cont2+1 
			sino 
				si zona1 entonces 
					grabarSec1()
					zona1:= falso
					contz1:= contz1+1  
				Fsi 
			Fsi 

			total:= total+1 
		FM 
		Cerrar(sec); Cerrar(sal1); Cerrar(sal2)

		verMayor()
		Esc("El tipo que mas circulo es:",obtTipo(tipo))

		porc1:= (contz1/total)*100; porc2:= (cont2/total)*100

		Esc("Cantidad de vehiculos zona 1:",contz1,"Porcentaje:",porc1,"%")
		Esc("Cantidad de vehiculos con infraccion y circulacion mayor a 2 horas:",cont2,"porcentaje:",porc2,"%")

	FProceso 
FAccion




		 


