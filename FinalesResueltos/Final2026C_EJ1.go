Accion ej1 es 
	Ambiente 

	sec: secuencia de caracteres
	s: caracter 

	nodo = registro 
		codigo: AN(10)
		suc: N(4)
		dir: AN(50)
		peso: AN(6)
		estado: ("P","T","E","D")
		prox: puntero a nodo 
	Freg 
	primR,primNR,p,q: puntero a nodo 
	cT,cR: entero 

	A: arreglo[1...4] de entero 

	Procedimiento indice(x:AN)
		segun x hacer 
			"P": i:= 1 
			"T": i:= 2 
			"E": i:= 3 
			"D": i:= 4
		Fsegun 
		A[i]:= A[i]+1 
	FP 

	Funcion estado(x:entero): AN 
		segun x hacer 
			1: estado:= "Pendiente"
			2: estado:= "Transito"
			3: estado:= "Entregado"
			4: estado:= "Devuelto"
		Fsegun 
	FF

	Procedimiento obtInfo()
		Para i:= 1 hasta 10 hacer 
			cod:= concatenar(cod,s)
			Avz(sec,s)
		FPara 

		Para i:= 1 hasta 4 hacer 
			suc:= concatenar(suc,s)
			Avz(sec,s)
		FPara 

		Para i:= 1 hasta 50 hacer 
			dir:= concatenar(dir,s)
			Avz(sec,s)
		FPara
		
		Para i:= 1 hasta 7 hacer 
			peso:= concatenar(peso,s)
			Avz(sec,s)
		FPara 

		estado:= s 
	FP 

	Procedimiento crearNodo()
		Nuevo(q); *q.codigo:= cod; *q.suc:= suc; *q.dir:= dir; *q.peso:= peso; *q.estado:= estado 
	FP 

	Procedimiento cargar(x:puntero a nodo)
		si ((estado = "P") o (estado = "T")) y EsEnvioRiesgoso(cod,peso) entonces 
			crearNodo()
			*q.prox:= primR 
			primR:= q 
			cR:= cR+1 
		sino 
			si (estado = "D") y NO EsEnvioRiesgoso(cod,peso) entonces 
				crearNodo()
				*q.prox:= primNR 
				primNR:= q 
			sino 
				si NO EsEnvioRiesgoso(cod,peso) entonces 
					indice()
				Fsi 
			Fsi 
		Fsi 
		cT:= cT+1 
	FP 

	Procedimiento init()
		Para i:= 1 hasta 4 hacer 
			A[i]:= 0 
		FPara 
		cT:= 0; cR:= 0
		primR:= nil; primNR:= nil 
		cod:= ""; suc:= ""; dir:= ""; peso:= ""; estado:= ""
	FP 

	Proceso 
		Arr(sec); Avz(sec,s)
		init()

		Mientras s <> "@" hacer 
			Mientras (s <> "#") y s <> "@" hacer
				Mientras (s <> "/") y s <> "#" hacer 
					obtInfo()
					cargar()
					cod:= ""; suc:= ""; dir:= ""; peso:= ""
				FM 
				Avz(sec,s)
			FM 
			Avz(sec,s)
		FM 
		Cerrar(sec)

		Esc("Porcentaje de envios riesgosos sobre el total:",(cR/cT)*100,"%")

		Para i:= 1 hasta 4 hacer 
			Esc("Estado de envio:",estado(i),"Cantidad de envios",A[i])
		FPara 

	FProceso 
FAccion 



